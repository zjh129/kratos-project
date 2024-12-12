package service

import (
	"kratos-project/app/command/internal/biz"

	"github.com/go-kratos/kratos/v2/log"

	"github.com/spf13/cobra"
)

type DemoService struct {
	uc  *biz.UserUseCase
	log *log.Helper
}

func NewDemoService(logger log.Logger, uc *biz.UserUseCase) *DemoService {
	return &DemoService{
		uc:  uc,
		log: log.NewHelper(log.With(logger, "module", "service")),
	}
}

func (s *DemoService) GetCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "demo",
		Short: "demo command",
		Run: func(cmd *cobra.Command, args []string) {
			s.log.Infof("demo command")
		},
	}
}
