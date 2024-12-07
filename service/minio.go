package service

import (
	"context"
	"fmt"
	"mime/multipart"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type FileService struct {
	client     *minio.Client
	bucketName string
}

// NewFileService initializes the MinIO client
func NewFileService(endpoint, accessKey, secretKey, bucketName string, useSSL bool) (*FileService, error) {
	client, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to initialize MinIO client: %w", err)
	}

	// Ensure the bucket exists
	ctx := context.Background()
	exists, err := client.BucketExists(ctx, bucketName)
	if err != nil {
		return nil, fmt.Errorf("error checking bucket existence: %w", err)
	}

	if !exists {
		err = client.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{})
		if err != nil {
			return nil, fmt.Errorf("failed to create bucket: %w", err)
		}
	}

	return &FileService{
		client:     client,
		bucketName: bucketName,
	}, nil
}

// GeneratePresignedPutURL generates a signed URL for uploading files (single or multiple)
func (fs *FileService) GeneratePresignedPutURL(ctx context.Context, objectNames ...string) (map[string]string, error) {
	urls := make(map[string]string)
	for _, objectName := range objectNames {
		url, err := fs.client.PresignedPutObject(ctx, fs.bucketName, objectName, time.Minute*10)
		if err != nil {
			return nil, fmt.Errorf("failed to generate presigned PUT URL for %s: %w", objectName, err)
		}
		urls[objectName] = url.String()
	}
	return urls, nil
}

// GeneratePresignedGetURL generates a signed URL for downloading files (single or multiple)
func (fs *FileService) GeneratePresignedGetURL(ctx context.Context, objectNames ...string) (map[string]string, error) {
	urls := make(map[string]string)
	for _, objectName := range objectNames {
		url, err := fs.client.PresignedGetObject(ctx, fs.bucketName, objectName, time.Minute*10, nil)
		if err != nil {
			return nil, fmt.Errorf("failed to generate presigned GET URL for %s: %w", objectName, err)
		}
		urls[objectName] = url.String()
	}
	return urls, nil
}

// UploadFiles uploads multiple files to MinIO and returns their URLs
func (fs *FileService) UploadFiles(ctx context.Context, files []*multipart.FileHeader) (map[string]string, error) {
	uploadedURLs := make(map[string]string)

	for _, fileHeader := range files {
		// Open the file
		file, err := fileHeader.Open()
		if err != nil {
			return nil, fmt.Errorf("failed to open file %s: %w", fileHeader.Filename, err)
		}
		defer file.Close()

		// Upload to MinIO
		_, err = fs.client.PutObject(ctx, fs.bucketName, fileHeader.Filename, file, fileHeader.Size, minio.PutObjectOptions{
			ContentType: fileHeader.Header.Get("Content-Type"),
		})
		if err != nil {
			return nil, fmt.Errorf("failed to upload file %s to MinIO: %w", fileHeader.Filename, err)
		}

		// Generate the file URL
		fileURL := fmt.Sprintf("https://%s/%s/%s", fs.client.EndpointURL().Host, fs.bucketName, fileHeader.Filename)
		uploadedURLs[fileHeader.Filename] = fileURL
	}

	return uploadedURLs, nil
}

// DeleteObject deletes a single object from MinIO
func (fs *FileService) DeleteObject(ctx context.Context, objectName string) error {
	err := fs.client.RemoveObject(ctx, fs.bucketName, objectName, minio.RemoveObjectOptions{})
	if err != nil {
		return fmt.Errorf("failed to delete object %s: %w", objectName, err)
	}
	return nil
}

// DeleteObjects deletes multiple objects from MinIO
func (fs *FileService) DeleteObjects(ctx context.Context, objectNames []string) error {
	for _, objectName := range objectNames {
		err := fs.DeleteObject(ctx, objectName)
		if err != nil {
			return fmt.Errorf("failed to delete object %s: %w", objectName, err)
		}
	}
	return nil
}
