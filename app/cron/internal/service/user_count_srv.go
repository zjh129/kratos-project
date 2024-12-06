package service

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"kratos-project/app/cron/internal/biz"
	"time"
)

type DemoService struct {
	name string
	spec string
	log  *log.Helper
	uc   *biz.UserUseCase
}

func NewUserCountService(logger log.Logger, uc *biz.UserUseCase) *DemoService {
	return &DemoService{
		name: "用户总数统计计划任务",
		spec: "*/1 * * * *",
		log:  log.NewHelper(log.With(logger, "module", "service")),
		uc:   uc,
	}
}

func (d *DemoService) GetName() string {
	return d.name
}

func (d *DemoService) GetSpec() string {
	return d.spec
}

func (d *DemoService) GetFunc() func() {
	return func() {
		fmt.Printf("当前时间 %v \n", time.Now().Format(time.DateTime))
		total, err := d.uc.Count(context.Background())
		if err != nil {
			d.log.Errorf("获取用户总数失败: %v", err)
		} else {
			fmt.Printf("当前用户总数 %d \n", total)
		}
	}
}
