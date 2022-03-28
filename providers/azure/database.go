// Copyright 2019 The Terraformer Authors.
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
	"strings"

	"github.com/Azure/azure-sdk-for-go/services/mariadb/mgmt/2018-06-01/mariadb"
	"github.com/Azure/azure-sdk-for-go/services/mysql/mgmt/2017-12-01/mysql"
	"github.com/Azure/azure-sdk-for-go/services/postgresql/mgmt/2017-12-01/postgresql"
	"github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/2017-03-01-preview/sql"
	"github.com/Azure/go-autorest/autorest"
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/hashicorp/go-azure-helpers/authentication"
)

type DatabasesGenerator struct {
	AzureService
}

func (g *DatabasesGenerator) getMariaDBServers() ([]mariadb.Server, error) {
	ctx := context.Background()
	SubscriptionID := g.Args["config"].(authentication.Config).SubscriptionID
	Authorizer := g.Args["authorizer"].(autorest.Authorizer)

	Client := mariadb.NewServersClient(SubscriptionID)
	Client.Authorizer = Authorizer

	var (
		Servers mariadb.ServerListResult
		err     error
	)
	if rg := g.Args["resource_group"].(string); rg != "" {
		Servers, err = Client.ListByResourceGroup(ctx, rg)
	} else {
		Servers, err = Client.List(ctx)
	}
	if err != nil {
		return nil, err
	}

	return *Servers.Value, nil
}

func (g *DatabasesGenerator) createMariaDBServerResources(servers []mariadb.Server) ([]terraformutils.Resource, error) {
	var resources []terraformutils.Resource

	for _, server := range servers {
		resources = append(resources, terraformutils.NewResource(
			*server.ID,
			*server.Name,
			"azurerm_mariadb_server",
			g.ProviderName,
			map[string]string{},
			[]string{},
			map[string]interface{}{
				"administrator_login_password": "",
			}))
	}

	return resources, nil
}

func (g *DatabasesGenerator) createMariaDBConfigurationResources(servers []mariadb.Server) ([]terraformutils.Resource, error) {
	var resources []terraformutils.Resource
	ctx := context.Background()
	SubscriptionID := g.Args["config"].(authentication.Config).SubscriptionID
	Authorizer := g.Args["authorizer"].(autorest.Authorizer)

	Client := mariadb.NewConfigurationsClient(SubscriptionID)
	Client.Authorizer = Authorizer

	for _, server := range servers {
		id, err := ParseAzureResourceID(*server.ID)
		if err != nil {
			return nil, err
		}
		configs, err := Client.ListByServer(ctx, id.ResourceGroup, *server.Name)
		if err != nil {
			return nil, err
		}

		for _, config := range *configs.Value {
			resources = append(resources, terraformutils.NewSimpleResource(
				*config.ID,
				*config.Name+"-"+*server.Name,
				"azurerm_mariadb_configuration",
				g.ProviderName,
				[]string{"value"}))
		}
	}

	return resources, nil
}

func (g *DatabasesGenerator) createMariaDBDatabaseResources(servers []mariadb.Server) ([]terraformutils.Resource, error) {
	var resources []terraformutils.Resource
	ctx := context.Background()
	SubscriptionID := g.Args["config"].(authentication.Config).SubscriptionID
	Authorizer := g.Args["authorizer"].(autorest.Authorizer)

	Client := mariadb.NewDatabasesClient(SubscriptionID)
	Client.Authorizer = Authorizer

	for _, server := range servers {
		id, err := ParseAzureResourceID(*server.ID)
		if err != nil {
			return nil, err
		}
		databases, err := Client.ListByServer(ctx, id.ResourceGroup, *server.Name)
		if err != nil {
			return nil, err
		}

		for _, database := range *databases.Value {
			resources = append(resources, terraformutils.NewSimpleResource(
				*database.ID,
				*database.Name+"-"+*server.Name,
				"azurerm_mariadb_database",
				g.ProviderName,
				[]string{}))
		}
	}

	return resources, nil
}

