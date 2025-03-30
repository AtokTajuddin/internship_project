# Evermos Virtual Internship Backend Project

## Overview

Proyek ini merupakan aplikasi backend komprehensif yang dikembangkan sebagai bagian dari program Magang Virtual Evermos. Proyek ini mengimplementasikan API e-commerce yang kuat dengan otentikasi pengguna, manajemen produk, dan integrasi data wilayah administratif Indonesia.

## Features

- **Authentication System**
  - User registration and login via HTTP
  - JWT token-based authentication
  - Password hashing and secure storage
- **Product Management**
  - Create, read, update, and delete (CRUD) operations for products
  - Product categorization and filtering
  - Store association for products
- **Indonesian Administrative Regions API**
  - Integration with Indonesia regional data
  - Hierarchical data: Provinces, Regencies, Districts, and Villages
  - Efficient data retrieval and caching
- **Store Management**
  - Store creation and association with users
  - Product listings per store

## Tech Stack

- **Language**: Go
- **Database**: PostgreSQL
- **Authentication**: JWT (JSON Web Tokens)
- **Testing**: Go testing package with testify suite
- **API Documentation**: HTTP request files for easy testing and documentation
- **External Data**: Integration with [Indonesia Regional API]

## Project Structure
```
project_virtual_internship_evermos
├───.vscode
├───app
└───internal
    ├───helper
    ├───httpserver
    │   ├───handler
    │   └───server
    ├───infra
    │   ├───container
    │   └───mysql
    ├───middleware
    ├───package
    │   ├───controller
    │   ├───entity
    │   ├───model
    │   ├───repository
    │   └───usecase
    └───utils
```

## API Endpoints

### Authentication
**Supported Method**: via `.http`

- **POST /api/v1/auth/register** - Register a new user
  - Required fields: `full_name`, `email`, `phone`, `password`
- **POST /api/v1/auth/login** - Login and receive authentication token
  - Required fields: `email`, `password`

### Products

- **POST /api/v1/products** - Create a new product
  - Required fields: `name`, `price`, `stock`, `category_id`, `store_id`
  - Optional fields: `description`, `brand_id`, `image_url`
- **GET /api/v1/products** - List all products
  - Supports filtering by `name`, `price range`, and `category`
- **GET /api/v1/products/:id** - Get a specific product by ID
- **PUT /api/v1/products/:id** - Update a product
  - Requires authentication as the store owner
- **DELETE /api/v1/products/:id** - Delete a product
  - Requires authentication as the store owner

### Indonesian Regions

- **GET /api/v1/regions/provinces** - List all provinces
- **GET /api/v1/regions/regencies/:provinceId** - List regencies in a province
- **GET /api/v1/regions/districts/:regencyId** - List districts in a regency
- **GET /api/v1/regions/villages/:districtId** - List villages in a district

## Architecture

Aplikasi ini mengikuti pendekatan arsitektur yang bersih dengan pemisahan yang jelas:

1. **Entity**: Model bisnis inti yang independen dari kerangka kerja eksternal.
2. **Repositori**: Antarmuka dan implementasi akses data.
3. **Layanan**: Logika bisnis yang mengatur aliran data dan mengimplementasikan aturan bisnis.
4. **Penanganan**: Penangan permintaan HTTP yang menerjemahkan antara API dan lapisan layanan.

## Testing

Proyek ini mencakup pengujian yang komprehensif:

- **Unit Tests**: Menguji komponen secara individu
- **HTTP Request Files**: Pengujian API manual menggunakan `.http` files
- **PowerShell Script**: Pengujian API otomatis menggunakan PowerShell

### Mock Testing

Lapisan repositori diuji menggunakan server HTTP tiruan yang mensimulasikan respons dari API eksternal, untuk memastikan hasil pengujian yang andal dan dapat diprediksi.

## Getting Started

### Prerequisites

- Go 1.16 atau lebih tinggi
- Database PostgreSQL
- Visual Studio Code (disarankan untuk bekerja dengan `.http` files)

### Setup

1. Clone repository
    ```sh
    git clone <repository-url>
    cd project_virtual_internship_evermos
    ```
2. Install dependencies
    ```sh
    go mod download
    ```
3. Set up database dan konfigurasi koneksi pada environment
4. Jalankan aplikasi
    ```sh
    go run main.go
    ```
5. API tersedia di `http://localhost:8080`

## Notes

1. **file `test_api.ps1`** berfungsi sebagai metode lain untuk memasukkan data ke dalam database selain metode HTTP.
2. **file `jwt_generator.ps1`** berfungsi sebagai generator JWT secret key, sehingga pengguna tidak perlu memasukkan secara manual saat melakukan build aplikasi.

## Acknowledgments
Evermos for providing the virtual internship opportunity
emsifa.com for the Indonesia Regional API
