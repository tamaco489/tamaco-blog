ENV      ?= dev
VAR_FILE := ./tfvars/${ENV}.tfvars
VAR_OPTS := -var-file "$(VAR_FILE)"

# default configuration
AWS_PROFILE           := ""
AWS_REGION            := "ap-northeast-1"
SOPS_KMS_ARN          := "arn:aws:kms:ap-northeast-1:<account_id>:key/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"

# environment specific overrides
ifeq ($(ENV),prd)
    AWS_PROFILE           := ""
    AWS_REGION            := "ap-northeast-1"
    SOPS_KMS_ARN          := "arn:aws:kms:ap-northeast-1:<account_id>:key/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"
endif


.PHONY: fmt validate lint lint-init init list show plan apply destroy import remove
fmt: ## format terraform files
	terraform fmt

validate: ## validate terraform configuration
	terraform validate

lint-init: ## initialize tflint plugins
	tflint --init

lint: ## run tflint checks
	tflint --recursive

init: ## initialize terraform workspace
	@AWS_PROFILE=$(AWS_PROFILE) AWS_DEFAULT_REGION=$(AWS_REGION) \
		terraform init -reconfigure $(if $(BACKEND_CONFIG),$(BACKEND_CONFIG),)

list: ## list terraform state resources
	@AWS_PROFILE=$(AWS_PROFILE) AWS_DEFAULT_REGION=$(AWS_REGION) terraform state list

show: ## show specific resource state (use with AWS_RESOURCE_TYPE and AWS_RESOURCE_NAME)
	@AWS_PROFILE=$(AWS_PROFILE) AWS_DEFAULT_REGION=$(AWS_REGION) terraform state show $(AWS_RESOURCE_TYPE).$(AWS_RESOURCE_NAME)

plan: ## plan terraform changes
	@AWS_PROFILE=$(AWS_PROFILE) AWS_DEFAULT_REGION=$(AWS_REGION) terraform plan $(VAR_OPTS)

apply: ## apply terraform changes
	@AWS_PROFILE=$(AWS_PROFILE) AWS_DEFAULT_REGION=$(AWS_REGION) terraform apply $(VAR_OPTS)

destroy: ## destroy terraform resources
	@AWS_PROFILE=$(AWS_PROFILE) AWS_DEFAULT_REGION=$(AWS_REGION) terraform destroy

remove: ## remove resource from state (use with AWS_RESOURCE_TYPE and AWS_RESOURCE_NAME)
	@AWS_PROFILE=$(AWS_PROFILE) AWS_DEFAULT_REGION=$(AWS_REGION) terraform state rm $(AWS_RESOURCE_TYPE).$(AWS_RESOURCE_NAME)

import: ## import existing aws resource (use with AWS_RESOURCE_TYPE, AWS_RESOURCE_NAME, AWS_RESOURCE_IDENTIFIER)
	@AWS_PROFILE=$(AWS_PROFILE) AWS_DEFAULT_REGION=$(AWS_REGION) \
		terraform import $(AWS_RESOURCE_TYPE).$(AWS_RESOURCE_NAME) $(AWS_RESOURCE_IDENTIFIER)


.PHONY: encrypt-secret decrypt-secret
encrypt-secret: ## encrypt secret file with sops (use with CREDENTIAL_FILE_NAME)
	@AWS_PROFILE=$(AWS_PROFILE) AWS_DEFAULT_REGION=$(AWS_REGION) \
		sops --kms $(SOPS_KMS_ARN) tfvars/$(ENV)_$(CREDENTIAL_FILE_NAME).yaml

decrypt-secret: ## decrypt secret file with sops (use with CREDENTIAL_FILE_NAME)
	@AWS_PROFILE=$(AWS_PROFILE) AWS_DEFAULT_REGION=$(AWS_REGION) \
		sops --kms $(SOPS_KMS_ARN) -d tfvars/$(ENV)_$(CREDENTIAL_FILE_NAME).yaml


.PHONY: help
help: ## show this help message
	@echo "Terraform Management Commands"
	@echo ""
	@echo "Usage: make [target] [ENV=stg|prd]"
	@echo ""
	@echo "Available targets:"
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-20s\033[0m %s\n", $$1, $$2}'

.DEFAULT_GOAL := help