func (g *DatabasesGenerator) createMariaDBFirewallRuleResources(servers []mariadb.Server) ([]terraformutils.Resource, error) {
	var resources []terraformutils.Resource
	ctx := context.Background()
	SubscriptionID := g.Args["config"].(authentication.Config).SubscriptionID
	Authorizer := g.Args["authorizer"].(autorest.Authorizer)

	Client := mariadb.NewFirewallRulesClient(SubscriptionID)
	Client.Authorizer = Authorizer
	for _, server := range servers {
		id, err := ParseAzureResourceID(*server.ID)
		if err != nil {
			return nil, err
		}

		rules, err := Client.ListByServer(ctx, id.ResourceGroup, *server.Name)
		if err != nil {
			return nil, err
		}
		for _, rule := range *rules.Value {
			resources = append(resources, terraformutils.NewSimpleResource(
				*rule.ID,
				*rule.Name,
				"azurerm_mariadb_firewall_rule",
				g.ProviderName,
				[]string{}))
		}
	}

	return resources, nil
}

func (g *DatabasesGenerator) createMariaDBVirtualNetworkRuleResources(servers []mariadb.Server) ([]terraformutils.Resource, error) {
	var resources []terraformutils.Resource
	ctx := context.Background()
	SubscriptionID := g.Args["config"].(authentication.Config).SubscriptionID
	Authorizer := g.Args["authorizer"].(autorest.Authorizer)

	Client := mariadb.NewVirtualNetworkRulesClient(SubscriptionID)
	Client.Authorizer = Authorizer

	for _, server := range servers {
		id, err := ParseAzureResourceID(*server.ID)
		if err != nil {
			return nil, err
		}
		iter, err := Client.ListByServerComplete(ctx, id.ResourceGroup, *server.Name)
		if err != nil {
			return nil, err
		}
		for iter.NotDone() {
			rule := iter.Value()
			resources = append(resources, terraformutils.NewSimpleResource(
				*rule.ID,
				*rule.Name,
				"azurerm_mariadb_virtual_network_rule",
				g.ProviderName,
				[]string{}))

			if err := iter.NextWithContext(ctx); err != nil {
				return nil, err
			}
		}
	}
	return resources, nil
}

func (g *DatabasesGenerator) getMySQLServers() ([]mysql.Server, error) {
	ctx := context.Background()
	SubscriptionID := g.Args["config"].(authentication.Config).SubscriptionID
	Authorizer := g.Args["authorizer"].(autorest.Authorizer)

	Client := mysql.NewServersClient(SubscriptionID)
	Client.Authorizer = Authorizer

	var (
		Servers mysql.ServerListResult
		err     error
	)

	if rg := g.Args["resource_group"].(string); rg != "" {
		Servers, err = Client.ListByResourceGroup(ctx, rg)
	} else {
		Servers, err = Client.List(ctx)
	}
	if err != nil {
		return nil, err
	}

	return *Servers.Value, nil
}

func (g *DatabasesGenerator) createMySQLServerResources(servers []mysql.Server) ([]terraformutils.Resource, error) {
	var resources []terraformutils.Resource

	for _, server := range servers {
		resources = append(resources, terraformutils.NewResource(
			*server.ID,
			*server.Name,
			"azurerm_mysql_server",
			g.ProviderName,
			map[string]string{},
			[]string{},
			map[string]interface{}{
				"administrator_login_password": "",
			}))
	}

	return resources, nil
}

func (g *DatabasesGenerator) createMySQLConfigurationResources(servers []mysql.Server) ([]terraformutils.Resource, error) {
	var resources []terraformutils.Resource
	ctx := context.Background()
	SubscriptionID := g.Args["config"].(authentication.Config).SubscriptionID
	Authorizer := g.Args["authorizer"].(autorest.Authorizer)

	Client := mysql.NewConfigurationsClient(SubscriptionID)
	Client.Authorizer = Authorizer

	for _, server := range servers {
		id, err := ParseAzureResourceID(*server.ID)
		if err != nil {
			return nil, err
		}

		configs, err := Client.ListByServer(ctx, id.ResourceGroup, *server.Name)
		if err != nil {
			return nil, err
		}
		for _, config := range *configs.Value {
			resources = append(resources, terraformutils.NewSimpleResource(
				*config.ID,
				*config.Name+"-"+*server.Name,
				"azurerm_mysql_configuration",
				g.ProviderName,
				[]string{"value"}))
		}
	}

	return resources, nil
}

