## Hướng dẫn chạy migration

Project này đang dùng các file `.sql` trong thư mục `migrations/` (ví dụ `1_create_table_up.sql`, `2_create_users_roles_up.sql`) để quản lý schema database.

Bạn có thể dùng một trong các tool phổ biến sau để apply migration.

---

### 1. Chuẩn bị MySQL

1. Chạy MySQL bằng Docker:

```bash
docker compose up -d mysql
```

2. Thông tin kết nối (trùng với `configs/local.yaml`):

- **Host**: `localhost`
- **Port**: `33306`
- **User**: `root`
- **Password**: `root`
- **Database**: `shopDev`

---

### 2. Dùng `goose` (khuyến nghị)

Tool chính mà project này hướng tới là `pressly/goose`, vì hỗ trợ tốt SQL thuần và dễ dùng với MySQL.

#### 2.1. Cài đặt CLI

- Repo: `github.com/pressly/goose/v3`

```bash
go install github.com/pressly/goose/v3/cmd/goose@latest
```

> Lưu ý: đảm bảo `$GOPATH/bin` nằm trong `PATH` để chạy được lệnh `goose`.

#### 2.2. Chạy migration `up`

Đứng ở root project:

```bash
cd /Users/kietle/Projects/go-ecommerce-backend-api

goose -dir ./migrations \
  mysql "root:root@tcp(localhost:33306)/shopDev" up
```

#### 2.3. Rollback / status

- Rollback 1 step:

```bash
goose -dir ./migrations \
  mysql "root:root@tcp(localhost:33306)/shopDev" down
```

- Xem trạng thái các migration:

```bash
goose -dir ./migrations \
  mysql "root:root@tcp(localhost:33306)/shopDev" status
```

---

### 3. Gợi ý dùng `golang-migrate` (tuỳ chọn)

Nếu bạn thích `golang-migrate` hơn, có thể tham khảo thêm:

- Repo: `github.com/golang-migrate/migrate`
- Cài CLI:

```bash
go install github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

- Ví dụ lệnh (đứng ở root project):

```bash
migrate -source "file://migrations" \
  -database "mysql://root:root@tcp(localhost:33306)/shopDev" up
```

Tuy nhiên, trong project này, `goose` là lựa chọn mặc định được khuyến nghị sử dụng hằng ngày.
