
package main

import (
    "encoding/xml"
    "fmt"
    "github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)


// v_ is appended before every variable so it doesn't give any conflict
// with any keyword in golang. ex - interface is keyword in golang
type xmlSecurityNatDestinationRule__SetName struct {
	XMLName xml.Name `xml:"configuration"`
	Groups  struct {
		XMLName	xml.Name	`xml:"groups"`
		Name	string	`xml:"name"`
		V_rule__set  struct {
			XMLName xml.Name `xml:"rule-set"`
			V_name  *string  `xml:"name,omitempty"`
		} `xml:"security>nat>destination>rule-set"`
	} `xml:"groups"`
	ApplyGroup string `xml:"apply-groups"`
}

// v_ is appended before every variable so it doesn't give any conflict
// with any keyword in golang. ex- interface is keyword in golang
func junosSecurityNatDestinationRule__SetNameCreate(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_name := d.Get("name").(string)


	config := xmlSecurityNatDestinationRule__SetName{}
	config.ApplyGroup = id
	config.Groups.Name = id
	config.Groups.V_rule__set.V_name = &V_name

    err = client.SendTransaction("", config, false)
    check(err)
    
    d.SetId(fmt.Sprintf("%s_%s", client.Host, id))
    
	return junosSecurityNatDestinationRule__SetNameRead(d,m)
}

func junosSecurityNatDestinationRule__SetNameRead(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
	config := &xmlSecurityNatDestinationRule__SetName{}

	err = client.MarshalGroup(id, config)
	check(err)
 	d.Set("name", config.Groups.V_rule__set.V_name)

	return nil
}

func junosSecurityNatDestinationRule__SetNameUpdate(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_name := d.Get("name").(string)


	config := xmlSecurityNatDestinationRule__SetName{}
	config.ApplyGroup = id
	config.Groups.Name = id
	config.Groups.V_rule__set.V_name = &V_name

    err = client.SendTransaction(id, config, false)
    check(err)
    
	return junosSecurityNatDestinationRule__SetNameRead(d,m)
}

func junosSecurityNatDestinationRule__SetNameDelete(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
    _, err = client.DeleteConfigNoCommit(id)
    check(err)

    d.SetId("")
    
	return nil
}

func junosSecurityNatDestinationRule__SetName() *schema.Resource {
	return &schema.Resource{
		Create: junosSecurityNatDestinationRule__SetNameCreate,
		Read: junosSecurityNatDestinationRule__SetNameRead,
		Update: junosSecurityNatDestinationRule__SetNameUpdate,
		Delete: junosSecurityNatDestinationRule__SetNameDelete,

        Schema: map[string]*schema.Schema{
            "resource_name": &schema.Schema{
                Type:    schema.TypeString,
                Required: true,
            },
			"name": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_rule__set. Rule-set name",
			},
		},
	}
}