func (g *DatabasesGenerator) createMySQLDatabaseResources(servers []mysql.Server) ([]terraformutils.Resource, error) {
	var resources []terraformutils.Resource
	ctx := context.Background()
	SubscriptionID := g.Args["config"].(authentication.Config).SubscriptionID
	Authorizer := g.Args["authorizer"].(autorest.Authorizer)

	Client := mysql.NewDatabasesClient(SubscriptionID)
	Client.Authorizer = Authorizer

	for _, server := range servers {
		id, err := ParseAzureResourceID(*server.ID)
		if err != nil {
			return nil, err
		}
		databases, err := Client.ListByServer(ctx, id.ResourceGroup, *server.Name)
		if err != nil {
			return nil, err
		}

		for _, database := range *databases.Value {
			resources = append(resources, terraformutils.NewSimpleResource(
				*database.ID,
				*database.Name+"-"+*server.Name,
				"azurerm_mysql_database",
				g.ProviderName,
				[]string{}))
		}
	}
	return resources, nil
}

func (g *DatabasesGenerator) createMySQLFirewallRuleResources(servers []mysql.Server) ([]terraformutils.Resource, error) {
	var resources []terraformutils.Resource
	ctx := context.Background()
	SubscriptionID := g.Args["config"].(authentication.Config).SubscriptionID
	Authorizer := g.Args["authorizer"].(autorest.Authorizer)

	Client := mysql.NewFirewallRulesClient(SubscriptionID)
	Client.Authorizer = Authorizer

	for _, server := range servers {
		id, err := ParseAzureResourceID(*server.ID)
		if err != nil {
			return nil, err
		}
		rules, err := Client.ListByServer(ctx, id.ResourceGroup, *server.Name)
		if err != nil {
			return nil, err
		}

		for _, rule := range *rules.Value {
			resources = append(resources, terraformutils.NewSimpleResource(
				*rule.ID,
				*rule.Name,
				"azurerm_mysql_firewall_rule",
				g.ProviderName,
				[]string{}))
		}
	}

	return resources, nil
}

func (g *DatabasesGenerator) createMySQLVirtualNetworkRuleResources(servers []mysql.Server) ([]terraformutils.Resource, error) {
	var resources []terraformutils.Resource
	ctx := context.Background()
	SubscriptionID := g.Args["config"].(authentication.Config).SubscriptionID
	Authorizer := g.Args["authorizer"].(autorest.Authorizer)

	Client := mysql.NewVirtualNetworkRulesClient(SubscriptionID)
	Client.Authorizer = Authorizer

	for _, server := range servers {
		id, err := ParseAzureResourceID(*server.ID)
		if err != nil {
			return nil, err
		}

		iter, err := Client.ListByServerComplete(ctx, id.ResourceGroup, *server.Name)
		if err != nil {
			return nil, err
		}

		for iter.NotDone() {
			rule := iter.Value()
			resources = append(resources, terraformutils.NewSimpleResource(
				*rule.ID,
				*rule.Name,
				"azurerm_mysql_virtual_network_rule",
				g.ProviderName,
				[]string{}))

			if err := iter.NextWithContext(ctx); err != nil {
				return nil, err
			}
		}
	}

	return resources, nil
}

func (g *DatabasesGenerator) getPostgreSQLServers() ([]postgresql.Server, error) {
	ctx := context.Background()
	SubscriptionID := g.Args["config"].(authentication.Config).SubscriptionID
	Authorizer := g.Args["authorizer"].(autorest.Authorizer)

	Client := postgresql.NewServersClient(SubscriptionID)
	Client.Authorizer = Authorizer

	var (
		Servers postgresql.ServerListResult
		err     error
	)

	if rg := g.Args["resource_group"].(string); rg != "" {
		Servers, err = Client.ListByResourceGroup(ctx, rg)
	} else {
		Servers, err = Client.List(ctx)
	}

	if err != nil {
		return nil, err
	}

	return *Servers.Value, nil
}

