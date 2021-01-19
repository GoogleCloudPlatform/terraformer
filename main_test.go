package main

import (
	"github.com/GoogleCloudPlatform/terraformer/cmd"
	"log"
	"os"
	"testing"
	"time"
)

func TestTerraformerMain(t1 *testing.T) {
	t1.Run("run main", func(t1 *testing.T) {
		tCommand := cmd.NewCmdRoot()
		tCommand.SetArgs([]string{
			"import",
			"aws",
			"--regions=sa-east-1",
			//"--regions=us-west-1",
			//"--regions=us-east-1",
			//"--regions=ap-northeast-1,eu-central-1,eu-west-1,us-east-1,af-south-1,ap-northeast-2,eu-west-2,us-west-1,us-west-2,ap-east-1,ap-south-1,ap-southeast-2,ca-central-1,sa-east-1,us-east-2,ap-southeast-1,aws-global,eu-north-1,eu-south-1,eu-west-3,me-south-1",
			"--resources=\"*\"",
			//"--resources=sg",
			//"--resources=cloudformation,sg,s3",
			//"--resources=cloudformation,sg",
			//"--resources=sg,vpc",
			//"--resources=kms",
			"--profile=nubank",
			//"--retry-number=5",
			//"--retry-sleep-ms=300",
		})
		start := time.Now()
		if err := tCommand.Execute(); err != nil {
			log.Println(err)
			os.Exit(1)
		}
		log.Printf("Importing took %s", time.Since(start))

	})
}
