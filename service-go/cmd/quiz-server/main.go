// Code generated by chaosmojo. DO NOT EDIT.
// Rerunning chaosmojo will overwrite this file.
// Version: 0.1.0
// Version Date:

package main

import (
	"flag"

	"github.com/chaos-io/chaos/config"
	nserver "github.com/chaos-io/chaos/gokit/server"
	"github.com/chaos-io/chaos/logs"

	// This Service
	"github.com/liankui/prenatal/service-go/internal/quiz-server"
)

var gitHash, gitBranch, buildTime string

func main() {
	flag.String("http_addr", ":20171", "default ip address for http server")
	flag.String("grpc_addr", ":20172", "default ip address for grpc server")
	flag.String("debug_addr", ":20173", "default ip address for debug server")

	// Update addresses if they have been overwritten by flags
	flag.Parse()
	if err := config.LoadFlag(); err != nil {
		logs.Warnw("failed to load flag", "error", err)
	}

	var conf nserver.Config
	err := config.ScanFrom(&conf, "quizServer", "server")
	if err != nil {
		logs.Warnw("failed to get the server config from config file", "error", err.Error())
	}

	// read the httpAddr value from flag, and will override the value from config file
	_ = config.ScanFrom(&conf.HttpAddr, "httpAddr")

	// read the grpcAddr value from flag, and will override the value from config file
	_ = config.ScanFrom(&conf.GrpcAddr, "grpcAddr")

	// read the debugAddr value from flag, and will override the value from config file
	_ = config.ScanFrom(&conf.DebugAddr, "debugAddr")

	if len(gitHash) == 0 {
		logs.Infow("starting QuizServer, you can use `-ldflags \"-X main.gitHash=xxx\" to generate gitHash")
	} else {
		logs.Infow("starting QuizServer", "gitHash", gitHash, "gitBranch", gitBranch, "buildTime", buildTime)
	}

	server.Run(conf)
}
