// https://github.com/awsdocs/aws-doc-sdk-examples
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	v4 "github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go/aws"
)

type Presigner struct {
	PresignClient *s3.PresignClient
}

func (presigner Presigner) PutObject(
	ctx context.Context,
	bucketName string,
	objectKey string,
	lifetimeDuration time.Duration) (*v4.PresignedHTTPRequest, error) {
	request, err := presigner.PresignClient.PresignPutObject(ctx, &s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objectKey),
	}, s3.WithPresignExpires(lifetimeDuration),
	)
	if err != nil {
		log.Printf("Couldn't get a presigned request to put %v:%v. Here's why: %v\n",
			bucketName, objectKey, err)
	}
	return request, err
}

func main() {
	bucketName := flag.String("b", "", "The bucket")
	objectKey := flag.String("k", "", "The object key")

	flag.Parse()

	if *bucketName == "" || *objectKey == "" {
		fmt.Println("You must supply a bucket name (-b BUCKET) and object key (-k KEY)")
		return
	}

	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		panic("configuration error, " + err.Error())
	}

	client := s3.NewFromConfig(cfg)
	presignedClient := s3.NewPresignClient(client)
	presigner := Presigner{PresignClient: presignedClient}
	req, err := presigner.PutObject(context.TODO(), *bucketName, *objectKey, time.Duration(2*time.Hour))

	if err != nil {
		panic("configuration error, " + err.Error())
	}

	fmt.Println("The URL: ")
	fmt.Println(req.URL)
}
