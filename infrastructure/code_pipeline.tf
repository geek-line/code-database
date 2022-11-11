resource "aws_codestarconnections_connection" "github" {
  name          = "code-database"
  provider_type = "GitHub"
}

resource "aws_codepipeline" "code-database" {
  name     = "code-database"
  role_arn = aws_iam_role.code_pipeline_role.arn

  artifact_store {
    location = aws_s3_bucket.code_pipeline_bucket.bucket
    type     = "S3"
  }

  stage {
    name = "Source"

    action {
      name             = "Source"
      category         = "Source"
      owner            = "AWS"
      provider         = "CodeStarSourceConnection"
      version          = "1"
      output_artifacts = ["SourceArtifact"]

      configuration = {
        ConnectionArn        = aws_codestarconnections_connection.github.arn
        FullRepositoryId     = "geek-line/code-database"
        BranchName           = "feature/terraform"
        OutputArtifactFormat = "CODEBUILD_CLONE_REF"
      }
    }
  }

  stage {
    name = "Build"

    action {
      name             = "Build"
      category         = "Build"
      owner            = "AWS"
      provider         = "CodeBuild"
      input_artifacts  = ["SourceArtifact"]
      output_artifacts = ["BuildArtifact"]
      version          = "1"

      configuration = {
        ProjectName = aws_codebuild_project.code-database.name
      }
    }
  }

  stage {
    name = "Deploy"

    action {
      name            = "Deploy"
      category        = "Deploy"
      owner           = "AWS"
      provider        = "CodeDeploy"
      input_artifacts = ["BuildArtifact"]
      version         = "1"
      configuration = {
        ApplicationName     = aws_codedeploy_app.code-database.name
        DeploymentGroupName = aws_codedeploy_deployment_group.code-database.app_name
      }
    }
  }
}

resource "aws_codebuild_project" "code-database" {
  name         = aws_codedeploy_app.code-database.name
  service_role = aws_iam_role.code_build_role.arn

  artifacts {
    type = "CODEPIPELINE"
  }

  environment {
    compute_type                = "BUILD_GENERAL1_LARGE"
    type                        = "LINUX_CONTAINER"
    image                       = "aws/codebuild/standard:5.0"
    image_pull_credentials_type = "CODEBUILD"
  }

  environment_variable {
      name  = "SQL_ENV"
      value = "root:password@tcp(127.0.0.1:3306)/code_database"
    }

    environment_variable {
      name  = "SOME_KEY1"
      value = "SOME_VALUE1"
    }

  source {
    type = "CODEPIPELINE"
  }
}

resource "aws_codedeploy_app" "code-database" {
  name = "code-database"
}

resource "aws_codedeploy_deployment_group" "code-database" {
  app_name              = aws_codedeploy_app.code-database.name
  deployment_group_name = "code-database"
  service_role_arn      = aws_iam_role.code_deploy_role.arn
  ec2_tag_set {
    ec2_tag_filter {
      key   = "Name"
      type  = "KEY_AND_VALUE"
      value = aws_instance.code-database_backend.tags.Name
    }
  }
}