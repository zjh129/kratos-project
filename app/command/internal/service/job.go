package service

import "github.com/spf13/cobra"

// 任务接口，所有任务都需要实现该接口
type JobInterface interface {
	GetCommand() *cobra.Command
}

// CommandService 任务服务
type CommandService struct {
	jobList []JobInterface
}

func NewCommandService(demo *DemoService) *CommandService {
	job := &CommandService{
		jobList: make([]JobInterface, 0),
	}
	// 添加示例任务
	job.AddJob(demo)

	return job
}

// AddJob 添加cron任务
func (s *CommandService) AddJob(job JobInterface) {
	s.jobList = append(s.jobList, job)
}

// GetJobList 获取cron任务列表
func (s *CommandService) GetJobList() []JobInterface {
	return s.jobList
}
