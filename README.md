# Ecommerce Sayur

[![Go Version](https://img.shields.io/badge/Go-1.25.5-blue.svg)](https://golang.org/)
[![Node Version](https://img.shields.io/badge/Node-18+-green.svg)](https://nodejs.org/)
[![Nuxt](https://img.shields.io/badge/Nuxt-3.16.2-00DC82.svg)](https://nuxt.com/)
[![Vue](https://img.shields.io/badge/Vue-3.5.13-4FC08D.svg)](https://vuejs.org/)
[![TypeScript](https://img.shields.io/badge/TypeScript-5.6.3-blue.svg)](https://www.typescriptlang.org/)
[![PostgreSQL](https://img.shields.io/badge/PostgreSQL-15-blue.svg)](https://www.postgresql.org/)
[![Redis](https://img.shields.io/badge/Redis-7-red.svg)](https://redis.io/)
[![RabbitMQ](https://img.shields.io/badge/RabbitMQ-3-orange.svg)](https://www.rabbitmq.com/)
[![Docker](https://img.shields.io/badge/Docker-Latest-2496ED.svg)](https://www.docker.com/)
[![Echo](https://img.shields.io/badge/Echo-4.13.4-blue.svg)](https://echo.labstack.com/)
[![Tailwind CSS](https://img.shields.io/badge/Tailwind_CSS-4.1.4-38B2AC.svg)](https://tailwindcss.com/)
[![Bootstrap](https://img.shields.io/badge/Bootstrap-5.3.5-7952B3.svg)](https://getbootstrap.com/)

Platform e-commerce berbasis microservices untuk penjualan sayuran dengan arsitektur terdistribusi yang scalable dan maintainable.

## Overview

Ecommerce Sayur adalah sistem e-commerce modern yang dirancang untuk mengelola penjualan produk sayuran secara online. Sistem ini dibangun menggunakan arsitektur microservices yang memungkinkan setiap modul beroperasi secara independen, memfasilitasi pengembangan, deployment, dan scaling yang lebih fleksibel.

Sistem bekerja dengan memisahkan concern ke dalam beberapa service terpisah yang berkomunikasi melalui message queue (RabbitMQ) dan HTTP API. API Gateway berfungsi sebagai single entry point yang menangani routing, authentication, rate limiting, dan logging untuk semua request dari frontend. Setiap service memiliki database PostgreSQL sendiri (database per service pattern) dan menggunakan Redis untuk caching serta session management.

## Features

### User Management

- Registrasi dan autentikasi pengguna dengan JWT
- Verifikasi email untuk aktivasi akun
- Reset password dengan token
- Manajemen profil pengguna
- Role-based access control (RBAC) dengan sistem role dan permission
- Upload avatar pengguna

### Product Management

- CRUD produk dengan kategori hierarkis
- Manajemen stok produk
- Pencarian dan filtering produk
- Upload gambar produk
- Status produk (DRAFT, PUBLISHED, dll)
- Variant produk dengan harga reguler dan sale price

### Order Management

- Pembuatan order dari cart
- Tracking status order (pending, processing, completed, cancelled)
- Manajemen order items
- Order history untuk user dan admin
- Worker untuk update status order secara asynchronous
- Product snapshot untuk menjaga konsistensi data historis

### Payment Integration

- Integrasi dengan Midtrans untuk payment gateway
- Payment callback handling
- Tracking status pembayaran
- Multiple payment methods support

### Notification System

- Real-time notification melalui WebSocket
- Notification history
- Mark as read/unread
- Integration dengan RabbitMQ untuk event-driven notifications

### Cart Management

- Add, update, delete items dari cart
- Session-based cart management
- Integration dengan order service

### Admin Dashboard

- Dashboard untuk monitoring dan manajemen
- Manajemen customers
- Manajemen orders
- Manajemen products dan categories
- Manajemen roles dan permissions

## Tech Stack

### Backend

- **Language**: Go 1.25.5
- **Framework**:
  - Echo v4 (HTTP framework untuk API Gateway dan services)
  - GORM (ORM untuk database operations)
  - Cobra (CLI framework untuk service commands)
- **Database**:
  - PostgreSQL 15 (primary database untuk setiap service)
  - Redis 7 (caching dan rate limiting)
- **Message Queue**: RabbitMQ 3 (asynchronous communication)
- **Authentication**: JWT (JSON Web Tokens)
- **Configuration**: Viper (configuration management)
- **Logging**: Logrus (structured logging)
- **Validation**: go-playground/validator

### Frontend

- **Framework**: Nuxt 3.16.2 (Vue.js meta-framework)
- **Language**: TypeScript 5.6.3
- **State Management**: Pinia 3.0.4
- **Styling**:
  - Tailwind CSS 4.1.4
  - Bootstrap 5.3.5
  - Nuxt UI 3.1.0
- **UI Libraries**:
  - Heroicons
  - Tabler Icons
  - Swiper (carousel/slider)
  - Quill (rich text editor)
  - Dropzone (file upload)
- **Utilities**:
  - VueUse (composition utilities)
  - Day.js (date manipulation)
- **Package Manager**: Bun

### Infrastructure & Tools

- **Containerization**: Docker & Docker Compose
- **Load Testing**: K6
- **Payment Gateway**: Midtrans (Snap integration)

## Project Structure

```
ecommerce-sayur/
├── backend/
│   ├── api-gateway/              # API Gateway service
│   │   ├── handlers/             # Route handlers untuk setiap service
│   │   ├── middleware/           # CORS, JWT, Logger, Rate Limit
│   │   └── utils/                # Proxy utilities
│   ├── user-service/             # User management service
│   │   ├── cmd/                  # CLI commands (start, root)
│   │   ├── config/               # Configuration (DB, Redis, RabbitMQ)
│   │   ├── database/             # Migrations dan seeds
│   │   ├── internal/
│   │   │   ├── adapter/          # HTTP handlers, repositories, message handlers
│   │   │   ├── app/              # Application bootstrap
│   │   │   └── core/             # Domain models dan business logic
│   │   └── utils/                # Utilities (validation, conversion)
│   ├── product-service/          # Product management service
│   │   ├── cmd/                  # CLI commands termasuk workers
│   │   ├── database/             # Migrations
│   │   └── internal/             # Clean architecture layers
│   ├── order-service/            # Order management service
│   │   ├── cmd/                  # CLI commands dan workers
│   │   ├── database/             # Migrations
│   │   └── internal/             # Clean architecture layers
│   ├── payment-service/          # Payment processing service
│   │   └── internal/             # Clean architecture layers
│   ├── notification-service/     # Notification service dengan WebSocket
│   │   └── internal/             # Clean architecture layers
│   ├── docker-compose.yml        # Docker orchestration
│   └── Makefile                  # Build automation untuk Go modules
│
├── frontend/
│   ├── components/               # Vue components
│   │   ├── admin/                # Admin dashboard components
│   │   ├── common/               # Shared components
│   │   ├── home/                 # Homepage components
│   │   └── modals/               # Modal components
│   ├── pages/                    # Nuxt pages (file-based routing)
│   │   ├── auth/                 # Authentication pages
│   │   ├── dashboard/            # Admin dashboard pages
│   │   ├── shop/                 # Shopping pages
│   │   └── account/              # User account pages
│   ├── stores/                   # Pinia stores (state management)
│   ├── composables/              # Vue composables
│   ├── middleware/               # Route middleware (auth, admin)
│   ├── layouts/                  # Layout components
│   ├── plugins/                  # Nuxt plugins
│   └── assets/                   # Static assets (CSS, images)
│
└── k6/                           # Load testing scripts
    ├── product_service/          # Product service tests
    └── user_service/             # User service tests
```

### Arsitektur Service

Setiap service mengikuti pola **Clean Architecture** dengan struktur:

- **cmd/**: Entry point dan CLI commands
- **config/**: Konfigurasi database, Redis, RabbitMQ
- **internal/adapter/**: Adapters untuk external interfaces (HTTP handlers, repositories, message consumers)
- **internal/core/**: Business logic dan domain models
- **internal/app/**: Application layer yang menginisialisasi dependencies

## Installation

### Prerequisites

- Go 1.25.5 atau lebih tinggi
- Node.js 18+ atau Bun
- Docker dan Docker Compose
- PostgreSQL 15 (atau gunakan Docker)
- Redis 7 (atau gunakan Docker)
- RabbitMQ 3 (atau gunakan Docker)

### Setup Backend

1. Clone repository:

```bash
git clone https://github.com/bintangnugrahaa/ecommerce-sayur
cd ecommerce-sayur/backend
```

2. Setup environment variables untuk setiap service. Buat file `.env` di root backend atau di setiap service directory dengan konfigurasi berikut:

**API Gateway (.env)**:

```env
PORT=8080
USER_SERVICE_URL=http://user-service:8090
PRODUCT_SERVICE_URL=http://product-service:8082
ORDER_SERVICE_URL=http://order-service:8083
PAYMENT_SERVICE_URL=http://payment-service:8084
NOTIFICATION_SERVICE_URL=http://notification-service:8081
REDIS_HOST=redis
REDIS_PORT=6379
JWT_SECRET=your_jwt_secret_key_here_change_me
```

**User Service**:

```env
DB_HOST=postgres
DB_PORT=5432
DB_NAME=user_service
DB_USER=postgres
DB_PASSWORD=lokal
REDIS_HOST=redis
REDIS_PORT=6379
RABBITMQ_HOST=rabbitmq
RABBITMQ_PORT=5672
RABBITMQ_USERNAME=guest
RABBITMQ_PASSWORD=guest
PORT=8090
JWT_SECRET=your_jwt_secret_key_here_change_me
```

**Product Service**:

```env
DB_HOST=postgres
DB_PORT=5432
DB_NAME=product_service
DB_USER=postgres
DB_PASSWORD=lokal
REDIS_HOST=redis
REDIS_PORT=6379
RABBITMQ_HOST=rabbitmq
RABBITMQ_PORT=5672
RABBITMQ_USERNAME=guest
RABBITMQ_PASSWORD=guest
PORT=8082
```

**Order Service**:

```env
DB_HOST=postgres
DB_PORT=5432
DB_NAME=order_service
DB_USER=postgres
DB_PASSWORD=lokal
REDIS_HOST=redis
REDIS_PORT=6379
RABBITMQ_HOST=rabbitmq
RABBITMQ_PORT=5672
RABBITMQ_USERNAME=guest
RABBITMQ_PASSWORD=guest
PORT=8083
USER_SERVICE_URL=http://user-service:8090
PRODUCT_SERVICE_URL=http://product-service:8082
PAYMENT_SERVICE_URL=http://payment-service:8084
```

**Payment Service**:

```env
DB_HOST=postgres
DB_PORT=5432
DB_NAME=payment_service
DB_USER=postgres
DB_PASSWORD=lokal
REDIS_HOST=redis
REDIS_PORT=6379
RABBITMQ_HOST=rabbitmq
RABBITMQ_PORT=5672
RABBITMQ_USERNAME=guest
RABBITMQ_PASSWORD=guest
PORT=8084
ORDER_SERVICE_URL=http://order-service:8083
NOTIFICATION_SERVICE_URL=http://notification-service:8081
```

**Notification Service**:

```env
DB_HOST=postgres
DB_PORT=5432
DB_NAME=notification_service
DB_USER=postgres
DB_PASSWORD=lokal
REDIS_HOST=redis
REDIS_PORT=6379
RABBITMQ_HOST=rabbitmq
RABBITMQ_PORT=5672
RABBITMQ_USERNAME=guest
RABBITMQ_PASSWORD=guest
PORT=8081
```

3. Install dependencies untuk semua services:

```bash
make mod-all
```

Atau manual untuk setiap service:

```bash
cd user-service && go mod download && cd ..
cd product-service && go mod download && cd ..
cd order-service && go mod download && cd ..
cd payment-service && go mod download && cd ..
cd notification-service && go mod download && cd ..
cd api-gateway && go mod download && cd ..
```

4. Jalankan database migrations (jika diperlukan):

```bash
# Migrations biasanya dijalankan otomatis saat service start
# Atau bisa dijalankan manual menggunakan migrate tool
```

5. Start semua services menggunakan Docker Compose:

```bash
docker-compose up -d
```

Atau jalankan secara lokal (tanpa Docker):

```bash
# Terminal 1: Start infrastructure
docker-compose up -d postgres redis rabbitmq

# Terminal 2-N: Start each service
cd user-service && go run main.go start
cd product-service && go run main.go start
cd order-service && go run main.go start
cd payment-service && go run main.go start
cd notification-service && go run main.go start
cd api-gateway && go run main.go
```

### Setup Frontend

1. Masuk ke directory frontend:

```bash
cd ../frontend
```

2. Install dependencies:

```bash
bun install
# atau
npm install
```

3. Buat file `.env` di root frontend:

```env
NUXT_USER_API_BASE_URL=http://localhost:8080/api/v1/users
NUXT_ORDER_API_BASE_URL=http://localhost:8080/api/v1/orders
NUXT_PRODUCT_API_BASE_URL=http://localhost:8080/api/v1/products
NUXT_PAYMENT_API_BASE_URL=http://localhost:8080/api/v1/payments
NUXT_NOTIFICATION_API_BASE_URL=http://localhost:8080/api/v1/notifications
MIDTRANS_CLIENT_KEY=your_midtrans_client_key
```

4. Jalankan development server:

```bash
bun run dev
# atau
npm run dev
```

Frontend akan berjalan di `http://localhost:3000`

## Configuration

### Environment Variables

#### Backend Services

Semua service memerlukan konfigurasi berikut:

- **Database**: `DB_HOST`, `DB_PORT`, `DB_NAME`, `DB_USER`, `DB_PASSWORD`
- **Redis**: `REDIS_HOST`, `REDIS_PORT`
- **RabbitMQ**: `RABBITMQ_HOST`, `RABBITMQ_PORT`, `RABBITMQ_USERNAME`, `RABBITMQ_PASSWORD`
- **Service Port**: `PORT` (default berbeda untuk setiap service)
- **JWT Secret**: `JWT_SECRET` (untuk user-service dan api-gateway)
- **Service URLs**: Untuk inter-service communication (ORDER_SERVICE_URL, USER_SERVICE_URL, dll)

#### Frontend

- **API Base URLs**: Konfigurasi URL untuk setiap service API
- **Midtrans Client Key**: Untuk integrasi payment gateway

### Database Configuration

Sistem menggunakan pattern **database per service**. Setiap service memiliki database PostgreSQL sendiri:

- `user_service`
- `product_service`
- `order_service`
- `payment_service`
- `notification_service`

Database diinisialisasi otomatis melalui Docker Compose dengan script `init-db.sh` (jika tersedia).

### Docker Compose

File `docker-compose.yml` mengatur:

- PostgreSQL dengan multiple databases
- Redis untuk caching
- RabbitMQ dengan management UI (port 15672)
- Semua microservices dengan health checks dan dependencies

## Usage

### Menjalankan Aplikasi

**Development mode dengan Docker**:

```bash
cd backend
docker-compose up
```

**Production build**:

```bash
cd backend
docker-compose -f docker-compose.yml build
docker-compose -f docker-compose.yml up -d
```

### API Endpoints

Semua API diakses melalui API Gateway di `http://localhost:8080/api/v1`

#### Public Endpoints (Tanpa Authentication)

**User Service**:

- `POST /api/v1/users/register` - Registrasi user baru
- `POST /api/v1/users/signin` - Login
- `POST /api/v1/users/verify-email` - Verifikasi email
- `POST /api/v1/users/forgot-password` - Request reset password
- `POST /api/v1/users/reset-password` - Reset password dengan token

**Product Service**:

- `GET /api/v1/products/shop` - List produk untuk shop (dengan pagination, filter)
- `GET /api/v1/products/home` - List produk untuk homepage
- `GET /api/v1/products/:id` - Detail produk
- `GET /api/v1/products/categories` - List kategori
- `GET /api/v1/products/categories/:id` - Detail kategori
- `GET /api/v1/products/search` - Pencarian produk

#### Protected Endpoints (Require JWT Token)

**User Service**:

- `GET /api/v1/users/profile` - Get user profile
- `PUT /api/v1/users/profile` - Update profile
- `PUT /api/v1/users/password` - Update password
- `POST /api/v1/users/upload-avatar` - Upload avatar
- `GET /api/v1/users/roles` - Get user roles

**Product Service (Admin)**:

- `GET /api/v1/admin/products` - List produk (admin)
- `POST /api/v1/admin/products` - Create produk
- `PUT /api/v1/admin/products/:id` - Update produk
- `DELETE /api/v1/admin/products/:id` - Delete produk
- `POST /api/v1/admin/products/upload-image` - Upload gambar produk

**Cart**:

- `GET /api/v1/cart` - Get cart items
- `POST /api/v1/cart` - Add item to cart
- `PUT /api/v1/cart/:id` - Update cart item
- `DELETE /api/v1/cart/:id` - Remove item from cart

**Order Service**:

- `GET /api/v1/orders` - List orders
- `GET /api/v1/orders/:id` - Get order detail
- `POST /api/v1/orders` - Create order
- `PUT /api/v1/orders/:id` - Update order
- `DELETE /api/v1/orders/:id` - Cancel order
- `PUT /api/v1/orders/:id/status` - Update order status
- `GET /api/v1/orders/:id/items` - Get order items

**Payment Service**:

- `GET /api/v1/payments` - List payments
- `GET /api/v1/payments/:id` - Get payment detail
- `POST /api/v1/payments` - Create payment
- `PUT /api/v1/payments/:id` - Update payment
- `POST /api/v1/payments/callback` - Payment callback (Midtrans)
- `GET /api/v1/payments/:id/status` - Get payment status

**Notification Service**:

- `GET /api/v1/notifications` - List notifications
- `GET /api/v1/notifications/:id` - Get notification detail
- `POST /api/v1/notifications` - Create notification
- `PUT /api/v1/notifications/:id` - Update notification
- `DELETE /api/v1/notifications/:id` - Delete notification
- `PUT /api/v1/notifications/:id/read` - Mark as read
- `GET /api/v1/notifications/unread` - Get unread notifications
- `GET /api/v1/notifications/ws` - WebSocket connection

### Contoh Request

**Login**:

```bash
curl -X POST http://localhost:8080/api/v1/users/signin \
  -H "Content-Type: application/json" \
  -d '{"email": "bintangnugraha.dev@gmail.com", "password": "password"}'
```

**Get Products (dengan authentication)**:

```bash
curl -X GET http://localhost:8080/api/v1/admin/products \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

**Create Order**:

```bash
curl -X POST http://localhost:8080/api/v1/orders?lat=-6.2088&lng=106.8456 \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "items": [
      {"product_id": 1, "quantity": 2, "price": 50000}
    ],
    "shipping_type": "DELIVERY",
    "remarks": "Please deliver in the morning"
  }'
```

## API Documentation

### Response Format

Semua API mengikuti format response standar:

**Success Response**:

```json
{
  "message": "success",
  "data": { ... },
  "pagination": {  // jika applicable
    "page": 1,
    "per_page": 10,
    "total_count": 100,
    "total_pages": 10
  }
}
```

**Error Response**:

```json
{
  "message": "error message",
  "errors": {
    // jika validation error
    "field": ["error message"]
  }
}
```

### Authentication

Semua protected endpoints memerlukan JWT token di header:

```
Authorization: Bearer <token>
```

Token diperoleh dari endpoint `/api/v1/users/signin` dan memiliki expiry time tertentu.

### Rate Limiting

API Gateway mengimplementasikan rate limiting menggunakan Redis. Limit default dapat dikonfigurasi di middleware rate limit.

## Testing

### Load Testing dengan K6

Sistem dilengkapi dengan script load testing menggunakan K6:

**Install K6**:

```bash
# macOS
brew install k6

# Linux
sudo gpg -k
sudo gpg --no-default-keyring --keyring /usr/share/keyrings/k6-archive-keyring.gpg --keyserver hkp://keyserver.ubuntu.com:80 --recv-keys C5AD17C747E3415A3642D57D77C6C491D6AC1D69
echo "deb [signed-by=/usr/share/keyrings/k6-archive-keyring.gpg] https://dl.k6.io/deb stable main" | sudo tee /etc/apt/sources.list.d/k6.list
sudo apt-get update
sudo apt-get install k6
```

**Jalankan Test**:

```bash
# Test product service
cd k6/product_service
k6 run product_home_test.js
k6 run product_shop_test.js
k6 run category_home_test.js

# Test user service
cd k6/user_service
k6 run signup_test.js
```

**Customize Test Parameters**:
Edit file test untuk mengubah `vus` (virtual users) dan `duration`:

```javascript
export const options = {
  vus: 20, // jumlah concurrent users
  duration: "1m", // durasi test
};
```

### Unit Testing

Untuk menjalankan unit test di setiap service (jika tersedia):

```bash
cd user-service
go test ./...
```

## Deployment

### Build Docker Images

Build semua service images:

```bash
cd backend
docker-compose build
```

Build individual service:

```bash
cd backend/user-service
docker build -t user-service:latest .
```

### Production Deployment

1. **Update environment variables** untuk production (database credentials, JWT secret, dll)

2. **Build production images**:

```bash
docker-compose -f docker-compose.yml build --no-cache
```

3. **Deploy dengan Docker Compose**:

```bash
docker-compose -f docker-compose.yml up -d
```

4. **Monitor logs**:

```bash
docker-compose logs -f
```

### Frontend Deployment

**Build production**:

```bash
cd frontend
bun run build
# atau
npm run build
```

**Preview production build**:

```bash
bun run preview
```

**Generate static site** (jika menggunakan static generation):

```bash
bun run generate
```

Output akan berada di folder `.output/` yang dapat di-deploy ke static hosting atau server.

### CI/CD Considerations

Untuk production, pertimbangkan:

- Setup CI/CD pipeline (GitHub Actions, GitLab CI, dll)
- Environment-specific configuration files
- Database migration strategy
- Health check endpoints untuk monitoring
- Log aggregation (ELK stack, Loki, dll)
- Monitoring dan alerting (Prometheus, Grafana)
- Load balancer untuk API Gateway
- Service discovery untuk inter-service communication

## Limitations & Assumptions

### Limitations

1. **Database Per Service**: Setiap service memiliki database terpisah, sehingga transaksi cross-service memerlukan pattern Saga atau eventual consistency
2. **Synchronous Communication**: Beberapa operasi masih menggunakan HTTP synchronous calls antar service, yang dapat menyebabkan latency
3. **Single API Gateway Instance**: API Gateway belum di-scale horizontal, dapat menjadi bottleneck
4. **No Service Discovery**: Service URLs dikonfigurasi secara statis, belum menggunakan service discovery
5. **File Storage**: Upload file masih menggunakan local storage atau Supabase, belum terintegrasi dengan object storage terdistribusi
6. **Limited Error Handling**: Error handling antar service masih basic, belum ada circuit breaker pattern
7. **No Distributed Tracing**: Belum ada implementasi distributed tracing untuk debugging request flow

### Assumptions

1. **Development Environment**: Konfigurasi default menggunakan Docker Compose untuk development
2. **Database Credentials**: Default credentials untuk development (harus diubah untuk production)
3. **JWT Secret**: Secret key harus diubah untuk production dan disimpan dengan aman
4. **Payment Gateway**: Menggunakan Midtrans sandbox untuk development
5. **Network**: Semua service dapat berkomunikasi dalam network yang sama (Docker network atau localhost)
6. **Message Queue**: RabbitMQ digunakan untuk asynchronous processing, dengan asumsi message delivery guarantee
7. **Caching Strategy**: Redis digunakan untuk caching, dengan asumsi data dapat di-refresh jika cache invalid

## Roadmap

### Short Term

- [ ] Implementasi service discovery (Consul atau Eureka)
- [ ] Setup distributed tracing (Jaeger atau Zipkin)
- [ ] Implementasi circuit breaker pattern
- [ ] Enhanced error handling dan retry mechanism
- [ ] API documentation dengan Swagger/OpenAPI
- [ ] Unit test coverage untuk semua services
- [ ] Integration test suite
- [ ] Performance optimization untuk database queries

### Medium Term

- [ ] Implementasi CQRS pattern untuk read-heavy operations
- [ ] Event sourcing untuk audit trail
- [ ] GraphQL API layer sebagai alternatif REST
- [ ] Multi-tenant support
- [ ] Advanced search dengan Elasticsearch
- [ ] Real-time inventory management
- [ ] Recommendation engine untuk produk
- [ ] Analytics dan reporting dashboard

### Long Term

- [ ] Kubernetes deployment configuration
- [ ] Auto-scaling berdasarkan metrics
- [ ] Multi-region deployment
- [ ] CDN integration untuk static assets
- [ ] Advanced monitoring dengan Prometheus dan Grafana
- [ ] Disaster recovery plan
- [ ] Backup dan restore automation
- [ ] Security audit dan penetration testing

## Contributing

### Development Workflow

1. **Fork repository** dan clone ke local
2. **Create feature branch**:

```bash
git checkout -b feature/nama-fitur
```

3. **Follow coding standards**:

   - Go: Ikuti [Effective Go](https://go.dev/doc/effective_go) guidelines
   - Frontend: Ikuti Vue.js style guide dan ESLint rules
   - Commit messages: Gunakan conventional commits format

4. **Test changes**:

   - Jalankan unit tests
   - Test secara manual di local environment
   - Pastikan tidak ada breaking changes

5. **Submit Pull Request**:
   - Deskripsi jelas tentang perubahan
   - Reference issue jika ada
   - Pastikan CI/CD checks pass

### Code Structure Guidelines

- **Backend Services**: Ikuti clean architecture pattern yang sudah ada
- **Frontend**: Gunakan composition API, avoid global state jika memungkinkan
- **Database**: Selalu buat migration untuk schema changes
- **API**: Maintain backward compatibility atau versioning

### Commit Message Format

Gunakan format conventional commits:

```
feat: add new payment method
fix: resolve cart calculation bug
docs: update API documentation
refactor: restructure order service
test: add unit tests for user service
```

## License

Lisensi untuk proyek ini belum ditentukan. Silakan hubungi maintainer untuk informasi lebih lanjut mengenai penggunaan dan distribusi kode ini.

---

**Note**: Dokumentasi ini dibuat berdasarkan analisis source code repository. Untuk informasi terbaru dan detail implementasi spesifik, silakan merujuk ke source code atau hubungi tim development.
