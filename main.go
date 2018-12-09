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

package main

import (
	"log"
	"os"

	"waze/terraformer/gcp_terraforming"
)

func main() {
	provider := os.Args[1]
	service := os.Args[2]
	args := []string{}
	if len(os.Args) > 2 {
		args = os.Args[3:]
	}
	var err error
	switch provider {
	case "aws":
	//	err = aws_terraforming.Generate(service, args)
	case "google":
		err = gcp_terraforming.Generate(service, args)
	}
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
