package main

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"log"
)

func main() {
	lambda.Start(handler)
}

// Handler s3 handler.
func handler(ctx context.Context, event events.S3Event) {
	// get bucket name and file name from event
	record := event.Records[0]
	sourceBucket := record.S3.Bucket.Name
	sourceKey := record.S3.Object.Key

	// create destination key
	destinationKey := "copied/" + sourceKey

	// create s3 client
	sess := session.Must(session.NewSession())
	s3Client := s3.New(sess)

	// create copy object input
	copyInput := &s3.CopyObjectInput{
		Bucket:     aws.String(sourceBucket),
		CopySource: aws.String(sourceBucket + "/" + sourceKey),
		Key:        aws.String(destinationKey),
	}

	// copy file
	_, err := s3Client.CopyObject(copyInput)
	if err != nil {
		log.Printf("Error copying file: %v", err)
		return
	}

	log.Printf("File copied from %s/%s to %s/%s", sourceBucket, sourceKey, "copied/"+sourceKey, destinationKey)
}
