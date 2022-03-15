terraform {
  required_providers {
    junos-vsrx = {
      source  = "cdot65/nops/junos-vsrx"
      version = ">= 21.3.0"
    }
  }
}

resource "junos-vsrx_SecurityAddress__BookAddressIp__Prefix" "cdot65_address_1_prefix" {
    resource_name = "cdot65"
    name          = "cdot65-book"
    name__1       = "server_lan1"
    ip__prefix    = "10.0.1.0/24"
}

resource "junos-vsrx_SecurityAddress__BookAddressName" "cdot65_address_1_name" {
    resource_name = "cdot65"
    name          = "cdot65-book"
    name__1       = "server_lan1"
}

resource "junos-vsrx_SecurityAddress__BookDescription" "cdot65_book_description" {
    resource_name = "cdot65"
    name          = "cdot65-book"
    description   = "This is an address book"
}

resource "junos-vsrx_SecurityAddress__BookName" "cdot65_book" {
    resource_name = "cdot65"
    name          = "cdot65-book"
}
