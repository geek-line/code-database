resource "aws_db_instance" "code-database_db" {
  identifier             = "code-database-db"
  db_name                = "code_database_db"
  allocated_storage      = 20 # 無料枠は20GBまでなので、初期値もこれに合わせる
  max_allocated_storage  = 100
  engine                 = "mysql"
  engine_version         = "8.0"
  instance_class         = "db.t2.micro"
  username               = "admin"
  password               = "password"
  vpc_security_group_ids = [aws_security_group.code-database_db.id]
  db_subnet_group_name   = aws_db_subnet_group.code-database_db.name
  skip_final_snapshot    = true
}

resource "aws_db_subnet_group" "code-database_db" {
  name        = "code-database-db"
  description = "rds subnet group for code-database"
  subnet_ids  = [aws_subnet.code-database_private_a.id, aws_subnet.code-database_private_c.id]
}

resource "aws_security_group" "code-database_db" {
  name        = "code-database-db"
  description = "rds service security group for code-database"
  vpc_id      = aws_vpc.code-database.id
}

resource "aws_security_group_rule" "code-database_db_to_backend" {
  type                     = "egress"
  from_port                = 3306
  to_port                  = 3306
  protocol                 = "tcp"
  source_security_group_id = aws_security_group.code-database_backend_ec2.id
  security_group_id        = aws_security_group.code-database_db.id
}