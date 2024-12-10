package timer

import "github.com/robfig/cron/v3"

type Timer interface {
	// FindCronList 寻找所有Cron
	FindCronList() map[string]*taskManager
	// AddTaskByFuncWithSecond 添加Task 方法形式以秒的形式加入
	AddTaskByFuncWithSecond(cronName string, spec string, fun func(), taskName string, option ...cron.Option) (cron.EntryID, error) // 添加Task Func以秒的形式加入
	// AddTaskByJobWithSeconds 添加Task 接口形式以秒的形式加入
	AddTaskByJobWithSeconds(cronName string, spec string, job interface{ Run() }, taskName string, option ...cron.Option) (cron.EntryID, error)
	// AddTaskByFunc 通过函数的方法添加任务
	AddTaskByFunc(cronName string, spec string, task func(), taskName string, option ...cron.Option) (cron.EntryID, error)
	// AddTaskByJob 通过接口的方法添加任务 要实现一个带有 Run方法的接口触发
	AddTaskByJob(cronName string, spec string, job interface{ Run() }, taskName string, option ...cron.Option) (cron.EntryID, error)
	// FindCron 获取对应taskName的cron 可能会为空
	FindCron(cronName string) (*taskManager, bool)
	// StartCron 指定cron开始执行
	StartCron(cronName string)
	// StopCron 指定cron停止执行
	StopCron(cronName string)
	// FindTask 查找指定cron下的指定task
	FindTask(cronName string, taskName string) (*task, bool)
	// RemoveTask 根据id删除指定cron下的指定task
	RemoveTask(cronName string, id int)
	// RemoveTaskByName 根据taskName删除指定cron下的指定task
	RemoveTaskByName(cronName string, taskName string)
	// Clear 清理掉指定cronName
	Clear(cronName string)
	// Close 停止所有的cron
	Close()
}
