package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/textract"
)

func HandleRequest(ctx context.Context, event events.S3Event) error {
	var bucket = event.Records[0].S3.Bucket.Name
	var key = event.Records[0].S3.Object.URLDecodedKey

	session := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	service := textract.New(session)

	input := &textract.DetectDocumentTextInput{
		Document: &textract.Document{
			S3Object: &textract.S3Object{
				Bucket: aws.String(bucket),
				Name:   aws.String(key),
			},
		},
	}

	response, err := service.DetectDocumentText(input)

	fmt.Println(response)

	return err
}

func main() {
	// lambda.Start(HandleRequest)
	HandleRequest(context.TODO(), events.S3Event{})
}
