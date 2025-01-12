package service

import (
	"NexaForm/internal/survey"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type FileService struct {
	client        *minio.Client
	bucketName    string
	redisClient   *redis.Client
	surveyOps     *survey.Ops
	loggerService *LoggerService
}

// Allowed file types and max size
var allowedFileTypes = []string{"image/", "video/", "audio/"}

// NewFileService initializes a new MinIO client
func NewFileService(surveyOps *survey.Ops, endpoint, accessKey, secretKey, bucketName string, useSSL bool, logger *LoggerService) (*FileService, error) {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6380",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	client, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		return nil, err
	}

	// Ensure bucket exists
	err = client.MakeBucket(context.Background(), bucketName, minio.MakeBucketOptions{Region: ""})
	if err != nil {
		exists, bucketErr := client.BucketExists(context.Background(), bucketName)
		if bucketErr != nil || !exists {
			return nil, err
		}
	}

	return &FileService{
		client:        client,
		bucketName:    bucketName,
		redisClient:   redisClient,
		surveyOps:     surveyOps,
		loggerService: logger,
	}, nil
}

// ValidateFileType checks if the file type is allowed
func ValidateFileType(contentType string) bool {
	for _, prefix := range allowedFileTypes {
		if strings.HasPrefix(contentType, prefix) {
			return true
		}
	}
	return false
}

// GeneratePresignedDownloadURLs retrieves all questions by survey ID and generates presigned download URLs for attachments
func (fs *FileService) GeneratePresignedDownloadURLs(ctx context.Context, surveyID uuid.UUID) (map[string]string, error) {
	// Get all questions for the given survey ID
	questions, err := fs.surveyOps.GetQuestionsBySurveyID(ctx, surveyID)
	if err != nil {
		return nil, err
	}

	// Prepare a map to hold presigned URLs for each attachment
	downloadURLs := make(map[string]string)

	// Iterate over each question and generate presigned URLs for their attachments
	for _, question := range questions {
		for _, attachment := range question.Attachments {
			// Check if the attachment is persisted in MinIO
			exists, err := fs.CheckObjectExists(ctx, attachment.FilePath)
			if err != nil {
				return nil, fmt.Errorf("failed to check existence of object %s: %w", attachment.FilePath, err)
			}
			if exists {
				// Generate presigned URL only if the object exists
				url, err := fs.client.PresignedGetObject(ctx, fs.bucketName, attachment.FilePath, time.Hour*24, nil)
				if err != nil {
					return nil, fmt.Errorf("failed to generate presigned URL for object %s: %w", attachment.FilePath, err)
				}
				downloadURLs[attachment.FilePath] = url.String()
			}
		}
	}

	return downloadURLs, nil
}
func (fs *FileService) CheckObjectExists(ctx context.Context, filePath string) (bool, error) {
	_, err := fs.client.StatObject(ctx, fs.bucketName, filePath, minio.StatObjectOptions{})
	if err != nil {
		if minio.ToErrorResponse(err).Code == "NoSuchKey" {
			return false, nil // Object does not exist
		}
		return false, err // Other errors
	}
	return true, nil // Object exists
}

// GeneratePresignedUploadURLs generates presigned URLs for uploading multiple files
func (fs *FileService) GeneratePresignedUploadURLs(ctx context.Context, objectPaths []string, contentTypes []string) (map[string]string, error) {
	if len(objectPaths) != len(contentTypes) {
		return nil, errors.New("mismatched object paths and content types")
	}

	urls := make(map[string]string)
	for i, objectPath := range objectPaths {
		if !ValidateFileType(contentTypes[i]) {
			return nil, errors.New("invalid file type")
		}

		url, err := fs.client.PresignedPutObject(ctx, fs.bucketName, objectPath, time.Hour*24)
		if err != nil {
			return nil, err
		}
		urls[objectPath] = url.String()
	}
	return urls, nil
}

// ListenForEvents listens for events in the Redis hash key and updates attachments accordingly
func (fs *FileService) ListenForEvents(ctx context.Context) {
	logger := fs.loggerService
	defer logger.Sync()
	redisKey := "bucketevents"
	logger.LogInfo(ctx, ServiceLogger, fmt.Sprintf("Listening for events in Redis hash key: %s\n", redisKey))

	for {
		// Fetch all entries in the Redis hash key
		events, err := fs.redisClient.HGetAll(ctx, redisKey).Result()
		if err != nil {
			logger.LogError(ctx, ServiceLogger, err.Error())
			time.Sleep(5 * time.Second)
			continue
		}

		// Iterate over each field in the hash and process the event
		for key, value := range events {
			var event struct {
				Records []struct {
					S3 struct {
						Object struct {
							Key string `json:"key"`
						} `json:"object"`
					} `json:"s3"`
				} `json:"Records"`
			}

			if err := json.Unmarshal([]byte(value), &event); err != nil {
				logger.LogError(ctx, ServiceLogger, err.Error())
				continue
			}
			logger.LogInfo(ctx, ServiceLogger, fmt.Sprintf("Processing event for key: %s", key))

			for _, record := range event.Records {
				filePath := record.S3.Object.Key

				// Decode the file path
				decodedFilePath, err := url.PathUnescape(filePath)
				if err != nil {
					logger.LogError(ctx, ServiceLogger, fmt.Sprintf("Error decoding file path %s: %v", filePath, err))
					continue
				}

				// Extract question_id from the file path (e.g., attachments/<question_id>/<file_name>)
				parts := strings.Split(decodedFilePath, "/")
				if len(parts) < 3 {
					logger.LogError(ctx, ServiceLogger, fmt.Sprintf("Invalid file path format: %s", decodedFilePath))
					continue
				}

				questionIDStr := parts[1]
				questionID, err := uuid.Parse(questionIDStr)
				if err != nil {
					logger.LogError(ctx, ServiceLogger, fmt.Sprintf("Invalid question ID in file path: %s", questionIDStr))
					continue
				}

				// Log the attachment details instead of modifying the database
				logger.LogInfo(ctx, ServiceLogger, fmt.Sprintf("Attachment processed - QuestionID: %s, FilePath: %s", questionID.String(), decodedFilePath))
			}

			// Delete the processed event from the Redis hash
			if _, err := fs.redisClient.HDel(ctx, redisKey, key).Result(); err != nil {
				logger.LogError(ctx, ServiceLogger, fmt.Sprintf("Error deleting processed event for key %s: %v", key, err))
			} else {
				logger.LogInfo(ctx, ServiceLogger, fmt.Sprintf("Successfully deleted processed event for key: %s", key))
			}
		}

		// Sleep for a few seconds before fetching again
		time.Sleep(5 * time.Second)
	}
}
