package main

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/lambda"
)

type S3Event struct {
	Records []struct {
		s3 struct {
			object struct {
				Key string `json:"key"`
			} `json:"object"`
		} `json:"s3"`
	} `json:"Records"`
}

func HandleRequest(ctx context.Context, s3Event S3Event) error {
	for _, record := range s3Event.Records {
		objectKey := record.s3.object.Key
		log.Printf("Object deleted: %s", objectKey)

		// Add your desired logic here, e.g.,
		// - Update a database to reflect the deletion
		// - Trigger another process or workflow
		// - Send notifications
	}
	return nil
}

func main() {
	lambda.Start(HandleRequest)
}