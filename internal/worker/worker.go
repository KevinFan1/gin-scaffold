package worker

import (
	"code/gin-scaffold/internal/global"
	"code/gin-scaffold/internal/settings"
	"log"
)

// Payload 发送到队列的内容
type Payload interface{}

// Job 单个job
type Job struct {
	Payload Payload
}

// JobQueue 一个可以发送工作请求的缓冲channel
var JobQueue chan Job

func init() {
	// 初始化job queue
	JobQueue = make(chan Job, settings.Setting.SystemBaseConfig.MaxQueue)
}

// Worker 工作worker
type Worker struct {
	WorkerPool chan chan Job //
	JobChannel chan Job      // 发送Job struct的channel
	Quit       chan bool     //结束标识
}

// NewWorker 生成一个新的worker
func NewWorker(workerPool chan chan Job) Worker {
	return Worker{
		WorkerPool: workerPool,
		JobChannel: make(chan Job),
		Quit:       make(chan bool),
	}
}

// Start worker处理任务
func (w Worker) Start() {
	go func() {
		for {
			w.WorkerPool <- w.JobChannel
			select {
			case job := <-w.JobChannel:
				//处理耗时任务task
				global.Logger.Infof("接收到任务, %v", job)
				//err := gredis.SSet("123", "!23", 10*time.Minute)
				//
				//if err != nil {
				//	logging.SugarLogger.Errorf("保存redis失败,job:%v, err: %v", job, err)
				//} else {
				//	logging.SugarLogger.Info("保存redis成功", job)
				//}
				//
				//rabbit.Emitter.Push("http://127.0.0.1:9000/test", "test1")
				//models.CreateUser(models.User{Name: "测试", Age: 18})
				//ret, err := celery.CeleryClient.Delay("worker.add", "http://127.0.0.1:9000/test")
				//if err != nil {
				//	logging.SugarLogger.Error("Celery任务失败. Err:", err)
				//} else {
				//	logging.SugarLogger.Info("Celery任务成功. 响应:", ret)
				//}
			case <-w.Quit:
				return
			}

		}
	}()
}

// Stop 发送停止的信号
func (w Worker) Stop() {
	go func() {
		w.Quit <- true
	}()
}

type Dispatcher struct {
	WorkerPool chan chan Job
}

func NewDispatcher(maxWorker int) Dispatcher {
	pool := make(chan chan Job, maxWorker)
	return Dispatcher{WorkerPool: pool}
}

func (d *Dispatcher) dispatch() {
	for {
		select {
		case job := <-JobQueue:
			//分发任务
			go func(job Job) {
				jobChannel := <-d.WorkerPool
				jobChannel <- job
			}(job)
		}
	}
}

func (d Dispatcher) Run() {
	for i := 0; i < settings.Setting.SystemBaseConfig.MaxWorker; i++ {
		// 生成max
		worker := NewWorker(d.WorkerPool)
		worker.Start()
	}
	//处理分发
	go d.dispatch()
}

func StartDispatcher() {
	log.Printf("启动worker成功. Worker数量: %v\n", settings.Setting.SystemBaseConfig.MaxWorker)
	d := NewDispatcher(settings.Setting.SystemBaseConfig.MaxWorker)
	d.Run()
}
