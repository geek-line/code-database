resource "aws_route53_zone" "code-database" {
  name          = var.domain
  force_destroy = false
}

resource "aws_route53_record" "code-database_record_verify_code" {
  name = "_d395d33dad203cf5531d043859133433.knowtfolio.com."
  records = [
    "_f87bd708fde05acc4defe491da3e4b77.btkxpdzscj.acm-validations.aws."
  ]
  ttl     = "300"
  type    = "CNAME"
  zone_id = aws_route53_zone.code-database.id
}

resource "aws_route53_record" "code-database_alb" {
  zone_id        = aws_route53_zone.code-database.id
  name           = "dev.${var.domain}"
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