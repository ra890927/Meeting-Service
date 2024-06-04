package mail

import (
	"fmt"
	"log"
	"meeting-center/src/repos"
	"sync"
	"time"

	"github.com/jasonlvhit/gocron"
)

var (
	scheduler     *gocron.Scheduler
	schedulerOnce sync.Once
)

func GetSchedulerInstance() *gocron.Scheduler {
	if scheduler == nil {
		schedulerOnce.Do(func() {
			initScheduler()
		})
	}
	return scheduler
}

func initScheduler() {
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		log.Fatal("[ERROR] SetupRoutineForNotification:", err)
		return
	}

	scheduler = gocron.NewScheduler()
	scheduler.ChangeLoc(loc)

	for i := 0; i < 24; i++ {
		timeStr25 := fmt.Sprintf("%02d:25", i)
		timeStr55 := fmt.Sprintf("%02d:55", i)
		scheduler.Every(1).Day().At(timeStr25).Do(getMeetingForNotification)
		scheduler.Every(1).Day().At(timeStr55).Do(getMeetingForNotification)
	}
}

func getMeetingForNotification() {
	log.Print("[INFO] Send notification mail start")

	meetingRepo := repos.NewMeetingRepo()

	dateFrom, dateTo := time.Now(), time.Now().Add(30*time.Minute)
	meetings, err := meetingRepo.GetMeetingsByDatePeriod(dateFrom, dateTo)
	if err != nil {
		log.Fatal("[ERROR] getMeetingForNotification:", err)
		return
	}

	for _, meeting := range meetings {
		SendEmailByMeeting(meeting)
	}

	log.Print("[INFO] Send notification mail end")
}
