package swift

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"

	"github.com/hashicorp/terraform/backend"
	"github.com/hashicorp/terraform/helper/schema"
	tf_openstack "github.com/terraform-providers/terraform-provider-openstack/openstack"
)

// New creates a new backend for Swift remote state.
func New() backend.Backend {
	s := &schema.Backend{
		Schema: map[string]*schema.Schema{
			"auth_url": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("OS_AUTH_URL", ""),
				Description: descriptions["auth_url"],
			},

			"user_id": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("OS_USER_ID", ""),
				Description: descriptions["user_name"],
			},

			"user_name": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("OS_USERNAME", ""),
				Description: descriptions["user_name"],
			},

			"application_credential_id": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("OS_APPLICATION_CREDENTIAL_ID", ""),
				Description: descriptions["application_credential_id"],
			},

			"application_credential_name": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("OS_APPLICATION_CREDENTIAL_NAME", ""),
				Description: descriptions["application_credential_name"],
			},

			"application_credential_secret": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("OS_APPLICATION_CREDENTIAL_SECRET", ""),
				Description: descriptions["application_credential_secret"],
			},

			"tenant_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{
					"OS_TENANT_ID",
					"OS_PROJECT_ID",
				}, ""),
				Description: descriptions["tenant_id"],
			},

			"tenant_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{
					"OS_TENANT_NAME",
					"OS_PROJECT_NAME",
				}, ""),
				Description: descriptions["tenant_name"],
			},

			"password": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("OS_PASSWORD", ""),
				Description: descriptions["password"],
			},

			"token": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{
					"OS_TOKEN",
					"OS_AUTH_TOKEN",
				}, ""),
				Description: descriptions["token"],
			},

			"user_domain_name": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("OS_USER_DOMAIN_NAME", ""),
				Description: descriptions["user_domain_name"],
			},

			"user_domain_id": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("OS_USER_DOMAIN_ID", ""),
				Description: descriptions["user_domain_id"],
			},

			"project_domain_name": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("OS_PROJECT_DOMAIN_NAME", ""),
				Description: descriptions["project_domain_name"],
			},

			"project_domain_id": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("OS_PROJECT_DOMAIN_ID", ""),
				Description: descriptions["project_domain_id"],
			},

			"domain_id": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("OS_DOMAIN_ID", ""),
				Description: descriptions["domain_id"],
			},

			"domain_name": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("OS_DOMAIN_NAME", ""),
				Description: descriptions["domain_name"],
			},

			"default_domain": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("OS_DEFAULT_DOMAIN", "default"),
				Description: descriptions["default_domain"],
			},

			"cloud": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("OS_CLOUD", ""),
				Description: descriptions["cloud"],
			},

			"region_name": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("OS_REGION_NAME", ""),
				Description: descriptions["region_name"],
			},

			"insecure": &schema.Schema{
				Type:        schema.TypeBool,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("OS_INSECURE", nil),
				Description: descriptions["insecure"],
			},

			"endpoint_type": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("OS_ENDPOINT_TYPE", ""),
			},

			"cacert_file": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("OS_CACERT", ""),
				Description: descriptions["cacert_file"],
			},

			"cert": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("OS_CERT", ""),
				Description: descriptions["cert"],
			},

			"key": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("OS_KEY", ""),
				Description: descriptions["key"],
			},

			"path": &schema.Schema{
				Type:          schema.TypeString,
				Optional:      true,
				Description:   descriptions["path"],
				Deprecated:    "Use container instead",
				ConflictsWith: []string{"container"},
			},

			"container": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: descriptions["container"],
			},

			"archive_path": &schema.Schema{
				Type:          schema.TypeString,
				Optional:      true,
				Description:   descriptions["archive_path"],
				Deprecated:    "Use archive_container instead",
				ConflictsWith: []string{"archive_container"},
			},

			"archive_container": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: descriptions["archive_container"],
			},

			"expire_after": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: descriptions["expire_after"],
			},

			"lock": &schema.Schema{
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Lock state access",
				Default:     true,
			},

			"state_name": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: descriptions["state_name"],
				Default:     "tfstate.tf",
			},
		},
	}

	result := &Backend{Backend: s}
	result.Backend.ConfigureFunc = result.configure
	return result
}

var descriptions map[string]string

