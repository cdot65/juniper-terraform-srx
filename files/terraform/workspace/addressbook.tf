terraform {
  required_providers {
    junos-vsrx = {
      source  = "cdot65/nops/junos-vsrx"
      version = ">= 21.3.0"
    }
  }
}

provider "junos-vsrx" {
    host = "192.168.105.196"
    port = 22
    username = "napalm"
    password = "aura-deceit-conform"
    sshkey = ""
}

module "vsrx_1" {
   source = "./vsrx_1"
   providers = { junos-vsrx = junos-vsrx }
   depends_on = [junos-vsrx_destroycommit.commit-main]
}

resource "junos-vsrx_commit" "commit-main" {
    resource_name = "commit"
    depends_on = [module.vsrx_1]
}

resource "junos-vsrx_destroycommit" "commit-main" {
  resource_name = "destroycommit"
}