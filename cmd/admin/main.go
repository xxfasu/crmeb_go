package main

import (
	"crmeb_go/cmd/admin/wire"
	"crmeb_go/internal/conf"
	"crmeb_go/internal/redis"
	"crmeb_go/pkg/logs"
	"github.com/gin-gonic/gin"
	"github.com/kardianos/service"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

// 定义程序结构体
type program struct {
	server    *gin.Engine
	svc       service.Service
	once      sync.Once
	clearFunc func()
}

// Start 方法在服务启动时调用
func (p *program) Start(s service.Service) error {
	// 在一个新的 goroutine 中启动服务
	go p.run()
	return nil
}

// Stop 方法在服务停止时调用
func (p *program) Stop(s service.Service) error {
	// 这里可以添加清理资源的代码
	p.once.Do(func() {
		if p.clearFunc != nil {
			p.clearFunc()
		}
	})
	return nil
}

// 运行 Gin 服务器
func (p *program) run() {
	err := conf.InitConfig()
	if err != nil {
		panic(err)
	}
	client, err := redis.InitRedis()
	if err != nil {
		panic(err)
	}
	rLock := redis.InitRedSync(client)
	logs.InitLog()
	wire, fn, err := wire.NewWire(client, rLock)
	p.clearFunc = fn
	if err != nil {
		panic(err)
	}
	p.server = wire

	if err = p.server.Run(conf.Config.System.Port); err != nil {
		panic(err)
	}
}

func main() {
	// 定义服务配置
	svcConfig := &service.Config{
		Name:        "Crmeb Admin",
		DisplayName: "Crmeb Admin",
		Description: "Crmeb Admin",
	}

	prg := &program{}
	s, err := service.New(prg, svcConfig)
	if err != nil {
		log.Fatal(err)
	}
	prg.svc = s

	// 设置日志
	errs := make(chan error)
	logger, err := s.Logger(errs)
	if err != nil {
		log.Fatal(err)
	}

	// 启动服务
	if err := s.Run(); err != nil {
		logger.Error(err)
	}

	// 处理错误
	go func() {
		for {
			select {
			case err := <-errs:
				log.Println("Error:", err)
			}
		}
	}()

	// 等待中断信号以优雅关闭
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs
	log.Println("Shutting down...")
}
