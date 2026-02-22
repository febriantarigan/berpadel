terraform {
  backend "s3" {
    bucket  = "berpadel-tf-state"
    key     = "berpadel/terraform.tfstate"
    region  = "eu-central-1"
    dynamodb_table = "berpadel-tf-state-lock"
    encrypt = true
  }
}
