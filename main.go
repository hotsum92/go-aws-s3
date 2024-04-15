package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func main() {

	accessKey := "minioadmin"
	secretKey := "minioadmin"

	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		panic(err)
	}

	s3Client := s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.Credentials = credentials.NewStaticCredentialsProvider(accessKey, secretKey, "")
		o.Region = "us-east-1"
		o.BaseEndpoint = aws.String("http://minio:9000")
	})

	count := 10

	fmt.Printf("Let's list up to %v buckets for your account.\n", count)

	result, err := s3Client.ListBuckets(context.TODO(), &s3.ListBucketsInput{})

	if err != nil {
		panic(err)
	}

	if len(result.Buckets) == 0 {
		fmt.Println("You don't have any buckets!")
	} else {
		if count > len(result.Buckets) {
			count = len(result.Buckets)
		}
		for _, bucket := range result.Buckets[:count] {
			fmt.Printf("\t%v\n", *bucket.Name)
		}
	}
}
