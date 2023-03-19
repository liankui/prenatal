module git.company.com/examples/hello-world/service-go

go 1.16

replace git.company.com/examples/hello-world/go => ../go

require (
	git.company.com/examples/hello-world/go v0.0.0
	github.com/etherlabsio/healthcheck/v2 v2.0.0
	github.com/go-kit/kit v0.12.0
	github.com/gorilla/mux v1.7.4
	github.com/json-iterator/go v1.1.12
	github.com/mojo-lang/core/go v0.0.0-20230316063425-d2ec3be7c3d2
	github.com/mojo-lang/http/go v0.0.0-20230317072936-ebf76a192974
	github.com/ncraft-io/ncraft-gokit v0.0.0-20230210073401-99185e709dd6
	github.com/ncraft-io/ncraft/go v0.0.0-20221226040705-cd865648f5a5
	github.com/opentracing/opentracing-go v1.2.0
	github.com/pkg/errors v0.9.1
	github.com/prometheus/client_golang v1.11.0
	github.com/rs/cors v1.7.0
	go.uber.org/automaxprocs v1.3.0
	google.golang.org/grpc v1.45.0
)
