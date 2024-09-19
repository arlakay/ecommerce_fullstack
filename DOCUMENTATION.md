
---
# Dokumentasi Proyek E-Commerce

## Latar Belakang dan Tujuan

### Latar Belakang
Proyek ini dikembangkan sebagai bagian dari materi pelajaran dalam kursus [Nama Kursus] dengan tujuan untuk memperdalam pemahaman tentang konsep-konsep utama dalam pengembangan aplikasi e-commerce. Melalui proyek ini, peserta diharapkan dapat mengaplikasikan teori-teori yang telah dipelajari, seperti desain API, struktur database, dan interaksi antara frontend dan backend.

### Tujuan
Tujuan dari aplikasi e-commerce ini adalah:
-   **Mempraktikkan Pengembangan API:** Membuat dan mengelola API untuk berbagai fungsionalitas seperti autentikasi, pengelolaan produk, dan transaksi pembelian.
-   **Memahami Struktur Database:** Merancang dan menerapkan struktur database yang efisien untuk mendukung aplikasi e-commerce.
-   **Implementasi Frontend dan Backend:** Mengembangkan antarmuka pengguna yang intuitif dan menghubungkannya dengan backend melalui API untuk memproses data dan transaksi.
-   **Latihan Integrasi Sistem:** Mengintegrasikan berbagai komponen aplikasi, termasuk frontend, backend, dan database, untuk membentuk sistem e-commerce yang fungsional.

## Daftar API Endpoints

### 1. **Autentikasi**

- **POST /login**
  - **Deskripsi:** Masuk ke aplikasi menggunakan kredensial pengguna.
  - **Request:**
    ```json
    {
      "email": "user@example.com",
      "password": "string"
    }
    ```
  - **Response:**
    ```json
    {
      "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOjcsImV4cCI6MTcyNjU5MDgxOH0.hWrgUemQ1aXJ0fBoQxg0jAyYbY9-XNDuZEKS6DbTwGA",
    }
    ```

- **POST /register**
  - **Deskripsi:** Mendaftar pengguna baru.
  - **Request:**
    ```json
    {
      "name": "string",
      "email": "user@example.com",
      "password": "string"
    }
    ```
  - **Response:**
    ```json
    {
      "message": "User registered successfully"
      "user_id": 9
    }
    ```

- **GET /users**
  - **Deskripsi:** Mendapatkan Semua User.
  - **Header:** Bearer token
  - **Response:**
    ```json
    [
        {
            "id": 1,
            "created_at": "2024-09-13T21:43:08.745452+07:00",
            "updated_at": "2024-09-13T21:43:08.745452+07:00",
            "deleted_at": null,
            "name": "admin",
            "email": "admin@admin.com",
            "password": "$2a$10$NrvYlmK4WwDMjUDgMvpA5ee.h9.hy69atGPzSanCp0EjlQa197w7S"
        },
        {
            "id": 5,
            "created_at": "2024-09-14T23:31:31.808792+07:00",
            "updated_at": "2024-09-14T23:31:31.808792+07:00",
            "deleted_at": null,
            "name": "ersa",
            "email": "ersa@ersa.com",
            "password": "ersa"
        },
        {
            "id": 9,
            "created_at": "2024-09-17T22:35:04.766732+07:00",
            "updated_at": "2024-09-17T22:35:04.766732+07:00",
            "deleted_at": null,
            "name": "z1",
            "email": "z1@example.com",
            "password": "z1"
        }
    ]
    ```

- **GET /user**
  - **Deskripsi:** Mendapatkan User data dengan token.
  - **Header:** Bearer token
  - **Response:**
    ```json
    {
        "id": 7,
        "created_at": "2024-09-14T23:34:41.921669+07:00",
        "updated_at": "2024-09-14T23:34:41.921669+07:00",
        "deleted_at": null,
        "name": "z",
        "email": "z@example.com",
        "password": "z"
    }
    ```

- **GET /users/{:user_id}**
  - **Deskripsi:** Mendapatkan user data berdasarkan id.
  - **Header:** Bearer token
  - **Response:**
    ```json
    {
        "id": 1,
        "created_at": "2024-09-13T21:43:08.745452+07:00",
        "updated_at": "2024-09-13T21:43:08.745452+07:00",
        "deleted_at": null,
        "name": "admin",
        "email": "admin@admin.com",
        "password": "$2a$10$NrvYlmK4WwDMjUDgMvpA5ee.h9.hy69atGPzSanCp0EjlQa197w7S"
    }
    ```

