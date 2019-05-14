package openstack

import (
	"fmt"
	"log"
	"time"

	"github.com/gophercloud/gophercloud/openstack/db/v1/configurations"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceDatabaseConfigurationV1() *schema.Resource {
	return &schema.Resource{
		Create: resourceDatabaseConfigurationV1Create,
		Read:   resourceDatabaseConfigurationV1Read,
		Delete: resourceDatabaseConfigurationV1Delete,

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"region": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},

			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"description": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"datastore": {
				Type:     schema.TypeList,
				Required: true,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"version": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"type": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
					},
				},
				MaxItems: 1,
			},

			"configuration": {
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"value": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
					},
				},
			},
		},
	}
}

func resourceDatabaseConfigurationV1Create(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	databaseV1Client, err := config.databaseV1Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating OpenStack database client: %s", err)
	}

	createOpts := &configurations.CreateOpts{
		Name:        d.Get("name").(string),
		Description: d.Get("description").(string),
	}

	var datastore configurations.DatastoreOpts
	if v, ok := d.GetOk("datastore"); ok {
		datastore = expandDatabaseConfigurationV1Datastore(v.([]interface{}))
	}
	createOpts.Datastore = &datastore

	values := make(map[string]interface{})
	if v, ok := d.GetOk("configuration"); ok {
		values = expandDatabaseConfigurationV1Values(v.([]interface{}))
	}
	createOpts.Values = values

	log.Printf("[DEBUG] openstack_db_configuration_v1 create options: %#v", createOpts)
	cgroup, err := configurations.Create(databaseV1Client, createOpts).Extract()

	if err != nil {
		return fmt.Errorf("Error creating openstack_db_configuration_v1: %s", err)
	}

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"BUILD"},
		Target:     []string{"ACTIVE"},
		Refresh:    databaseConfigurationV1StateRefreshFunc(databaseV1Client, cgroup.ID),
		Timeout:    d.Timeout(schema.TimeoutCreate),
		Delay:      10 * time.Second,
		MinTimeout: 3 * time.Second,
	}

	_, err = stateConf.WaitForState()
	if err != nil {
		return fmt.Errorf("Error waiting for openstack_db_configuration_v1 %s to become ready: %s", cgroup.ID, err)
	}

	// Store the ID now
	d.SetId(cgroup.ID)

	return resourceDatabaseConfigurationV1Read(d, meta)
}

func resourceDatabaseConfigurationV1Read(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	databaseV1Client, err := config.databaseV1Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating OpenStack database client: %s", err)
	}

	cgroup, err := configurations.Get(databaseV1Client, d.Id()).Extract()
	if err != nil {
		return CheckDeleted(d, err, "Error retrieving openstack_db_configuration_v1")
	}

	log.Printf("[DEBUG] Retrieved openstack_db_configuration_v1 %s: %#v", d.Id(), cgroup)

	d.Set("name", cgroup.Name)
	d.Set("description", cgroup.Description)
	d.Set("region", GetRegion(d, config))

	return nil
}

func resourceDatabaseConfigurationV1Delete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	databaseV1Client, err := config.databaseV1Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating OpenStack database client: %s", err)
	}

	err = configurations.Delete(databaseV1Client, d.Id()).ExtractErr()
	if err != nil {
		return fmt.Errorf("Error deleting openstack_db_configuration_v1 %s: %s", d.Id(), err)
	}

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"ACTIVE", "SHUTOFF"},
		Target:     []string{"DELETED"},
		Refresh:    databaseConfigurationV1StateRefreshFunc(databaseV1Client, d.Id()),
		Timeout:    d.Timeout(schema.TimeoutDelete),
		Delay:      10 * time.Second,
		MinTimeout: 3 * time.Second,
	}

	_, err = stateConf.WaitForState()
	if err != nil {
		return fmt.Errorf("Error waiting for openstack_db_configuration_v1 %s to delete: %s", d.Id(), err)
	}

	return nil
}
