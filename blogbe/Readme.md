# BlogBe
```
sudo docker run -d --name mysql -p 3306:3306 mysql/mysql-server:latest
docker logs mysql 2>&1 | grep GENERATE
sudo docker exec -it 530e3bac9282 bash
docker exec -it your_container_name_or_id bash
mysql -u your_user -p (2 lan)
```
```
CREATE USER 'dangnk'@'%' IDENTIFIED BY '12345678';
GRANT ALL PRIVILEGES ON *.* TO 'dangnk'@'%' WITH GRANT OPTION;
FLUSH PRIVILEGES;
```
```
sudo docker run -d --name phpmyadmin --link mysql:db -p 8090:80 phpmyadmin
```