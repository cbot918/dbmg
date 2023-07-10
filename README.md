# DBMG
a simple and minimal db migrate tool using in small personal side-project 

# TestVia
1. linux ubuntu
2. postgresql

# Install
```
bash <(curl -s https://raw.githubusercontent.com/cbot918/dbmg/main/installer/install_in_linux.sh)
```

# Usecase
1. run db
```
./create_db.sh
```

2. put your config in .env
example .env
```
DB_TYPE="postgres"
DB_URL="postgres://postgres:12345@localhost:5433/dby?sslmode=disable"
```

3. run dbmg
create table using sql/create.sql
```
dbmg sql/create.sql
```
drop table
```
dbmg sql.drop.sql
```

# Notice
1. use 5433 port in case not the same with your local 5432 port
2. only very minimal implement 