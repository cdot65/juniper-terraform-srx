
package main

import (
    "encoding/xml"
    "fmt"
    "github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)


// v_ is appended before every variable so it doesn't give any conflict
// with any keyword in golang. ex - interface is keyword in golang
type xmlSecurityPoliciesPolicyPolicyThenPermitTunnelIpsec__Vpn struct {
	XMLName xml.Name `xml:"configuration"`
	Groups  struct {
		XMLName	xml.Name	`xml:"groups"`
		Name	string	`xml:"name"`
		V_policy  struct {
			XMLName xml.Name `xml:"policy"`
			V_from__zone__name  *string  `xml:"from-zone-name,omitempty"`
			V_to__zone__name  *string  `xml:"to-zone-name,omitempty"`
			V_policy__1  struct {
				XMLName xml.Name `xml:"policy"`
				V_name  *string  `xml:"name,omitempty"`
				V_tunnel  struct {
					XMLName xml.Name `xml:"tunnel"`
					V_ipsec__vpn  *string  `xml:"ipsec-vpn,omitempty"`
				} `xml:"then>permit>tunnel"`
			} `xml:"policy"`
		} `xml:"security>policies>policy"`
	} `xml:"groups"`
	ApplyGroup string `xml:"apply-groups"`
}

// v_ is appended before every variable so it doesn't give any conflict
// with any keyword in golang. ex- interface is keyword in golang
func junosSecurityPoliciesPolicyPolicyThenPermitTunnelIpsec__VpnCreate(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_from__zone__name := d.Get("from__zone__name").(string)
	V_to__zone__name := d.Get("to__zone__name").(string)
	V_name := d.Get("name").(string)
	V_ipsec__vpn := d.Get("ipsec__vpn").(string)


	config := xmlSecurityPoliciesPolicyPolicyThenPermitTunnelIpsec__Vpn{}
	config.ApplyGroup = id
	config.Groups.Name = id
	config.Groups.V_policy.V_from__zone__name = &V_from__zone__name
	config.Groups.V_policy.V_to__zone__name = &V_to__zone__name
	config.Groups.V_policy.V_policy__1.V_name = &V_name
	config.Groups.V_policy.V_policy__1.V_tunnel.V_ipsec__vpn = &V_ipsec__vpn

    err = client.SendTransaction("", config, false)
    check(err)
    
    d.SetId(fmt.Sprintf("%s_%s", client.Host, id))
    
	return junosSecurityPoliciesPolicyPolicyThenPermitTunnelIpsec__VpnRead(d,m)
}

func junosSecurityPoliciesPolicyPolicyThenPermitTunnelIpsec__VpnRead(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
	config := &xmlSecurityPoliciesPolicyPolicyThenPermitTunnelIpsec__Vpn{}

	err = client.MarshalGroup(id, config)
	check(err)
 	d.Set("from__zone__name", config.Groups.V_policy.V_from__zone__name)
	d.Set("to__zone__name", config.Groups.V_policy.V_to__zone__name)
	d.Set("name", config.Groups.V_policy.V_policy__1.V_name)
	d.Set("ipsec__vpn", config.Groups.V_policy.V_policy__1.V_tunnel.V_ipsec__vpn)

	return nil
}

func junosSecurityPoliciesPolicyPolicyThenPermitTunnelIpsec__VpnUpdate(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_from__zone__name := d.Get("from__zone__name").(string)
	V_to__zone__name := d.Get("to__zone__name").(string)
	V_name := d.Get("name").(string)
	V_ipsec__vpn := d.Get("ipsec__vpn").(string)


	config := xmlSecurityPoliciesPolicyPolicyThenPermitTunnelIpsec__Vpn{}
	config.ApplyGroup = id
	config.Groups.Name = id
	config.Groups.V_policy.V_from__zone__name = &V_from__zone__name
	config.Groups.V_policy.V_to__zone__name = &V_to__zone__name
	config.Groups.V_policy.V_policy__1.V_name = &V_name
	config.Groups.V_policy.V_policy__1.V_tunnel.V_ipsec__vpn = &V_ipsec__vpn

    err = client.SendTransaction(id, config, false)
    check(err)
    
	return junosSecurityPoliciesPolicyPolicyThenPermitTunnelIpsec__VpnRead(d,m)
}

func junosSecurityPoliciesPolicyPolicyThenPermitTunnelIpsec__VpnDelete(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
    _, err = client.DeleteConfigNoCommit(id)
    check(err)

    d.SetId("")
    
	return nil
}

func junosSecurityPoliciesPolicyPolicyThenPermitTunnelIpsec__Vpn() *schema.Resource {
	return &schema.Resource{
		Create: junosSecurityPoliciesPolicyPolicyThenPermitTunnelIpsec__VpnCreate,
		Read: junosSecurityPoliciesPolicyPolicyThenPermitTunnelIpsec__VpnRead,
		Update: junosSecurityPoliciesPolicyPolicyThenPermitTunnelIpsec__VpnUpdate,
		Delete: junosSecurityPoliciesPolicyPolicyThenPermitTunnelIpsec__VpnDelete,

        Schema: map[string]*schema.Schema{
            "resource_name": &schema.Schema{
                Type:    schema.TypeString,
                Required: true,
            },
			"from__zone__name": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_policy",
			},
			"to__zone__name": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_policy",
			},
			"name": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_policy.V_policy__1",
			},
			"ipsec__vpn": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_policy.V_policy__1.V_tunnel. Enable VPN with name",
			},
		},
	}
}