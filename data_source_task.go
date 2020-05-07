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

	_, okId := d.GetOk("id")
	_, okName := d.GetOk("name")

	if !okId && !okName {
		return fmt.Errorf("define task id or name")
	}

	var task *hrvst.Task

	if okId {
		id, err := cast.ToIntE(d.Get("id"))
		if err != nil {
			return fmt.Errorf("task id should be convertable integer")
		}

		task, err = api.GetTask(id, hrvst.Defaults())

		if err != nil {
			return fmt.Errorf("on getting task - %s", err)
		}
	}

	if okName {
		name, err := cast.ToStringE(d.Get("name"))
		if err != nil {
			return fmt.Errorf("task name should be convertable to string")
		}

		if task == nil {
			task, err = getTaskByName(name, api)
			if err != nil {
				return err
			}
		} else if task.Name != name {
			return fmt.Errorf("task found by id, but has different name - %s", task.Name)
		}
	}

	if task == nil {
		return fmt.Errorf("task not found")
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
			return nil, fmt.Errorf("on task finding by name - %s", err)
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

	return nil, nil
}