### 2. **Produk**

- **GET /products**
  - **Deskripsi:** Mendapatkan list semua produk.
  - **Response:**
    ```json
    [
        {
            "id": 6,
            "created_at": "2024-09-12T20:56:37.592469+07:00",
            "updated_at": "2024-09-12T20:56:37.592469+07:00",
            "deleted_at": null,
            "name": "Headphones",
            "description": "Noise-cancelling over-ear headphones with Bluetooth",
            "image_url": "",
            "price": 129.99,
            "stock": 60
        },
        {
            "id": 7,
            "created_at": "2024-09-12T20:56:37.592469+07:00",
            "updated_at": "2024-09-12T20:56:37.592469+07:00",
            "deleted_at": null,
            "name": "External Hard Drive",
            "description": "1TB external hard drive with USB 3.0",
            "image_url": "",
            "price": 89.99,
            "stock": 80
        }
    ]
    ```

- **GET /products/{:product_id}**
  - **Deskripsi:** Mendapatkan detail produk berdasarkan id.
  - **Response:**
    ```json
    {
        "id": 5,
        "created_at": "2024-09-12T20:56:37.592469+07:00",
        "updated_at": "2024-09-12T20:56:37.592469+07:00",
        "deleted_at": null,
        "name": "Monitor",
        "description": "24-inch Full HD monitor with slim bezel",
        "image_url": "https://images.tokopedia.net/img/cache/300/VqbcmM/2024/3/28/f680e25b-8435-46d0-a81c-bec8dd060c58.jpg",
        "price": 179.99,
        "stock": 40
    }
    ```

- **GET /products/category/{:category_id}**
  - **Deskripsi:** Mendapatkan list produk berdasarkan kategory id.
  - **Response:**
    ```json
    [
        {
            "id": 6,
            "created_at": "2024-09-12T20:56:37.592469+07:00",
            "updated_at": "2024-09-12T20:56:37.592469+07:00",
            "deleted_at": null,
            "name": "Headphones",
            "description": "Noise-cancelling over-ear headphones with Bluetooth",
            "image_url": "",
            "price": 129.99,
            "stock": 60
        },
        {
            "id": 7,
            "created_at": "2024-09-12T20:56:37.592469+07:00",
            "updated_at": "2024-09-12T20:56:37.592469+07:00",
            "deleted_at": null,
            "name": "External Hard Drive",
            "description": "1TB external hard drive with USB 3.0",
            "image_url": "",
            "price": 89.99,
            "stock": 80
        }
    ]
    ```

- **POST /products**
  - **Deskripsi:** Menambahkan produk baru.
  - **Header:** Bearer token
  - **Request:**
    ```json
    {
        "name":"Test Produk",
        "description":"Ini adalah test produk",
        "image_url":"https://images.tokopedia.net/img/cache/900/hDjmkQ/2024/5/16/c8c354c7-5ac7-4f95-87cb-3dcfed78a8ba.jpg",
        "price":80000,
        "stock":20
    }
    ```
  - **Response:**
    ```json
    {
        "id": 77,
        "created_at": "2024-09-17T22:51:09.8948709+07:00",
        "updated_at": "2024-09-17T22:51:09.8948709+07:00",
        "deleted_at": null,
        "name": "Test Produk",
        "description": "Ini adalah test produk",
        "image_url": "https://images.tokopedia.net/img/cache/900/hDjmkQ/2024/5/16/c8c354c7-5ac7-4f95-87cb-3dcfed78a8ba.jpg",
        "price": 80000,
        "stock": 20
    }
    ```

