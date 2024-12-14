package main

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/lambda"
)

// Incoming type of s3 upload event
type S3Event struct {
	Records []struct {
		S3 struct {
			Object struct {
				Key   string `json:"key"`
				Size  int64  `json:"size"`
				ETag  string `json:"etag"`
				Owner struct {
					DisplayName string `json:"DisplayName"`
					ID          string `json:"ID"`
				} `json:"owner"`
			} `json:"object"`
		} `json:"s3"`
	} `json:"Records"`
}

// Handling the lambda request
func HandleRequest(ctx context.Context, s3Event S3Event) error {
	for _, record := range s3Event.Records {
		objectKey := record.S3.Object.Key
		objectSize := record.S3.Object.Size
		objectETag := record.S3.Object.ETag
		objectOwner := record.S3.Object.Owner

		log.Printf("File uploaded: %s", objectKey)
		log.Printf("File size: %d bytes", objectSize)
		log.Printf("File ETag: %s", objectETag)
		log.Printf("File owner: %s (%s)", objectOwner.DisplayName, objectOwner.ID)
	}
	return nil
}

func main() {
	lambda.Start(HandleRequest)
}
