package graph

import (
	"fmt"
	// "net/smtp"
	"github.com/yolla-kargo/kargo-trucks/graph/model"
	// "os"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Trucks []*model.Truck
	Mails  []*model.Mail
}

func (r *Resolver) Init() {
	for i := 0; i < 20; i++ {
		truck := &model.Truck{
			ID:      fmt.Sprintf("TRUCK-%d", len(r.Trucks)+1),
			PlateNo: fmt.Sprintf("B %d CD", len(r.Trucks)+1),
		}
		r.Trucks = append(r.Trucks, truck)
	}
}

// func sendEmail(r *Resolver) {
//     // Sender data.
// 	from := "yollaafrdl@gmail.com"
// 	password := os.Getenv("EMAILPASSWORD")

// 	// Receiver email address.
// 	to := []string{
// 	"yollafrdl@gmail.com" ,
// 	}

// 	// smtp server configuration
// 	smtpHost := "smtp.gmail.com"
// 	smtpPort := "587"

// 	// Message
// 	message := []byte("coba")

// 	// Authentication
// 	auth := smtp.PlainAuth ("", from, password,
// 	smtpHost)

// 	// Sending email
// 	err := smtp.SendMail (smtpHost +":"+smtpPort, auth,
// 	from, to, message)
// 	if err != nil {
// 	fmt.Println(err)
// 	return
// 	}
// 	fmt.Println("Email Sent Successfully!" )
// }
func main(r *Resolver) {
	fmt.Println(r.Trucks)
}
