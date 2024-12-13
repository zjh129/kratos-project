package service

import (
	"context"
	"fmt"
	"github.com/schollz/progressbar/v3"
	"kratos-project/app/command/internal/biz"
	"os"
	"time"

	"github.com/go-kratos/kratos/v2/log"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

type DemoService struct {
	uc      *biz.UserUseCase
	log     *log.Helper
	rootCmd *cobra.Command
}

// NewDemoService init demo service
func NewDemoService(logger log.Logger, rootCmd *cobra.Command, uc *biz.UserUseCase) *DemoService {
	return &DemoService{
		uc:      uc,
		log:     log.NewHelper(log.With(logger, "module", "service")),
		rootCmd: rootCmd,
	}
}

// 设置命令
func (s *DemoService) GetCommand() *cobra.Command {
	var printUser bool
	cmd := &cobra.Command{
		Use:   "demo",
		Short: "demo command",
		Run: func(cmd *cobra.Command, args []string) {
			s.log.Infof("demo command")
			if printUser {
				users, err := s.uc.List(context.Background(), 1, 10)
				if err != nil {
					cmd.PrintErrf("list user error: %v\n", err)
				}
				table := tablewriter.NewWriter(os.Stdout)
				table.SetHeader([]string{"ID", "ACCOUNT", "PASSWORD"})
				for _, user := range users {
					table.Append([]string{fmt.Sprintf("%d", user.Id), user.Account, user.Passwd})
				}
				table.Render() // Send output
			} else {
				bar := progressbar.NewOptions(100,
					progressbar.OptionEnableColorCodes(true),
					progressbar.OptionShowBytes(true),
					progressbar.OptionSetWidth(15),
					progressbar.OptionSetDescription("[cyan]Downloading..."),
					progressbar.OptionSetTheme(progressbar.Theme{
						Saucer:        "[green]=[reset]",
						SaucerHead:    "[green]>[reset]",
						SaucerPadding: " ",
						BarStart:      "[",
						BarEnd:        "]",
					}))
				//bar := progressbar.Default(100)
				for i := 0; i < 100; i++ {
					bar.Add(1)
					time.Sleep(40 * time.Millisecond)
				}
			}
		},
	}
	// 添加命令行参数
	cmd.Flags().BoolVarP(&printUser, "output-user", "o", false, "output user")
	return cmd
}
