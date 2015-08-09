package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/rjocoleman/lambda-ses/Godeps/_workspace/src/github.com/aws/aws-sdk-go/aws"
	"github.com/rjocoleman/lambda-ses/Godeps/_workspace/src/github.com/aws/aws-sdk-go/service/ses"
)

// Request - inwards JSON
type Request struct {
	FromAddress   string
	ToAddress     string
	SenderAddress string

	Subject string
	Message string
}

func die(err error) {
	fmt.Fprintf(os.Stderr, "error: %s\n", err)
	os.Exit(1)
}

func main() {
	if len(os.Args) < 2 {
		die(fmt.Errorf("must specify event as argument"))
	}

	data := []byte(os.Args[1])

	var req Request

	err := json.Unmarshal(data, &req)

	if err != nil {
		die(err)
	}

	fmt.Printf("req = %+v\n", req)

	SES := ses.New(nil)

	_, err = SES.SendEmail(&ses.SendEmailInput{
		Destination: &ses.Destination{
			ToAddresses: []*string{
				aws.String(req.ToAddress),
			},
		},
		Message: &ses.Message{
			Body: &ses.Body{
				Text: &ses.Content{
					Data: aws.String(req.Message),
				},
			},
			Subject: &ses.Content{
				Data: aws.String(req.Subject),
			},
		},
		Source: aws.String(req.SenderAddress),
		ReplyToAddresses: []*string{
			aws.String(req.FromAddress),
		},
	})

	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		return
	}
}
