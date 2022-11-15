--Membuat database reza_50420900_pert5
CREATE DATABASE reza_50420900_pert5;

--menggunakan database reza_50420900_pert5
use reza_50420900_pert5;

--membuat tabel employee
CREATE TABLE employee (
    id int(6) AUTO_INCREMENT primary key,
    name varchar(30),
    city varchar(30));

--menampilkan tabel employee
select * from employee;

--menampilkan database employee
desc employee;


--output
Setting environment for using XAMPP for Windows.
R@DESKTOP-Q3HDCLA c:\xampp
# mysql -u root -p
Enter password:
Welcome to the MariaDB monitor.  Commands end with ; or \g.
Your MariaDB connection id is 10
Server version: 10.4.25-MariaDB mariadb.org binary distribution

Copyright (c) 2000, 2018, Oracle, MariaDB Corporation Ab and others.

Type 'help;' or '\h' for help. Type '\c' to clear the current input statement.

MariaDB [(none)]> CREATE DATABASE reza_50420900_pert5;
Query OK, 1 row affected (0.001 sec)

MariaDB [reza_50420900_pert5]> CREATE TABLE employee (
    ->     id int(6) AUTO_INCREMENT primary key,
    ->     name varchar(30),
    ->     city varchar(30));
Query OK, 0 rows affected (0.020 sec)

MariaDB [reza_50420900_pert5]> desc employee;
+-------+-------------+------+-----+---------+----------------+
| Field | Type        | Null | Key | Default | Extra          |
+-------+-------------+------+-----+---------+----------------+
| id    | int(6)      | NO   | PRI | NULL    | auto_increment |
| name  | varchar(30) | YES  |     | NULL    |                |
| city  | varchar(30) | YES  |     | NULL    |                |
+-------+-------------+------+-----+---------+----------------+
3 rows in set (0.009 sec)