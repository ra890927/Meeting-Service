package mail

import (
	"log"
	"meeting-center/src/models"
	"meeting-center/src/repos"
	"sync"

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
	subject := job.ArgString("subject")
	from := sgmail.NewEmail("Meeting Center", viper.GetString("mail.sender"))
	to := sgmail.NewEmail(job.ArgString("user_name"), job.ArgString("user_email"))
	plainTextContent := "and easy to do anywhere, even with Go"
	htmlContent := "<strong>and easy to do anywhere, even with Go</strong>"
	message := sgmail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)

	client := GetMailInstance()
	response, err := client.Send(message)
	if err != nil {
		log.Fatal("[ERROR] SendEmail:", err)
	} else {
		log.Print("[INFO] SendEmail")
		log.Print("[INFO]Status Code:", response.StatusCode)
		log.Print("[INFO]Body:", response.Body)
		log.Print("[INFO]Headers:", response.Headers)
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

func SendEmailByMeeting(meeting models.Meeting) {
	userRepo := repos.NewUserRepo()
	for _, uid := range meeting.Participants {
		user, err := userRepo.GetUserByID(uid)
		if err != nil {
			log.Fatal("[ERROR] getMeetingForNotification:", err)
			continue
		}

		_, err = enqueuer.Enqueue("send_email", work.Q{
			"recipient": user.Email,
			"subject":   meeting.Title,
			"body":      "",
		})
		if err != nil {
			log.Fatal("[ERROR] getMeetingForNotification:", err)
		}
	}
}
