package mail

import (
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
			conn, err := redigo.Dial(
				"tcp", viper.GetString("redis.host")+":"+viper.GetString("redis.port"),
				redigo.DialPassword(viper.GetString("redis.password")),
				redigo.DialDatabase(viper.GetInt("redis.database")),
			)
			if err != nil {
				panic(err)
			}
			return conn, err
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
