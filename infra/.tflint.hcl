plugin "terraform" {
  enabled = true
  preset  = "recommended"
}

plugin "aws" {
  enabled = true
  version = "0.31.0"
  source  = "github.com/terraform-linters/tflint-ruleset-aws"
}

config {
  # Disable some rules that might be too strict for development
  disabled_by_default = false

  # Call module
  call_module_type = "all"

  # Force module
  force = false
}

rule "aws_instance_invalid_type" {
  enabled = true
}

rule "aws_instance_previous_type" {
  enabled = true
}

rule "terraform_deprecated_interpolation" {
  enabled = true
}

rule "terraform_documented_outputs" {
  enabled = true
}

rule "terraform_documented_variables" {
  enabled = true
}

rule "terraform_naming_convention" {
  enabled = true
  format  = "snake_case"
}

rule "terraform_standard_module_structure" {
  enabled = true
}

rule "terraform_typed_variables" {
  enabled = true
}

rule "terraform_unused_declarations" {
  enabled = true
}

rule "terraform_unused_required_providers" {
  enabled = true
}

# AWS specific rules
rule "aws_ecr_repository_image_tag_mutability" {
  enabled = true
}

rule "aws_ecr_repository_encryption_configuration" {
  enabled = false  # We might not always need encryption
}

rule "aws_resource_missing_tags" {
  enabled = true
  tags = [
    "Name",
    "Environment",
    "Project"
  ]
}
