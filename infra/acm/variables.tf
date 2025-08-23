variable "env" {
  description = "The environment in which the acm will be created"
  type        = string
}

variable "project" {
  description = "The project name"
  type        = string
  default     = "tamaco-blog"
}

locals {
  fqn = "${var.env}-${var.project}"
}
