// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package fkpulsarexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/pulsarexporter"

import (
	"bytes"

	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/gogo/protobuf/jsonpb"
	jaegerproto "github.com/jaegertracing/jaeger/model"
	"go.opentelemetry.io/collector/pdata/ptrace"
	"go.uber.org/multierr"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/translator/jaeger"
)

type jaegerMarshaler struct {
	marshaler jaegerSpanMarshaler
}

var _ TracesMarshaler = (*jaegerMarshaler)(nil)

func (j jaegerMarshaler) Marshal(traces ptrace.Traces, _ string) ([]*pulsar.ProducerMessage, error) {
	batches, err := jaeger.ProtoFromTraces(traces)
	if err != nil {
		return nil, err
	}

	var errs error
	var messages []*pulsar.ProducerMessage

	for _, batch := range batches {
		for _, span := range batch.Spans {
			span.Process = batch.Process
			bts, err := j.marshaler.marshal(span)
			// continue to process spans that can be serialized
			if err != nil {
				errs = multierr.Append(errs, err)
				continue
			}
			key := span.TraceID.String()
			messages = append(messages, &pulsar.ProducerMessage{
				Payload: bts,
				Key:     key,
			})
		}
	}

	return messages, errs
}

func (j jaegerMarshaler) Encoding() string {
	return j.marshaler.encoding()
}

type jaegerSpanMarshaler interface {
	marshal(span *jaegerproto.Span) ([]byte, error)
	encoding() string
}

type jaegerProtoSpanMarshaler struct {
}

var _ jaegerSpanMarshaler = (*jaegerProtoSpanMarshaler)(nil)

func (p jaegerProtoSpanMarshaler) marshal(span *jaegerproto.Span) ([]byte, error) {
	return span.Marshal()
}

func (p jaegerProtoSpanMarshaler) encoding() string {
	return "jaeger_proto"
}

type jaegerJSONSpanMarshaler struct {
	pbMarshaler *jsonpb.Marshaler
}

var _ jaegerSpanMarshaler = (*jaegerJSONSpanMarshaler)(nil)

func newJaegerJSONMarshaler() *jaegerJSONSpanMarshaler {
	return &jaegerJSONSpanMarshaler{
		pbMarshaler: &jsonpb.Marshaler{},
	}
}

func (p jaegerJSONSpanMarshaler) marshal(span *jaegerproto.Span) ([]byte, error) {
	out := new(bytes.Buffer)
	err := p.pbMarshaler.Marshal(out, span)
	return out.Bytes(), err
}

func (p jaegerJSONSpanMarshaler) encoding() string {
	return "jaeger_json"
}
