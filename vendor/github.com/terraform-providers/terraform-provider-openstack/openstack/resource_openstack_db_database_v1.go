package openstack

import (
	"fmt"
	"strings"
	"time"

	"github.com/gophercloud/gophercloud/openstack/db/v1/databases"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceDatabaseDatabaseV1() *schema.Resource {
	return &schema.Resource{
		Create: resourceDatabaseDatabaseV1Create,
		Read:   resourceDatabaseDatabaseV1Read,
		Delete: resourceDatabaseDatabaseV1Delete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

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

			"instance_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceDatabaseDatabaseV1Create(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	databaseV1Client, err := config.databaseV1Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating OpenStack database client: %s", err)
	}

	dbName := d.Get("name").(string)
	instanceID := d.Get("instance_id").(string)

	var dbs databases.BatchCreateOpts
	dbs = append(dbs, databases.CreateOpts{
		Name: dbName,
	})

	exists, err := databaseDatabaseV1Exists(databaseV1Client, instanceID, dbName)
	if err != nil {
		return fmt.Errorf("Error checking openstack_db_database_v1 %s status on %s: %s", dbName, instanceID, err)
	}

	if exists {
		return fmt.Errorf("openstack_db_database_v1 %s already exists on instance %s", dbName, instanceID)
	}

	err = databases.Create(databaseV1Client, instanceID, dbs).ExtractErr()
	if err != nil {
		return fmt.Errorf("Error creating openstack_db_database_v1 %s on %s: %s", dbName, instanceID, err)
	}

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"BUILD"},
		Target:     []string{"ACTIVE"},
		Refresh:    databaseDatabaseV1StateRefreshFunc(databaseV1Client, instanceID, dbName),
		Timeout:    d.Timeout(schema.TimeoutCreate),
		Delay:      10 * time.Second,
		MinTimeout: 3 * time.Second,
	}

	_, err = stateConf.WaitForState()
	if err != nil {
		return fmt.Errorf("Error waiting for openstack_db_database_v1 %s on %s to become ready: %s", dbName, instanceID, err)
	}

	// Store the ID now
	d.SetId(fmt.Sprintf("%s/%s", instanceID, dbName))

	return resourceDatabaseDatabaseV1Read(d, meta)
}

func resourceDatabaseDatabaseV1Read(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	databaseV1Client, err := config.databaseV1Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating OpenStack database client: %s", err)
	}

	dbID := strings.SplitN(d.Id(), "/", 2)
	if len(dbID) != 2 {
		return fmt.Errorf("Invalid openstack_db_database_v1 ID: %s", d.Id())
	}

	instanceID := dbID[0]
	dbName := dbID[1]

	exists, err := databaseDatabaseV1Exists(databaseV1Client, instanceID, dbName)
	if err != nil {
		return fmt.Errorf("Error checking if openstack_db_database_v1 %s exists: %s", d.Id(), err)
	}

	if !exists {
		d.SetId("")
		return nil
	}

	d.Set("instance_id", instanceID)
	d.Set("name", dbName)

	return nil
}

func resourceDatabaseDatabaseV1Delete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	databaseV1Client, err := config.databaseV1Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating OpenStack database client: %s", err)
	}

	dbID := strings.SplitN(d.Id(), "/", 2)
	if len(dbID) != 2 {
		return fmt.Errorf("Invalid openstack_db_database_v1 ID: %s", d.Id())
	}

	instanceID := dbID[0]
	dbName := dbID[1]

	exists, err := databaseDatabaseV1Exists(databaseV1Client, instanceID, dbName)
	if err != nil {
		return fmt.Errorf("Error checking if openstack_db_database_v1 %s exists: %s", d.Id(), err)
	}

	if !exists {
		return nil
	}

	err = databases.Delete(databaseV1Client, instanceID, dbName).ExtractErr()
	if err != nil {
		return fmt.Errorf("Error deleting openstack_db_database_v1 %s: %s", dbName, err)
	}

	return nil
}