func init() {
	descriptions = map[string]string{
		"auth_url": "The Identity authentication URL.",

		"user_name": "Username to login with.",

		"user_id": "User ID to login with.",

		"application_credential_id": "Application Credential ID to login with.",

		"application_credential_name": "Application Credential name to login with.",

		"application_credential_secret": "Application Credential secret to login with.",

		"tenant_id": "The ID of the Tenant (Identity v2) or Project (Identity v3)\n" +
			"to login with.",

		"tenant_name": "The name of the Tenant (Identity v2) or Project (Identity v3)\n" +
			"to login with.",

		"password": "Password to login with.",

		"token": "Authentication token to use as an alternative to username/password.",

		"user_domain_name": "The name of the domain where the user resides (Identity v3).",

		"user_domain_id": "The ID of the domain where the user resides (Identity v3).",

		"project_domain_name": "The name of the domain where the project resides (Identity v3).",

		"project_domain_id": "The ID of the domain where the proejct resides (Identity v3).",

		"domain_id": "The ID of the Domain to scope to (Identity v3).",

		"domain_name": "The name of the Domain to scope to (Identity v3).",

		"default_domain": "The name of the Domain ID to scope to if no other domain is specified. Defaults to `default` (Identity v3).",

		"cloud": "An entry in a `clouds.yaml` file to use.",

		"region_name": "The name of the Region to use.",

		"insecure": "Trust self-signed certificates.",

		"cacert_file": "A Custom CA certificate.",

		"endpoint_type": "The catalog endpoint type to use.",

		"cert": "A client certificate to authenticate with.",

		"key": "A client private key to authenticate with.",

		"path": "Swift container path to use.",

		"container": "Swift container to create",

		"archive_path": "Swift container path to archive state to.",

		"archive_container": "Swift container to archive state to.",

		"expire_after": "Archive object expiry duration.",

		"state_name": "Name of state object in container",
	}
}

type Backend struct {
	*schema.Backend

	// Fields below are set from configure
	client           *gophercloud.ServiceClient
	archive          bool
	archiveContainer string
	expireSecs       int
	container        string
	lock             bool
	stateName        string
}

func (b *Backend) configure(ctx context.Context) error {
	if b.client != nil {
		return nil
	}

	// Grab the resource data
	data := schema.FromContextBackendConfig(ctx)
	config := &tf_openstack.Config{
		CACertFile:                  data.Get("cacert_file").(string),
		ClientCertFile:              data.Get("cert").(string),
		ClientKeyFile:               data.Get("key").(string),
		Cloud:                       data.Get("cloud").(string),
		DefaultDomain:               data.Get("default_domain").(string),
		DomainID:                    data.Get("domain_id").(string),
		DomainName:                  data.Get("domain_name").(string),
		EndpointType:                data.Get("endpoint_type").(string),
		IdentityEndpoint:            data.Get("auth_url").(string),
		Password:                    data.Get("password").(string),
		ProjectDomainID:             data.Get("project_domain_id").(string),
		ProjectDomainName:           data.Get("project_domain_name").(string),
		Token:                       data.Get("token").(string),
		TenantID:                    data.Get("tenant_id").(string),
		TenantName:                  data.Get("tenant_name").(string),
		UserDomainID:                data.Get("user_domain_id").(string),
		UserDomainName:              data.Get("user_domain_name").(string),
		Username:                    data.Get("user_name").(string),
		UserID:                      data.Get("user_id").(string),
		ApplicationCredentialID:     data.Get("application_credential_id").(string),
		ApplicationCredentialName:   data.Get("application_credential_name").(string),
		ApplicationCredentialSecret: data.Get("application_credential_secret").(string),
	}

	if v, ok := data.GetOkExists("insecure"); ok {
		insecure := v.(bool)
		config.Insecure = &insecure
	}

	if err := config.LoadAndValidate(); err != nil {
		return err
	}

	// Assign state name
	b.stateName = data.Get("state_name").(string)

	// Assign Container
	b.container = data.Get("container").(string)
	if b.container == "" {
		// Check deprecated field
		b.container = data.Get("path").(string)
	}

	// Store the lock information
	b.lock = data.Get("lock").(bool)

	// Enable object archiving?
	if archiveContainer, ok := data.GetOk("archive_container"); ok {
		log.Printf("[DEBUG] Archive_container set, enabling object versioning")
		b.archive = true
		b.archiveContainer = archiveContainer.(string)
	} else if archivePath, ok := data.GetOk("archive_path"); ok {
		log.Printf("[DEBUG] Archive_path set, enabling object versioning")
		b.archive = true
		b.archiveContainer = archivePath.(string)
	}

	// Enable object expiry?
	if expireRaw, ok := data.GetOk("expire_after"); ok {
		expire := expireRaw.(string)
		log.Printf("[DEBUG] Requested that remote state expires after %s", expire)

		if strings.HasSuffix(expire, "d") {
			log.Printf("[DEBUG] Got a days expire after duration. Converting to hours")
			days, err := strconv.Atoi(expire[:len(expire)-1])
			if err != nil {
				return fmt.Errorf("Error converting expire_after value %s to int: %s", expire, err)
			}

			expire = fmt.Sprintf("%dh", days*24)
			log.Printf("[DEBUG] Expire after %s hours", expire)
		}

		expireDur, err := time.ParseDuration(expire)
		if err != nil {
			log.Printf("[DEBUG] Error parsing duration %s: %s", expire, err)
			return fmt.Errorf("Error parsing expire_after duration '%s': %s", expire, err)
		}
		log.Printf("[DEBUG] Seconds duration = %d", int(expireDur.Seconds()))
		b.expireSecs = int(expireDur.Seconds())
	}

	objClient, err := openstack.NewObjectStorageV1(config.OsClient, gophercloud.EndpointOpts{
		Region: data.Get("region_name").(string),
	})
	if err != nil {
		return err
	}

	b.client = objClient

	return nil
}
