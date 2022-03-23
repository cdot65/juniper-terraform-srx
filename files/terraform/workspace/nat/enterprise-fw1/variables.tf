
variable "nat_apply_group_name" {
  type        = string
  description = "Name of our apply-group for the security -> nat configuration stanza"
  default     = "awx1_nat"
}

variable "nat_pool_name" {
  type        = string
  description = "Name of our nat pool"
  default     = "awx1_pool"
}

variable "nat_ruleset_name" {
  type        = string
  description = "Name of our NAT rule set"
  default     = "awx1_ruleset"
}

variable "nat_ruleset_rule_name" {
  type        = string
  description = "Name of our NAT rule"
  default     = "awx1"
}
