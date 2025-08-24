resource "aws_kms_key" "this" {
  description             = "kms key for sops encryption"
  enable_key_rotation     = true
  rotation_period_in_days = 90
  deletion_window_in_days = 7
}

resource "aws_kms_alias" "this" {
  name          = "alias/${var.env}/${var.project}/sops"
  target_key_id = aws_kms_key.this.key_id
}
