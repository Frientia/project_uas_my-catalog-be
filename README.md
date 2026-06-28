# My Catalog Book Backend API

<div align="center">
  <img width="300" height="301" alt="Institut Teknologi dan Bisnis Bina Sarana Global" src="https://github.com/user-attachments/assets/1e84f66a-135b-4cf2-b07a-b2a9098ce119" width="200"/>
</div>

<div align="center">
Institut Teknologi dan Bisnis Bina Sarana Global <br>
FAKULTAS TEKNOLOGI INFORMASI & KOMUNIKASI <br>
https://global.ac.id/
</div>

## Project UAS
- Nim : 1123150114
- Nama : Muhamad Yajid Rizky
- Mata Kuliah : Aplikasi Mobile
- Kelas : TI-SE 23 SH

## Deskripsi Project
Project ini adalah backend API untuk sistem My Catalog Book yang digunakan untuk mendukung autentikasi pengguna, manajemen katalog produk, keranjang belanja, serta proses checkout dan pemesanan buku secara aman. Aplikasi ini dibangun menggunakan Go dengan framework Gin, MySQL, Firebase Authentication, dan JWT untuk autentikasi backend.

## Demo Video
Lihat demo aplikasi dan alur fitur yang tersedia dalam video berikut.

**[Watch Full Demo on YouTube]()**

Alternative link: **[Google Drive Demo]()**

## Fitur Utama
- Autentikasi pengguna dengan Firebase Authentication
- Verifikasi token Firebase dan penerbitan JWT backend
- Pengambilan data profil pengguna yang terautentikasi
- Manajemen katalog produk dengan pagination dan filter kategori
- CRUD produk untuk role admin
- Keranjang belanja pengguna
- Checkout order dan riwayat pesanan
- Middleware autentikasi dan role-based access control

## Teknologi yang Digunakan
- **[Go](https://go.dev/)** - Bahasa pemrograman backend
- **[Gin](https://gin-gonic.com/)** - Web framework
- **[GORM](https://gorm.io/)** - ORM untuk MySQL
- **[MySQL](https://www.mysql.com/)** - Database relasional
- **[Firebase](https://firebase.google.com/)** - Authentication
- **[JWT](https://jwt.io/)** - Token autentikasi
- **[godotenv](https://github.com/joho/godotenv)** - Manajemen environment variable

## Persyaratan Sistem
Pastikan perangkat Anda sudah memiliki:
- Go (versi terbaru yang kompatibel dengan modul ini)
- MySQL Server
- Firebase project dengan service account
- Git
- Postman (opsional untuk testing API)

## Cara Menjalankan Project

### 1. Clone Repository
```bash
git clone https://github.com/Frientia/my-catalog-book-backend.git
cd my_catalog_book
```

### 2. Install Dependency
```bash
go mod tidy
```

### 3. Siapkan Environment
Buat file `.env` berdasarkan konfigurasi yang dibutuhkan, contohnya:
```env
APP_PORT=8080
APP_ENV=development
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASSWORD=
DB_NAME=book_store
JWT_SECRET=1234567890abcdef1234567890abcdef
JWT_EXPIRE_HOURS=24
FIREBASE_CREDENTIALS_PATH=./firebase-service-account.json
GOOGLE_APPLICATION_CREDENTIALS=./firebase-service-account.json
```

### 4. Siapkan Firebase
- Buat project Firebase
- Aktifkan Authentication
- Download file service account JSON
- Simpan sebagai `firebase-service-account.json` di root project

### 5. Siapkan Database MySQL
Pastikan MySQL sudah berjalan dan database yang digunakan sudah tersedia sebelum menjalankan server.

### 6. Jalankan Server
```bash
go run main.go
```

Server akan berjalan di:
```bash
http://localhost:8080
```

### 7. Seed Data Produk (Opsional)
Untuk mengisi data produk contoh ke database:
```bash
go run ./seed
```

## Struktur Project
```bash
.
├── config/             # Konfigurasi aplikasi dan koneksi database/Firebase
├── handlers/           # Handler HTTP untuk endpoint API
├── middleware/         # Middleware autentikasi dan role guard
├── models/             # Struktur data model untuk user, product, cart, order
├── repositories/       # Layer akses data ke database
├── routes/             # Routing API
├── seed/               # Script seeding data produk contoh
├── services/           # Logika bisnis aplikasi
├── main.go             # Entry point aplikasi
├── go.mod              # Modul Go
├── .env                # Konfigurasi environment
└── firebase-service-account.json # Kredensial Firebase Admin SDK
```

## Dokumentasi API
Base URL:
```bash
http://localhost:8080/v1
```

### Authentication
- `POST /v1/auth/verify-token` - Verifikasi Firebase ID Token dan menghasilkan JWT backend

### Profile
- `GET /v1/profile` - Ambil data profil user (butuh token)

### Products
- `GET /v1/products` - Ambil daftar produk dengan pagination dan filter kategori
- `GET /v1/products/:id` - Ambil detail produk
- `POST /v1/products` - Tambah produk (admin only)
- `PUT /v1/products/:id` - Update produk (admin only)
- `DELETE /v1/products/:id` - Hapus produk (admin only)

### Cart
- `GET /v1/cart` - Ambil isi keranjang pengguna
- `POST /v1/cart` - Tambah produk ke keranjang
- `PUT /v1/cart/:id` - Update quantity item keranjang
- `DELETE /v1/cart/:id` - Hapus satu item keranjang
- `DELETE /v1/cart` - Kosongkan keranjang

### Orders
- `POST /v1/orders/checkout` - Checkout keranjang menjadi pesanan
- `GET /v1/orders` - Ambil riwayat pesanan pengguna
- `GET /v1/orders/:id` - Ambil detail pesanan

### Health Check
- `GET /v1/health` - Cek status service

## Lisensi
Project ini dilisensikan di bawah MIT License.

## Ucapan Terima Kasih
- [Flutter Community](https://flutter.dev/community)
- [Firebase](https://firebase.google.com/)
- [Gin Gonic](https://gin-gonic.com/)
- [GORM](https://gorm.io/)
- [MySQL](https://www.mysql.com/)

---
<div align="center">
  <p>© 2026 My Catalog Book Backend API. All rights reserved.</p>
</div>
