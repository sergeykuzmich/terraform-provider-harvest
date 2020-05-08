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

	filter, err := getInputArgument(d)
	if err != nil {
		return err
	}

	var task *hrvst.Task

	switch filter.Type {
	case "ID":
		task, err = api.GetTask(cast.ToInt(filter.Value), hrvst.Defaults())
	case "Name":
		task, err = getTaskByName(cast.ToString(filter.Value), api)
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

type argumment struct {
	Value interface{}
	Type  string
}

func getInputArgument(d *schema.ResourceData) (argumment, error) {
	id, idOk := d.GetOk("id")
	name, nameOk := d.GetOk("name")

	output := argumment{}

	if !idOk && !nameOk {
		return output, fmt.Errorf("Missing required argument")
	}

	if idOk && nameOk {
		return output, fmt.Errorf("Whether ID or Name should be provided but not both of them")
	}

	if idOk {
		output.Value = id
		output.Type = "ID"
	} else {
		output.Value = name
		output.Type = "Name"
	}

	return output, nil
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
