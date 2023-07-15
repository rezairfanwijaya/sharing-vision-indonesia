# Simple Article Management

Simple article management merupakan website sederhana yang digunakan untuk mengatur artikel dengan sangat mudah. Pengguna bisa membuat arikel, mengedit dan juga melihat list artikel dengan status artikel `Publish` atau `Draft` ataupun `Trash`. Berikut merupakan persayaratan untuk menjalankan aplikasi di dalam environtment lokal anda.

## Requirements
- [Golang ^1.18](https://go.dev/dl/)
- [Nodejs ^18.12.0](https://nodejs.org/en)
- [XAMPP 8.0.28](https://www.apachefriends.org/download.html)

Berikut merupakan teknologi yang dipakai dalam aplikasi ini

## What's inside?
- [GIN v1.9.1](https://github.com/gin-gonic/gin)
- [GORM v1.25.2 ](https://gorm.io)
- [Testify v1.8.4](https://github.com/stretchr/testify)
- [Godotenv v1.5.1](https://github.com/joho/godotenv)
- [React Router v1.5.1](https://reactrouter.com/en/main)
- [React Router ^6.14.1](https://reactrouter.com/en/main)
- [Sweet Alert ^11.7.16](https://sweetalert2.github.io)
- [Prime React ^9.6.0](https://primereact.org)
- [Flowbite ^1.7.0](https://flowbite.com)
- [Tailwind](https://tailwindcss.com)

Untuk dapat menjalankan program ini dapat mengikuti intsruksi berikut

## Run Locally For Setup Server (Backend)

#### Clone the project

```bash
  git clone https://github.com/rezairfanwijaya/sharing-vision-indonesia.git
```

#### Go to the project directory

```bash
  cd sharing-vision-indonesia/backend
```

#### Get Dependenvy

```bash
  go mod tidy
```


#### Setup ENV
##### Rename .env.example file menjadi .env dan sesuaikan dengan database yang dipilih
##### contoh :
```bash
DATABASE_USERNAME = "root"
DATABASE_PASSWORD = ""
DATABASE_HOST = "127.0.0.1"
DATABASE_PORT = "3306"
DATABASE_NAME = "article"
```

#### Run application
##### Terdapat executable file bernama `main.exe` dan jalankan file tersebut sehingga terbuka terminal. Jangan menutup terminal selama program dijalankan.


![main file](image/sample/main%20file.png)
![server](image/sample/terminal%20server.png)

Ketika server berhasil dijalankan maka akan dibuatkan tabel artikel secara otomatis di database dengan sample sejumlah tiga artikel.

## Run Locally For Setup Client (Frontend)
#### Go to the project directory

```bash
  cd sharing-vision-indonesia/frontend
```

#### Get Dependency
```bash
  npm install
```

#### Serve Application
```bash
  npm start
```
![halaman utama](image/sample/halaman%20utama.png)
![halaman update](image/sample/halaman%20update.png)
![halaman preview](image/sample/halaman%20preview.png)

## API DOCUMENTATION
#### [POSTMAN](https://documenter.getpostman.com/view/11940636/2s946cfuJE)

