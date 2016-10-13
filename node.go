package main

import (
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"net/http"
)

type AgentResponse struct {
	Config struct {
		Client struct {
			NodeID string
		}
	} `json:"config"`
}

func resourceNode() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"node_addr": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},

		Create: resourceNodeCreate,
		Read:   resourceNodeRead,
		Update: resourceNodeUpdate,
		Delete: resourceNodeDelete,
	}
}

func resourceNodeCreate(d *schema.ResourceData, meta interface{}) error {
	url := d.Get("node_addr").(string) + "/v1/agent/self"

	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("could not query agent: %s", err)
	}
	defer resp.Body.Close()

	var response AgentResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return fmt.Errorf("could not parse agent response: %s", err)
	}

	d.SetId(string(response.Config.Client.NodeID))
	return nil
}

func resourceNodeRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceNodeUpdate(d *schema.ResourceData, meta interface{}) error {
	err := resourceNodeDelete(d, meta)
	if err != nil {
		return err
	}

	return resourceNodeCreate(d, meta)
}

func resourceNodeDelete(d *schema.ResourceData, meta interface{}) error {
	d.SetId("")
	return nil
}
