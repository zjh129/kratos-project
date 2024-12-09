package service

type JobInterface interface {
	GetName() string //获取cron名称
	GetSpec() string //获取cron表达式
	GetFunc() func() //获取cron执行函数
}

type CronService struct {
	jobList []JobInterface
}

func NewJobService(userCountSrc *UserCountService) *CronService {
	job := &CronService{
		jobList: make([]JobInterface, 0),
	}
	// 添加示例任务
	job.AddJob(userCountSrc)
	return job
}

// AddJob 添加cron任务
func (s *CronService) AddJob(job JobInterface) {
	s.jobList = append(s.jobList, job)
}

// GetJobList 获取cron任务列表
func (s *CronService) GetJobList() []JobInterface {
	return s.jobList
}
