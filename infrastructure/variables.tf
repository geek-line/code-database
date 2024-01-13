variable "domain" {
  default = "code-database.com"
}

variable "execution_year" {
  description = "awsのアカウント移行の際にs3ファイル名などのグローバルでユニークにする必要のある変数に対応するための変数"
  default     = "2024"
}