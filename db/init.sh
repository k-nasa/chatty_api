mysql -uroot -e "DROP DATABASE IF EXISTS chatty; CREATE DATABASE chatty;"
mysql -uroot chatty < ./chatty.sql
