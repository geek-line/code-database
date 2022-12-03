resource "aws_cognito_identity_pool" "file_upload" {
  identity_pool_name               = "cognito-file-upload"
  allow_unauthenticated_identities = true
}

resource "aws_cognito_identity_pool_roles_attachment" "knowtfolio" {
  identity_pool_id = aws_cognito_identity_pool.file_upload.id
  roles = {
    "authenticated"   = aws_iam_role.admin.arn
    "unauthenticated" = aws_iam_role.admin.arn
  }
}