package main

import (
	"github.com/adlio/harvest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/spf13/cast"
)

func dataSourceTask() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceTaskRead,

		Schema: map[string]*schema.Schema{
			"task_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"billable_by_default": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
			},
			"default_hourly_rate": &schema.Schema{
				Type:     schema.TypeFloat,
				Computed: true,
			},
			"is_default": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
			},
			"is_active": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
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

func dataSourceTaskRead(d *schema.ResourceData, m interface{}) error {
	api := harvest.NewTokenAPI(m.(*Config).AccountId, m.(*Config).AccessToken)

	task_id := cast.ToInt64(d.Get("task_id"))
	task, _ := api.GetTask(task_id, harvest.Defaults())

	d.SetId(cast.ToString(task.ID))
	d.Set("name", task.Name)
	d.Set("billable_by_default", task.BillableByDefault)
	d.Set("default_hourly_rate", task.DefaultHourlyRate)
	d.Set("is_default", task.IsDefault)
	d.Set("is_active", !task.Deactivated)
	d.Set("created_at", cast.ToString(task.CreatedAt))
	d.Set("updated_at", cast.ToString(task.UpdatedAt))

	return nil
}
