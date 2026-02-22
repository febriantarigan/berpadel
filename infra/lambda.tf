resource "aws_lambda_function" "berpadel" {
  function_name    = local.function_name
  description      = "Lambda function to handle berpadel tournaments, matches, and leaderboard"
  filename         = local.archive_path # Ensure the Go binary is packaged
  source_code_hash = filebase64sha256(local.archive_path)
  handler          = local.binary_name
  runtime          = "provided.al2023"
  role             = aws_iam_role.lambda_role.arn
  timeout          = 30
  architectures = ["arm64"]

  depends_on = [aws_cloudwatch_log_group.log_group]
}

resource "aws_iam_role" "lambda_role" {
  name = "${local.function_name}-lambda-iam-role"
  assume_role_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [{
      Action = "sts:AssumeRole"
      Effect = "Allow"
      Principal = {
        Service = "lambda.amazonaws.com"
      }
    }]
  })
}

resource "aws_iam_policy" "lambda_logging_policy" {
  name        = "${local.function_name}-logging-policy"
  description = "IAM policy for logging from a Lambda function"

  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Effect = "Allow"
        Action = [
          "logs:CreateLogStream",
          "logs:PutLogEvents"
        ]
        Resource = [
          "arn:aws:logs:*:*:log-group:/aws/lambda/functionname:${local.function_name}"
        ]
      }
    ]
  })
}

// create log group in cloudwatch to gather logs of our lambda function
resource "aws_cloudwatch_log_group" "log_group" {
  name              = "/aws/lambda/${local.function_name}"
  retention_in_days = local.retention_in_days
  tags              = local.tags
}

