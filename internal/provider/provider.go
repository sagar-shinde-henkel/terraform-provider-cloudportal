package provider

import (
	"fmt"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"net/http"

	"github.com/terraform-provider-cloudportal/cloudportal/internal/logger"
)

// CloudportalAPIClient represents a custom API client that communicates with the API
type CloudportalAPIClient struct {
	BaseURL     string
	APIKey      string
	Client      *http.Client
	aziclient   *azidentity.ClientSecretCredential
	isdebug     bool
	tenantID    string
	cp_clientID string
}

// NewCloudportalAPIClient initializes a new API client
func NewCloudportalAPIClient(azidentity *azidentity.ClientSecretCredential, apiKey, baseURL string, tenID string, cpID string) *CloudportalAPIClient {
	return &CloudportalAPIClient{
		BaseURL:     baseURL,
		APIKey:      apiKey,
		Client:      &http.Client{},
		aziclient:   azidentity,
		tenantID:    tenID,
		cp_clientID: cpID,
	}
}

// providerConfigure initializes the custom API client
func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	apiKey := d.Get("api_key").(string)
	baseURL := d.Get("base_url").(string)
	debugInfo := d.Get("debug_info").(bool)

	if debugInfo {
		// Create a new logger with debug enabled
		// Initialize the logger once, using debugEnabled=true
		_, err := logger.NewLogger(true)
		if err != nil {
			log.Fatal("Error initializing logger:", err)
		}
		defer logger.Close()
	}
	logger.Info("start")
	if apiKey == "" || baseURL == "" {
		logger.Error("API key and base URL must be provided")
		return nil, fmt.Errorf("API key and base URL must be provided")
	}

	// Define your Azure credentials
	clientID := d.Get("clientid").(string)
	clientSecret := d.Get("clientsecret").(string)
	tenantID := d.Get("tenantid").(string)
	cp_clientID := d.Get("cp_clientid").(string)

	// Use azidentity to authenticate using client credentials
	client, err := azidentity.NewClientSecretCredential(tenantID, clientID, clientSecret, nil)
	if err != nil {
		logger.Error(err.Error())
	}

	apiclient := NewCloudportalAPIClient(client, apiKey, baseURL, tenantID, cp_clientID)

	return apiclient, nil
}

func Provider() *schema.Provider {
	return &schema.Provider{
		// Define the provider schema (inputs from Terraform)
		Schema: map[string]*schema.Schema{
			"api_key": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "API key for authenticating with the custom API",
			},
			"base_url": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Base URL of the custom API",
			},
			"debug_info": {
				Type:        schema.TypeBool,
				Required:    true,
				Description: "Debug infor mation logging",
			},
			"clientid": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "clientID key for authenticating with the custom API",
			},
			"clientsecret": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "clientSecret key for authenticating with the custom API",
			},
			"tenantid": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "tenantID key for authenticating with the custom API",
			},
			"cp_clientid": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "cloud portal client id  key for authenticating with the custom API",
			},
		},
		// Configure the provider with API credentials
		ConfigureFunc: providerConfigure,

		// Define the resources and data sources
		/*ResourcesMap: map[string]*schema.Resource{
			"cloudportal_api_resource": ResourceCustom(),
		},

		// Map resources and data sources
		ResourcesMap: map[string]*schema.Resource{
			"cloudportal_resource": ResourceCustom(),
		},*/
		DataSourcesMap: map[string]*schema.Resource{
			"cloudportal_datasource":                  dataSourceTicket(),          // Add data source here
			"cloudportal_datasource_ticket_inventory": dataSourceTicketInventory(), // Add data source here
			"cloudportal_datasource_ticket_search":    dataSourceTicketSearch(),
		},
	}
}
