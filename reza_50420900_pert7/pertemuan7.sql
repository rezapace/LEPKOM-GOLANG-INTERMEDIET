--Membuat database reza_50420900_pert7
CREATE DATABASE reza_50420900_pert7;

--menggunakan database reza_50420900_pert7
use reza_50420900_pert7;

--membuat tabel mahasiswa
CREATE TABLE mahasiswa (
    id int(6) unsigned AUTO_INCREMENT primary key,
    npm char(8) NOT NULL,
    nama varchar(30) NOT NULL,
    kelas char(5) NOT NULL,
    profile varchar(30)NOT NULL);

--menampilkan tabel employee
select * from mahasiswa;

--menampilkan database employee
desc mahasiswa;


--output

Setting environment for using XAMPP for Windows.
R@DESKTOP-Q3HDCLA c:\xampp
# mysql -u root -p
Enter password:
Welcome to the MariaDB monitor.  Commands end with ; or \g.
Your MariaDB connection id is 36
Server version: 10.4.25-MariaDB mariadb.org binary distribution

Copyright (c) 2000, 2018, Oracle, MariaDB Corporation Ab and others.

Type 'help;' or '\h' for help. Type '\c' to clear the current input statement.

MariaDB [(none)]> CREATE DATABASE reza_50420900_pert7;
Query OK, 1 row affected (0.001 sec)

MariaDB [(none)]> use reza_50420900_pert7;
Database changed
MariaDB [reza_50420900_pert7]> CREATE TABLE mahasiswa (
    ->     id int(6) unsigned AUTO_INCREMENT primary key,
    ->     npm char(8) NOT NULL,
    ->     nama varchar(30) NOT NULL,
    ->     kelas char(5) NOT NULL,
    ->     profile varchar(30)NOT NULL);
Query OK, 0 rows affected (0.018 sec)

MariaDB [reza_50420900_pert7]> desc mahasiswa;
+---------+-----------------+------+-----+---------+----------------+
| Field   | Type            | Null | Key | Default | Extra          |
+---------+-----------------+------+-----+---------+----------------+
| id      | int(6) unsigned | NO   | PRI | NULL    | auto_increment |
| npm     | char(8)         | NO   |     | NULL    |                |
| nama    | varchar(30)     | NO   |     | NULL    |                |
| kelas   | char(5)         | NO   |     | NULL    |                |
| profile | varchar(30)     | NO   |     | NULL    |                |
+---------+-----------------+------+-----+---------+----------------+
5 rows in set (0.007 sec)

MariaDB [reza_50420900_pert7]>

--output mahasiswa

Setting environment for using XAMPP for Windows.
R@DESKTOP-Q3HDCLA c:\xampp
# mysql -u root -p
Enter password:
Welcome to the MariaDB monitor.  Commands end with ; or \g.
Your MariaDB connection id is 61
Server version: 10.4.25-MariaDB mariadb.org binary distribution

Copyright (c) 2000, 2018, Oracle, MariaDB Corporation Ab and others.

Type 'help;' or '\h' for help. Type '\c' to clear the current input statement.

MariaDB [(none)]> use reza_50420900_pert7;
Database changed
MariaDB [reza_50420900_pert7]> select * from mahasiswa;
+----+----------+---------------------+-------+-------------+
| id | npm      | nama                | kelas | profile     |
+----+----------+---------------------+-------+-------------+
|  1 | 51416277 | Azman Agung Nugraha | 4IA17 | gambar1.jpg |
|  3 | 52417067 | Fahmi Ardhiansyah   | 3IA03 | gambar1.jpg |
|  4 | 50420900 | Muhammad Reza       | 3IA20 | gambar1.jpg |
+----+----------+---------------------+-------+-------------+
3 rows in set (0.000 sec)

MariaDB [reza_50420900_pert7]>