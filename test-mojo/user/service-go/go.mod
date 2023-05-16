module github.com/liankui/prenatal-server/service-go

go 1.16

replace github.com/liankui/prenatal-server/go => ../go

require (
	github.com/etherlabsio/healthcheck/v2 v2.0.0
	github.com/go-kit/kit v0.12.0
	github.com/gorilla/mux v1.7.4
	github.com/json-iterator/go v1.1.12
	github.com/liankui/prenatal-server/go v0.0.0
	github.com/mojo-lang/core/go v0.0.0-20230427095043-961105c650c8
	github.com/mojo-lang/http/go v0.0.0-20230427095618-4ac8e30ae2b4
	github.com/ncraft-io/ncraft-gokit v0.0.0-20230427122020-1685fef4143f
	github.com/ncraft-io/ncraft/go v0.0.0-20230427121501-11df4525d1d8
	github.com/opentracing/opentracing-go v1.2.0
	github.com/pkg/errors v0.9.1
	github.com/prometheus/client_golang v1.11.0
	github.com/rs/cors v1.7.0
	go.uber.org/automaxprocs v1.3.0
	google.golang.org/grpc v1.54.0
)
