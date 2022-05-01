package main

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/textract"
)

func ExtractData(sess *session.Session, bucket string, key string) (*textract.DetectDocumentTextOutput, error) {
	log.Print("ExtractData")
	svc := textract.New(sess)

	input := &textract.DetectDocumentTextInput{
		Document: &textract.Document{
			S3Object: &textract.S3Object{
				Bucket: aws.String(bucket),
				Name:   aws.String(key),
			},
		},
	}

	response, err := svc.DetectDocumentText(input)

	responseJson, _ := json.MarshalIndent(response, "", "  ")
	log.Printf("RESPONSE: %s", responseJson)

	errorJson, _ := json.MarshalIndent(err, "", "  ")
	log.Printf("ERROR: %s", errorJson)

	return response, err
}

func HandleRequest(ctx context.Context, event events.S3Event) error {
	log.Print("HandleRequest")

	eventJson, _ := json.MarshalIndent(event, "", "  ")
	log.Printf("EVENT: %s", eventJson)

	var bucket = event.Records[0].S3.Bucket.Name
	var key = event.Records[0].S3.Object.URLDecodedKey

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	_, err := ExtractData(sess, bucket, key)

	return err
}

func main() {
	// lambda.Start(HandleRequest)
	HandleRequest(context.TODO(), events.S3Event{})
}
