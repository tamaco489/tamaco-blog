variable "env" {
  description = "The environment in which the ecr will be created"
  type        = string
  default     = "stg"
}

variable "project" {
  description = "The project name"
  type        = string
  default     = "tamaco-blog"
}

locals {
  fqn         = "${var.env}-${var.project}"
  article_api = "${var.env}-article-api"
}
