resource "aws_dynamodb_table" "padel" {
  name         = "berpadel-table"
  billing_mode = "PAY_PER_REQUEST"
  hash_key     = "PK"
  range_key    = "SK"

  attribute { 
    name = "PK"  
    type = "S" 
  }

  attribute {
    name = "SK"
    type = "S"
  }

  # GSI1: List Tournaments / Users
  attribute {
    name = "GSI1PK"
    type = "S"
  }

  attribute {
    name = "GSI1SK"
    type = "S"
  }

  global_secondary_index {
    name            = "GSI1"
    hash_key        = "GSI1PK"
    range_key       = "GSI1SK"
    projection_type = "ALL"
  }

  # GSI4: Username search
  attribute {
    name = "GSI4PK"
    type = "S"
  }
  attribute {
    name = "GSI4SK"
    type = "S"
  }

  global_secondary_index {
    name            = "GSI4"
    hash_key        = "GSI4PK"
    range_key       = "GSI4SK"
    projection_type = "ALL"
  }

  tags = {
    Project = "padel"
    Env     = "dev"
  }
}

output "padel_table_name" {
  value = aws_dynamodb_table.padel.name
}