func (g *DatabasesGenerator) createPostgreSQLServerResources(servers []postgresql.Server) ([]terraformutils.Resource, error) {
	var resources []terraformutils.Resource

	for _, server := range servers {
		resources = append(resources, terraformutils.NewResource(
			*server.ID,
			*server.Name,
			"azurerm_postgresql_server",
			g.ProviderName,
			map[string]string{},
			[]string{},
			map[string]interface{}{
				"administrator_login_password": "",
			}))
	}

	return resources, nil
}

func (g *DatabasesGenerator) createPostgreSQLDatabaseResources(servers []postgresql.Server) ([]terraformutils.Resource, error) {
	var resources []terraformutils.Resource
	ctx := context.Background()
	SubscriptionID := g.Args["config"].(authentication.Config).SubscriptionID
	Authorizer := g.Args["authorizer"].(autorest.Authorizer)

	Client := postgresql.NewDatabasesClient(SubscriptionID)
	Client.Authorizer = Authorizer

	for _, server := range servers {
		id, err := ParseAzureResourceID(*server.ID)
		if err != nil {
			return nil, err
		}
		databases, err := Client.ListByServer(ctx, id.ResourceGroup, *server.Name)
		if err != nil {
			return nil, err
		}

		for _, database := range *databases.Value {
			resources = append(resources, terraformutils.NewSimpleResource(
				*database.ID,
				*database.Name+"-"+*server.Name,
				"azurerm_postgresql_database",
				g.ProviderName,
				[]string{}))
		}
	}
	return resources, nil
}

func (g *DatabasesGenerator) createPostgreSQLConfigurationResources(servers []postgresql.Server) ([]terraformutils.Resource, error) {
	var resources []terraformutils.Resource
	ctx := context.Background()
	SubscriptionID := g.Args["config"].(authentication.Config).SubscriptionID
	Authorizer := g.Args["authorizer"].(autorest.Authorizer)
	Client := postgresql.NewConfigurationsClient(SubscriptionID)
	Client.Authorizer = Authorizer

	for _, server := range servers {
		id, err := ParseAzureResourceID(*server.ID)
		if err != nil {
			return nil, err
		}
		configs, err := Client.ListByServer(ctx, id.ResourceGroup, *server.Name)
		if err != nil {
			return nil, err
		}

		for _, config := range *configs.Value {
			resources = append(resources, terraformutils.NewSimpleResource(
				*config.ID,
				*config.Name+"-"+*server.Name,
				"azurerm_postgresql_configuration",
				g.ProviderName,
				[]string{"value"}))
		}
	}
	return resources, nil
}

func (g *DatabasesGenerator) createPostgreSQLFirewallRuleResources(servers []postgresql.Server) ([]terraformutils.Resource, error) {
	var resources []terraformutils.Resource
	ctx := context.Background()
	SubscriptionID := g.Args["config"].(authentication.Config).SubscriptionID
	Authorizer := g.Args["authorizer"].(autorest.Authorizer)

	Client := postgresql.NewFirewallRulesClient(SubscriptionID)
	Client.Authorizer = Authorizer

	for _, server := range servers {
		id, err := ParseAzureResourceID(*server.ID)
		if err != nil {
			return nil, err
		}
		rules, err := Client.ListByServer(ctx, id.ResourceGroup, *server.Name)
		if err != nil {
			return nil, err
		}

		for _, rule := range *rules.Value {
			resources = append(resources, terraformutils.NewSimpleResource(
				*rule.ID,
				*rule.Name,
				"azurerm_postgresql_firewall_rule",
				g.ProviderName,
				[]string{}))
		}
	}
	return resources, nil
}

