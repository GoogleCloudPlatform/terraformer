package main

import (
	"context"
	"log"
	"os"
	"strings"

	"cloud.google.com/go/storage"
)

type bucket struct{}

const bucketStateName = "waze-terraform-state"

func bucketGetTfData(path string) interface{} {
	bucketStateData := map[string]interface{}{
		"terraform": map[string]interface{}{
			"backend": map[string]interface{}{
				"gcs": map[string]interface{}{
					"bucket": bucketStateName,
					"prefix": bucketPrefix(path),
				},
			},
		},
	}
	return bucketStateData
}

func bucketPrefix(path string) string {
	rootPath, _ := os.Getwd()
	return strings.Replace(path, rootPath+"/imported/infra/", "", -1)
}

func bucketUpload(path string, file []byte) error {
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	wc := client.Bucket(bucketStateName).Object(bucketPrefix(path) + "/default.tfstate").NewWriter(ctx)
	if _, err = wc.Write(file); err != nil {
		return err
	}
	if err := wc.Close(); err != nil {
		return err
	}
	return nil
}
