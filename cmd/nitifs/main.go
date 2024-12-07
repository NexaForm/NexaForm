package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

// Record represents the structure of a MinIO event record
type Record struct {
	EventName string `json:"eventName"`
	EventTime string `json:"eventTime"`
	S3        struct {
		Bucket struct {
			Name string `json:"name"`
		} `json:"bucket"`
		Object struct {
			Key  string `json:"key"`
			Size int    `json:"size"`
		} `json:"object"`
	} `json:"s3"`
}

// Event represents the structure of a MinIO bucket event
type Event struct {
	Records []Record `json:"Records"`
}

func main() {
	// Set up Redis connection
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6380",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	ctx := context.Background()
	redisKey := "bucketevents"

	fmt.Printf("Listening for events in Redis hash key: %s\n", redisKey)

	for {
		// Fetch all entries in the Redis hash key
		events, err := rdb.HGetAll(ctx, redisKey).Result()
		if err != nil {
			log.Printf("Error fetching events from Redis: %v", err)
			time.Sleep(5 * time.Second)
			continue
		}

		// Iterate over each field in the hash and print the event, then delete it
		for key, value := range events {
			var event Event
			if err := json.Unmarshal([]byte(value), &event); err != nil {
				log.Printf("Error decoding event for key %s: %v", key, err)
				continue
			}
			for _, record := range event.Records {
				log.Printf("New event for object '%s': %+v", key, record)
			}

			// Delete the processed event from the Redis hash
			if _, err := rdb.HDel(ctx, redisKey, key).Result(); err != nil {
				log.Printf("Error deleting processed event for key %s: %v", key, err)
			}
		}

		// Sleep for a few seconds before fetching again
		time.Sleep(5 * time.Second)
	}
}
