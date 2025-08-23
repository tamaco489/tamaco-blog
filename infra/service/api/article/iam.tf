# =================================================================
# basic iam policy
# =================================================================
data "aws_iam_policy_document" "lambda_execution_assume_role" {
  statement {
    effect = "Allow"
    principals {
      type        = "Service"
      identifiers = ["lambda.amazonaws.com"]
    }
    actions = ["sts:AssumeRole"]
  }
}

resource "aws_iam_role" "article_api" {
  name               = local.service
  assume_role_policy = data.aws_iam_policy_document.lambda_execution_assume_role.json

  tags = { Name = local.service }
}

# https://docs.aws.amazon.com/ja_jp/aws-managed-policy/latest/reference/AWSLambdaVPCAccessExecutionRole.html
resource "aws_iam_role_policy_attachment" "article_api_execution_role" {
  role       = aws_iam_role.article_api.name
  policy_arn = "arn:aws:iam::aws:policy/service-role/AWSLambdaVPCAccessExecutionRole"
}

# =================================================================
# secrets manager iam policy
# =================================================================
data "aws_iam_policy_document" "article_api_secrets_manager" {
  statement {
    effect = "Allow"
    actions = [
      "secretsmanager:GetSecretValue",
    ]
    resources = ["*"]
    # resources = [data.terraform_remote_state.article_core_db.outputs.article_core_db.arn]
  }
  statement {
    effect = "Allow"
    actions = [
      "secretsmanager:BatchGetSecretValue",
    ]
    resources = ["*"]
  }
  statement {
    effect = "Allow"
    actions = [
      "kms:Decrypt",
    ]
    resources = [data.aws_kms_key.secretsmanager.arn]
  }
}

resource "aws_iam_policy" "article_api_secrets_manager" {
  name        = local.service
  description = "Allows Lambda to access Secrets Manager secrets"
  policy      = data.aws_iam_policy_document.article_api_secrets_manager.json
}

resource "aws_iam_role_policy_attachment" "article_api_secrets_manager" {
  role       = aws_iam_role.article_api.name
  policy_arn = aws_iam_policy.article_api_secrets_manager.arn
}
