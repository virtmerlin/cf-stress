package main

import (
	"fmt"
	"log"

	"github.com/minio/minio-go"
)

func fn_s3_put(accessKeyID string, secretAccessKey string, bucketName string, objectName string, filePath string, contentType string) {
	endpoint := "s3.amazonaws.com"
	useSSL := true

	minioClient, err := minio.New(endpoint, accessKeyID, secretAccessKey, useSSL)
	if err != nil {
		log.Fatalln(err)
	}

	p, err := minioClient.FPutObject(bucketName, objectName, filePath, contentType)
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("Successfully uploaded %s of size %d\n", objectName, p)
}

func fn_s3_get(accessKeyID string, secretAccessKey string, bucketName string, objectName string, filePath string) {
	endpoint := "s3.amazonaws.com"
	useSSL := true

	minioClient, err := minio.New(endpoint, accessKeyID, secretAccessKey, useSSL)
	if err != nil {
		log.Fatalln(err)
	}

	err = minioClient.FGetObject(bucketName, objectName, filePath)
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("Successfully read %v", objectName)
}

func fn_s3_rm(accessKeyID string, secretAccessKey string, bucketName string, objectName string) {
	endpoint := "s3.amazonaws.com"
	useSSL := true

	minioClient, err := minio.New(endpoint, accessKeyID, secretAccessKey, useSSL)
	if err != nil {
		log.Fatalln(err)
	}

	err = minioClient.RemoveObject(bucketName, objectName)
	if err != nil {
		fmt.Println(err)
		return
	}
}
