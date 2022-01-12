# Markicabs-Hotel-Information-System
<div id="top"></div>
<!-- PROJECT LOGO -->
<br/>
<div align="center">
  <a href="https://github.com/alta-be4-andri/Project-2">
    <img src="images/header.gif" alt="Logo" width="700" height="300">
  </a>

  <h3 align="center">Project#2 "AirBnb" Reservation App </h3>

  <p align="center">
    Projek Kedua Pembangunan RESTful API Program Immersive Back End Batch 4
    <br />
    <a href="https://github.com/alta-be4-andri/Project-2"><strong>Kunjungi kami »</strong></a>
    <br />
  </p>
</div>

### 🛠 &nbsp;Build App & Database

![JSON](https://img.shields.io/badge/-JSON-05122A?style=flat&logo=json&logoColor=000000)&nbsp;
![GitHub](https://img.shields.io/badge/-GitHub-05122A?style=flat&logo=github)&nbsp;
![Visual Studio Code](https://img.shields.io/badge/-Visual%20Studio%20Code-05122A?style=flat&logo=visual-studio-code&logoColor=007ACC)&nbsp;
![MySQL](https://img.shields.io/badge/-MySQL-05122A?style=flat&logo=mysql&logoColor=4479A1)&nbsp;
![Golang](https://img.shields.io/badge/-Golang-05122A?style=flat&logo=go&logoColor=4479A1)&nbsp;
![AWS](https://img.shields.io/badge/-AWS-05122A?style=flat&logo=amazon)&nbsp;
![Postman](https://img.shields.io/badge/-Postman-05122A?style=flat&logo=postman)&nbsp;
![Docker](https://img.shields.io/badge/-Docker-05122A?style=flat&logo=docker)&nbsp;

<!-- ABOUT THE PROJECT -->
### 💻 &nbsp;About The Project

AirBnb merupakan projek kedua untuk membangun sebuah RESTful API Reservation App dengan menggunakan bahasa Golang.    
dilengkapi dengan berbagai fitur yang memungkinkan user untuk mengakses data yang ada didalam server. mulai dari membuat akun hingga hosting tempat penginapan. Adapun fitur yang ada dalam RESTful API kami antara lain :
<div>
      <details>
<summary>🙎 Users</summary>
  
  <!---
  | Command | Description |
| --- | --- |
  --->
  
 Di User terdapat fitur untuk membuat Akun dan Login agar mendapat legalitas untuk mengakses berbagai fitur lain di aplikasi, 
 terdapat juga fitur Update untuk mengedit data yang berkaitan dengan user, serta fitur delete berfungsi jika user menginginkan hapus akun.
 
<div>
  
| Feature User | Endpoint | Param | JWT Token | Fungsi |
| --- | --- | --- | --- | --- |
| POST | /signup  | - | NO | Melakukan proses registrasi user |
| POST | /signin | - | NO | Melakukan proses login user |
| GET | /users | - | YES | Mendapatkan informasi user yang sedang login |
| PUT | /users | - | YES | Melakukan update informasi user yang sedang login | 
| DEL | /users | - | YES | Menghapus user yang sedang login |

</details>  

<details>
<summary>🏢 &nbsp;Homestay or Hotel</summary>
  
  <!---
  | Command | Description |
| --- | --- |
  --->
  
  Homestay pada tabel ini user dapat mencari homestay atau hotel, untuk memudahkan pencarian, developer menambahkan beberapa fitur seperti getHomestayByKota, getHomestayByRating. 
  selain mencari hotel, tentunya user dapat menghosting hotel, Homestay atau villa untuk dibooking oleh user lain.
  
| Feature Homestay | Endpoint | Param | JWT Token | Fungsi |
| --- | --- | --- | --- | --- |
| POST | /homestays  | - | YES | Membuat homestay baru |
| GET | /homestays | - | YES | Mendapatkan informasi seluruh homestay |
| GET | /homestays/:id | id | NO | Mendapatkan informasi homestay berdasarkan homestay id |
| GET | /homestays/nama_kota | id_kota | NO | Mendapatkan informasi homestay berdasarkan kota |
| PUT | /homestays/:id | id | YES | Melakukan update homestay tertentu berdasarkan id homestay |
| DEL | /homestays/:id | id | YES | Melakukan delete homestay tertentu berdasarkan id homestay |

</details>


<details>
<summary>🛌 &nbsp;Rooms</summary>

<img src="images/fasilitas.gif" width="700" height="300">
  
  <!---
  | Command | Description |
| --- | --- |
  --->
  User dapat mem-posting berbagai room untuk disewakan kepada user lain, sebagai hoster user dapat menambahkan berbagai fasilitas persetiap room, guna mempermudah user lain dalam memilih room sesuai dengan fasilitas yang dikehendaki, adapun fitur GetUserByHomestayId, sebagai hoster user bisa mengupload, edit serta delete.
  
| Feature Room | Endpoint | Param | JWT Token | Fungsi |
| --- | --- | --- | --- | --- |
| POST | /rooms:id | id | YES | Membuat room baru di homestay tertentu |
| GET | /rooms | - | NO | Mendapatkan informasi seluruh room |
| GET | /rooms/homestays/:id | id | NO | Mendapatkan informasi room di homestay berdasarkan homestay id |
| POST | /rooms:id | id | NO | Mendapatkan detail room tertentu berdasarkan room id |
| POST | /rooms/upload | - | YES | Melakukan upload foto room |
| PUT | /homestays/:id | id | YES | Melakukan update room tertentu berdasarkan room id |
| DEL | /homestays/:id | id | YES | Melakukan delete room tertentu berdasarkan room id |

</details>

<details>
<summary>🗓&nbsp;Reservation</summary>
  
  <!---
  | Command | Description |
| --- | --- |
  --->
Setelah user melakukan pemilihan homestay dan melihat room dengan berbagai fasilitas yang ditawarkan, user melakukan reservation dengan melakukan pengecekan tanggal diawal, jika sistem merespon room yang dimaksud "avalaible", user baru dapat melakukan reservation. Adapun fitur dalam reservation CheckAvailable, GetReservationById   
  
| Feature Reservaton | Endpoint | Param | JWT Token | Fungsi |
| --- | --- | --- | --- | --- |
| POST | /rooms/check/:id | id | YES | Melakukan cek ketersediaan room tertentu berdasarkan tanggal check-in dan check-out |
| POST | /reservations | - | YES | Membuat reservasi room |
| GET | /reservations/:id | id | YES | Mendapatkan informasi reservasi berdasarkan reservation id |
| DEL | /reservations/:id | id | YES | Melakukan cancel reservasi berdasarkan reservation id |

</details>

<details>
<summary>💟&nbsp;Review</summary>
  
  <!---
  | Command | Description |
| --- | --- |
  --->
  Review adalah fitur untuk user menuliskan komentar terkait pelayanan dan kondisi kelayakan kamar, sebagai feedback user kepada homestay terkait yang bertujuan untuk memperbaiki serta mempertahan kualitas pelayanan kamar.
  
| Feature Reservaton | Endpoint | Param | JWT Token | Fungsi |
| --- | --- | --- | --- | --- |
| POST | /reviews | - | YES | Membuat review homestay |

</details>
      

<!-- ERD -->
### ERD
<img src="images/Project2 (3).jpg">

<!-- CONTACT -->
### Contact

[![Gmail: Alfy](https://img.shields.io/badge/-Alfy-maroon?style=flat&logo=gmail)](https://mail.google.com/mail/u/0/#inbox?compose=CllgCHrjmjRlSpLttDDmhqnRQTQVTSQCjFvQxCSSqGDHvQjrjJvvzKMvnlWTrWwkcGdSzfJPXnV)
[![GitHub Alfy](https://img.shields.io/badge/-Alfy-white?style=flat&logo=github&logoColor=black)](https://github.com/alfiancikoa)

[![Gmail: Andri](https://img.shields.io/badge/-Andri-maroon?style=flat&logo=gmail)](https://mail.google.com/mail/u/0/#inbox?compose=DmwnWslzCnrLrhrlnrRWdpHqsBmRtbbtZSKxXFrdGHmhLVLjLDmVfNRxdBShrxQNTBBHFgDdLfKQ)
[![GitHub Andri](https://img.shields.io/badge/-Andri-white?style=flat&logo=github&logoColor=black)](https://github.com/DylanRipper)

[![Gmail: Fafa](https://img.shields.io/badge/-Fafa-maroon?style=flat&logo=gmail)](https://mail.google.com/mail/u/0/#inbox?compose=DmwnWslzCnrLrhrlnrRWdpHqsBmRtbbtZSKxXFrdGHmhLVLjLDmVfNRxdBShrxQNTBBHFgDdLfKQ)
[![GitHub FAfa](https://img.shields.io/badge/-Fafa-white?style=flat&logo=github&logoColor=black)](https://github.com/DylanRipper)


<p align="center">:copyright: 2021 | AAF</p>
</h3>
