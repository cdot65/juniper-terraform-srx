
package main

import (
    "encoding/xml"
    "fmt"
    "github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)


// v_ is appended before every variable so it doesn't give any conflict
// with any keyword in golang. ex - interface is keyword in golang
type xmlSecurityAddress__BookDescription struct {
	XMLName xml.Name `xml:"configuration"`
	Groups  struct {
		XMLName	xml.Name	`xml:"groups"`
		Name	string	`xml:"name"`
		V_address__book  struct {
			XMLName xml.Name `xml:"address-book"`
			V_name  *string  `xml:"name,omitempty"`
			V_description  *string  `xml:"description,omitempty"`
		} `xml:"security>address-book"`
	} `xml:"groups"`
	ApplyGroup string `xml:"apply-groups"`
}

// v_ is appended before every variable so it doesn't give any conflict
// with any keyword in golang. ex- interface is keyword in golang
func junosSecurityAddress__BookDescriptionCreate(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_name := d.Get("name").(string)
	V_description := d.Get("description").(string)


	config := xmlSecurityAddress__BookDescription{}
	config.ApplyGroup = id
	config.Groups.Name = id
	config.Groups.V_address__book.V_name = &V_name
	config.Groups.V_address__book.V_description = &V_description

    err = client.SendTransaction("", config, false)
    check(err)
    
    d.SetId(fmt.Sprintf("%s_%s", client.Host, id))
    
	return junosSecurityAddress__BookDescriptionRead(d,m)
}

func junosSecurityAddress__BookDescriptionRead(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
	config := &xmlSecurityAddress__BookDescription{}

	err = client.MarshalGroup(id, config)
	check(err)
 	d.Set("name", config.Groups.V_address__book.V_name)
	d.Set("description", config.Groups.V_address__book.V_description)

	return nil
}

func junosSecurityAddress__BookDescriptionUpdate(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_name := d.Get("name").(string)
	V_description := d.Get("description").(string)


	config := xmlSecurityAddress__BookDescription{}
	config.ApplyGroup = id
	config.Groups.Name = id
	config.Groups.V_address__book.V_name = &V_name
	config.Groups.V_address__book.V_description = &V_description

    err = client.SendTransaction(id, config, false)
    check(err)
    
	return junosSecurityAddress__BookDescriptionRead(d,m)
}

func junosSecurityAddress__BookDescriptionDelete(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
    _, err = client.DeleteConfigNoCommit(id)
    check(err)

    d.SetId("")
    
	return nil
}

func junosSecurityAddress__BookDescription() *schema.Resource {
	return &schema.Resource{
		Create: junosSecurityAddress__BookDescriptionCreate,
		Read: junosSecurityAddress__BookDescriptionRead,
		Update: junosSecurityAddress__BookDescriptionUpdate,
		Delete: junosSecurityAddress__BookDescriptionDelete,

        Schema: map[string]*schema.Schema{
            "resource_name": &schema.Schema{
                Type:    schema.TypeString,
                Required: true,
            },
			"name": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_address__book",
			},
			"description": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_address__book. Text description of address book",
			},
		},
	}
}