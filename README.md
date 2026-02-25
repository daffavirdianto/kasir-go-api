# 🧾 Kasir Go API

RESTful API aplikasi **Kasir (Point of Sale)** yang dibangun menggunakan **Golang (Go)**.  
Project ini mengelola sistem penjualan seperti produk, kategori, transaksi (checkout), dan laporan penjualan.

---

## 🚀 Fitur Utama

- ✅ CRUD Produk
- ✅ CRUD Kategori
- ✅ Checkout / Transaksi
- ✅ Laporan Penjualan (Report)
- ✅ Clean Architecture (Handler → Service → Repository)
- ✅ Environment Configuration dengan Viper
- ✅ Health Check Endpoint
- ✅ RESTful API Design

---

## 🏗️ Struktur Project

```
kasir-go-api/
│
├── database/        # Konfigurasi & koneksi database
├── handlers/        # HTTP Handler (Controller)
├── models/          # Struct / Model data
├── repositories/    # Query & akses database
├── services/        # Business logic
├── main.go          # Entry point aplikasi
├── go.mod
└── go.sum
```

---

## ⚙️ Requirements

Pastikan sudah menginstall:

- Go 1.20+
- PostgreSQL / MySQL
- Git

Cek versi Go:

```bash
go version
```

---

## 🔧 Konfigurasi Environment

Project ini menggunakan **Viper** untuk membaca environment variable.

Buat file `.env` di root project:

```env
PORT=8080
DB_CONN=postgresql://user:password@localhost:5432/kasir_db?sslmode=disable
```

---

## 📦 Instalasi & Menjalankan Project

### 1️⃣ Clone Repository

```bash
git clone https://github.com/daffavirdianto/kasir-go-api.git
cd kasir-go-api
```

### 2️⃣ Install Dependencies

```bash
go mod tidy
```

### 3️⃣ Jalankan Server

```bash
go run main.go
```

---

# 📡 API Endpoints

---

## 📦 Products

| Method | Endpoint             | Deskripsi            |
|--------|----------------------|----------------------|
| GET    | /api/products        | Ambil semua produk   |
| GET    | /api/products/{id}   | Detail produk        |
| POST   | /api/products        | Tambah produk        |
| PUT    | /api/products/{id}   | Update produk        |
| DELETE | /api/products/{id}   | Hapus produk         |

---

## 🗂️ Categories

| Method | Endpoint               | Deskripsi              |
|--------|------------------------|------------------------|
| GET    | /api/categories        | Ambil semua kategori   |
| GET    | /api/categories/{id}   | Detail kategori        |
| POST   | /api/categories        | Tambah kategori        |
| PUT    | /api/categories/{id}   | Update kategori        |
| DELETE | /api/categories/{id}   | Hapus kategori         |

---

## 💳 Checkout / Transaction

| Method | Endpoint        | Deskripsi               |
|--------|-----------------|-------------------------|
| POST   | /api/checkout   | Proses transaksi        |

---

## 📊 Reports

| Method | Endpoint              | Deskripsi                      |
|--------|-----------------------|--------------------------------|
| GET    | /api/report           | Laporan seluruh transaksi      |
| GET    | /api/report/today     | Laporan transaksi hari ini     |

---

## ❤️ Health Check

| Method | Endpoint | Deskripsi              |
|--------|----------|------------------------|
| GET    | /health  | Cek status service     |


## 🛠️ Arsitektur

Project ini menggunakan pola:

```
Handler → Service → Repository → Database
```

- **Handler** → Mengatur HTTP request/response
- **Service** → Business logic
- **Repository** → Query database
- **Database** → Connection & driver

---