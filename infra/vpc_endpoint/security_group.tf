resource "aws_security_group" "vpce_secrets_manager" {
  name        = "${local.fqn}-secrets-manager-vpce"
  description = "security group for secrets manager vpc endpoint"
  vpc_id      = data.terraform_remote_state.network.outputs.vpc.id

  tags = { Name = "${local.fqn}-secrets-manager-vpce" }
}