func (g *DatabasesGenerator) createPostgreSQLVirtualNetworkRuleResources(servers []postgresql.Server) ([]terraformutils.Resource, error) {
	var resources []terraformutils.Resource
	ctx := context.Background()
	SubscriptionID := g.Args["config"].(authentication.Config).SubscriptionID
	Authorizer := g.Args["authorizer"].(autorest.Authorizer)

	Client := postgresql.NewVirtualNetworkRulesClient(SubscriptionID)
	Client.Authorizer = Authorizer

	for _, server := range servers {
		id, err := ParseAzureResourceID(*server.ID)
		if err != nil {
			return nil, err
		}
		rulePages, err := Client.ListByServerComplete(ctx, id.ResourceGroup, *server.Name)
		if err != nil {
			return nil, err
		}

		for rulePages.NotDone() {
			rule := rulePages.Value()
			resources = append(resources, terraformutils.NewSimpleResource(
				*rule.ID,
				*rule.Name,
				"azurerm_postgresql_virtual_network_rule",
				g.ProviderName,
				[]string{}))
		}
	}
	return resources, nil
}

func (g *DatabasesGenerator) getSQLServers() ([]sql.Server, error) {
	var servers []sql.Server
	ctx := context.Background()
	SubscriptionID := g.Args["config"].(authentication.Config).SubscriptionID
	Authorizer := g.Args["authorizer"].(autorest.Authorizer)

	Client := sql.NewServersClient(SubscriptionID)
	Client.Authorizer = Authorizer

	var (
		ServerPages sql.ServerListResultPage
		err         error
	)

	if rg := g.Args["resource_group"].(string); rg != "" {
		ServerPages, err = Client.ListByResourceGroup(ctx, rg)
	} else {
		ServerPages, err = Client.List(ctx)
	}
	if err != nil {
		return nil, err
	}
	for ServerPages.NotDone() {
		servers = append(servers, ServerPages.Values()...)
		if err := ServerPages.NextWithContext(ctx); err != nil {
			return nil, err
		}
	}

	return servers, nil
}

func (g *DatabasesGenerator) createSQLServerResources(servers []sql.Server) ([]terraformutils.Resource, error) {
	var resources []terraformutils.Resource

	for _, server := range servers {
		resources = append(resources, terraformutils.NewResource(
			*server.ID,
			*server.Name,
			"azurerm_sql_server",
			g.ProviderName,
			map[string]string{},
			[]string{},
			map[string]interface{}{
				"administrator_login_password": "",
			}))
	}

	return resources, nil
}

func (g *DatabasesGenerator) createSQLDatabaseResources(servers []sql.Server) ([]terraformutils.Resource, error) {
	var resources []terraformutils.Resource
	ctx := context.Background()
	SubscriptionID := g.Args["config"].(authentication.Config).SubscriptionID
	Authorizer := g.Args["authorizer"].(autorest.Authorizer)

	Client := sql.NewDatabasesClient(SubscriptionID)
	Client.Authorizer = Authorizer

	for _, server := range servers {
		id, err := ParseAzureResourceID(*server.ID)
		if err != nil {
			return nil, err
		}
		databases, err := Client.ListByServer(ctx, id.ResourceGroup, *server.Name, "", "")
		if err != nil {
			return nil, err
		}

		for _, database := range *databases.Value {
			resources = append(resources, terraformutils.NewSimpleResource(
				*database.ID,
				*database.Name+"-"+*server.Name,
				"azurerm_sql_database",
				g.ProviderName,
				[]string{}))
		}
	}
	return resources, nil
}

func (g *DatabasesGenerator) createSQLFirewallRuleResources(servers []sql.Server) ([]terraformutils.Resource, error) {
	var resources []terraformutils.Resource
	ctx := context.Background()
	SubscriptionID := g.Args["config"].(authentication.Config).SubscriptionID
	Authorizer := g.Args["authorizer"].(autorest.Authorizer)

	Client := sql.NewFirewallRulesClient(SubscriptionID)
	Client.Authorizer = Authorizer

	for _, server := range servers {
		id, err := ParseAzureResourceID(*server.ID)
		if err != nil {
			return nil, err
		}
		rules, err := Client.ListByServer(ctx, id.ResourceGroup, *server.Name)
		if err != nil {
			return nil, err
		}

		for _, rule := range *rules.Value {
			resources = append(resources, terraformutils.NewSimpleResource(
				*rule.ID,
				*rule.Name,
				"azurerm_sql_firewall_rule",
				g.ProviderName,
				[]string{}))
		}
	}
	return resources, nil
}

