package main

import (
	"context"
	"fmt"
	"os"

	"kratos-project/app/command/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/tracing"

	"github.com/spf13/cobra"
	_ "go.uber.org/automaxprocs"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name of the compiled software.
	Name string = "command"
	// Version is the version of the compiled software.
	Version string = "v0.0.1"
	// flagconf is the config flag.
	flagconf string

	id, _ = os.Hostname()

	rootCmd = &cobra.Command{
		Use:   Name,
		Short: "command service",
	}
)

func init() {
	rootCmd.PersistentFlags().StringVar(&flagconf, "conf", flagconf, "config path, eg: --conf=app/cron/configs")
}

func main() {
	// 由于结合了cobra，所以需要执行rootCmd.Execute()，提前解析出整个程序执行的配置文件目录
	_ = rootCmd.ParseFlags(os.Args[1:])
	// 由于cobra的rootCmd是一个空的命令，所以需要在rootCmd.Run中执行rootCmd.Help()，输出帮助信息
	rootCmd.Run = func(cmd *cobra.Command, args []string) {
		cmd.Help()
	}
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
	ctx := context.Background()
	// init registry
	rConf := initRegistryConf()

	app, cleanup, err := wireApp(bc.Server, bc.Data, rConf, rootCmd, logger)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	// start and wait for stop signal
	if err := app.Start(ctx); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
