ENV      := stg
VAR_FILE := ./tfvars/${ENV}.tfvars
VAR_OPTS := -var-file "$(VAR_FILE)"


ifeq ($(ENV),stg)
	AWS_PROFILE           := ""
	AWS_REGION            := "ap-northeast-1"
	SOPS_KMS_ARN          := "arn:aws:kms:ap-northeast-1:<account_id>:key/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"
else ifeq ($(ENV),prd)
	AWS_PROFILE           := ""
	AWS_REGION            := "ap-northeast-1"
	SOPS_KMS_ARN          := "arn:aws:kms:ap-northeast-1:<account_id>:key/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"
else
	AWS_PROFILE           := ""
	AWS_REGION            := "ap-northeast-1"
	SOPS_KMS_ARN          := "arn:aws:kms:ap-northeast-1:<account_id>:key/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"
endif


.PHONY: fmt init list show plan apply destroy import remove
fmt: ## Format Terraform files
	terraform fmt

init: ## Initialize Terraform workspace
	@AWS_PROFILE=$(AWS_PROFILE) AWS_DEFAULT_REGION=$(AWS_REGION) terraform init -reconfigure

list: ## List Terraform state resources
	@AWS_PROFILE=$(AWS_PROFILE) AWS_DEFAULT_REGION=$(AWS_REGION) terraform state list

show: ## Show specific resource state (use with AWS_RESOURCE_TYPE and AWS_RESOURCE_NAME)
	@AWS_PROFILE=$(AWS_PROFILE) AWS_DEFAULT_REGION=$(AWS_REGION) terraform state show $(AWS_RESOURCE_TYPE).$(AWS_RESOURCE_NAME)

plan: ## Plan Terraform changes
	@AWS_PROFILE=$(AWS_PROFILE) AWS_DEFAULT_REGION=$(AWS_REGION) terraform plan $(VAR_OPTS)

apply: ## Apply Terraform changes
	@AWS_PROFILE=$(AWS_PROFILE) AWS_DEFAULT_REGION=$(AWS_REGION) terraform apply $(VAR_OPTS)

destroy: ## Destroy Terraform resources
	@AWS_PROFILE=$(AWS_PROFILE) AWS_DEFAULT_REGION=$(AWS_REGION) terraform destroy

remove: ## Remove resource from state (use with AWS_RESOURCE_TYPE and AWS_RESOURCE_NAME)
	@AWS_PROFILE=$(AWS_PROFILE) AWS_DEFAULT_REGION=$(AWS_REGION) terraform state rm $(AWS_RESOURCE_TYPE).$(AWS_RESOURCE_NAME)

import: ## Import existing AWS resource (use with AWS_RESOURCE_TYPE, AWS_RESOURCE_NAME, AWS_RESOURCE_IDENTIFIER)
	@AWS_PROFILE=$(AWS_PROFILE) AWS_DEFAULT_REGION=$(AWS_REGION) \
		terraform import $(AWS_RESOURCE_TYPE).$(AWS_RESOURCE_NAME) $(AWS_RESOURCE_IDENTIFIER)


.PHONY: encrypt-secret decrypt-secret
encrypt-secret: ## Encrypt secret file with SOPS (use with CREDENTIAL_FILE_NAME)
	@AWS_PROFILE=$(AWS_PROFILE) AWS_DEFAULT_REGION=$(AWS_REGION) \
		sops --kms $(SOPS_KMS_ARN) tfvars/$(ENV)_$(CREDENTIAL_FILE_NAME).yaml

decrypt-secret: ## Decrypt secret file with SOPS (use with CREDENTIAL_FILE_NAME)
	@AWS_PROFILE=$(AWS_PROFILE) AWS_DEFAULT_REGION=$(AWS_REGION) \
		sops --kms $(SOPS_KMS_ARN) -d tfvars/$(ENV)_$(CREDENTIAL_FILE_NAME).yaml


.PHONY: help
help: ## Show this help message
	@echo "Terraform Management Commands"
	@echo ""
	@echo "Usage: make [target] [ENV=stg|prd]"
	@echo ""
	@echo "Available targets:"
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-20s\033[0m %s\n", $$1, $$2}'

.DEFAULT_GOAL := help