- **PUT /products/{:product_id}**
  - **Deskripsi:** Mengubah produk berdasarkan id.
  - **Header:** Bearer token
  - **Request:**
    ```json
    {
        "name":"Test Produk",
        "description":"Ini adalah test produk",
        "image_url":"https://images.tokopedia.net/img/cache/900/hDjmkQ/2024/5/16/c8c354c7-5ac7-4f95-87cb-3dcfed78a8ba.jpg",
        "price":80000,
        "stock":20
    }
    ```
  - **Response:**
    ```json
    {
        "id": 77,
        "created_at": "2024-09-17T22:51:09.8948709+07:00",
        "updated_at": "2024-09-17T22:51:09.8948709+07:00",
        "deleted_at": null,
        "name": "Test Produk",
        "description": "Ini adalah test produk",
        "image_url": "https://images.tokopedia.net/img/cache/900/hDjmkQ/2024/5/16/c8c354c7-5ac7-4f95-87cb-3dcfed78a8ba.jpg",
        "price": 80000,
        "stock": 20
    }
    ```

- **DELETE /products/{:product_id}**
  - **Deskripsi:** Menghapus produk berdasarkan id.
  - **Header:** Bearer token
  - **Response:**
    ```json
    {
        "message": "product item removed successfully"
    }
    ```

### 3. **Kategori**

- **GET /categories**
  - **Deskripsi:** Mendapatkan daftar semua kategori.
  - **Response:**
    ```json
    [
      {
        "id": 1,
        "created_at": "0001-01-01T00:00:00Z",
        "updated_at": "0001-01-01T00:00:00Z",
        "deleted_at": null,
        "name": "Electronics",
        "description": "Devices and gadgets including phones, computers, and accessories."
      },
      {
        "id": 2,
        "created_at": "0001-01-01T00:00:00Z",
        "updated_at": "0001-01-01T00:00:00Z",
        "deleted_at": null,
        "name": "Clothing",
        "description": "Apparel for men, women, and children including shirts, pants, and dresses."
      }
    ]
    ```

- **GET /categories/{:category_id}**
  - **Deskripsi:** Mendapatkan detail kategori berdasarkan ID.
  - **Response:**
    ```json
    {
        "id": 1,
        "created_at": "0001-01-01T00:00:00Z",
        "updated_at": "0001-01-01T00:00:00Z",
        "deleted_at": null,
        "name": "Electronics",
        "description": "Devices and gadgets including phones, computers, and accessories."
    }
    ```

- **POST /categories**
  - **Deskripsi:** Menambahkan kategori baru.
  - **Header:** Bearer token
  - **Request:**
    ```json
    {
        "name":"test category",
        "description":"ini adalah test category"
    }
    ```
  - **Response:**
    ```json
    {
        "id": 11,
        "created_at": "2024-09-17T22:59:14.6814128+07:00",
        "updated_at": "2024-09-17T22:59:14.6814128+07:00",
        "deleted_at": null,
        "name": "test category",
        "description": "ini adalah test category"
    }
    ```

- **PUT /categories/{:category_id}**
  - **Deskripsi:** Mengubah kategori berdasarkan id.
  - **Header:** Bearer token
  - **Request:**
    ```json
    {
        "name":"Test Category Edit",
        "description":"Ini adalah test category edited version"
    }
    ```
  - **Response:**
    ```json
    {
        "id": 1,
        "created_at": "0001-01-01T00:00:00Z",
        "updated_at": "2024-09-17T23:00:36.0519184+07:00",
        "deleted_at": null,
        "name": "Test Category Edit",
        "description": "Ini adalah test category edited version"
    }
    ```

- **DELETE /categories/{:category_id}**
  - **Deskripsi:** Menghapus kategori berdasarkan id.
  - **Header:** Bearer token
  - **Response:**
    ```json
    {
        "message": "category removed successfully"
    }
    ```

### 4. **Keranjang Belanja**

- **POST /cart**
  - **Deskripsi:** Menambahkan produk ke keranjang.
  - **Header:** Bearer token
  - **Request:**
    ```json
    {
        "user_id":1,
        "product_id":4,
        "quantity":1
    }
    ```
  - **Response:**
    ```json
    {
        "id": 18,
        "created_at": "2024-09-17T23:03:13.5188853+07:00",
        "updated_at": "2024-09-17T23:03:13.5188853+07:00",
        "deleted_at": null,
        "user_id": 1,
        "product_id": 4,
        "quantity": 1,
        "added_at": "2024-09-17T23:03:13.51913+07:00",
        "product": {
            "id": 0,
            "created_at": "0001-01-01T00:00:00Z",
            "updated_at": "0001-01-01T00:00:00Z",
            "deleted_at": null,
            "name": "",
            "description": "",
            "image_url": "",
            "price": 0,
            "stock": 0
        }
    }
    ```

