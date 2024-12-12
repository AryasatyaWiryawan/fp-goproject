# Aplikasi E-Commerce dengan Golang dan Gin

## Plan

### Tujuan
Membangun aplikasi e-commerce sederhana menggunakan framework Golang dan Gin, dengan fitur operasi CRUD dasar untuk dua entitas terkait: `Users` dan `Products`.

### Langkah Utama

1. **Persiapan Proyek**
   - Inisialisasi proyek Go baru dengan `go mod init`.
   - Instal framework Gin dan driver database (misalnya, SQLite atau PostgreSQL).
   - Buat struktur proyek dengan direktori untuk models, controllers, routes, dan database.

2. **Konfigurasi Database**
   - Atur koneksi ke database.
   - Definisikan model `Users` dan `Products` dengan field dan relasi yang sesuai.
   - Konfigurasikan migrasi otomatis untuk model.

3. **Implementasi CRUD**
   - Buat controller untuk menangani operasi CRUD pada user dan produk.
   - Implementasikan route API untuk endpoint seperti `GET /users`, `POST /users`, `GET /products`, dan `POST /products`.

4. **Routing**
   - Definisikan dan kelompokkan route untuk user dan produk.
   - Daftarkan route di file aplikasi utama.

5. **Pengujian**
   - Jalankan aplikasi secara lokal.
   - Uji semua endpoint menggunakan Postman atau `curl` untuk memastikan fungsionalitas.

### Hasil yang Diharapkan
- Aplikasi e-commerce yang berfungsi penuh dengan operasi CRUD dasar.
- Struktur proyek yang bersih dan terorganisir.
- Database SQLite untuk setup dan pengujian yang cepat.

### Langkah Selanjutnya
- Pertimbangkan untuk menambahkan autentikasi dan fitur yang lebih kompleks di masa mendatang.
- Dokumentasikan penggunaan dan endpoint untuk pengembangan lebih lanjut.

---
