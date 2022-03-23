terraform {
  cloud {
    organization = "cdot65"

    workspaces {
      name = "juniper-terraform-srx"
    }
  }
  required_providers {
    junos-nat = {
      source  = "cdot65/juniper-terraform-srx/junos-nat"
      version = ">= 21.3.0"
    }
  }
}

// NAT pool
resource "junos-nat_SecurityNatDestinationPoolName" "nat_pool_name" {
  resource_name = var.nat_apply_group_name
  name          = var.nat_pool_name
}

resource "junos-nat_SecurityNatDestinationPoolDescription" "nat_pool_description" {
  resource_name = var.nat_apply_group_name
  name          = var.nat_pool_name
  description   = "NAT pool for awx1.dmz.home on port 30071"
}

resource "junos-nat_SecurityNatDestinationPoolAddressIpaddr" "nat_pool_ipaddress" {
  resource_name = var.nat_apply_group_name
  name          = var.nat_pool_name
  ipaddr        = "192.168.203.2/32"
}

resource "junos-nat_SecurityNatDestinationPoolAddressPort" "nat_pool_description" {
  resource_name = var.nat_apply_group_name
  name          = var.nat_pool_name
  port          = "30071"
}

// NAT Rule Set
resource "junos-nat_SecurityNatDestinationRule__SetName" "nat_ruleset_name" {
  resource_name = var.nat_apply_group_name
  name          = var.nat_ruleset_name
}

resource "junos-nat_SecurityNatDestinationRule__SetFromZone" "nat_ruleset_name" {
  resource_name = var.nat_apply_group_name
  name          = var.nat_ruleset_name
  zone          = "internet"
}

resource "junos-nat_SecurityNatDestinationRule__SetRuleDest__Nat__Rule__MatchSource__Address" "nat_ruleset_match_source" {
  resource_name   = var.nat_apply_group_name
  name            = var.nat_ruleset_name
  name__1         = var.nat_ruleset_rule_name
  source__address = "0.0.0.0/0"
}

resource "junos-nat_SecurityNatDestinationRule__SetRuleDest__Nat__Rule__MatchDestination__AddressDst__Addr" "nat_ruleset_match_dstaddr" {
  resource_name = var.nat_apply_group_name
  name          = var.nat_ruleset_name
  name__1       = var.nat_ruleset_rule_name
  dst__addr     = "0.0.0.0/0"
}

resource "junos-nat_SecurityNatDestinationRule__SetRuleDest__Nat__Rule__MatchDestination__PortName" "nat_ruleset_match_dstport" {
  resource_name = var.nat_apply_group_name
  name          = var.nat_ruleset_name
  name__1       = var.nat_ruleset_rule_name
  name__2       = "30071"
}

resource "junos-nat_SecurityNatDestinationRule__SetRuleThenDestination__NatPoolPool__Name" "nat_ruleset_then" {
  resource_name = var.nat_apply_group_name
  name          = var.nat_ruleset_name
  name__1       = var.nat_ruleset_rule_name
  pool__name    = var.nat_pool_name
}
