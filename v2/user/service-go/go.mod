module service-go

go 1.16

replace (
    go => ../go
)

require (
    go v0.0.0
    github.com/etherlabsio/healthcheck/v2 v2.0.0
    github.com/mojo-lang/core/go v0.0.0-20211228010257-772a79853c5e
    github.com/mojo-lang/geom/go v0.0.0-20211217060203-19a1f652e101
    github.com/ncraft-io/ncraft-gokit v0.0.0-20220322120959-b7d2795d6943
	github.com/HdrHistogram/hdrhistogram-go v1.0.1 // indirect
	github.com/gin-gonic/gin v1.6.3 // indirect
	github.com/go-kit/kit v0.10.0
	github.com/google/uuid v1.1.1
	github.com/gorilla/mux v1.7.4
	github.com/json-iterator/go v1.1.9
	github.com/konsorten/go-windows-terminal-sequences v1.0.2 // indirect
	github.com/nacos-group/nacos-sdk-go v1.0.8 // indirect
	github.com/opentracing/opentracing-go v1.1.0
	github.com/pkg/errors v0.9.1
	github.com/prometheus/client_golang v1.6.0
	github.com/rs/cors v1.7.0
	github.com/uber/jaeger-client-go v2.25.0+incompatible // indirect
	github.com/uber/jaeger-lib v2.4.0+incompatible // indirect
	go.uber.org/automaxprocs v1.3.0
	google.golang.org/grpc v1.42.0
)
