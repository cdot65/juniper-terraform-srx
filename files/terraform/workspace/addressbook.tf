terraform {
  required_providers {
    junos-addressbook = {
      source  = "cdot65/juniper-terraform-srx/junos-addressbook"
      version = ">= 21.3.0"
    }
  }
}

provider "junos-addressbook" {
    host = "192.168.105.196"
    port = 22
    username = "terraform"
    password = "juniper123"
    sshkey = ""
}

module "enterprise-fw1" {
   source = "./enterprise-fw1"
   providers = { junos-addressbook = junos-addressbook }
   depends_on = [junos-addressbook_destroycommit.commit-main]
}

resource "junos-addressbook_commit" "commit-main" {
    resource_name = "commit"
    depends_on = [module.enterprise-fw1]
}

resource "junos-addressbook_destroycommit" "commit-main" {
  resource_name = "destroycommit"
}