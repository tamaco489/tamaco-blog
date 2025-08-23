output "article_api" {
  description = "ecr repository for article api"
  value = {
    arn  = aws_ecr_repository.article_api.arn
    id   = aws_ecr_repository.article_api.id
    name = aws_ecr_repository.article_api.name
    url  = aws_ecr_repository.article_api.repository_url
  }
}
