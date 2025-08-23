provider "aws" {
  default_tags {
    tags = {
      Project = var.project
      Env     = var.env
      Managed = "terraform"
    }
  }
}

terraform {
  required_version = "~> 1.13.0"

  backend "s3" {
    # backend configuration will be provided via -backend-config flag
  }

  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = ">= 5.0"
    }
  }
}
