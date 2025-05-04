// https://github.com/awsdocs/aws-doc-sdk-examples
package presigner

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	v4 "github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type Presigner struct {
	PresignClient PresignClient
}

type PresignClient interface {
	PresignPutObject(ctx context.Context, params *s3.PutObjectInput, optFuncs ...func(*s3.PresignOptions)) (*v4.PresignedHTTPRequest, error)
	PresignGetObject(ctx context.Context, params *s3.GetObjectInput, optFuncs ...func(*s3.PresignOptions)) (*v4.PresignedHTTPRequest, error)
	PresignDeleteObject(ctx context.Context, params *s3.DeleteObjectInput, optFuncs ...func(*s3.PresignOptions)) (*v4.PresignedHTTPRequest, error)
}

func NewPresigner(presignClient PresignClient) *Presigner {
	return &Presigner{PresignClient: presignClient}
}

func (presigner Presigner) PutObject(
	ctx context.Context,
	bucketName string,
	objectKey string,
	lifetimeDuration time.Duration) error {
	request, err := presigner.PresignClient.PresignPutObject(ctx, &s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objectKey),
	}, s3.WithPresignExpires(lifetimeDuration),
	)
	if err != nil {
		log.Printf("Couldn't get a presigned request to put %v:%v. Here's why: %v\n",
			bucketName, objectKey, err)
		return err
	}

	fmt.Println("The URL: ")
	fmt.Println(request.URL)

	return nil
}

func (presigner Presigner) GetObject(
	ctx context.Context,
	bucketName string,
	objectKey string,
	lifetimeDuration time.Duration) error {
	request, err := presigner.PresignClient.PresignGetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objectKey),
	}, s3.WithPresignExpires(lifetimeDuration),
	)
	if err != nil {
		log.Printf("Couldn't get a presigned request to get %v:%v. Here's why: %v\n",
			bucketName, objectKey, err)
		return err
	}

	fmt.Println("The URL: ")
	fmt.Println(request.URL)

	return nil
}

func (presigner Presigner) DeleteObject(
	ctx context.Context,
	bucketName string,
	objectKey string,
	lifetimeDuration time.Duration) error {
	request, err := presigner.PresignClient.PresignDeleteObject(ctx, &s3.DeleteObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objectKey),
	}, s3.WithPresignExpires(lifetimeDuration),
	)
	if err != nil {
		log.Printf("Couldn't get a presigned request to delete %v:%v. Here's why: %v\n",
			bucketName, objectKey, err)
		return err
	}

	fmt.Println("The URL: ")
	fmt.Println(request.URL)

	return nil
}
