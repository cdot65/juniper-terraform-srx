
package main

import (
    "encoding/xml"
    "fmt"
    "github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)


// v_ is appended before every variable so it doesn't give any conflict
// with any keyword in golang. ex - interface is keyword in golang
type xmlSecurityFlowTcp__MssIpsec__VpnMss struct {
	XMLName xml.Name `xml:"configuration"`
	Groups  struct {
		XMLName	xml.Name	`xml:"groups"`
		Name	string	`xml:"name"`
		V_ipsec__vpn  struct {
			XMLName xml.Name `xml:"ipsec-vpn"`
			V_mss  *string  `xml:"mss,omitempty"`
		} `xml:"security>flow>tcp-mss>ipsec-vpn"`
	} `xml:"groups"`
	ApplyGroup string `xml:"apply-groups"`
}

// v_ is appended before every variable so it doesn't give any conflict
// with any keyword in golang. ex- interface is keyword in golang
func junosSecurityFlowTcp__MssIpsec__VpnMssCreate(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_mss := d.Get("mss").(string)


	config := xmlSecurityFlowTcp__MssIpsec__VpnMss{}
	config.ApplyGroup = id
	config.Groups.Name = id
	config.Groups.V_ipsec__vpn.V_mss = &V_mss

    err = client.SendTransaction("", config, false)
    check(err)
    
    d.SetId(fmt.Sprintf("%s_%s", client.Host, id))
    
	return junosSecurityFlowTcp__MssIpsec__VpnMssRead(d,m)
}

func junosSecurityFlowTcp__MssIpsec__VpnMssRead(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
	config := &xmlSecurityFlowTcp__MssIpsec__VpnMss{}

	err = client.MarshalGroup(id, config)
	check(err)
 	d.Set("mss", config.Groups.V_ipsec__vpn.V_mss)

	return nil
}

func junosSecurityFlowTcp__MssIpsec__VpnMssUpdate(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_mss := d.Get("mss").(string)


	config := xmlSecurityFlowTcp__MssIpsec__VpnMss{}
	config.ApplyGroup = id
	config.Groups.Name = id
	config.Groups.V_ipsec__vpn.V_mss = &V_mss

    err = client.SendTransaction(id, config, false)
    check(err)
    
	return junosSecurityFlowTcp__MssIpsec__VpnMssRead(d,m)
}

func junosSecurityFlowTcp__MssIpsec__VpnMssDelete(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
    _, err = client.DeleteConfigNoCommit(id)
    check(err)

    d.SetId("")
    
	return nil
}

func junosSecurityFlowTcp__MssIpsec__VpnMss() *schema.Resource {
	return &schema.Resource{
		Create: junosSecurityFlowTcp__MssIpsec__VpnMssCreate,
		Read: junosSecurityFlowTcp__MssIpsec__VpnMssRead,
		Update: junosSecurityFlowTcp__MssIpsec__VpnMssUpdate,
		Delete: junosSecurityFlowTcp__MssIpsec__VpnMssDelete,

        Schema: map[string]*schema.Schema{
            "resource_name": &schema.Schema{
                Type:    schema.TypeString,
                Required: true,
            },
			"mss": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_ipsec__vpn. MSS value",
			},
		},
	}
}