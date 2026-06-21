# Golang Starterkit (Gin Gonic + GORM)

Starterkit minimalis, modern, dan siap production (_production-ready_) untuk membangun RESTful API menggunakan bahasa pemrograman Go, framework **Gin Gonic**, dan **GORM** ORM.

Starterkit ini dikonfigurasi menggunakan driver SQLite _pure-Go_ (CGO-free) secara bawaan sehingga mempermudah jalannya proses development di sistem operasi Windows tanpa perlu instalasi compiler C (GCC) tambahan. Starterkit ini juga mendukung PostgreSQL dan MySQL jika diperlukan.

---

## 📁 Struktur Folder

Struktur folder starterkit didesain menggunakan pendekatan modular berbasis _feature/domain package_:

```text
├── config/              # Manajemen konfigurasi environment
│   └── config.go        # Loader variabel .env ke struct Go
├── database/            # Setup koneksi database & pool
│   └── database.go      # Manajemen koneksi GORM (SQLite/Postgres/MySQL)
├── middleware/          # HTTP Middlewares global/grup
│   ├── cors.go          # Middleware CORS (Cross-Origin Resource Sharing)
│   └── auth.go          # Middleware Autentikasi JWT & Otorisasi Role
├── migrations/          # File migrasi database (opsional jika manual)
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

---

## 📡 Endpoint API (Blog Post CRUD)

Endpoint berada di bawah group `/api/v1`:

| Method     | Endpoint            | Proteksi / Otorisasi           | Keterangan                                  |
| :--------- | :------------------ | :----------------------------- | :------------------------------------------ |
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
