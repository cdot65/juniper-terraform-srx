terraform {
  required_providers {
    junos-nat = {
      source  = "cdot65/juniper-terraform-srx/junos-nat"
      version = ">= 21.3.0"
    }
  }
}

provider "junos-nat" {
  host     = var.juniper_host_name
  port     = var.juniper_ssh_port
  sshkey   = var.juniper_ssh_key
  username = var.juniper_user_name
  password = var.juniper_user_password
}

module "enterprise-fw1" {
  source     = "./enterprise-fw1"
  providers  = { junos-nat = junos-nat }
  depends_on = [junos-nat_destroycommit.commit-main]
}

resource "junos-nat_commit" "commit-main" {
  resource_name = "commit"
  depends_on    = [module.enterprise-fw1]
}

resource "junos-nat_destroycommit" "commit-main" {
  resource_name = "destroycommit"
}