func (g *DatabasesGenerator) createSQLVirtualNetworkRuleResources(servers []sql.Server) ([]terraformutils.Resource, error) {
	var resources []terraformutils.Resource
	ctx := context.Background()
	SubscriptionID := g.Args["config"].(authentication.Config).SubscriptionID
	Authorizer := g.Args["authorizer"].(autorest.Authorizer)

	Client := sql.NewVirtualNetworkRulesClient(SubscriptionID)
	Client.Authorizer = Authorizer

	for _, server := range servers {
		id, err := ParseAzureResourceID(*server.ID)
		if err != nil {
			return nil, err
		}
		ruleIter, err := Client.ListByServerComplete(ctx, id.ResourceGroup, *server.Name)
		if err != nil {
			return nil, err
		}

		for ruleIter.NotDone() {
			rule := ruleIter.Value()
			resources = append(resources, terraformutils.NewSimpleResource(
				*rule.ID,
				*rule.Name,
				"azurerm_sql_virtual_network_rule",
				g.ProviderName,
				[]string{}))

			if err := ruleIter.NextWithContext(ctx); err != nil {
				return nil, err
			}
		}
	}
	return resources, nil
}

func (g *DatabasesGenerator) createSQLElasticPoolResources(servers []sql.Server) ([]terraformutils.Resource, error) {
	var resources []terraformutils.Resource
	ctx := context.Background()
	SubscriptionID := g.Args["config"].(authentication.Config).SubscriptionID
	Authorizer := g.Args["authorizer"].(autorest.Authorizer)

	Client := sql.NewElasticPoolsClient(SubscriptionID)
	Client.Authorizer = Authorizer

	for _, server := range servers {
		id, err := ParseAzureResourceID(*server.ID)
		if err != nil {
			return nil, err
		}
		pools, err := Client.ListByServer(ctx, id.ResourceGroup, *server.Name)
		if err != nil {
			return nil, err
		}

		for _, pool := range *pools.Value {
			resources = append(resources, terraformutils.NewSimpleResource(
				*pool.ID,
				*pool.Name,
				"azurerm_sql_elasticpool",
				g.ProviderName,
				[]string{}))
		}
	}
	return resources, nil
}

func (g *DatabasesGenerator) createSQLFailoverResources(servers []sql.Server) ([]terraformutils.Resource, error) {
	var resources []terraformutils.Resource
	ctx := context.Background()
	SubscriptionID := g.Args["config"].(authentication.Config).SubscriptionID
	Authorizer := g.Args["authorizer"].(autorest.Authorizer)

	Client := sql.NewFailoverGroupsClient(SubscriptionID)
	Client.Authorizer = Authorizer

	for _, server := range servers {
		id, err := ParseAzureResourceID(*server.ID)
		if err != nil {
			return nil, err
		}

		iter, err := Client.ListByServerComplete(ctx, id.ResourceGroup, *server.Name)
		if err != nil {
			return nil, err
		}

		for iter.NotDone() {
			failoverGroup := iter.Value()

			resources = append(resources, terraformutils.NewSimpleResource(
				*failoverGroup.ID,
				*failoverGroup.Name,
				"azurerm_sql_failover_group",
				g.ProviderName,
				[]string{}))
		}
	}
	return resources, nil
}

