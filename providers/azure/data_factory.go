// Copyright 2021 The Terraformer Authors.
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

package azure

import (
	"context"
	"fmt"
	"log"
	"reflect"
	"strings"

	"github.com/Azure/azure-sdk-for-go/services/datafactory/mgmt/2018-06-01/datafactory"
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type DataFactoryGenerator struct {
	AzureService
}

// Maps item.Properties.Type -> terraform.ResoruceType
// Information extracted from
//   SupportedResources are in:
//   @ github.com/azure/azure-sdk-for-go@v42.3.0+incompatible/services/datafactory/mgmt/2018-06-01/datafactory/models.go
//   PossibleTypeBasicDatasetValues, PossibleTypeBasicIntegrationRuntimeValues, PossibleTypeBasicLinkedServiceValues, PossibleTypeBasicTriggerValues
//   TypeBasicDataset,TypeBasicIntegrationRuntime, TypeBasicLinkedService, TypeBasicTrigger, TypeBasicDataFlow

var (
	SupportedResources = map[string]string{
		"AzureBlob":                "azurerm_data_factory_dataset_azure_blob",
		"Binary":                   "azurerm_data_factory_dataset_binary",
		"CosmosDbSqlApiCollection": "azurerm_data_factory_dataset_cosmosdb_sqlapi",
		"CustomDataset":            "azurerm_data_factory_custom_dataset",
		"DelimitedText":            "azurerm_data_factory_dataset_delimited_text",
		"HttpFile":                 "azurerm_data_factory_dataset_http",
		"Json":                     "azurerm_data_factory_dataset_json",
		"MySqlTable":               "azurerm_data_factory_dataset_mysql",
		"Parquet":                  "azurerm_data_factory_dataset_parquet",
		"PostgreSqlTable":          "azurerm_data_factory_dataset_postgresql",
		"SnowflakeTable":           "azurerm_data_factory_dataset_snowflake",
		"SqlServerTable":           "azurerm_data_factory_dataset_sql_server_table",
		"IntegrationRuntime":       "azurerm_data_factory_integration_runtime_azure",
		"Managed":                  "azurerm_data_factory_integration_runtime_azure_ssis",
		"SelfHosted":               "azurerm_data_factory_integration_runtime_self_hosted",
		"AzureBlobStorage":         "azurerm_data_factory_linked_service_azure_blob_storage",
		"AzureDatabricks":          "azurerm_data_factory_linked_service_azure_databricks",
		"AzureFileStorage":         "azurerm_data_factory_linked_service_azure_file_storage",
		"AzureFunction":            "azurerm_data_factory_linked_service_azure_function",
		"AzureSearch":              "azurerm_data_factory_linked_service_azure_search",
		"AzureSqlDatabase":         "azurerm_data_factory_linked_service_azure_sql_database",
		"AzureTableStorage":        "azurerm_data_factory_linked_service_azure_table_storage",
		"CosmosDb":                 "azurerm_data_factory_linked_service_cosmosdb",
		"CustomDataSource":         "azurerm_data_factory_linked_custom_service",
		"AzureBlobFS":              "azurerm_data_factory_linked_service_data_lake_storage_gen2",
		"AzureKeyVault":            "azurerm_data_factory_linked_service_key_vault",
		"AzureDataExplore":         "azurerm_data_factory_linked_service_kusto",
		"MySql":                    "azurerm_data_factory_linked_service_mysql",
		"OData":                    "azurerm_data_factory_linked_service_odata",
		"PostgreSql":               "azurerm_data_factory_linked_service_postgresql",
		"Sftp":                     "azurerm_data_factory_linked_service_sftp",
		"Snowflake":                "azurerm_data_factory_linked_service_snowflake",
		"SqlServer":                "azurerm_data_factory_linked_service_sql_server",
		"AzureSqlDW":               "azurerm_data_factory_linked_service_synapse",
		"Web":                      "azurerm_data_factory_linked_service_web",
		"BlobEventsTrigger":        "azurerm_data_factory_trigger_blob_event",
		"ScheduleTrigger":          "azurerm_data_factory_trigger_schedule",
		"TumblingWindowTrigger":    "azurerm_data_factory_trigger_tumbling_window",
	}
)

func getResourceTypeFrom(azureResourceName string) string {
	return SupportedResources[azureResourceName]
}

func getFieldFrom(v interface{}, field string) reflect.Value {
	reflected := reflect.ValueOf(v)
	if reflected.IsValid() {
		indirected := reflect.Indirect(reflected)
		if indirected.Kind() == reflect.Struct {
			fieldValue := indirected.FieldByName(field)
			return fieldValue
		}
	}
	return reflect.Value{}
}

func getFieldAsString(v interface{}, field string) string {
	fieldValue := getFieldFrom(v, field)
	if fieldValue.IsValid() {
		return fieldValue.String()
	}
	return ""
}

func (az *AzureService) appendResourceAs(resources []terraformutils.Resource, itemID string, itemName string, resourceType string, abbreviation string) []terraformutils.Resource {
	prefix := strings.ReplaceAll(resourceType, resourceType, abbreviation)
	suffix := strings.ReplaceAll(itemName, "-", "_")
	resourceName := prefix + "_" + suffix
	res := terraformutils.NewSimpleResource(itemID, resourceName, resourceType, az.ProviderName, []string{})
	resources = append(resources, res)
	return resources
}

func (az *DataFactoryGenerator) appendResourceFrom(resources []terraformutils.Resource, id string, name string, properties interface{}) []terraformutils.Resource {
	azureType := getFieldAsString(properties, "Type")
	if azureType != "" {
		resourceType := getResourceTypeFrom(azureType)
		if resourceType == "" {
			msg := fmt.Sprintf(`azurerm_data_factory: resource "%s" id: %s type: %s not handled yet by terraform or terraformer`, name, id, azureType)
			log.Println(msg)
		} else {
			resources = az.appendResourceAs(resources, id, name, resourceType, "adf")
		}
	}
	return resources
}

func (az *DataFactoryGenerator) listFactories() ([]datafactory.Factory, error) {
	subscriptionID, resourceGroup, authorizer := az.getClientArgs()
	client := datafactory.NewFactoriesClient(subscriptionID)
	client.Authorizer = authorizer
	var (
		iterator datafactory.FactoryListResponseIterator
		err      error
	)
	ctx := context.Background()
	if resourceGroup != "" {
		iterator, err = client.ListByResourceGroupComplete(ctx, resourceGroup)
	} else {
		iterator, err = client.ListComplete(ctx)
	}
	if err != nil {
		return nil, err
	}
	var resources []datafactory.Factory
	for iterator.NotDone() {
		item := iterator.Value()
		resources = append(resources, item)
		if err := iterator.NextWithContext(ctx); err != nil {
			log.Println(err)
			return resources, err
		}
	}
	return resources, nil
}

func (az *DataFactoryGenerator) createDataFactoryResources(dataFactories []datafactory.Factory) ([]terraformutils.Resource, error) {
	var resources []terraformutils.Resource
	for _, item := range dataFactories {
		resources = az.appendResourceAs(resources, *item.ID, *item.Name, "azurerm_data_factory", "adf")
	}
	return resources, nil
}

func getIntegrationRuntimeType(properties interface{}) string {
	azureType := getFieldAsString(properties, "Type")
	if azureType == "SelfHosted" {
		return "azurerm_data_factory_integration_runtime_self_hosted"
	}
	// item.Properties.ManagedIntegrationRuntimeTypeProperties.SsisProperties
	if typeProperties := getFieldFrom(properties, "ManagedIntegrationRuntimeTypeProperties"); typeProperties.IsValid() {
		managedRuntime := typeProperties.Interface()
		SsisProperties := getFieldFrom(managedRuntime, "SsisProperties")
		if SsisProperties.IsNil() {
			return "azurerm_data_factory_integration_runtime_azure"
		}
	}
	return "azurerm_data_factory_integration_runtime_azure_ssis"
}

func (az *DataFactoryGenerator) createIntegrationRuntimesResources(dataFactories []datafactory.Factory) ([]terraformutils.Resource, error) {
	subscriptionID, _, authorizer := az.getClientArgs()
	client := datafactory.NewIntegrationRuntimesClient(subscriptionID)
	client.Authorizer = authorizer
	ctx := context.Background()
	var resources []terraformutils.Resource
	for _, factory := range dataFactories {
		id, err := ParseAzureResourceID(*factory.ID)
		if err != nil {
			return nil, err
		}
		iterator, err := client.ListByFactoryComplete(ctx, id.ResourceGroup, *factory.Name)
		if err != nil {
			return nil, err
		}
		for iterator.NotDone() {
			item := iterator.Value()
			resourceType := getIntegrationRuntimeType(item.Properties)
			resources = az.appendResourceAs(resources, *item.ID, *item.Name, resourceType, "adfr")
			if err := iterator.NextWithContext(ctx); err != nil {
				log.Println(err)
				return resources, err
			}
		}
	}
	return resources, nil
}

func (az *DataFactoryGenerator) createLinkedServiceResources(dataFactories []datafactory.Factory) ([]terraformutils.Resource, error) {
	subscriptionID, _, authorizer := az.getClientArgs()
	client := datafactory.NewLinkedServicesClient(subscriptionID)
	client.Authorizer = authorizer
	ctx := context.Background()
	var resources []terraformutils.Resource
	for _, factory := range dataFactories {
		id, err := ParseAzureResourceID(*factory.ID)
		if err != nil {
			return nil, err
		}
		iterator, err := client.ListByFactoryComplete(ctx, id.ResourceGroup, *factory.Name)
		if err != nil {
			return nil, err
		}
		for iterator.NotDone() {
			item := iterator.Value()
			resources = az.appendResourceFrom(resources, *item.ID, *item.Name, item.Properties)
			if err = iterator.NextWithContext(ctx); err != nil {
				log.Println(err)
				return resources, err
			}
		}
	}
	return resources, nil
}

func (az *DataFactoryGenerator) createPipelineResources(dataFactories []datafactory.Factory) ([]terraformutils.Resource, error) {
	subscriptionID, _, authorizer := az.getClientArgs()
	client := datafactory.NewPipelinesClient(subscriptionID)
	client.Authorizer = authorizer
	ctx := context.Background()
	var resources []terraformutils.Resource
	for _, factory := range dataFactories {
		id, err := ParseAzureResourceID(*factory.ID)
		if err != nil {
			return nil, err
		}
		iterator, err := client.ListByFactoryComplete(ctx, id.ResourceGroup, *factory.Name)
		if err != nil {
			return nil, err
		}
		for iterator.NotDone() {
			item := iterator.Value()
			resources = az.appendResourceAs(resources, *item.ID, *item.Name, "azurerm_data_factory_pipeline", "adfp")
			if err := iterator.NextWithContext(ctx); err != nil {
				log.Println(err)
				return resources, err
			}
		}
	}
	return resources, nil
}

func (az *DataFactoryGenerator) createPipelineTriggerScheduleResources(dataFactories []datafactory.Factory) ([]terraformutils.Resource, error) {
	subscriptionID, _, authorizer := az.getClientArgs()
	client := datafactory.NewTriggersClient(subscriptionID)
	client.Authorizer = authorizer
	ctx := context.Background()
	var resources []terraformutils.Resource
	for _, factory := range dataFactories {
		id, err := ParseAzureResourceID(*factory.ID)
		if err != nil {
			return nil, err
		}
		iterator, err := client.ListByFactoryComplete(ctx, id.ResourceGroup, *factory.Name)
		if err != nil {
			return nil, err
		}
		for iterator.NotDone() {
			item := iterator.Value()
			resources = az.appendResourceFrom(resources, *item.ID, *item.Name, item.Properties)
			if err := iterator.NextWithContext(ctx); err != nil {
				log.Println(err)
				return resources, err
			}
		}
	}
	return resources, nil
}

func (az *DataFactoryGenerator) createDataFlowResources(dataFactories []datafactory.Factory) ([]terraformutils.Resource, error) {
	subscriptionID, _, authorizer := az.getClientArgs()
	client := datafactory.NewDataFlowsClient(subscriptionID)
	client.Authorizer = authorizer
	ctx := context.Background()
	var resources []terraformutils.Resource
	for _, factory := range dataFactories {
		id, err := ParseAzureResourceID(*factory.ID)
		if err != nil {
			return nil, err
		}
		iterator, err := client.ListByFactoryComplete(ctx, id.ResourceGroup, *factory.Name)
		if err != nil {
			return nil, err
		}
		for iterator.NotDone() {
			item := iterator.Value()
			resources = az.appendResourceAs(resources, *item.ID, *item.Name, "azurerm_data_factory_data_flow", "adfl")
			if err := iterator.NextWithContext(ctx); err != nil {
				log.Println(err)
				return resources, err
			}
		}
	}
	return resources, nil
}

func (az *DataFactoryGenerator) createPipelineDatasetResources(dataFactories []datafactory.Factory) ([]terraformutils.Resource, error) {
	subscriptionID, _, authorizer := az.getClientArgs()
	client := datafactory.NewDatasetsClient(subscriptionID)
	client.Authorizer = authorizer
	ctx := context.Background()
	var resources []terraformutils.Resource
	for _, factory := range dataFactories {
		id, err := ParseAzureResourceID(*factory.ID)
		if err != nil {
			return nil, err
		}
		iterator, err := client.ListByFactoryComplete(ctx, id.ResourceGroup, *factory.Name)
		if err != nil {
			return nil, err
		}
		for iterator.NotDone() {
			item := iterator.Value()
			resources = az.appendResourceFrom(resources, *item.ID, *item.Name, item.Properties)
			if err := iterator.NextWithContext(ctx); err != nil {
				log.Println(err)
				return resources, err
			}
		}
	}
	return resources, nil
}

func (az *DataFactoryGenerator) InitResources() error {

	dataFactories, err := az.listFactories()
	if err != nil {
		return err
	}

	factoriesFunctions := []func([]datafactory.Factory) ([]terraformutils.Resource, error){
		az.createDataFactoryResources,
		az.createIntegrationRuntimesResources,
		az.createLinkedServiceResources,
		az.createPipelineResources,
		az.createPipelineTriggerScheduleResources,
		az.createPipelineDatasetResources,
		az.createDataFlowResources,
	}

	for _, f := range factoriesFunctions {
		resources, ero := f(dataFactories)
		if ero != nil {
			return ero
		}
		az.Resources = append(az.Resources, resources...)
	}
	return nil
}

// PostGenerateHook for formatting json properties as heredoc
// - azurerm_data_factory_pipeline property activities_json
func (az *DataFactoryGenerator) PostConvertHook() error {
	for i, resource := range az.Resources {
		if resource.InstanceInfo.Type == "azurerm_data_factory_pipeline" {
			if val, ok := az.Resources[i].Item["activities_json"]; ok {
				if val != nil {
					json := val.(string)
					// json := asJson(val)
					hereDoc := asHereDoc(json)
					az.Resources[i].Item["activities_json"] = hereDoc
				}
			}
		}
	}
	return nil
}
