
variable "ike_apply_group_name" {
  type        = string
  description = "Name of our apply-group for the security -> IKE configuration stanza"
  default     = "cdot65_ike"
}

variable "ike_proposal_name" {
  type        = string
  description = "Name of our IKE proposal"
  default     = "cdot65_ike_proposal"
}

variable "ike_policy_name" {
  type        = string
  description = "Name of our IKE policy"
  default     = "cdot65_ike_policy"
}

variable "ike_gateway_name" {
  type        = string
  description = "Name of our IKE gateway"
  default     = "cdot65_ike_gateway"
}
