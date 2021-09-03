package azure

import (
	"context"
	"fmt"
	"log"
	"reflect"

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

func (g *DataFactoryGenerator) appendResourceFrom(resources []terraformutils.Resource, id string, name string, properties interface{}) []terraformutils.Resource {
	azureType := getFieldAsString(properties, "Type")
	if azureType != "" {
		resourceType := getResourceTypeFrom(azureType)
		if resourceType == "" {
			msg := fmt.Sprintf(`azurerm_data_factory: resource "%s" id: %s type: %s not handled yet by terraform or terraformer`, name, id, azureType)
			log.Println(msg)
		} else {
			resources = g.appendResourceAs(resources, id, name, resourceType, "adf")
		}
	}
	return resources
}

func (g *DataFactoryGenerator) listFactories() ([]datafactory.Factory, error) {
	subscriptionID, authorizer := g.getArgsProperties()
	client := datafactory.NewFactoriesClient(subscriptionID)
	client.Authorizer = authorizer
	var (
		iterator datafactory.FactoryListResponseIterator
		err      error
	)
	ctx := context.Background()
	if rg := g.Args["resource_group"].(string); rg != "" {
		iterator, err = client.ListByResourceGroupComplete(ctx, rg)
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

func (g *DataFactoryGenerator) createDataFactoryResources(dataFactories []datafactory.Factory) ([]terraformutils.Resource, error) {
	var resources []terraformutils.Resource
	for _, item := range dataFactories {
		resources = g.appendResourceAs(resources, *item.ID, *item.Name, "azurerm_data_factory", "adf")
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

func (g *DataFactoryGenerator) createIntegrationRuntimesResources(dataFactories []datafactory.Factory) ([]terraformutils.Resource, error) {
	subscriptionID, authorizer := g.getArgsProperties()
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
			resources = g.appendResourceAs(resources, *item.ID, *item.Name, resourceType, "adf")
			if err := iterator.NextWithContext(ctx); err != nil {
				log.Println(err)
				return resources, err
			}
		}
	}
	return resources, nil
}

func (g *DataFactoryGenerator) createLinkedServiceResources(dataFactories []datafactory.Factory) ([]terraformutils.Resource, error) {
	subscriptionID, authorizer := g.getArgsProperties()
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
			resources = g.appendResourceFrom(resources, *item.ID, *item.Name, item.Properties)
			if err = iterator.NextWithContext(ctx); err != nil {
				log.Println(err)
				return resources, err
			}
		}
	}
	return resources, nil
}

func (g *DataFactoryGenerator) createPipelineResources(dataFactories []datafactory.Factory) ([]terraformutils.Resource, error) {
	subscriptionID, authorizer := g.getArgsProperties()
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
			resources = g.appendResourceAs(resources, *item.ID, *item.Name, "azurerm_data_factory_pipeline", "adf")
			if err := iterator.NextWithContext(ctx); err != nil {
				log.Println(err)
				return resources, err
			}
		}
	}
	return resources, nil
}

func (g *DataFactoryGenerator) createPipelineTriggerScheduleResources(dataFactories []datafactory.Factory) ([]terraformutils.Resource, error) {
	subscriptionID, authorizer := g.getArgsProperties()
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
			resources = g.appendResourceFrom(resources, *item.ID, *item.Name, item.Properties)
			if err := iterator.NextWithContext(ctx); err != nil {
				log.Println(err)
				return resources, err
			}
		}
	}
	return resources, nil
}

func (g *DataFactoryGenerator) createDataFlowResources(dataFactories []datafactory.Factory) ([]terraformutils.Resource, error) {
	subscriptionID, authorizer := g.getArgsProperties()
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
			resources = g.appendResourceAs(resources, *item.ID, *item.Name, "azurerm_data_factory_data_flow", "adf")
			if err := iterator.NextWithContext(ctx); err != nil {
				log.Println(err)
				return resources, err
			}
		}
	}
	return resources, nil
}

func (g *DataFactoryGenerator) createPipelineDatasetResources(dataFactories []datafactory.Factory) ([]terraformutils.Resource, error) {
	subscriptionID, authorizer := g.getArgsProperties()
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
			resources = g.appendResourceFrom(resources, *item.ID, *item.Name, item.Properties)
			if err := iterator.NextWithContext(ctx); err != nil {
				log.Println(err)
				return resources, err
			}
		}
	}
	return resources, nil
}

func (g *DataFactoryGenerator) InitResources() error {

	dataFactories, err := g.listFactories()
	if err != nil {
		return err
	}

	factoriesFunctions := []func([]datafactory.Factory) ([]terraformutils.Resource, error){
		g.createDataFactoryResources,
		g.createIntegrationRuntimesResources,
		g.createLinkedServiceResources,
		g.createPipelineResources,
		g.createPipelineTriggerScheduleResources,
		g.createPipelineDatasetResources,
		g.createDataFlowResources,
	}

	for _, f := range factoriesFunctions {
		resources, ero := f(dataFactories)
		if ero != nil {
			return ero
		}
		g.Resources = append(g.Resources, resources...)
	}
	return nil
}

// PostGenerateHook for formatting json properties as heredoc
// - azurerm_data_factory_pipeline property activities_json
func (g *DataFactoryGenerator) PostConvertHook() error {
	for i, resource := range g.Resources {
		if resource.InstanceInfo.Type == "azurerm_data_factory_pipeline" {
			if val, ok := g.Resources[i].Item["activities_json"]; ok {
				if val != nil {
					json := val.(string)
					// json := asJson(val)
					hereDoc := asHereDoc(json)
					g.Resources[i].Item["activities_json"] = hereDoc
				}
			}
		}
	}
	return nil
}
