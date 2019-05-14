package openstack

import (
	"fmt"
	"strings"
	"time"

	"github.com/gophercloud/gophercloud/openstack/db/v1/users"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceDatabaseUserV1() *schema.Resource {
	return &schema.Resource{
		Create: resourceDatabaseUserV1Create,
		Read:   resourceDatabaseUserV1Read,
		Delete: resourceDatabaseUserV1Delete,

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

			"password": {
				Type:      schema.TypeString,
				Required:  true,
				ForceNew:  true,
				Sensitive: true,
			},

			"host": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},

			"databases": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Set:      schema.HashString,
			},
		},
	}
}

func resourceDatabaseUserV1Create(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	databaseV1Client, err := config.databaseV1Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating OpenStack database client: %s", err)
	}

	userName := d.Get("name").(string)
	rawDatabases := d.Get("databases").(*schema.Set).List()
	instanceID := d.Get("instance_id").(string)

	var usersList users.BatchCreateOpts
	usersList = append(usersList, users.CreateOpts{
		Name:      userName,
		Password:  d.Get("password").(string),
		Host:      d.Get("host").(string),
		Databases: expandDatabaseUserV1Databases(rawDatabases),
	})

	err = users.Create(databaseV1Client, instanceID, usersList).ExtractErr()
	if err != nil {
		return fmt.Errorf("Error creating openstack_db_user_v1: %s", err)
	}

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"BUILD"},
		Target:     []string{"ACTIVE"},
		Refresh:    databaseUserV1StateRefreshFunc(databaseV1Client, instanceID, userName),
		Timeout:    d.Timeout(schema.TimeoutCreate),
		Delay:      10 * time.Second,
		MinTimeout: 3 * time.Second,
	}

	_, err = stateConf.WaitForState()
	if err != nil {
		return fmt.Errorf("Error waiting for openstack_db_user_v1 %s to be created: %s", userName, err)
	}

	// Store the ID now
	d.SetId(fmt.Sprintf("%s/%s", instanceID, userName))

	return resourceDatabaseUserV1Read(d, meta)
}

func resourceDatabaseUserV1Read(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	databaseV1Client, err := config.databaseV1Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating OpenStack database client: %s", err)
	}

	userID := strings.SplitN(d.Id(), "/", 2)
	if len(userID) != 2 {
		return fmt.Errorf("Invalid openstack_db_user_v1 ID: %s", d.Id())
	}

	instanceID := userID[0]
	userName := userID[1]

	exists, userObj, err := databaseUserV1Exists(databaseV1Client, instanceID, userName)
	if err != nil {
		return fmt.Errorf("Error checking if openstack_db_user_v1 %s exists: %s", d.Id(), err)
	}

	if !exists {
		d.SetId("")
		return nil
	}

	d.Set("name", userName)

	databases := flattenDatabaseUserV1Databases(userObj.Databases)
	if err := d.Set("databases", databases); err != nil {
		return fmt.Errorf("Unable to set databases: %s", err)
	}

	return nil
}

func resourceDatabaseUserV1Delete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	databaseV1Client, err := config.databaseV1Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating OpenStack database client: %s", err)
	}

	userID := strings.SplitN(d.Id(), "/", 2)
	if len(userID) != 2 {
		return fmt.Errorf("Invalid openstack_db_user_v1 ID: %s", d.Id())
	}

	instanceID := userID[0]
	userName := userID[1]

	exists, _, err := databaseUserV1Exists(databaseV1Client, instanceID, userName)
	if err != nil {
		return fmt.Errorf("Error checking if openstack_db_user_v1 %s exists: %s", d.Id(), err)
	}

	if !exists {
		return nil
	}

	err = users.Delete(databaseV1Client, instanceID, userName).ExtractErr()
	if err != nil {
		return fmt.Errorf("Error deleting openstack_db_user_v1 %s: %s", d.Id(), err)
	}

	return nil
}
