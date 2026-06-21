# Golang Starterkit (Gin Gonic + GORM)

Starterkit minimalis, modern, dan siap production (_production-ready_) untuk membangun RESTful API menggunakan bahasa pemrograman Go, framework **Gin Gonic**, dan **GORM** ORM.

Starterkit ini dikonfigurasi menggunakan driver SQLite _pure-Go_ (CGO-free) secara bawaan sehingga mempermudah jalannya proses development di sistem operasi Windows tanpa perlu instalasi compiler C (GCC) tambahan. Starterkit ini juga mendukung PostgreSQL dan MySQL jika diperlukan.

---

## 📁 Struktur Folder

Struktur folder starterkit didesain menggunakan pendekatan modular berbasis _feature/domain package_:

```text
├── config/              # Manajemen konfigurasi environment
│   └── config.go        # Loader variabel .env ke struct Go
├── database/            # Setup database, migrasi & seeder
│   ├── database.go      # Manajemen koneksi GORM (SQLite/Postgres/MySQL)
│   ├── migrations/      # Kumpulan file auto-migrasi skema database
│   │   ├── migration.go     # Runner/Pusat migrasi
│   │   └── post_migration.go# Registrasi model Post untuk migrasi
│   └── seeders/         # Kumpulan file data seed awal
│       ├── seeder.go        # Runner/Pusat seeder
│       └── post_seeder.go   # Data seed untuk model Post
├── middleware/          # HTTP Middlewares global/grup
│   ├── cors.go          # Middleware CORS (Cross-Origin Resource Sharing)
│   └── auth.go          # Middleware Autentikasi JWT & Otorisasi Role
├── packages/            # Lokasi modul domain/fitur (Modular Architecture)
│   └── post/            # Modul fitur Blog Post (CRUD)
│       ├── model.go      # Struktur tabel & representasi data GORM
│       ├── repository.go # Interaksi query langsung ke database
│       ├── service.go    # Validasi input & logika bisnis utama
│       └── handler.go    # Handler HTTP untuk request/response Gin
├── routes/              # Konfigurasi routing endpoint
│   └── api.go           # Definisi & routing HTTP endpoints (/api/v1)
├── .env                 # File environment aktif (lokal)
├── .env.example         # Template environment default
├── go.mod               # Definisi modul & dependencies
├── main.go              # Entry point bootstrap aplikasi
└── README.md            # Dokumentasi project
```

---

## 🚀 Panduan Memulai

### 1. Prasyarat

Pastikan komputer Anda sudah terinstall **Go** versi **1.24** atau yang lebih baru.

### 2. Konfigurasi Environment (`.env`)

Salin file `.env.example` menjadi `.env` (file `.env` default sudah dibuatkan dengan konfigurasi SQLite):

```bash
# Untuk Windows (Command Prompt)
copy .env.example .env

# Untuk Bash/PowerShell
cp .env.example .env
```

### 3. Mengunduh Dependencies

Jalankan perintah berikut untuk mengunduh semua library/packages yang dibutuhkan:

```bash
go mod tidy
```

### 4. Menjalankan Server API

Jalankan aplikasi dengan perintah:

```bash
go run main.go
```

Aplikasi akan secara otomatis mendeteksi koneksi database, menjalankan auto-migrasi tabel `posts`, dan menyalakan server di **http://localhost:8080**.

---

## 🔑 Autentikasi & Otorisasi (JWT & Roles)

Starterkit ini dilengkapi dengan sistem keamanan bawaan di folder `middleware/auth.go`:

1. **Authentication**: Melindungi rute menggunakan `AuthMiddleware`. Header `Authorization: Bearer <JWT_TOKEN>` wajib disertakan.
2. **Authorization**: Membatasi hak akses menggunakan role melalui `AuthorizeRoles("role_name")`.

### Cara Menguji API Terproteksi lewat Swagger UI
1. Kirim request `POST /api/v1/auth/login` menggunakan input kredensial berikut:
   ```json
   {
     "username": "admin",
     "password": "password",
     "role": "admin"
   }
   ```
2. Salin token JWT yang dihasilkan di response.
3. Klik tombol hijau **"Authorize"** di kanan atas halaman Swagger UI.
4. Masukkan value dengan format: `Bearer <token_jwt_anda>` (contoh: `Bearer eyJhbGciOi...`).
5. Klik **"Authorize"**. Sekarang Anda bisa memanggil endpoint yang membutuhkan role admin.

---

## 📡 Endpoint API (Blog Post CRUD)

Endpoint berada di bawah group `/api/v1`:

| Method     | Endpoint            | Proteksi / Otorisasi           | Keterangan                                  |
| :--------- | :------------------ | :----------------------------- | :------------------------------------------ |
| **POST**   | `/api/v1/auth/login`| Public                         | Autentikasi user & mendapatkan JWT Token    |
| **GET**    | `/api/v1/posts`     | Public                         | Menampilkan seluruh list blog post          |
| **GET**    | `/api/v1/posts/:id` | Public                         | Menampilkan detail blog post berdasarkan ID |
| **POST**   | `/api/v1/posts`     | Bearer Token (Role: **admin**) | Membuat blog post baru                      |
| **PUT**    | `/api/v1/posts/:id` | Bearer Token (Role: **admin**) | Memperbarui blog post berdasarkan ID        |
| **DELETE** | `/api/v1/posts/:id` | Bearer Token (Role: **admin**) | Menghapus blog post berdasarkan ID          |

### Contoh Request Payload untuk Membuat Post (`POST /api/v1/posts`)

```json
{
  "title": "Belajar Golang Starterkit",
  "content": "Ini adalah konten artikel pertama saya menggunakan Gin Gonic dan GORM.",
  "status": "published"
}
```

_(Catatan: Field `slug` akan otomatis terbuat secara otomatis di level service berdasarkan `title`)_.

---

## 📝 Dokumentasi API (Swagger)

Aplikasi ini menggunakan **Swagger** untuk dokumentasi API interaktif.

### Cara Mengakses Swagger UI

Setelah menjalankan server (`go run main.go`), buka browser dan akses alamat berikut:
**`http://localhost:8080/swagger/index.html`**

### Cara Memperbarui Dokumentasi API

Setiap kali Anda membuat atau memperbarui anotasi Swagger pada handler, jalankan perintah ini di root folder untuk men-generate ulang dokumentasi:

```bash
# Pastikan tool swag sudah terinstall (go install github.com/swaggo/swag/cmd/swag@latest)
# Jalankan perintah ini di root project:
swag init
```

File dokumentasi baru akan dibuat secara otomatis di dalam folder `/docs` dan langsung siap disajikan saat server dijalankan.
