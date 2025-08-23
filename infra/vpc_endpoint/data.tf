data "terraform_remote_state" "network" {
  backend = "s3"
  config = {
    bucket = "${var.env}-tamaco-blog-tfstate"
    key    = "network/terraform.tfstate"
  }
}
