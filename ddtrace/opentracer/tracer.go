// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2016-2019 Datadog, Inc.

// Package opentracer provides a wrapper on top of the Datadog tracer that can be used with Opentracing.
// It also provides a set of opentracing.StartSpanOption that are specific to Datadog's APM product.
// To use it, simply call "New".
//
// Note that there are currently some small incompatibilities between the Opentracing spec and the Datadog
// APM product, which we are in the process of addressing on the long term. When using Datadog, the
// Opentracing operation name is what is called resource in Datadog's terms and the Opentracing "component"
// tag is Datadog's operation name. Meaning that in order to define (in Opentracing terms) a span that
// has the operation name "/user/profile" and the component "http.request", one would do:
//  opentracing.StartSpan("http.request", opentracer.ResourceName("/user/profile"))
//
// Some libraries and frameworks are supported out-of-the-box by using our integrations. You can see a list
// of supported integrations here: https://godoc.org/gopkg.in/DataDog/dd-trace-go.v1/contrib. They are fully
// compatible with the Opentracing implementation.
package opentracer

import (
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"

	opentracing "github.com/opentracing/opentracing-go"
)

// New creates, instantiates and returns an Opentracing compatible version of the
// Datadog tracer using the provided set of options.
func New(opts ...tracer.StartOption) opentracing.Tracer {
	return tracer.NewOpenTracer(opts...)
}