- **GET /cart/user/{:user_id}**
  - **Deskripsi:** Mendapatkan list item di keranjang berdasarkan user id.
  - **Header:** Bearer token
  - **Response:**
    ```json
    [
        {
            "id": 7,
            "created_at": "2024-09-16T17:43:27.410455+07:00",
            "updated_at": "2024-09-16T17:43:27.410455+07:00",
            "deleted_at": null,
            "user_id": 1,
            "product_id": 5,
            "quantity": 1,
            "added_at": "2024-09-16T17:43:27.4107+07:00",
            "product": {
                "id": 5,
                "created_at": "2024-09-12T20:56:37.592469+07:00",
                "updated_at": "2024-09-12T20:56:37.592469+07:00",
                "deleted_at": null,
                "name": "Monitor",
                "description": "24-inch Full HD monitor with slim bezel",
                "image_url": "https://images.tokopedia.net/img/cache/300/VqbcmM/2024/3/28/f680e25b-8435-46d0-a81c-bec8dd060c58.jpg",
                "price": 179.99,
                "stock": 40
            }
        },
        {
            "id": 8,
            "created_at": "2024-09-16T18:36:21.657712+07:00",
            "updated_at": "2024-09-16T18:36:21.657712+07:00",
            "deleted_at": null,
            "user_id": 1,
            "product_id": 4,
            "quantity": 1,
            "added_at": "2024-09-16T18:36:21.657804+07:00",
            "product": {
                "id": 4,
                "created_at": "2024-09-12T20:56:37.592469+07:00",
                "updated_at": "2024-09-12T20:56:37.592469+07:00",
                "deleted_at": null,
                "name": "Keyboard",
                "description": "Mechanical keyboard with RGB backlighting",
                "image_url": "https://images.tokopedia.net/img/cache/300/VqbcmM/2024/3/28/f680e25b-8435-46d0-a81c-bec8dd060c58.jpg",
                "price": 89.99,
                "stock": 75
            }
        },
        {
            "id": 18,
            "created_at": "2024-09-17T23:03:13.518885+07:00",
            "updated_at": "2024-09-17T23:03:13.518885+07:00",
            "deleted_at": null,
            "user_id": 1,
            "product_id": 4,
            "quantity": 1,
            "added_at": "2024-09-17T23:03:13.51913+07:00",
            "product": {
                "id": 4,
                "created_at": "2024-09-12T20:56:37.592469+07:00",
                "updated_at": "2024-09-12T20:56:37.592469+07:00",
                "deleted_at": null,
                "name": "Keyboard",
                "description": "Mechanical keyboard with RGB backlighting",
                "image_url": "https://images.tokopedia.net/img/cache/300/VqbcmM/2024/3/28/f680e25b-8435-46d0-a81c-bec8dd060c58.jpg",
                "price": 89.99,
                "stock": 75
            }
        }
    ]
    ```

- **PUT /cart/{:cart_id}**
  - **Deskripsi:** Merubah jumlah item di keranjang berdasarkan id.
  - **Header:** Bearer token
  - **Request:**
    ```json
    {
        "user_id":1,
        "product_id":4,
        "quantity":2
    }
    ```
  - **Response:**
    ```json
    {
        "id": 18,
        "created_at": "2024-09-17T23:03:13.518885+07:00",
        "updated_at": "2024-09-17T23:08:21.4078132+07:00",
        "deleted_at": null,
        "user_id": 1,
        "product_id": 4,
        "quantity": 2,
        "added_at": "2024-09-17T23:03:13.51913+07:00",
        "product": {
            "id": 0,
            "created_at": "0001-01-01T00:00:00Z",
            "updated_at": "0001-01-01T00:00:00Z",
            "deleted_at": null,
            "name": "",
            "description": "",
            "image_url": "",
            "price": 0,
            "stock": 0
        }
    }
    ```

- **DELETE /cart/{:cart_id}**
  - **Deskripsi:** Menghapus item di keranjang berdasarkan id.
  - **Header:** Bearer token
  - **Response:**
    ```json
    {
        "message": "Item removed from cart successfully"
    }
    ```

