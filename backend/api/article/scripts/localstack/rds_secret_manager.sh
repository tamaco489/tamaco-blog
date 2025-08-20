#!/bin/bash

# Load values from .env_localstack
set -a # Enable automatic export of environment variables
source /etc/localstack/init/ready.d/.env_localstack
set +a # Disable automatic export of environment variables

# Create a secret in Secrets Manager
awslocal secretsmanager create-secret \
  --name 'dev/tamaco-blog/article/core/rds' \
  --region ap-northeast-1 \
  --secret-string "{
    \"host\":\"${POSTGRES_HOST}\",
    \"port\":\"${POSTGRES_PORT}\",
    \"username\":\"${POSTGRES_USER}\",
    \"password\":\"${POSTGRES_PASSWORD}\",
    \"dbname\":\"${POSTGRES_DB}\"
  }"
