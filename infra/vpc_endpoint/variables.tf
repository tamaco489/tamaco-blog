variable "env" {
  description = "The environment in which the vpce will be created"
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