### 5. **Orders**

- **POST /orders**
  - **Deskripsi:** Membuat transaksi pembelian.
  - **Header:** Bearer token
  - **Request:**
    ```json
    {
      "user_id": 1,
      "order_items": [
        {
          "product_id": 5,
          "quantity": 1,
          "price": 179.99
        },
        {
          "product_id": 4,
          "quantity": 1,
          "price": 89.99
        }
      ],
      "total": 269.98,
      "status": "pending"
    }
    ```
  - **Response:**
    ```json
    {
        "id": 3,
        "created_at": "2024-09-18T21:14:07.9185832+07:00",
        "updated_at": "2024-09-18T21:14:07.9185832+07:00",
        "deleted_at": null,
        "user_id": 1,
        "order_items": [
            {
                "id": 5,
                "created_at": "2024-09-18T21:14:08.1054748+07:00",
                "updated_at": "2024-09-18T21:14:08.1054748+07:00",
                "deleted_at": null,
                "order_id": 3,
                "product_id": 5,
                "quantity": 1,
                "price": 179.99
            },
            {
                "id": 6,
                "created_at": "2024-09-18T21:14:08.1054748+07:00",
                "updated_at": "2024-09-18T21:14:08.1054748+07:00",
                "deleted_at": null,
                "order_id": 3,
                "product_id": 4,
                "quantity": 1,
                "price": 89.99
            }
        ],
        "total": 269.98,
        "status": "pending"
    }
    ```

- **GET /orders/user/{:user_id}**
  - **Deskripsi:** Melihat transaksi pembelian berdasarkan user id.
  - **Header:** Bearer token
  - **Response:**
    ```json
    {
      "message": "Payment successful",
      "order_id": "integer"
    }
    ```

- **GET /orders/{:order_id}**
  - **Deskripsi:** Mendapatkan detil transaksi pembelian berdasarkan order id.
  - **Header:** Bearer token
  - **Response:**
    ```json
    {
      "message": "Payment successful",
      "order_id": "integer"
    }
    ```

## Struktur Database

### Tabel users
| Kolom        | Tipe Data | Deskripsi                |
|--------------|-----------|--------------------------|
| id           | INTEGER   | ID Pengguna (Primary Key)|
| name         | VARCHAR   | Nama Pengguna            |
| email        | VARCHAR   | Email Pengguna (Unique)  |
| password     | VARCHAR   | Password Ter-enkripsi    |
| created_at   | TIMESTAMP | Waktu Pembuatan          |
| updated_at   | TIMESTAMP | Waktu Perubahan          |
| deleted_at   | TIMESTAMP | Waktu Penghapusan        |

### Tabel products
| Kolom        | Tipe Data | Deskripsi               |
|--------------|-----------|-------------------------|
| id           | INTEGER   | ID Produk (Primary Key) |
| name         | VARCHAR   | Nama Produk             |
| description  | TEXT      | Deskripsi Produk        |
| image_url    | VARCHAR   | Gambar Produk           |
| price        | DECIMAL   | Harga Produk            |
| stock        | INTEGER   | Jumlah Produk           |
| created_at   | TIMESTAMP | Waktu Pembuatan         |
| updated_at   | TIMESTAMP | Waktu Perubahan         |
| deleted_at   | TIMESTAMP | Waktu Penghapusan       |

### Tabel categories
| Kolom        | Tipe Data | Deskripsi                |
|--------------|-----------|--------------------------|
| id           | INTEGER   | ID Kategori (Primary Key)|
| name         | VARCHAR   | Nama Kategori            |
| description  | TEXT      | Deskripsi Kategori       |
| created_at   | TIMESTAMP | Waktu Pembuatan          |
| updated_at   | TIMESTAMP | Waktu Perubahan          |
| deleted_at   | TIMESTAMP | Waktu Penghapusan        |

### Tabel product_categories
| Kolom        | Tipe Data | Deskripsi               |
|--------------|-----------|-------------------------|
| product_id   | INTEGER   | ID Kategori             |
| category_id  | INTEGER   | ID Produk               |

