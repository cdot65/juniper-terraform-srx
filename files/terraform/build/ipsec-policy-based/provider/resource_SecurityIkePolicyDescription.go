
package main

import (
    "encoding/xml"
    "fmt"
    "github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)


// v_ is appended before every variable so it doesn't give any conflict
// with any keyword in golang. ex - interface is keyword in golang
type xmlSecurityIkePolicyDescription struct {
	XMLName xml.Name `xml:"configuration"`
	Groups  struct {
		XMLName	xml.Name	`xml:"groups"`
		Name	string	`xml:"name"`
		V_policy  struct {
			XMLName xml.Name `xml:"policy"`
			V_name  *string  `xml:"name,omitempty"`
			V_description  *string  `xml:"description,omitempty"`
		} `xml:"security>ike>policy"`
	} `xml:"groups"`
	ApplyGroup string `xml:"apply-groups"`
}

// v_ is appended before every variable so it doesn't give any conflict
// with any keyword in golang. ex- interface is keyword in golang
func junosSecurityIkePolicyDescriptionCreate(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_name := d.Get("name").(string)
	V_description := d.Get("description").(string)


	config := xmlSecurityIkePolicyDescription{}
	config.ApplyGroup = id
	config.Groups.Name = id
	config.Groups.V_policy.V_name = &V_name
	config.Groups.V_policy.V_description = &V_description

    err = client.SendTransaction("", config, false)
    check(err)
    
    d.SetId(fmt.Sprintf("%s_%s", client.Host, id))
    
	return junosSecurityIkePolicyDescriptionRead(d,m)
}

func junosSecurityIkePolicyDescriptionRead(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
	config := &xmlSecurityIkePolicyDescription{}

	err = client.MarshalGroup(id, config)
	check(err)
 	d.Set("name", config.Groups.V_policy.V_name)
	d.Set("description", config.Groups.V_policy.V_description)

	return nil
}

func junosSecurityIkePolicyDescriptionUpdate(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_name := d.Get("name").(string)
	V_description := d.Get("description").(string)


	config := xmlSecurityIkePolicyDescription{}
	config.ApplyGroup = id
	config.Groups.Name = id
	config.Groups.V_policy.V_name = &V_name
	config.Groups.V_policy.V_description = &V_description

    err = client.SendTransaction(id, config, false)
    check(err)
    
	return junosSecurityIkePolicyDescriptionRead(d,m)
}

func junosSecurityIkePolicyDescriptionDelete(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
    _, err = client.DeleteConfigNoCommit(id)
    check(err)

    d.SetId("")
    
	return nil
}

func junosSecurityIkePolicyDescription() *schema.Resource {
	return &schema.Resource{
		Create: junosSecurityIkePolicyDescriptionCreate,
		Read: junosSecurityIkePolicyDescriptionRead,
		Update: junosSecurityIkePolicyDescriptionUpdate,
		Delete: junosSecurityIkePolicyDescriptionDelete,

        Schema: map[string]*schema.Schema{
            "resource_name": &schema.Schema{
                Type:    schema.TypeString,
                Required: true,
            },
			"name": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_policy",
			},
			"description": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_policy. Text description of IKE policy",
			},
		},
	}
}