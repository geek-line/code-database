resource "aws_iam_role" "code_build_role" {
  name               = "code-build-role"
  assume_role_policy = file("${path.module}/template/iam/code_build/assume_role.json")
}

resource "aws_iam_role_policy" "code_build_policy" {
  name = "code-build-policy"
  role = aws_iam_role.code_build_role.id
  policy = templatefile("${path.module}/template/iam/code_build/code_build_policy.json", {
    bucket     = aws_s3_bucket.code_pipeline_bucket.arn
    connection = aws_codestarconnections_connection.github.arn
  })
}

resource "aws_iam_role" "code_pipeline_role" {
  name = "code-pipeline-role"

  assume_role_policy = file("${path.module}/template/iam/code_pipeline/assume_role.json")
}

resource "aws_iam_role_policy" "code_pipeline_policy" {
  name = "code-pipeline-policy"
  role = aws_iam_role.code_pipeline_role.id

  policy = templatefile("${path.module}/template/iam/code_pipeline/code_pipeline_policy.json", {
    bucket     = aws_s3_bucket.code_pipeline_bucket.arn
    connection = aws_codestarconnections_connection.github.arn
    }
  )
}

resource "aws_iam_role" "code_deploy_role" {
  name = "code-deploy-role"

  assume_role_policy = file("${path.module}/template/iam/code_deploy/assume_role.json")
}

resource "aws_iam_role_policy_attachment" "aws_code_deploy_role" {
  policy_arn = "arn:aws:iam::aws:policy/service-role/AWSCodeDeployRole"
  role       = aws_iam_role.code_deploy_role.name
}

resource "aws_iam_role" "code-database_backend_role" {
  name = "code-database-backend"

  assume_role_policy = file("${path.module}/template/iam/ec2/assume_policy.json")
}

resource "aws_iam_role_policy" "code-database_backend_policy" {
  name = "code-database-backend-policy"
  role = aws_iam_role.code-database_backend_role.id

  policy = templatefile("${path.module}/template/iam/ec2/code_database_backend.json", {
    bucket = aws_s3_bucket.code_pipeline_bucket.arn
    }
  )
}

resource "aws_iam_role_policy_attachment" "backend_ssm" {
  role       = aws_iam_role.code-database_backend_role.name
  policy_arn = "arn:aws:iam::aws:policy/AmazonSSMManagedInstanceCore"
}
resource "aws_iam_role" "admin" {
  name = "code-database-admin"
  assume_role_policy = templatefile("${path.module}/template/iam/cognito/assume_role_policy.json", {
    aud = aws_cognito_identity_pool.file_upload.id
  })
}

resource "aws_iam_role_policy" "upload_images" {
  name = "upload-image"
  role = aws_iam_role.admin.id

  policy = templatefile("${path.module}/template/iam/cognito/admin_policy.json", {
    resource = aws_s3_bucket.code-database_images.arn
  })
}

resource "aws_iam_role" "xml_helper_bot" {
  name               = "xml-helper-bot"
  assume_role_policy = file("${path.module}/template/iam/lambda/assume_role_policy.json")
}

resource "aws_iam_policy" "basic_lambda" {
  name   = "basic-lambda"
  policy = file("${path.module}/template/iam/lambda/basic_lambda.json")
}

resource "aws_iam_role_policy_attachment" "helper_basic" {
  role       = aws_iam_role.xml_helper_bot.name
  policy_arn = aws_iam_policy.basic_lambda.arn
}

resource "aws_iam_role_policy_attachment" "helper_ec2_read" {
  role       = aws_iam_role.xml_helper_bot.name
  policy_arn = "arn:aws:iam::aws:policy/AmazonEC2ReadOnlyAccess"
}

resource "aws_iam_role_policy_attachment" "helper_ssm" {
  role       = aws_iam_role.xml_helper_bot.name
  policy_arn = "arn:aws:iam::aws:policy/AmazonSSMFullAccess"
}