func (g *DatabasesGenerator) createSQLADAdministratorResources(servers []sql.Server) ([]terraformutils.Resource, error) {
	var resources []terraformutils.Resource
	ctx := context.Background()
	SubscriptionID := g.Args["config"].(authentication.Config).SubscriptionID
	Authorizer := g.Args["authorizer"].(autorest.Authorizer)

	Client := sql.NewServerAzureADAdministratorsClient(SubscriptionID)
	Client.Authorizer = Authorizer

	for _, server := range servers {
		id, err := ParseAzureResourceID(*server.ID)
		if err != nil {
			return nil, err
		}

		administrators, err := Client.ListByServer(ctx, id.ResourceGroup, *server.Name)
		if err != nil {
			return nil, err
		}

		for _, administrator := range *administrators.Value {
			resources = append(resources, terraformutils.NewSimpleResource(
				*administrator.ID,
				*administrator.Name,
				"azurerm_sql_active_directory_administrator",
				g.ProviderName,
				[]string{}))
		}
	}
	return resources, nil
}

func (g *DatabasesGenerator) InitResources() error {
	mariadbServers, err := g.getMariaDBServers()
	if err != nil {
		return err
	}

	mysqlServers, err := g.getMySQLServers()
	if err != nil {
		return err
	}

	postgresqlServers, err := g.getPostgreSQLServers()
	if err != nil {
		return err
	}

	sqlServers, err := g.getSQLServers()
	if err != nil {
		return err
	}

	mariadbFunctions := []func([]mariadb.Server) ([]terraformutils.Resource, error){
		g.createMariaDBServerResources,
		g.createMariaDBDatabaseResources,
		g.createMariaDBConfigurationResources,
		g.createMariaDBFirewallRuleResources,
		g.createMariaDBVirtualNetworkRuleResources,
	}

	mysqlFunctions := []func([]mysql.Server) ([]terraformutils.Resource, error){
		g.createMySQLServerResources,
		g.createMySQLDatabaseResources,
		g.createMySQLConfigurationResources,
		g.createMySQLFirewallRuleResources,
		g.createMySQLVirtualNetworkRuleResources,
	}

	postgresqlFunctions := []func([]postgresql.Server) ([]terraformutils.Resource, error){
		g.createPostgreSQLServerResources,
		g.createPostgreSQLDatabaseResources,
		g.createPostgreSQLConfigurationResources,
		g.createPostgreSQLFirewallRuleResources,
		g.createPostgreSQLVirtualNetworkRuleResources,
	}

	sqlFunctions := []func([]sql.Server) ([]terraformutils.Resource, error){
		g.createSQLServerResources,
		g.createSQLDatabaseResources,
		g.createSQLADAdministratorResources,
		g.createSQLElasticPoolResources,
		g.createSQLFailoverResources,
		g.createSQLFirewallRuleResources,
		g.createSQLVirtualNetworkRuleResources,
	}

	for _, f := range mariadbFunctions {
		resources, err := f(mariadbServers)
		if err != nil {
			return err
		}
		g.Resources = append(g.Resources, resources...)
	}

	for _, f := range mysqlFunctions {
		resources, err := f(mysqlServers)
		if err != nil {
			return err
		}
		g.Resources = append(g.Resources, resources...)
	}

	for _, f := range postgresqlFunctions {
		resources, err := f(postgresqlServers)
		if err != nil {
			return err
		}
		g.Resources = append(g.Resources, resources...)
	}

	for _, f := range sqlFunctions {
		resources, err := f(sqlServers)
		if err != nil {
			return err
		}
		g.Resources = append(g.Resources, resources...)
	}

	return nil
}

func (g *DatabasesGenerator) PostConvertHook() error {
	dbEngines := []string{
		"mariadb",
		"mysql",
		"postgresql",
		"sql",
	}

	for _, engineName := range dbEngines {
		for _, resource := range g.Resources {
			dbServerResourceType := fmt.Sprintf("azurerm_%s_server", engineName)
			if resource.InstanceInfo.Type == dbServerResourceType {
				dbName := resource.Item["name"]
				for rIdx, r := range g.Resources {
					if r.InstanceInfo.Type != dbServerResourceType &&
						strings.Contains(r.InstanceInfo.Type, engineName) &&
						r.Item["server_name"] == dbName {
						g.Resources[rIdx].Item["server_name"] = fmt.Sprintf("${%s.%s}", resource.InstanceInfo.Id, "name")
					}
				}
			}
		}
	}

	return nil
}
