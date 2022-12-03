data "archive_file" "xml_helper_bot" {
  type        = "zip"
  source_file = "${path.module}/function_scripts/xml_helper_bot.py"
  output_path = "${path.module}/function_scripts/xml_helper_bot.zip"
}

resource "aws_lambda_function" "xml_helper_bot" {
  function_name    = "xml_helper_bot"
  role             = aws_iam_role.xml_helper_bot.arn
  filename         = data.archive_file.xml_helper_bot.output_path
  source_code_hash = data.archive_file.xml_helper_bot.output_base64sha256
  handler          = "xml_helper_bot.main"
  runtime          = "python3.9"
  timeout          = 60

  environment {
    variables = {
      INSTANCE_ID = aws_instance.code-database_backend.id
    }
  }
}

resource "aws_lambda_permission" "cloudwatch" {
  statement_id  = "AllowExecutionFromCloudWatch"
  function_name = aws_lambda_function.xml_helper_bot.function_name
  action        = "lambda:InvokeFunction"
  principal     = "events.amazonaws.com"
  source_arn    = aws_cloudwatch_event_rule.event_rule.arn
}

resource "aws_cloudwatch_event_rule" "event_rule" {
  name                = "everyday"
  schedule_expression = "cron(0 0 * * ? *)"
}

resource "aws_cloudwatch_event_target" "event_target" {
  rule = aws_cloudwatch_event_rule.event_rule.name
  arn  = aws_lambda_function.xml_helper_bot.arn
}