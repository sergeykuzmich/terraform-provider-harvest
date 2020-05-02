package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/sergeykuzmich/harvestapp-sdk"
	"github.com/spf13/cast"
)

func resourceTask() *schema.Resource {
	return &schema.Resource{
		Create: resourceTaskCreate,
		Read:   resourceTaskRead,
		Update: resourceTaskUpdate,
		Delete: resourceTaskDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"billable_by_default": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"default_hourly_rate": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
				Default:  0,
			},
			"is_default": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"is_active": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
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
	api := hrvst.Client(m.(*Config).AccountId, m.(*Config).AccessToken)

	task_data := hrvst.Task{
		Name:              d.Get("name").(string),
		BillableByDefault: d.Get("billable_by_default").(bool),
		IsActive:          d.Get("is_active").(bool),
		DefaultHourlyRate: d.Get("default_hourly_rate").(float64),
		IsDefault:         d.Get("is_default").(bool),
	}

	task, error := api.CreateTask(&task_data, hrvst.Defaults())
	if error != nil {
		return error
	}

	d.SetId(cast.ToString(task.ID))
	return resourceTaskRead(d, m)
}

func resourceTaskRead(d *schema.ResourceData, m interface{}) error {
	api := hrvst.Client(m.(*Config).AccountId, m.(*Config).AccessToken)
	task, error := api.GetTask(cast.ToInt(d.Id()), hrvst.Defaults())
	if error != nil {
		return error
	}

	d.Set("name", task.Name)
	d.Set("billable_by_default", task.BillableByDefault)
	d.Set("is_active", task.IsActive)
	d.Set("default_hourly_rate", task.DefaultHourlyRate)
	d.Set("is_default", task.IsDefault)
	d.Set("created_at", cast.ToString(task.CreatedAt))
	d.Set("updated_at", cast.ToString(task.UpdatedAt))

	return nil
}

func resourceTaskUpdate(d *schema.ResourceData, m interface{}) error {
	api := hrvst.Client(m.(*Config).AccountId, m.(*Config).AccessToken)

	task_data := hrvst.Task{
		ID:                cast.ToInt(d.Id()),
		Name:              d.Get("name").(string),
		BillableByDefault: d.Get("billable_by_default").(bool),
		IsActive:          d.Get("is_active").(bool),
		DefaultHourlyRate: d.Get("default_hourly_rate").(float64),
		IsDefault:         d.Get("is_default").(bool),
	}

	_, error := api.UpdateTask(&task_data, hrvst.Defaults())
	if error != nil {
		return error
	}

	return resourceTaskRead(d, m)
}

func resourceTaskDelete(d *schema.ResourceData, m interface{}) error {
	api := hrvst.Client(m.(*Config).AccountId, m.(*Config).AccessToken)

	error := api.DeleteTask(cast.ToInt(d.Id()), hrvst.Defaults())
	if error != nil {
		return error
	}

	d.SetId("")
	return nil
}
