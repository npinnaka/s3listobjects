package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func main() {
	// Configure the AWS session using the default credentials chain
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		panic(err)
	}

	// Create an S3 client
	client := s3.NewFromConfig(cfg)

	// List objects in a specific bucket
	bucketName := "your-bucket-name" // Replace with your bucket name

	input := &s3.ListObjectsV2Input{
		Bucket: &bucketName,
	}

	result, err := client.ListObjectsV2(context.TODO(), input)
	if err != nil {
		panic(err)
	}

	for _, object := range result.Contents {
		fmt.Println(*object.Key)
	}
}
