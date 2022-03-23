
package main

import (
    "encoding/xml"
    "fmt"
    "github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)


// v_ is appended before every variable so it doesn't give any conflict
// with any keyword in golang. ex - interface is keyword in golang
type xmlSecurityPoliciesPolicyFrom__Zone__Name struct {
	XMLName xml.Name `xml:"configuration"`
	Groups  struct {
		XMLName	xml.Name	`xml:"groups"`
		Name	string	`xml:"name"`
		V_policy  struct {
			XMLName xml.Name `xml:"policy"`
			V_to__zone__name  *string  `xml:"to-zone-name,omitempty"`
			V_from__zone__name  *string  `xml:"from-zone-name,omitempty"`
		} `xml:"security>policies>policy"`
	} `xml:"groups"`
	ApplyGroup string `xml:"apply-groups"`
}

// v_ is appended before every variable so it doesn't give any conflict
// with any keyword in golang. ex- interface is keyword in golang
func junosSecurityPoliciesPolicyFrom__Zone__NameCreate(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_to__zone__name := d.Get("to__zone__name").(string)
	V_from__zone__name := d.Get("from__zone__name").(string)


	config := xmlSecurityPoliciesPolicyFrom__Zone__Name{}
	config.ApplyGroup = id
	config.Groups.Name = id
	config.Groups.V_policy.V_to__zone__name = &V_to__zone__name
	config.Groups.V_policy.V_from__zone__name = &V_from__zone__name

    err = client.SendTransaction("", config, false)
    check(err)
    
    d.SetId(fmt.Sprintf("%s_%s", client.Host, id))
    
	return junosSecurityPoliciesPolicyFrom__Zone__NameRead(d,m)
}

func junosSecurityPoliciesPolicyFrom__Zone__NameRead(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
	config := &xmlSecurityPoliciesPolicyFrom__Zone__Name{}

	err = client.MarshalGroup(id, config)
	check(err)
 	d.Set("to__zone__name", config.Groups.V_policy.V_to__zone__name)
	d.Set("from__zone__name", config.Groups.V_policy.V_from__zone__name)

	return nil
}

func junosSecurityPoliciesPolicyFrom__Zone__NameUpdate(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_to__zone__name := d.Get("to__zone__name").(string)
	V_from__zone__name := d.Get("from__zone__name").(string)


	config := xmlSecurityPoliciesPolicyFrom__Zone__Name{}
	config.ApplyGroup = id
	config.Groups.Name = id
	config.Groups.V_policy.V_to__zone__name = &V_to__zone__name
	config.Groups.V_policy.V_from__zone__name = &V_from__zone__name

    err = client.SendTransaction(id, config, false)
    check(err)
    
	return junosSecurityPoliciesPolicyFrom__Zone__NameRead(d,m)
}

func junosSecurityPoliciesPolicyFrom__Zone__NameDelete(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
    _, err = client.DeleteConfigNoCommit(id)
    check(err)

    d.SetId("")
    
	return nil
}

func junosSecurityPoliciesPolicyFrom__Zone__Name() *schema.Resource {
	return &schema.Resource{
		Create: junosSecurityPoliciesPolicyFrom__Zone__NameCreate,
		Read: junosSecurityPoliciesPolicyFrom__Zone__NameRead,
		Update: junosSecurityPoliciesPolicyFrom__Zone__NameUpdate,
		Delete: junosSecurityPoliciesPolicyFrom__Zone__NameDelete,

        Schema: map[string]*schema.Schema{
            "resource_name": &schema.Schema{
                Type:    schema.TypeString,
                Required: true,
            },
			"to__zone__name": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_policy",
			},
			"from__zone__name": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_policy. Source zone",
			},
		},
	}
}