# API Gateway (HTTP API - The cheapest type)
resource "aws_apigatewayv2_api" "http_api" {
  name          = "berpadel-api"
  protocol_type = "HTTP"
}

resource "aws_apigatewayv2_stage" "default" {
  api_id      = aws_apigatewayv2_api.http_api.id
  name        = "$default"
  auto_deploy = true
}

# The "Catch-all" Route
resource "aws_apigatewayv2_route" "proxy_route" {
  api_id    = aws_apigatewayv2_api.http_api.id
  route_key = "ANY /{proxy+}" # Redirects everything to Gin
  target    = "integrations/${aws_apigatewayv2_integration.lambda_handler.id}"
}

resource "aws_apigatewayv2_integration" "lambda_handler" {
  api_id           = aws_apigatewayv2_api.http_api.id
  integration_type = "AWS_PROXY"
  integration_uri  = aws_lambda_function.berpadel.invoke_arn
  payload_format_version = "2.0"
}