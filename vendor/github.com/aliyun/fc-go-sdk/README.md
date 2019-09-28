Aliyun FunctionCompute Go SDK
=================================

API Reference :

[FC API](https://help.aliyun.com/document_detail/52877.html)

[![GitHub Version](https://badge.fury.io/gh/aliyun%2Ffc-go-sdk.svg)](https://badge.fury.io/gh/aliyun%2Ffc-go-sdk)
[![Build Status](https://travis-ci.org/aliyun/fc-go-sdk.svg?branch=master)](https://travis-ci.org/aliyun/fc-go-sdk)
<!-- [![Coverage Status](https://coveralls.io/repos/github/aliyun/fc-go-sdk/badge.svg?branch=master&foo=bar)](https://coveralls.io/github/aliyun/fc-go-sdk?branch=master&foo=bar) -->

VERSION
--------
go >= 1.8

Overview
--------
Aliyun FunctionCompute Go SDK, sample
```go
package main

import (
	"fmt"
	"os"
)

func main() {
	serviceName := "service555"
	client, _ := fc.NewClient(os.Getenv("ENDPOINT"), "2016-08-15", os.Getenv("ACCESS_KEY_ID"), os.Getenv("ACCESS_KEY_SECRET"))

	fmt.Println("Creating service")
	createServiceOutput, err := client.CreateService(fc.NewCreateServiceInput().
		WithServiceName(serviceName).
		WithDescription("this is a smoke test for go sdk"))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	if createServiceOutput != nil {
		fmt.Printf("CreateService response: %s \n", createServiceOutput)
	}

	// GetService
	fmt.Println("Getting service")
	getServiceOutput, err := client.GetService(fc.NewGetServiceInput(serviceName))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	} else {
		fmt.Printf("GetService response: %s \n", getServiceOutput)
	}

	// UpdateService
	fmt.Println("Updating service")
	updateServiceInput := fc.NewUpdateServiceInput(serviceName).WithDescription("new description")
	updateServiceOutput, err := client.UpdateService(updateServiceInput)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	} else {
		fmt.Printf("UpdateService response: %s \n", updateServiceOutput)
	}

	// UpdateService with IfMatch
	fmt.Println("Updating service with IfMatch")
	updateServiceInput2 := fc.NewUpdateServiceInput(serviceName).WithDescription("new description2").
		WithIfMatch(updateServiceOutput.Header.Get("ETag"))
	updateServiceOutput2, err := client.UpdateService(updateServiceInput2)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	} else {
		fmt.Printf("UpdateService response: %s \n", updateServiceOutput2)
	}

	// UpdateService with wrong IfMatch
	fmt.Println("Updating service with wrong IfMatch")
	updateServiceInput3 := fc.NewUpdateServiceInput(serviceName).WithDescription("new description2").
		WithIfMatch("1234")
	updateServiceOutput3, err := client.UpdateService(updateServiceInput3)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	} else {
		fmt.Printf("UpdateService response: %s \n", updateServiceOutput3)
	}

	// ListServices
	fmt.Println("Listing services")
	listServicesOutput, err := client.ListServices(fc.NewListServicesInput().WithLimit(100))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	} else {
		fmt.Printf("ListServices response: %s \n", listServicesOutput)
	}

	// CreateFunction
	fmt.Println("Creating function1")
	createFunctionInput1 := fc.NewCreateFunctionInput(serviceName).WithFunctionName("testf1").
        		WithDescription("go sdk test function").
        		WithHandler("main.my_handler").WithRuntime("python2.7").
        		WithCode(fc.NewCode().WithFiles("./testCode/hello_world.zip")).
        		WithTimeout(5)

	createFunctionOutput, err := client.CreateFunction(createFunctionInput1)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	} else {
		fmt.Printf("CreateFunction response: %s \n", createFunctionOutput)
	}
	fmt.Println("Creating function2")
	createFunctionOutput2, err := client.CreateFunction(createFunctionInput1.WithFunctionName("testf2"))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	} else {
		fmt.Printf("CreateFunction response: %s \n", createFunctionOutput2)
	}

	// ListFunctions
	fmt.Println("Listing functions")
	listFunctionsOutput, err := client.ListFunctions(fc.NewListFunctionsInput(serviceName).WithPrefix("test"))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	} else {
		fmt.Printf("ListFunctions response: %s \n", listFunctionsOutput)
	}

	// UpdateFunction
	fmt.Println("Updating function")
	updateFunctionOutput, err := client.UpdateFunction(fc.NewUpdateFunctionInput(serviceName, "testf1").
		WithDescription("newdesc"))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	} else {
		fmt.Printf("UpdateFunction response: %s \n", updateFunctionOutput)
	}

	// InvokeFunction
	fmt.Println("Invoking function, log type Tail")
	invokeInput := fc.NewInvokeFunctionInput(serviceName, "testf1").WithLogType("Tail")
	invokeOutput, err := client.InvokeFunction(invokeInput)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	} else {
		fmt.Printf("InvokeFunction response: %s \n", invokeOutput)
		logResult, err := invokeOutput.GetLogResult()
		if err != nil {
			fmt.Printf("Failed to get LogResult due to %v\n", err)
		} else {
			fmt.Printf("Invoke function LogResult %s \n", logResult)
		}
	}

	fmt.Println("Invoking function, log type None")
	invokeInput = fc.NewInvokeFunctionInput(serviceName, "testf1").WithLogType("None")
	invokeOutput, err = client.InvokeFunction(invokeInput)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	} else {
		fmt.Printf("InvokeFunction response: %s \n", invokeOutput)
	}

    // PublishServiceVersion
	fmt.Println("Publishing service version")
	publishServiceVersionInput := fc.NewPublishServiceVersionInput(serviceName)
	publishServiceVersionOutput, err := client.PublishServiceVersion(publishServiceVersionInput)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	} else {
		fmt.Printf("PublishServiceVersion response: %s \n", publishServiceVersionOutput)
	}

    // PublishServiceVersion with IfMatch
    fmt.Println("Publishing service version with IfMatch")
    publishServiceVersionInput2 := fc.NewPublishServiceVersionInput(serviceName).
	       WithIfMatch(getServiceOutput.Header.Get("ETag"))
	publishServiceVersionOutput2, err := client.PublishServiceVersion(publishServiceVersionInput2)
	if err != nil {
	     fmt.Fprintln(os.Stderr, err)
	} else {
	      fmt.Printf("PublishServiceVersion response: %s \n", publishServiceVersionOutput2)
	}

	// PublishServiceVersion with wrong IfMatch
	fmt.Println("Publishing service with wrong IfMatch")
	publishServiceVersionInput3 := fc.NewPublishServiceVersionInput(serviceName).
	       WithIfMatch("1234")
	publishServiceVersionOutput3, err := client.PublishServiceVersion(publishServiceVersionInput3)
	if err != nil {
	       fmt.Fprintln(os.Stderr, err)
	} else {
	       fmt.Printf("PublishServiceVersion response: %s \n", publishServiceVersionOutput3)
	}

	// ListServiceVersions
	fmt.Println("Listing service versions")
	listServiceVersionsOutput, err := client.ListServiceVersions(fc.NewListServiceVersionsInput(serviceName).WithLimit(10))
	if err != nil {
	       fmt.Fprintln(os.Stderr, err)
	} else {
	       fmt.Printf("ListServiceVersions response: %s \n", listServiceVersionsOutput)
	}

	// GetService with qualifier
	fmt.Println("Getting service with qualifier")
	getServiceOutput2, err := client.GetService(fc.NewGetServiceInput(serviceName).WithQualifier(publishServiceVersionOutput.VersionID))
	if err != nil {
	       fmt.Fprintln(os.Stderr, err)
	} else {
	       fmt.Printf("GetService with qualifier response: %s \n", getServiceOutput2)
	}

	// CreateAlias
	aliasName := "alias"
	fmt.Println("Creating alias")
	createAliasOutput, err := client.CreateAlias(fc.NewCreateAliasInput(serviceName).WithAliasName(aliasName).WithVersionID(publishServiceVersionOutput.VersionID))
	if err != nil {
	       fmt.Fprintln(os.Stderr, err)
	} else {
	       fmt.Printf("CreateAlias response: %s \n", createAliasOutput)
	}

	// GetAlias
	fmt.Println("Getting alias")
	getAliasOutput, err := client.GetAlias(fc.NewGetAliasInput(serviceName, aliasName))
	if err != nil {
	       fmt.Fprintln(os.Stderr, err)
	} else {
	       fmt.Printf("GetAlias response: %s \n", getAliasOutput)
	}

	// UpdateAlias
	fmt.Println("Updating alias")
	updateAliasOutput, err := client.UpdateAlias(fc.NewUpdateAliasInput(serviceName, aliasName).WithVersionID(publishServiceVersionOutput2.VersionID))
	if err != nil {
	       fmt.Fprintln(os.Stderr, err)
	} else {
	       fmt.Printf("UpdateAlias response: %s \n", updateAliasOutput)
	}

	// ListAliases
	fmt.Println("Listing aliases")
	listAliasesOutput, err := client.ListAliases(fc.NewListAliasesInput(serviceName))
	if err != nil {
	       fmt.Fprintln(os.Stderr, err)
	} else {
	       fmt.Printf("ListAliases response: %s \n", listAliasesOutput)
	}

	// DeleteAlias
	fmt.Println("Deleting aliases")
	deleteAliasOutput, err := client.DeleteAlias(fc.NewDeleteAliasInput(serviceName, aliasName))
	if err != nil {
	       fmt.Fprintln(os.Stderr, err)
	} else {
	       fmt.Printf("DeleteAlias response: %s \n", deleteAliasOutput)
	}

	// DeleteServiceVersion
	fmt.Println("Deleting service version")
	deleteServiceVersionOutput, err := client.DeleteServiceVersion(fc.NewDeleteServiceVersionInput(serviceName, publishServiceVersionOutput.VersionID))
	if err != nil {
	       fmt.Fprintln(os.Stderr, err)
	} else {
	       fmt.Printf("DeleteServiceVersion response: %s \n", deleteServiceVersionOutput)
	}

	deleteServiceVersionOutput2, err := client.DeleteServiceVersion(fc.NewDeleteServiceVersionInput(serviceName, publishServiceVersionOutput2.VersionID))
	if err != nil {
	       fmt.Fprintln(os.Stderr, err)
	} else {
	       fmt.Printf("DeleteServiceVersion response: %s \n", deleteServiceVersionOutput2)
	}

	// DeleteFunction
	fmt.Println("Deleting functions")
	listFunctionsOutput, err = client.ListFunctions(fc.NewListFunctionsInput(serviceName).WithLimit(10))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	} else {
		fmt.Printf("ListFunctions response: %s \n", listFunctionsOutput)
		for _, fuc := range listFunctionsOutput.Functions {
			fmt.Printf("Deleting function %s \n", *fuc.FunctionName)
			if output, err := client.DeleteFunction(fc.NewDeleteFunctionInput(serviceName, *fuc.FunctionName)); err != nil {
				fmt.Fprintln(os.Stderr, err)
			} else {
				fmt.Printf("DeleteFunction response: %s \n", output)
			}

		}
	}

	// DeleteService
	fmt.Println("Deleting service")
	deleteServiceOutput, err := client.DeleteService(fc.NewDeleteServiceInput(serviceName))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	} else {
		fmt.Printf("DeleteService response: %s \n", deleteServiceOutput)
	}
}

```


More resources
--------------
- [Aliyun FunctionCompute docs](https://help.aliyun.com/product/50980.html)

Contacting us
-------------
- [Links](https://help.aliyun.com/document_detail/53087.html)

License
-------
- [MIT](https://github.com/aliyun/fc-python-sdk/blob/master/LICENSE)
