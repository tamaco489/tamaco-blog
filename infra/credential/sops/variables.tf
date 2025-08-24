variable "env" {
  description = "The environment in which the credential sops will be created"
  type        = string
}

variable "project" {
  description = "The project name"
  type        = string
  default     = "tamaco-blog"
}
