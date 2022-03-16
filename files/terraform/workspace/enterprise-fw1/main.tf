terraform {
  required_providers {
    junos-addressbook = {
      source  = "cdot65/juniper-terraform-srx/junos-addressbook"
      version = ">= 21.3.0"
    }
  }
}

resource "junos-addressbook_SecurityAddress__BookName" "cdot65_book" {
    resource_name = "cdot65"
    name          = "cdot65-book"
}

resource "junos-addressbook_SecurityAddress__BookDescription" "cdot65_book_description" {
    resource_name = "cdot65"
    name          = "cdot65-book"
    description   = "This is an address book"
}

resource "junos-addressbook_SecurityAddress__BookAddressName" "cdot65_address_1_name" {
    resource_name = "cdot65"
    name          = "cdot65-book"
    name__1       = "server_lan1"
}

resource "junos-addressbook_SecurityAddress__BookAddressIp__Prefix" "cdot65_address_1_prefix" {
    resource_name = "cdot65"
    name          = "cdot65-book"
    name__1       = "server_lan1"
    ip__prefix    = "10.0.1.0/24"
}
