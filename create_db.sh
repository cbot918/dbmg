NETWORK=qcode-server_default
HOST=db
DB_USER=root
DB_PASSWORD=12345
DB=qcode

MY_IP=$(docker inspect -f '{{range.NetworkSettings.Networks}}{{.IPAddress}}{{end}}' dby)

# docker run -p 3307:3306 --name dby -d -e MYSQL_ROOT_PASSWORD=12345 -e POSTGRES_DB=dby postgres


echo $MY_IP
# docker run --name dby -e MYSQL_ROOT_PASSWORD=12345 -d mysql
docker run -it --rm mysql mysql -h $MY_IP -u root -p$"12345" -e "create database dby"

# docker run -it --rm mysql mysql -h <container_host> -P <container_port> -u <username> -p<password> <database_name>




# docker run -it --rm --network $NETWORK mysql mysql -h db -u $DB_USER -p$"$DB_PASSWORD" $DB -e "$1"