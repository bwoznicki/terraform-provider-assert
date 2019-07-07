package main

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
)

func dataSource() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceRead,

		Schema: map[string]*schema.Schema{
			"test": {
				Type:     schema.TypeBool,
				Required: true,
			},
			"throw": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func dataSourceRead(d *schema.ResourceData, meta interface{}) error {

	test := d.Get("test")
	err := d.Get("throw")
	if err == "" {
		return fmt.Errorf("the \"throw\" value cannot be empty")
	}

	if test == false {
		return fmt.Errorf("%s", err)
	}

	return nil
}
