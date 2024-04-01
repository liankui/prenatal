package handlers

import (
	"github.com/chaos-io/chaos/gokit/middleware"
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	"github.com/go-kit/kit/tracing/opentracing"
	"github.com/mojo-lang/core/go/pkg/mojo/core"
	stdopentracing "github.com/opentracing/opentracing-go"

	"github.com/liankui/prenatal/go/pkg/prenatal"
	"github.com/liankui/prenatal/service-go/pkg/quiz-service/svc"

	// this service api
	pb "github.com/liankui/prenatal/go/pkg/prenatal/v1"
)

var (
	_ = prenatal.Question{}
	_ = core.Null{}
)

// WrapEndpoints accepts the service's entire collection of endpoints, so that a
// set of middlewares can be wrapped around every middleware (e.g., access
// logging and instrumentation), and others wrapped selectively around some
// endpoints and not others (e.g., endpoints requiring authenticated access).
// Note that the final middleware wrapped will be the outermost middleware
// (i.e. applied first)
func WrapEndpoints(in svc.Endpoints, options map[string]interface{}) svc.Endpoints {

	// Pass a middleware you want applied to every endpoint.
	// optionally pass in endpoints by name that you want to be excluded
	// e.g.
	// in.WrapAllExcept(authMiddleware, "Status", "Ping")

	// Pass in a svc.LabeledMiddleware you want applied to every endpoint.
	// These middlewares get passed the endpoints name as their first argument when applied.
	// This can be used to write generic metric gathering middlewares that can
	// report the endpoint name for free.
	// github.com/ncraft-io//_example/middlewares/labeledmiddlewares.go for examples.
	// in.WrapAllLabeledExcept(errorCounter(statsdCounter), "Status", "Ping")

	// How to apply a middleware to a single endpoint.
	// in.ExampleEndpoint = authMiddleware(in.ExampleEndpoint)

	var tracer stdopentracing.Tracer
	if value, ok := options["tracer"]; ok && value != nil {
		tracer = value.(stdopentracing.Tracer)
	}
	var count *kitprometheus.Counter
	if value, ok := options["count"]; ok && value != nil {
		count = value.(*kitprometheus.Counter)
	}
	var latency *kitprometheus.Histogram
	if value, ok := options["latency"]; ok && value != nil {
		latency = value.(*kitprometheus.Histogram)
	}
	// var validator *middleware.Validator
	// if value, ok := options["validator"]; ok && value != nil {
	//	validator = value.(*middleware.Validator)
	// }

	{ // create_question
		if tracer != nil {
			in.CreateQuestionEndpoint = opentracing.TraceServer(tracer, "create_question")(in.CreateQuestionEndpoint)
		}
		if count != nil && latency != nil {
			in.CreateQuestionEndpoint = middleware.Instrumenting(latency.With("method", "create_question"), count.With("method", "create_question"))(in.CreateQuestionEndpoint)
		}
		// if validator != nil {
		//	in.CreateQuestionEndpoint = validator.Validate()(in.CreateQuestionEndpoint)
		// }
	}
	{ // get_question
		if tracer != nil {
			in.GetQuestionEndpoint = opentracing.TraceServer(tracer, "get_question")(in.GetQuestionEndpoint)
		}
		if count != nil && latency != nil {
			in.GetQuestionEndpoint = middleware.Instrumenting(latency.With("method", "get_question"), count.With("method", "get_question"))(in.GetQuestionEndpoint)
		}
		// if validator != nil {
		//	in.GetQuestionEndpoint = validator.Validate()(in.GetQuestionEndpoint)
		// }
	}
	{ // update_question
		if tracer != nil {
			in.UpdateQuestionEndpoint = opentracing.TraceServer(tracer, "update_question")(in.UpdateQuestionEndpoint)
		}
		if count != nil && latency != nil {
			in.UpdateQuestionEndpoint = middleware.Instrumenting(latency.With("method", "update_question"), count.With("method", "update_question"))(in.UpdateQuestionEndpoint)
		}
		// if validator != nil {
		//	in.UpdateQuestionEndpoint = validator.Validate()(in.UpdateQuestionEndpoint)
		// }
	}
	{ // delete_question
		if tracer != nil {
			in.DeleteQuestionEndpoint = opentracing.TraceServer(tracer, "delete_question")(in.DeleteQuestionEndpoint)
		}
		if count != nil && latency != nil {
			in.DeleteQuestionEndpoint = middleware.Instrumenting(latency.With("method", "delete_question"), count.With("method", "delete_question"))(in.DeleteQuestionEndpoint)
		}
		// if validator != nil {
		//	in.DeleteQuestionEndpoint = validator.Validate()(in.DeleteQuestionEndpoint)
		// }
	}

	return in
}

func WrapService(in pb.QuizServer, options map[string]interface{}) pb.QuizServer {
	return in
}
