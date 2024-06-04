package router

import (
	"context"
	"log"
	"meeting-center/src/clients"
	"meeting-center/src/mail"
	"net/http"
	"time"

	"github.com/gocraft/work"
	"github.com/jasonlvhit/gocron"
	"github.com/spf13/viper"
)

var (
	HttpSrvHandler *http.Server
	Scheduler      *gocron.Scheduler
	WorkerPool     *work.WorkerPool
)

func HttpServerRun() {
	r := InitRouter()
	Scheduler = mail.GetSchedulerInstance()
	WorkerPool = mail.GetWorkerPoolInstance()

	addr := ":" + viper.GetString("app.port")
	HttpSrvHandler = &http.Server{
		Addr:    addr,
		Handler: r,
	}
	go func() {
		log.Print("[INFO] Woker pool start")
		WorkerPool.Start()

		log.Print("[INFO] Scheduler start")
		Scheduler.Start()

		log.Print("[INFO] HttpServerRun:", addr)
		if err := HttpSrvHandler.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("[ERROR] HttpServerRun: %s err: %v", addr, err)
		}
	}()
}

func HttpServerStop() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	WorkerPool.Stop()

	if err := HttpSrvHandler.Shutdown(ctx); err != nil {
		log.Fatal("[ERROR] HttpServerStop err:", err)
	} else {
		log.Print("[INFO] HttpServerStop stopped")
	}

	db, err := clients.GetDBInstance().DB()
	if err != nil {
		log.Fatal("[ERROR] Get DB instance error:", err)
	} else {
		if err := db.Close(); err != nil {
			log.Fatal("[ERROR] DB close error:", err)
		} else {
			log.Print("[INFO] DB closed")
		}
	}

	redis := clients.GetRedisInstance()
	if err := redis.Close(); err != nil {
		log.Fatal("[ERROR] Redis close error:", err)
	} else {
		log.Print("[INFO] Redis closed")
	}

	gcs := clients.GetGCSInstance()
	if err := gcs.Close(); err != nil {
		log.Fatal("[ERROR] Gcs close error:", err)
	} else {
		log.Print("[INFO] Gcs closed")
	}
}
