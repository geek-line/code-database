locals {
  UBUNTU_20_AMI = "ami-09b18720cb71042df"
}

resource "aws_key_pair" "code-database" {
  key_name   = "code-database"
  public_key = "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCI32rVwydl3GufuH+9ewXyn0jtQUc9GSisHehILOvg5GKLmt7rC9dHkGPuTwMGAQFoYHyleB+k5RLi/3cu9TkNUwAAMxreeKiQvR1DJTiFzxbGivgFko4e6BMgrtp3QcpsXxqBtn55+VUlqMB7jVrVi2iZVIyyBSzm727wJZM/G3j5oQw9VMyqmqD3sOV5444vWR9MolDQwK160VrUqxWy+hM6mIO9jvb2/SIiHU808BnxLqBC7zjHkgqRU7soCAeZCMpZ7DXXh6Sw2ie8Y2Qbb3RS1uzFk+ey6+6zSBf7d8pZHvF5N6C7e4/EPwXPAiNqrnyi/nAo+R4vTIePUyJj"
}

resource "aws_eip" "code-database_backend" {
  instance = aws_instance.code-database_backend.id
  vpc      = true
}

resource "aws_instance" "code-database_backend" {
  ami                         = local.UBUNTU_20_AMI
  instance_type               = "t2.micro"
  subnet_id                   = aws_subnet.code-database_public_a.id
  key_name                    = aws_key_pair.code-database.id
  associate_public_ip_address = true

  root_block_device {
    volume_type = "gp2"
    volume_size = 30
  }

  vpc_security_group_ids = [aws_security_group.code-database_backend_ec2.id]

  tags = {
    "Name" = "dev-code-database-backend"
  }

  user_data = file("${path.module}/user_data/code-database_backend.sh")
}

resource "aws_security_group" "code-database_backend_ec2" {
  name   = "code-database-backend-ec2"
  vpc_id = aws_vpc.code-database.id

  ingress {
    from_port   = 8080
    to_port     = 8080
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  ingress {
    from_port   = 22
    to_port     = 22
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  egress {
    from_port   = 80
    to_port     = 80
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  egress {
    from_port   = 22
    to_port     = 22
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  egress {
    from_port   = 443
    to_port     = 443
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  egress {
    from_port       = 3306
    to_port         = 3306
    protocol        = "tcp"
    security_groups = [aws_security_group.code-database_db.id]
  }
}