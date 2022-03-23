
package main

import (
    "encoding/xml"
    "fmt"
    "github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)


// v_ is appended before every variable so it doesn't give any conflict
// with any keyword in golang. ex - interface is keyword in golang
type xmlSecurityNatDestinationPoolDescription struct {
	XMLName xml.Name `xml:"configuration"`
	Groups  struct {
		XMLName	xml.Name	`xml:"groups"`
		Name	string	`xml:"name"`
		V_pool  struct {
			XMLName xml.Name `xml:"pool"`
			V_name  *string  `xml:"name,omitempty"`
			V_description  *string  `xml:"description,omitempty"`
		} `xml:"security>nat>destination>pool"`
	} `xml:"groups"`
	ApplyGroup string `xml:"apply-groups"`
}

// v_ is appended before every variable so it doesn't give any conflict
// with any keyword in golang. ex- interface is keyword in golang
func junosSecurityNatDestinationPoolDescriptionCreate(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_name := d.Get("name").(string)
	V_description := d.Get("description").(string)


	config := xmlSecurityNatDestinationPoolDescription{}
	config.ApplyGroup = id
	config.Groups.Name = id
	config.Groups.V_pool.V_name = &V_name
	config.Groups.V_pool.V_description = &V_description

    err = client.SendTransaction("", config, false)
    check(err)
    
    d.SetId(fmt.Sprintf("%s_%s", client.Host, id))
    
	return junosSecurityNatDestinationPoolDescriptionRead(d,m)
}

func junosSecurityNatDestinationPoolDescriptionRead(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
	config := &xmlSecurityNatDestinationPoolDescription{}

	err = client.MarshalGroup(id, config)
	check(err)
 	d.Set("name", config.Groups.V_pool.V_name)
	d.Set("description", config.Groups.V_pool.V_description)

	return nil
}

func junosSecurityNatDestinationPoolDescriptionUpdate(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_name := d.Get("name").(string)
	V_description := d.Get("description").(string)


	config := xmlSecurityNatDestinationPoolDescription{}
	config.ApplyGroup = id
	config.Groups.Name = id
	config.Groups.V_pool.V_name = &V_name
	config.Groups.V_pool.V_description = &V_description

    err = client.SendTransaction(id, config, false)
    check(err)
    
	return junosSecurityNatDestinationPoolDescriptionRead(d,m)
}

func junosSecurityNatDestinationPoolDescriptionDelete(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
    _, err = client.DeleteConfigNoCommit(id)
    check(err)

    d.SetId("")
    
	return nil
}

func junosSecurityNatDestinationPoolDescription() *schema.Resource {
	return &schema.Resource{
		Create: junosSecurityNatDestinationPoolDescriptionCreate,
		Read: junosSecurityNatDestinationPoolDescriptionRead,
		Update: junosSecurityNatDestinationPoolDescriptionUpdate,
		Delete: junosSecurityNatDestinationPoolDescriptionDelete,

        Schema: map[string]*schema.Schema{
            "resource_name": &schema.Schema{
                Type:    schema.TypeString,
                Required: true,
            },
			"name": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_pool",
			},
			"description": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_pool. Text description of pool",
			},
		},
	}
}