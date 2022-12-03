resource "aws_acm_certificate" "code-database" {
  domain_name               = var.domain
  subject_alternative_names = ["*.${var.domain}"]
  validation_method         = "DNS"

  provider = aws.virginia
}

resource "aws_acm_certificate" "code-database-ap" {
  domain_name               = var.domain
  subject_alternative_names = ["*.${var.domain}"]
  validation_method         = "DNS"
}