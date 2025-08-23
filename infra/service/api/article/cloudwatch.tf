resource "aws_cloudwatch_log_group" "article_api" {
  name              = "/aws/lambda/${aws_lambda_function.article_api.function_name}"
  retention_in_days = 3

  tags = { Name = local.service }
}
