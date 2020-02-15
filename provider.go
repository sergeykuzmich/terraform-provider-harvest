package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"access_token": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("HARVEST_ACCESS_TOKEN", nil),
				Description: "Personal Access Token for Harvest API V2",
			},
			"account_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("HARVEST_ACCOUNT_ID", nil),
				Description: "Account ID for Harvest API V2",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"harvest_task": resourceTask(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"harvest_task": dataSourceTask(),
		},
		ConfigureFunc: configureProvider,
	}
}

func configureProvider(d *schema.ResourceData) (interface{}, error) {
	config := &Config{
		AccessToken: d.Get("access_token").(string),
		AccountId:   d.Get("account_id").(string),
	}
	return config, nil
}
