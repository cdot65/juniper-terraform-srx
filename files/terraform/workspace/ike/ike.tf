terraform {
  required_providers {
    junos-ike = {
      source  = "cdot65/juniper-terraform-srx/junos-ike"
      version = ">= 21.3.0"
    }
  }
}

provider "junos-ike" {
    host     = "192.168.105.196"
    port     = 22
    username = "terraform"
    password = "juniper123"
    sshkey   = ""
}

module "enterprise-fw1" {
   source     = "./enterprise-fw1"
   providers  = { junos-ike = junos-ike }
   depends_on = [junos-ike_destroycommit.commit-main]
}

resource "junos-ike_commit" "commit-main" {
    resource_name = "commit"
    depends_on    = [module.enterprise-fw1]
}

resource "junos-ike_destroycommit" "commit-main" {
    resource_name = "destroycommit"
}