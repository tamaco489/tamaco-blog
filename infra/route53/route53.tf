resource "aws_route53_zone" "main" {
  name    = var.domain
  comment = "tamaco-blog"

  tags = { Name = local.fqn }
}
