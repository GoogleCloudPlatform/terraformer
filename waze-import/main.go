package main

/*
├── infra
│   ├── aws
│   │   ├── waze
│   │   │   ├── iam
│   │   │   │   └── global
│   │   │   └── sg
│   │   │       └── us-east1
│   │   └── waze-mapreduce
│   │       ├── iam
│   │       │   └── global
│   │       └── sg
│   │           └── us-east1
│   └── gcp
│       ├── waze-ci
│       │   ├── firewall
│       │   │   ├── europe-west1
│       │   │   └── us-east1
│       │   ├── iam
│       │   │   └── global
│       │   └── subnets
│       │       ├── europe-west1
│       │       └── us-east1
│       ├── waze-development
│       │   ├── firewall
│       │   │   ├── europe-west1
│       │   │   └── us-east1
│       │   ├── iam
│       │   │   └── global
│       │   └── subnets
│       │       ├── europe-west1
│       │       └── us-east1
│       └── waze-prod
*/

func main() {
	importGCP()
}
