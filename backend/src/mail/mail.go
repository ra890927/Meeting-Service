package mail

import (
	"fmt"
	"log"
	"meeting-center/src/models"
	"meeting-center/src/repos"
	"os"
	"sync"
	"time"

	"github.com/gocraft/work"
	"github.com/sendgrid/sendgrid-go"
	sgmail "github.com/sendgrid/sendgrid-go/helpers/mail"
	"github.com/spf13/viper"
)

var (
	mailInstance *sendgrid.Client
	mailOnce     sync.Once
)

type Context struct{}

func (c *Context) SendEmail(job *work.Job) error {
	log.Print("[INFO] SendEmail to", job.ArgString("user_name"))

	subject := job.ArgString("subject")
	from := sgmail.NewEmail("Meeting Center", viper.GetString("mail.sender"))
	to := sgmail.NewEmail(job.ArgString("user_name"), job.ArgString("user_email"))

	template, err := os.ReadFile("./docs/template.txt")
	if err != nil {
		log.Fatal("[ERROR] SendEmail:", err)
	}

	plainText := fmt.Sprintf(
		string(template),
		job.ArgString("user_name"),
		job.ArgString("first_line"),
		job.ArgString("date"),
		job.ArgString("time"),
	)

	webTemplate, err := os.ReadFile("./docs/web_template.txt")
	if err != nil {
		log.Fatal("[ERROR] SendEmail:", err)
	}

	webText := fmt.Sprintf(
		string(webTemplate),
		job.ArgString("user_name"),
		job.ArgString("first_line"),
		job.ArgString("date"),
		job.ArgString("time"),
	)

	message := sgmail.NewSingleEmail(from, subject, to, plainText, webText)

	client := GetMailInstance()
	response, err := client.Send(message)
	if err != nil {
		log.Fatal("[ERROR] SendEmail:", err)
	} else {
		log.Print("[INFO] SendEmail")
		log.Print("[INFO] Status Code:", response.StatusCode)
	}

	return err
}

func GetMailInstance() *sendgrid.Client {
	if mailInstance == nil {
		mailOnce.Do(func() {
			apiKey := viper.GetString("mail.apiKey")
			mailInstance = sendgrid.NewSendClient(apiKey)
		})
	}
	return mailInstance
}

func SendEmailByMeeting(meeting models.Meeting, mode uint) {
	userRepo := repos.NewUserRepo()
	for _, uid := range meeting.Participants {
		user, err := userRepo.GetUserByID(uid)
		if err != nil {
			log.Fatal("[ERROR] getMeetingForNotification:", err)
			continue
		}

		var firstLine string
		if mode == NOTICE {
			firstLine = fmt.Sprintf("您被獲邀參與 %s 會議", user.Username)
		} else {
			firstLine = fmt.Sprintf("提醒您待會 %s 有會議要參加", user.Username)
		}

		utcPlus8 := time.FixedZone("UTC+8", 8*60*60)
		utcPlus8Time := meeting.StartTime.In(utcPlus8)

		_, err = enqueuer.Enqueue("send_email", work.Q{
			"user_email": user.Email,
			"user_name":  user.Username,
			"subject":    meeting.Title,
			"date":       utcPlus8Time.Format("2006-01-02"),
			"time":       utcPlus8Time.Format("15:04"),
			"first_line": firstLine,
		})
		if err != nil {
			log.Fatal("[ERROR] getMeetingForNotification:", err)
		}
	}
}
