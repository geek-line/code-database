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