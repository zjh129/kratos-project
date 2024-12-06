package service

import "fmt"

type DemoService struct {
	name string
	spec string
}

func NewDemoService() *DemoService {
	return &DemoService{
		name: "示例计划任务",
		spec: "*/5 * * * * *",
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
		fmt.Println("执行示例任务")
	}
}
