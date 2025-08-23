variable "env" {
  description = "The environment in which the ecr will be created"
  type        = string
}

variable "project" {
  description = "The project name"
  type        = string
  default     = "tamaco-blog"
}

locals {
  article_api = "${var.env}-article-api"
}
