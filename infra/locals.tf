locals {
  region = "eu-central-1"

  function_name     = "berpadel"
  api_name          = "berpadel-api-gw"
  retention_in_days = 7

  binary_name  = "bootstrap"
  archive_path = "./artifacts/berpadel.zip"

  tags = {
    cc_id     = "216127929236"
    app       = "berpadel"
  }
}

