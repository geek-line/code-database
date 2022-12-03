resource "aws_cloudfront_distribution" "code-database-secondary" {
  origin {
    domain_name = aws_s3_bucket.code-database_secondary.bucket_regional_domain_name
    origin_id   = aws_s3_bucket.code-database_secondary.id

    s3_origin_config {
      origin_access_identity = aws_cloudfront_origin_access_identity.secondary.cloudfront_access_identity_path
    }
  }

  comment             = "CDN for code-database static files hosting"
  enabled             = true
  is_ipv6_enabled     = false
  default_root_object = "503.html"

  aliases = [var.domain]

  default_cache_behavior {
    allowed_methods        = ["GET", "HEAD"]
    cached_methods         = ["GET", "HEAD"]
    viewer_protocol_policy = "redirect-to-https"
    target_origin_id       = aws_s3_bucket.code-database_secondary.id
    cache_policy_id        = "658327ea-f89d-4fab-a63d-7e88639e58f6" // CachingOptimized
    compress               = true
  }

  # TODO: SPAのルーティングの方法について考え直す
  custom_error_response {
    error_code         = 403
    response_code      = 200
    response_page_path = "/503.html"
  }

  restrictions {
    geo_restriction {
      restriction_type = "none"
    }
  }

  viewer_certificate {
    cloudfront_default_certificate = false
    acm_certificate_arn            = aws_acm_certificate.code-database.arn
    minimum_protocol_version       = "TLSv1.2_2018"
    ssl_support_method             = "sni-only"
  }
}

resource "aws_cloudfront_origin_access_identity" "secondary" {
  comment = "origin access identity for s3 code-database-secondary"
}

resource "aws_cloudfront_distribution" "upload_files" {
  origin {
    domain_name = aws_s3_bucket.code-database_images.bucket_regional_domain_name
    origin_id   = aws_s3_bucket.code-database_images.id

    s3_origin_config {
      origin_access_identity = aws_cloudfront_origin_access_identity.images.cloudfront_access_identity_path
    }
  }

  comment             = "CDN for code-database uploaded files hosting"
  enabled             = true
  is_ipv6_enabled     = false
  default_root_object = "notfound.png"

  aliases = ["image.${var.domain}"]

  default_cache_behavior {
    allowed_methods        = ["GET", "HEAD"]
    cached_methods         = ["GET", "HEAD"]
    viewer_protocol_policy = "redirect-to-https"
    target_origin_id       = aws_s3_bucket.code-database_images.id
    cache_policy_id        = "658327ea-f89d-4fab-a63d-7e88639e58f6" // CachingOptimized
    compress               = true
  }
  ordered_cache_behavior {
    path_pattern           = "/uploads"
    allowed_methods        = ["GET", "HEAD"]
    cached_methods         = ["GET", "HEAD"]
    viewer_protocol_policy = "redirect-to-https"
    target_origin_id       = aws_s3_bucket.code-database_images.id
    cache_policy_id        = "658327ea-f89d-4fab-a63d-7e88639e58f6" // CachingOptimized
    compress               = true
  }

  ordered_cache_behavior {
    path_pattern           = "/eyecatches"
    allowed_methods        = ["GET", "HEAD"]
    cached_methods         = ["GET", "HEAD"]
    viewer_protocol_policy = "redirect-to-https"
    target_origin_id       = aws_s3_bucket.code-database_images.id
    cache_policy_id        = "658327ea-f89d-4fab-a63d-7e88639e58f6" // CachingOptimized
    compress               = true
  }

  custom_error_response {
    error_code         = 403
    response_code      = 404
    response_page_path = "/notfound.png"
  }


  restrictions {
    geo_restriction {
      restriction_type = "none"
    }
  }

  viewer_certificate {
    cloudfront_default_certificate = false
    acm_certificate_arn            = aws_acm_certificate.code-database.arn
    minimum_protocol_version       = "TLSv1.2_2018"
    ssl_support_method             = "sni-only"
  }
}

resource "aws_cloudfront_origin_access_identity" "images" {
  comment = "origin access identity for s3 code-database-images"
}