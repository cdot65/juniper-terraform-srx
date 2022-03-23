
package main

import (
    "encoding/xml"
    "fmt"
    "github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)


// v_ is appended before every variable so it doesn't give any conflict
// with any keyword in golang. ex - interface is keyword in golang
type xmlSecurityNatDestinationRule__SetFromZone struct {
	XMLName xml.Name `xml:"configuration"`
	Groups  struct {
		XMLName	xml.Name	`xml:"groups"`
		Name	string	`xml:"name"`
		V_rule__set  struct {
			XMLName xml.Name `xml:"rule-set"`
			V_name  *string  `xml:"name,omitempty"`
			V_from  struct {
				XMLName xml.Name `xml:"from"`
				V_zone  *string  `xml:"zone,omitempty"`
			} `xml:"from"`
		} `xml:"security>nat>destination>rule-set"`
	} `xml:"groups"`
	ApplyGroup string `xml:"apply-groups"`
}

// v_ is appended before every variable so it doesn't give any conflict
// with any keyword in golang. ex- interface is keyword in golang
func junosSecurityNatDestinationRule__SetFromZoneCreate(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_name := d.Get("name").(string)
	V_zone := d.Get("zone").(string)


	config := xmlSecurityNatDestinationRule__SetFromZone{}
	config.ApplyGroup = id
	config.Groups.Name = id
	config.Groups.V_rule__set.V_name = &V_name
	config.Groups.V_rule__set.V_from.V_zone = &V_zone

    err = client.SendTransaction("", config, false)
    check(err)
    
    d.SetId(fmt.Sprintf("%s_%s", client.Host, id))
    
	return junosSecurityNatDestinationRule__SetFromZoneRead(d,m)
}

func junosSecurityNatDestinationRule__SetFromZoneRead(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
	config := &xmlSecurityNatDestinationRule__SetFromZone{}

	err = client.MarshalGroup(id, config)
	check(err)
 	d.Set("name", config.Groups.V_rule__set.V_name)
	d.Set("zone", config.Groups.V_rule__set.V_from.V_zone)

	return nil
}

func junosSecurityNatDestinationRule__SetFromZoneUpdate(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_name := d.Get("name").(string)
	V_zone := d.Get("zone").(string)


	config := xmlSecurityNatDestinationRule__SetFromZone{}
	config.ApplyGroup = id
	config.Groups.Name = id
	config.Groups.V_rule__set.V_name = &V_name
	config.Groups.V_rule__set.V_from.V_zone = &V_zone

    err = client.SendTransaction(id, config, false)
    check(err)
    
	return junosSecurityNatDestinationRule__SetFromZoneRead(d,m)
}

func junosSecurityNatDestinationRule__SetFromZoneDelete(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
    _, err = client.DeleteConfigNoCommit(id)
    check(err)

    d.SetId("")
    
	return nil
}

func junosSecurityNatDestinationRule__SetFromZone() *schema.Resource {
	return &schema.Resource{
		Create: junosSecurityNatDestinationRule__SetFromZoneCreate,
		Read: junosSecurityNatDestinationRule__SetFromZoneRead,
		Update: junosSecurityNatDestinationRule__SetFromZoneUpdate,
		Delete: junosSecurityNatDestinationRule__SetFromZoneDelete,

        Schema: map[string]*schema.Schema{
            "resource_name": &schema.Schema{
                Type:    schema.TypeString,
                Required: true,
            },
			"name": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_rule__set",
			},
			"zone": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_rule__set.V_from. Source zone list",
			},
		},
	}
}