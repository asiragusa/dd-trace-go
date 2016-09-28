package tracer

import (
	"golang.org/x/net/context"
)

const (
	datadogSpanKey = "datadog_trace_span"
)

// ContextWithSpan populates the given Context with a Span using an internal
// datadogActiveSpanKey. This method is a good helper to pass parent spans
// in a new function call, to simplify the creation of a child span.
func ContextWithSpan(ctx context.Context, span *Span) context.Context {
	return context.WithValue(ctx, datadogSpanKey, span)
}

// SpanFromContext returns the stored *Span from the Context if it's available.
// This helper returns also the ok value that is true if the span is present.
func SpanFromContext(ctx context.Context) (*Span, bool) {
	span, ok := ctx.Value(datadogSpanKey).(*Span)
	return span, ok
}

// SpanFromContextDefault returns the stored *Span from the Context. If not, it
// will return an empty span that will do nothing.
func SpanFromContextDefault(ctx context.Context) *Span {

	// FIXME[matt] is it better to return a singleton empty span?
	if ctx == nil {
		return &Span{}
	}

	span, ok := ctx.Value(datadogSpanKey).(*Span)
	if !ok {
		return &Span{}
	}
	return span
}