package server

import (
	"context"

	"kratos-project/app/command/internal/service"

	"github.com/go-kratos/kratos/v2/log"

	"github.com/spf13/cobra"
)

type CommandServer struct {
	log     *log.Helper
	rootCmd *cobra.Command
}

func NewCommandServer(jobSrc *service.CommandService, _ log.Logger) *CommandServer {
	ser := &CommandServer{
		rootCmd: &cobra.Command{
			Use:   "kratos-project",
			Short: "",
			Run: func(cmd *cobra.Command, args []string) {
				cmd.Help()
			},
		},
	}
	if jobSrc != nil {
		for _, job := range jobSrc.GetJobList() {
			ser.rootCmd.AddCommand(job.GetCommand())
		}
	}
	return ser
}

// Start start the server
func (s *CommandServer) Start(c context.Context) error {
	return s.rootCmd.Execute()
}

// Stop stop the server
func (s *CommandServer) Stop(c context.Context) error {
	s.rootCmd.ResetCommands()
	return nil
}
