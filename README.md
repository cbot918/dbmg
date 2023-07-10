# DBMG
a simple and minimal db migrate tool using in small personal side-project 

# Test Via
1. linux ubuntu (unix like system should work)
2. postgresql (mysql should work)

# Pre Requesties
```
docker or local postgres server
```

# Install
```
bash <(curl -s https://raw.githubusercontent.com/cbot918/dbmg/main/installer/install_in_linux.sh)
```

# Getting Started
1. clone
```
git clone https://github.com/cbot918/dbmg && cd dbmg
```

2. run db
```
./create_db.sh
```

3. put your config in .env, there is an example

.env
```
DB_TYPE="postgres"
DB_URL="postgres://postgres:12345@localhost:5433/dby?sslmode=disable"
```

4. run dbmg

create table with sql/create.sql
```
dbmg sql/create.sql
```
drop table with sql/drop.sql
```
dbmg sql/drop.sql
```

# Verify
```
docker exec -it dby bash
psql -U postgres -W
// password: 12345
\c dby
// password: 12345
\dt
```

# Notice
1. use 5433 port in case not conflict with your local 5432 port
2. only very minimal implement 