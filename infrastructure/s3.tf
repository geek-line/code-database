resource "aws_s3_bucket" "code_pipeline_bucket" {
  bucket        = "code-database-pipeline"
  force_destroy = false
}

resource "aws_s3_bucket_acl" "code_pipeline_bucket" {
  bucket = aws_s3_bucket.code_pipeline_bucket.id
  acl    = "private"
}

resource "aws_s3_bucket" "code-database_secondary" {
  bucket        = "code-database-secondary"
  force_destroy = false
}

resource "aws_s3_bucket_acl" "code-database_secondary" {
  bucket = aws_s3_bucket.code-database_secondary.id
  acl    = "private"
}

resource "aws_s3_bucket_policy" "secondary" {
  bucket = aws_s3_bucket.code-database_secondary.id
  policy = templatefile("${path.module}/template/s3/hosting_bucket_policy.json", {
    bucket     = aws_s3_bucket.code-database_secondary.bucket
    identifier = aws_cloudfront_origin_access_identity.secondary.iam_arn
  })
}

resource "aws_s3_bucket" "code-database_images" {
  bucket        = "code-database-images-ver2"
  force_destroy = false
}

resource "aws_s3_bucket_acl" "code-database_images" {
  bucket = aws_s3_bucket.code-database_images.id
  acl    = "private"
}

resource "aws_s3_bucket_policy" "images" {
  bucket = aws_s3_bucket.code-database_images.id
  policy = templatefile("${path.module}/template/s3/hosting_bucket_policy.json", {
    bucket     = aws_s3_bucket.code-database_images.bucket
    identifier = aws_cloudfront_origin_access_identity.images.iam_arn
  })
}

resource "aws_s3_bucket_cors_configuration" "images" {
  bucket = aws_s3_bucket.code-database_images.bucket

  cors_rule {
    allowed_headers = ["*"]
    allowed_methods = ["PUT", "POST"]
    allowed_origins = ["https://code-database.com", "http://localhost:8080"]
    expose_headers  = ["ETag"]
    max_age_seconds = 3000
  }

  cors_rule {
    allowed_methods = ["GET"]
    allowed_origins = ["*"]
  }
}