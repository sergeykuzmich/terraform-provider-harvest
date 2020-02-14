package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceTask() *schema.Resource {
	return &schema.Resource{
		Create: resourceTaskCreate,
		Read:   resourceTaskRead,
		Update: resourceTaskUpdate,
		Delete: resourceTaskDelete,

		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"billable_by_default": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"default_hourly_rate": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"is_default": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"is_active": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"created_at": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"updated_at": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceTaskCreate(d *schema.ResourceData, m interface{}) error {
	return resourceTaskRead(d, m)
}

func resourceTaskRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceTaskUpdate(d *schema.ResourceData, m interface{}) error {
	return resourceTaskRead(d, m)
}

func resourceTaskDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
