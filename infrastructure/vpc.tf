resource "aws_vpc" "code-database" {
  cidr_block = "10.0.0.0/16"
  tags = {
    "Name" = "code-database"
  }
}

resource "aws_subnet" "code-database_public_a" {
  vpc_id                  = aws_vpc.code-database.id
  cidr_block              = "10.0.1.0/24"
  availability_zone       = "ap-northeast-1a"
  map_public_ip_on_launch = true
  tags = {
    "Name" = "code-database-public-a"
  }
}

resource "aws_subnet" "code-database_public_c" {
  vpc_id                  = aws_vpc.code-database.id
  cidr_block              = "10.0.2.0/24"
  availability_zone       = "ap-northeast-1c"
  map_public_ip_on_launch = true
  tags = {
    "Name" = "code-database-public-c"
  }
}

resource "aws_subnet" "code-database_private_a" {
  vpc_id            = aws_vpc.code-database.id
  cidr_block        = "10.0.16.0/20"
  availability_zone = "ap-northeast-1a"
  tags = {
    "Name" = "code-database-private-a"
  }
}

resource "aws_subnet" "code-database_private_c" {
  vpc_id            = aws_vpc.code-database.id
  cidr_block        = "10.0.32.0/20"
  availability_zone = "ap-northeast-1c"
  tags = {
    "Name" = "code-database-private-c"
  }
}

resource "aws_internet_gateway" "code-database" {
  vpc_id = aws_vpc.code-database.id

  tags = {
    "Name" = "code-database-igw"
  }
}

resource "aws_route_table" "code-database_public" {
  vpc_id = aws_vpc.code-database.id

  route {
    cidr_block = "0.0.0.0/0"
    gateway_id = aws_internet_gateway.code-database.id
  }

  tags = {
    "Name" = "code-database"
  }
}

resource "aws_route_table" "code-database_private_a" {
  vpc_id = aws_vpc.code-database.id

  tags = {
    "Name" = "code-database-private-a"
  }
}

resource "aws_route_table" "code-database_private_c" {
  vpc_id = aws_vpc.code-database.id

  tags = {
    "Name" = "code-database-private-c"
  }
}

resource "aws_route_table_association" "code-database_public_a" {
  route_table_id = aws_route_table.code-database_public.id
  subnet_id      = aws_subnet.code-database_public_a.id
}

resource "aws_route_table_association" "code-database_public_c" {
  route_table_id = aws_route_table.code-database_public.id
  subnet_id      = aws_subnet.code-database_public_c.id
}

resource "aws_main_route_table_association" "code-database" {
  vpc_id         = aws_vpc.code-database.id
  route_table_id = aws_route_table.code-database_public.id
}

resource "aws_route_table_association" "code-database_private_a" {
  route_table_id = aws_route_table.code-database_private_a.id
  subnet_id      = aws_subnet.code-database_private_a.id
}

resource "aws_route_table_association" "code-database_private_c" {
  route_table_id = aws_route_table.code-database_private_c.id
  subnet_id      = aws_subnet.code-database_private_c.id
}