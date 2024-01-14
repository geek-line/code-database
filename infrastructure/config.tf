terraform {
  required_version = ">= 1.1.3"
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "4.11.0"
    }
  }

  backend "s3" {
    bucket                  = "code-database-terraform-2024"
    key                     = "alpha/terraform.tfstate"
    region                  = "ap-northeast-1"
    profile                 = "code-database"
    shared_credentials_file = "./.aws/credentials"
    encrypt                 = true
  }

}

provider "aws" {
  region                   = "ap-northeast-1"
  profile                  = "code-database"
  shared_credentials_files = ["./.aws/credentials"]
  shared_config_files      = ["./.aws/config"]
}

provider "aws" {
  region                   = "us-east-1"
  alias                    = "virginia"
  profile                  = "code-database"
  shared_credentials_files = ["./.aws/credentials"]
  shared_config_files      = ["./.aws/config"]
}