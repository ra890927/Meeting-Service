package mail

import (
	"log"
	"sync"

	"github.com/gocraft/work"
	redigo "github.com/gomodule/redigo/redis"
	"github.com/spf13/viper"
)

var (
	workerPool *work.WorkerPool
	enqueuer   *work.Enqueuer
	workerOnce sync.Once
)

type Context struct{}

func (c *Context) SendEmail(job *work.Job) error {
	recipient := job.ArgString("recipient")
	subject := job.ArgString("subject")
	body := job.ArgString("body")
	if err := job.ArgError(); err != nil {
		return err
	}

	// 這裡添加實際的寄信邏輯
	log.Printf("Sending email to %s: %s\n%s\n", recipient, subject, body)
	return nil
}

func GetWorkerPoolInstance() *work.WorkerPool {
	if workerPool == nil {
		initWorker()
	}
	return workerPool
}

func GetEnqueuerInstance() *work.Enqueuer {
	if enqueuer == nil {
		initWorker()
	}
	return enqueuer
}

func getRedigoPool() *redigo.Pool {
	return &redigo.Pool{
		MaxActive: viper.GetInt("mail.redigo.maxActive"),
		MaxIdle:   viper.GetInt("mail.redigo.maxIdle"),
		Wait:      viper.GetBool("mail.redigo.wait"),
		Dial: func() (redigo.Conn, error) {
			return redigo.Dial("tcp", ":"+viper.GetString("mail.redigo.port"))
		},
	}
}

func initWorker() {
	workerOnce.Do(func() {
		redisPool := getRedigoPool()
		namespace := viper.GetString("mail.namespace")
		concurrency := viper.GetUint("mail.concurrency")

		workerPool = work.NewWorkerPool(Context{}, concurrency, namespace, redisPool)
		workerPool.Job("send_email", (*Context).SendEmail)
		enqueuer = work.NewEnqueuer(namespace, redisPool)
	})
}
