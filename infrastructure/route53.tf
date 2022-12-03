resource "aws_route53_zone" "code-database" {
  name          = var.domain
  force_destroy = false
}

resource "aws_route53_record" "code-database_alb" {
  zone_id        = aws_route53_zone.code-database.id
  name           = var.domain
  type           = "A"
  set_identifier = "primary"

  alias {
    evaluate_target_health = true
    name                   = aws_lb.code-database_backend.dns_name
    zone_id                = aws_lb.code-database_backend.zone_id
  }

  failover_routing_policy {
    type = "PRIMARY"
  }
}

resource "aws_route53_record" "secondary" {
  zone_id        = aws_route53_zone.code-database.id
  name           = var.domain
  type           = "A"
  set_identifier = "secondary"

  alias {
    evaluate_target_health = true
    name                   = aws_cloudfront_distribution.code-database-secondary.domain_name
    zone_id                = aws_cloudfront_distribution.code-database-secondary.hosted_zone_id
  }

  failover_routing_policy {
    type = "SECONDARY"
  }
}

resource "aws_route53_record" "images" {
  zone_id = aws_route53_zone.code-database.id
  name    = "image.${var.domain}"
  type    = "A"

  alias {
    evaluate_target_health = false
    name                   = aws_cloudfront_distribution.upload_files.domain_name
    zone_id                = aws_cloudfront_distribution.upload_files.hosted_zone_id
  }
}

resource "aws_route53_record" "google_site_verification" {
  zone_id = aws_route53_zone.code-database.id
  name    = var.domain
  type    = "TXT"
  ttl     = 300

  records = ["google-site-verification=yqt2lIBUEELAKUS-P_FmHfkZm0v4e6a9eD7sHg2Jvtw"]
}