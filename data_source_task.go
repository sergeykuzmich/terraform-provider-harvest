package main

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/sergeykuzmich/harvestapp-sdk"
	"github.com/spf13/cast"
)

func dataSourceTask() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceTaskRead,

		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
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
	api := hrvst.Client(m.(*Config).AccountId, m.(*Config).AccessToken)

	id, idOk := d.GetOk("id")
	name, nameOk := d.GetOk("name")

	if !idOk && !nameOk {
		return fmt.Errorf("Missing required argument")
	}

	if idOk && nameOk {
		return fmt.Errorf("Whether ID or Name should be provided but not both of them")
	}

	var task *hrvst.Task
	var err error

	if idOk {
		id := cast.ToInt(id)
		task, err = api.GetTask(id, hrvst.Defaults())
	} else {
		name := cast.ToString(name)
		task, err = getTaskByName(name, api)
	}

	if err != nil {
		return err
	}

	d.SetId(cast.ToString(task.ID))
	d.Set("name", task.Name)
	d.Set("billable_by_default", task.BillableByDefault)
	d.Set("default_hourly_rate", task.DefaultHourlyRate)
	d.Set("is_default", task.IsDefault)
	d.Set("is_active", task.IsActive)
	d.Set("created_at", cast.ToString(task.CreatedAt))
	d.Set("updated_at", cast.ToString(task.UpdatedAt))

	return nil
}

func getTaskByName(name string, api *hrvst.API) (*hrvst.Task, error) {
	tasks, next, err := api.GetTasks(hrvst.Defaults())

	for {
		if err != nil {
			return nil, err
		}

		for i := 0; i < len(tasks); i++ {
			if tasks[i].Name == name {
				return tasks[i], nil
			}
		}

		if next == nil {
			break
		}

		tasks, next, err = next()
	}

	return nil, fmt.Errorf("Not Found: %s", name)
}
