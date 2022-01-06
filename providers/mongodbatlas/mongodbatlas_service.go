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

package mongodbatlas

import (
	"log"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/mongodb-forks/digest"
	mongodbatlas "go.mongodb.org/atlas/mongodbatlas"
)

type MongoDBAtlasService struct { //nolint
	terraformutils.Service
}

func (s *MongoDBAtlasService) generateClient() *mongodbatlas.Client {
	t := digest.NewTransport(s.Args["public_key"].(string), s.Args["private_key"].(string))
	tc, err := t.Client()
	if err != nil {
		log.Fatalf(err.Error())
	}

	return mongodbatlas.NewClient(tc)
}
