package progress

import (
	"context"
	"regexp"
	"strings"
	"time"

	"github.com/docker/buildx/util/metricutil"
	"github.com/moby/buildkit/client"
	"github.com/opencontainers/go-digest"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
)

type metricWriter struct {
	recorders []metricRecorder
	attrs     attribute.Set
}

func newMetrics(mp metric.MeterProvider, attrs attribute.Set) *metricWriter {
	meter := metricutil.Meter(mp)
	return &metricWriter{
		recorders: []metricRecorder{
			newLocalSourceTransferMetricRecorder(meter, attrs),
		},
		attrs: attrs,
	}
}

func (mw *metricWriter) Write(ss *client.SolveStatus) {
	for _, recorder := range mw.recorders {
		recorder.Record(ss)
	}
}

type metricRecorder interface {
	Record(ss *client.SolveStatus)
}

type (
	localSourceTransferState struct {
		// Attributes holds the attributes specific to this context transfer.
		Attributes attribute.Set

		// LastTransferSize contains the last byte count for the transfer.
		LastTransferSize int64
	}
	localSourceTransferMetricRecorder struct {
		// BaseAttributes holds the set of base attributes for all metrics produced.
		BaseAttributes attribute.Set

		// State contains the state for individual digests that are being processed.
		State map[digest.Digest]*localSourceTransferState

		// TransferSize holds the metric for the number of bytes transferred.
		TransferSize metric.Int64Counter

		// Duration holds the metric for the total time taken to perform the transfer.
		Duration metric.Float64Counter
	}
)

func newLocalSourceTransferMetricRecorder(meter metric.Meter, attrs attribute.Set) *localSourceTransferMetricRecorder {
	mr := &localSourceTransferMetricRecorder{
		BaseAttributes: attrs,
		State:          make(map[digest.Digest]*localSourceTransferState),
	}
	mr.TransferSize, _ = meter.Int64Counter("source.local.transfer.io",
		metric.WithDescription("Measures the number of bytes transferred between the client and server for the context."),
		metric.WithUnit("By"))

	mr.Duration, _ = meter.Float64Counter("source.local.transfer.time",
		metric.WithDescription("Measures the length of time spent transferring the context."),
		metric.WithUnit("ms"))
	return mr
}

func (mr *localSourceTransferMetricRecorder) Record(ss *client.SolveStatus) {
	for _, v := range ss.Vertexes {
		state, ok := mr.State[v.Digest]
		if !ok {
			attr := detectLocalSourceType(v.Name)
			if !attr.Valid() {
				// Not a context transfer operation so just ignore.
				continue
			}

			state = &localSourceTransferState{
				Attributes: attribute.NewSet(attr),
			}
			mr.State[v.Digest] = state
		}

		if v.Started != nil && v.Completed != nil {
			dur := float64(v.Completed.Sub(*v.Started)) / float64(time.Millisecond)
			mr.Duration.Add(context.Background(), dur,
				metric.WithAttributeSet(mr.BaseAttributes),
				metric.WithAttributeSet(state.Attributes),
			)
		}
	}

	for _, status := range ss.Statuses {
		state, ok := mr.State[status.Vertex]
		if !ok {
			continue
		}

		if strings.HasPrefix(status.Name, "transferring") {
			diff := status.Current - state.LastTransferSize
			if diff > 0 {
				mr.TransferSize.Add(context.Background(), diff,
					metric.WithAttributeSet(mr.BaseAttributes),
					metric.WithAttributeSet(state.Attributes),
				)
			}
		}
	}
}

var reLocalSourceType = regexp.MustCompile(
	strings.Join([]string{
		`(?P<context>\[internal] load build context)`,
		`(?P<dockerfile>load build definition)`,
		`(?P<dockerignore>load \.dockerignore)`,
		`(?P<namedcontext>\[context .+] load from client)`,
	}, "|"),
)

func detectLocalSourceType(vertexName string) attribute.KeyValue {
	match := reLocalSourceType.FindStringSubmatch(vertexName)
	if match == nil {
		return attribute.KeyValue{}
	}

	for i, source := range reLocalSourceType.SubexpNames() {
		if len(source) == 0 {
			// Not a subexpression.
			continue
		}

		// Did we find a match for this subexpression?
		if len(match[i]) > 0 {
			// Use the match name which corresponds to the name of the source.
			return attribute.String("source.local.type", source)
		}
	}
	// No matches found.
	return attribute.KeyValue{}
}
