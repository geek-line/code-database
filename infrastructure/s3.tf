resource "aws_s3_bucket" "code_pipeline_bucket" {
  bucket        = "code-database-pipeline"
  force_destroy = false
}


resource "aws_s3_bucket_acl" "code_pipeline_bucket" {
  bucket = aws_s3_bucket.code_pipeline_bucket.id
  acl    = "private"
}