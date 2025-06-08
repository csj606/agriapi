terraform{
    required_providers {
        aws = {
            source = "hashicorp/aws"
            version = "~> 3.27"
        }
    }
}

provider "aws" {
    profile = "default"
    region = "us-east-1"
}

resource "aws_dynamodb_table" "in-season-table" {
    name = "InSeasonTableIA"
    billing_mode = "PROVISIONED"
    read_capacity = 5
    write_capacity = 5
    hash_key = "month"

    attribute {
        name = "month"
        type = "S"
    }

    tags = {
        Name = "in-season-table-1"
        Environment = "production"
    }
}