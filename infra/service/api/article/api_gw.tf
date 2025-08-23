resource "aws_apigatewayv2_api" "article_api" {
  name          = local.service
  description   = "article api service"
  protocol_type = "HTTP"

  tags = { Name = local.service }
}

resource "aws_apigatewayv2_integration" "lambda_integration" {
  api_id                 = aws_apigatewayv2_api.article_api.id
  integration_type       = "AWS_PROXY"
  integration_uri        = aws_lambda_function.article_api.invoke_arn
  payload_format_version = "2.0"
  timeout_milliseconds   = 30000 # 30 seconds
}

resource "aws_apigatewayv2_route" "article_api" {
  api_id    = aws_apigatewayv2_api.article_api.id
  route_key = "ANY /article/v1/{proxy+}"
  target    = "integrations/${aws_apigatewayv2_integration.lambda_integration.id}"
}

resource "aws_apigatewayv2_stage" "article_api" {
  api_id      = aws_apigatewayv2_api.article_api.id
  name        = "$default"
  auto_deploy = true
}

resource "aws_apigatewayv2_api_mapping" "article_api" {
  api_id      = aws_apigatewayv2_api.article_api.id
  domain_name = data.terraform_remote_state.acm.outputs.article_api_apigatewayv2_domain_name.id
  stage       = aws_apigatewayv2_stage.article_api.id
}
