// Copyright 2018 The Terraformer Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package terraformoutput

import (
	"context"
	"log"
	"strings"

	"cloud.google.com/go/storage"
)

type BucketState struct {
	Name string
}

func (b BucketState) BucketGetTfData(path string) interface{} {
	name := strings.ReplaceAll(b.Name, "gs://", "")
	bucketStateData := map[string]interface{}{
		"terraform": map[string]interface{}{
			"backend": []map[string]interface{}{
				{
					"gcs": map[string]interface{}{
						"bucket": name,
						"prefix": b.BucketPrefix(path),
					},
				},
			},
		},
	}
	return bucketStateData
}

func (b BucketState) BucketPrefix(path string) string {
	return strings.TrimSuffix(path, "/")
}

func (b BucketState) BucketUpload(path string, file []byte) error {
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	name := strings.ReplaceAll(b.Name, "gs://", "")
	wc := client.Bucket(name).Object(b.BucketPrefix(path) + "/default.tfstate").NewWriter(ctx)
	if _, err = wc.Write(file); err != nil {
		return err
	}
	if err := wc.Close(); err != nil {
		return err
	}
	return nil
}
