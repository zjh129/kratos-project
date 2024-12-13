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
	Name string
	// Version is the version of the compiled software.
	Version string
	// flagconf is the config flag.
	flagconf string

	id, _ = os.Hostname()

	rootCmd = &cobra.Command{
		Use:   Name,
		Short: "",
		PreRun: func(cmd *cobra.Command, args []string) {
			// Parse flags here
			if err := cmd.Flags().Parse(args); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		},
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}
)

func init() {
	rootCmd.PersistentFlags().StringVar(&flagconf, "conf", flagconf, "config path, eg: --conf=app/cron/configs")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
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
