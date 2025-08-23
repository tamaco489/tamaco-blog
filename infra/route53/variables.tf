variable "env" {
  description = "The environment in which the route53 will be created"
  type        = string
}

variable "project" {
  description = "The project name"
  type        = string
  default     = "tamaco-blog"
}

variable "domain" {
  description = "The domain name"
  type        = string
  default     = "example.com"
}

locals {
  fqn = "${var.env}-${var.project}"
}
