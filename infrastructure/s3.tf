resource "aws_s3_bucket" "code_pipeline_bucket" {
  bucket        = "code-database-pipeline-${var.execution_year}"
  force_destroy = false
}


resource "aws_s3_bucket" "code-database_secondary" {
  bucket        = "code-database-secondary-${var.execution_year}"
  force_destroy = false
}

resource "aws_s3_bucket_policy" "secondary" {
  bucket = aws_s3_bucket.code-database_secondary.id
  policy = templatefile("${path.module}/template/s3/hosting_bucket_policy.json", {
    bucket     = aws_s3_bucket.code-database_secondary.bucket
    identifier = aws_cloudfront_origin_access_identity.secondary.iam_arn
  })
}

resource "aws_s3_bucket" "code-database_images" {
  bucket        = "code-database-images-${var.execution_year}"
  force_destroy = false
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