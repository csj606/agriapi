terraform{
    required_providers {
        aws = {
            source = "hashicorp/aws"
            version = "~> 4.16"
        }
    }
    required_version = ">= 1.2.0"
}

provider "aws" {
    region = "us-east-1"
}

resource "aws_instance" "agriapi" {
    instance_type = "t2.micro"
    key_name = "cloud-private-key"
    ami = "ami-09e6f87a47903347c"

    tags ={
        Name = "agriapi"
    }

    connection {
        type = "ssh"
        user = "ec2-user"
        private_key = file("~/.ssh/id_rsa")
        host = self.public_ip
    }

    provisioner "remote-exec" {
        inline = [
            "sudo yum update",
            "sudo yum install golang-go git -y",
            "git clone https://github.com/csj606/agriapi",
            "cd api",
            "nohup go run ingress.go &"
        ]
    }
}