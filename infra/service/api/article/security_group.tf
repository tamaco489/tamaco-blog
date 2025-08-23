resource "aws_security_group" "article_api" {
  name        = local.service
  description = "security group for ${local.service}"
  vpc_id      = data.terraform_remote_state.network.outputs.vpc.id

  tags = { Name = local.service }
}

# article api SG -> secrets manager endpoint sg (egress)
resource "aws_vpc_security_group_egress_rule" "article_api_to_secrets_manager_endpoint" {
  security_group_id            = aws_security_group.article_api.id
  description                  = "allow https to secrets manager endpoint sg"
  from_port                    = 443
  to_port                      = 443
  ip_protocol                  = "tcp"
  referenced_security_group_id = data.terraform_remote_state.vpc_endpoint.outputs.vpce_secrets_manager_sg
}

# article api SG -> secrets manager endpoint sg (ingress)
resource "aws_vpc_security_group_ingress_rule" "secrets_manager_endpoint_from_article_api" {
  security_group_id            = data.terraform_remote_state.vpc_endpoint.outputs.vpce_secrets_manager_sg
  description                  = "allow https from article api sg"
  from_port                    = 443
  to_port                      = 443
  ip_protocol                  = "tcp"
  referenced_security_group_id = aws_security_group.article_api.id
}
