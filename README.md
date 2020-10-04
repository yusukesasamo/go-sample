# go-sample

## go-sample url rooting
baseurl(/go-sample/api/v1) and add each path

| Path       | HTTP Method |function|
| ---------- | ------ | -------------- |
| /user     | GET    | get lists of user |
| /user     | POST   | register user     |
| /user/    | PATCH  | patch password of user     |
| /userAuth | POST   | auth user     |
| /userPurchaseHistory     | GET   | get lists of user purchase history    |
| /userPurchaseHistory     | POST   | register user purchase history     |
| /item     | GET   | get lists of item     |
| /item     | POST   | register item     |
| /item/:id     | PATCH   | patch item     |
| /item/:id     | DELETE   | delete item     |
| /purchase     | POST   | purchase item     |



### install
```shell
$ go get github.com/gin-gonic/gin
$ go get github.com/go-sql-driver/mysql
$ go get github.com/yusukesasamo/go-sample/src/controller
$ go get github.com/yusukesasamo/go-sample/src/model

```


### DB
please make gosampledb in mysql.

```shell
$ mysql -u root -p
mysql> create database gosampledb;
mysql> create user 'gosample'@'localhost' identified by 'gosample';
mysql> grant all on `gosampledb`.* to 'gosample'@'localhost';
```

```sql
USE gosampledb
CREATE TABLE IF NOT EXISTS user (
    id INT UNSIGNED NOT NULL AUTO_INCREMENT,
	mail VARCHAR(255) NOT NULL,
	password VARCHAR(255) NOT NULL,
	authkey VARCHAR(255) NOT NULL,
	point INT UNSIGNED NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    PRIMARY KEY(id)
);

CREATE TABLE IF NOT EXISTS user_purchase_history (
    id INT UNSIGNED NOT NULL AUTO_INCREMENT,
	user_id INT NOT NULL,
	item_id INT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    PRIMARY KEY(id)
);

CREATE TABLE IF NOT EXISTS item (
    id INT UNSIGNED NOT NULL AUTO_INCREMENT,
	user_id INT NOT NULL,
	name VARCHAR(255) NOT NULL,
	price INT UNSIGNED NOT NULL,
	stock_flg INT UNSIGNED NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    PRIMARY KEY(id)
);
```


## TODO
Refactor code to use swagger.

Refactor code to use sqlmock for unit test.