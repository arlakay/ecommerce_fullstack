# Rencana Proyek Capstone Fullstack AI Integration - E-Commerce

## Tema Aplikasi: Platform E-Commerce

Kita akan membuat platform e-commerce sederhana yang memungkinkan pengguna untuk menjual, membeli, dan mengelola produk.

## Komponen Teknis

### 1. API Endpoints

#### Produk (Products)
- `GET /products`: Mendapatkan semua produk
- `GET /products/:id`: Mendapatkan produk berdasarkan ID
- `POST /products`: Menambahkan produk baru
- `PUT /products/:id`: Memperbarui produk
- `DELETE /products/:id`: Menghapus produk

#### Pengguna (Users)
- `GET /users`: Mendapatkan semua pengguna
- `GET /users/:id`: Mendapatkan pengguna berdasarkan ID
- `POST /register`: Mendaftarkan pengguna baru
- `POST /login`: Melakukan login pengguna

#### Keranjang Belanja (Cart)
- `GET /cart/:user_id`: Mendapatkan isi keranjang pengguna
- `POST /cart`: Menambahkan item ke keranjang
- `PUT /cart/:id`: Mengubah jumlah item di keranjang
- `DELETE /cart/:id`: Menghapus item dari keranjang

#### Pesanan (Orders)
- `GET /orders/:user_id`: Mendapatkan pesanan pengguna
- `POST /orders`: Membuat pesanan baru
- `GET /orders/:id`: Mendapatkan detail pesanan

### 2. Database (PostgreSQL)

Tables:
1. Users
2. Products
3. Cart_Items (many-to-many relationship antara Users dan Products)
4. Orders
5. Order_Items (many-to-many relationship antara Orders dan Products)

### 3. Routing (React Router)

Halaman:
- Home
- Login/Register
- Product List
- Product Details
- Cart
- Checkout
- User Dashboard (protected)
- Order History (protected)

### 4. Styling (ChakraUI)

Menggunakan ChakraUI untuk membuat interface yang responsif dan user-friendly, termasuk komponen seperti product cards, shopping cart, dan form checkout.

### 5. Fitur

- Autentikasi (JWT)
- Otorisasi untuk halaman dan aksi tertentu
- CRUD operasi untuk produk
- Manajemen keranjang belanja
- Proses checkout
- Riwayat pesanan

### 6. Testing

- Backend: Unit testing dengan Gin (Go)
- Frontend: Unit testing dengan Jest (React)

### 7. Dokumentasi

README.md akan mencakup:
- Latar belakang dan tujuan aplikasi e-commerce
- Daftar API Endpoints
- Struktur database
- Screenshot UI (termasuk halaman produk, keranjang, dan checkout)
- Flow interaksi frontend-backend (misalnya, proses pembelian)