### Tabel cart_items
| Kolom        | Tipe Data | Deskripsi                 |
|--------------|-----------|---------------------------|
| id           | INTEGER   | ID Keranjang (Primary Key)|
| user_id      | INTEGER   | ID Pengguna (Foreign Key) |
| product_id   | INTEGER   | ID Produk (Foreign Key)   |
| quantity     | INTEGER   | Jumlah Produk             |
| added_at     | TIMESTAMP | Waktu menambahkan item    |
| created_at   | TIMESTAMP | Waktu Pembuatan           |
| updated_at   | TIMESTAMP | Waktu Perubahan           |
| deleted_at   | TIMESTAMP | Waktu Penghapusan         |

### Tabel orders
| Kolom        | Tipe Data | Deskripsi                 |
|--------------|-----------|---------------------------|
| id           | INTEGER   | ID Pesanan (Primary Key)  |
| user_id      | INTEGER   | ID Pengguna (Foreign Key) |
| total        | DECIMAL   | Jumlah Total              |
| status       | VARCHAR   | Status Pesanan            |
| created_at   | TIMESTAMP | Waktu Pembuatan           |
| updated_at   | TIMESTAMP | Waktu Perubahan           |
| deleted_at   | TIMESTAMP | Waktu Penghapusan         |

### Tabel order_items
| Kolom        | Tipe Data | Deskripsi                |
|--------------|-----------|--------------------------|
| id           | INTEGER   | ID Pesanan (Primary Key) |
| order_id     | INTEGER   | ID Pengguna (Foreign Key)|
| product_id   | INTEGER   | Jumlah Total             |
| quantity     | INTEGER   | Status Pesanan           |
| price        | DECIMAL   | Status Pesanan           |
| created_at   | TIMESTAMP | Waktu Pembuatan          |
| updated_at   | TIMESTAMP | Waktu Perubahan          |
| deleted_at   | TIMESTAMP | Waktu Penghapusan        |


## Screenshot UI

![Screenshot UI](./screenshot/register.png)
![Screenshot UI](./screenshot/login.png)
![Screenshot UI](./screenshot/home.png)
![Screenshot UI](./screenshot/product_detail.png)
![Screenshot UI](./screenshot/cart.png)
![Screenshot UI](./screenshot/order_summary.png)
![Screenshot UI](./screenshot/order_place.png)
![Screenshot UI](./screenshot/my_order.png)
![Screenshot UI](./screenshot/my_order_details.png)

## Dokumentasi Flow Interaksi Frontend dengan Backend

### 1. **Authentication**
   - **Frontend:** Pengguna memasukkan email dan password pada formulir login.
   - **Backend:** Mengirimkan data ke endpoint `/login`.
   - **Frontend:** Menerima token dan informasi pengguna untuk disimpan dalam sesi.

### 2. **Produk List**
   - **Frontend:** Menampilkan daftar produk dari endpoint `/products`.
   - **Backend:** Mengambil data produk dari database dan mengirimkan ke frontend.

### 3. **Product Details**
   - **Frontend:** Pengguna melihat detil produk menggunakan endpoint `/products/:id`.
   - **Backend:** Mengupdate data keranjang pengguna dan mengirimkan konfirmasi ke frontend.
   
### 4. **Add to Cart**
   - **Frontend:** Pengguna menambahkan produk ke keranjang menggunakan endpoint `/cart`.
   - **Backend:** Mengupdate data keranjang pengguna dan mengirimkan konfirmasi ke frontend.

### 5. **Place Order**
   - **Frontend:** Pengguna melakukan checkout menggunakan endpoint `/orders`.
   - **Backend:** Memproses data keranjang menjadi pesanan ke frontend.

### 5. **My Order**
   - **Frontend:** Pengguna dapat melihat list pembelian menggunakan endpoint `/orders/user/${user.id}`.
   - **Backend:** Memproses daftar pesanan berdasarkan user yang login dan mengirimkannya ke frontend.


---

Dokumentasi ini diharapkan dapat memberikan gambaran yang jelas tentang aplikasi e-commerce Anda dan memudahkan pengembang lain dalam memahami serta berkontribusi pada proyek ini. Jika ada perubahan atau penambahan, pastikan untuk memperbarui dokumentasi ini agar tetap relevan.