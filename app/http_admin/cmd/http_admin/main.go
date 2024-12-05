package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"kratos-project/app/http_admin/internal/conf"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/http"

	_ "go.uber.org/automaxprocs"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name of the compiled software.
	Name string = "kratos-project.http_admin"
	// Version is the version of the compiled software.
	Version string = "v0.0.1"
	// flagconf is the config flag.
	flagconf string
	// host name
	hostname, _ = os.Hostname()
	// id
	id = fmt.Sprintf("%s/%s", hostname, Name)
)

func init() {
	flag.StringVar(&flagconf, "conf", "../../configs", "config path, eg: -conf config.yaml")
}

func newApp(logger log.Logger, hs *http.Server, rr registry.Registrar) *kratos.App {
	return kratos.New(
		kratos.ID(id),
		kratos.Name(Name),
		kratos.Version(Version),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Server(
			hs,
		),
		kratos.Registrar(rr),
	)
}

func main() {
	flag.Parse()
	logger := log.With(log.NewStdLogger(os.Stdout),
		"ts", log.DefaultTimestamp,
		"caller", log.DefaultCaller,
		"service.id", id,
		"service.name", Name,
		"service.version", Version,
		"trace.id", tracing.TraceID(),
		"span.id", tracing.SpanID(),
	)

	// init config
	c := initConfig()
	defer c.Close()

	if err := c.Load(); err != nil {
		panic(err)
	}

	var bc conf.Bootstrap
	if err := c.Scan(&bc); err != nil {
		panic(err)
	}

	// 注册链路追踪
	ctx := context.Background()
	err := setTracerProvider(ctx, bc.Trace)
	if err != nil {
		log.Error(err)
	}

	// init registry
	rConf := initRegistryConf()
	rr := initRegistry(rConf)

	app, cleanup, err := wireApp(bc.Server, bc.Auth, bc.Data, rConf, logger, rr)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	// start and wait for stop signal
	if err := app.Run(); err != nil {
		panic(err)
	}
}
