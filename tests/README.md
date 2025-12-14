## Hướng dẫn chạy test

### 1. Chạy toàn bộ test

Từ thư mục gốc của project:

```bash
go test ./...
```

Lệnh này sẽ chạy tất cả các test trong mọi package (bao gồm `internal/...`, `pkg/...`, v.v.).

### 2. Chạy test cho một package cụ thể

Ví dụ chỉ chạy test cho `internal/service`:

```bash
go test ./internal/service
```

### 3. Chạy một (hoặc một nhóm) test theo tên

Ví dụ với các test trong `internal/service/user_service_test.go`:

- `TestGetUserByID_Found`
- `TestGetUserByID_NotFound`

Bạn có thể chạy cả hai bằng:

```bash
go test ./internal/service -run TestGetUserByID
```

Hoặc chỉ một test chính xác:

```bash
go test ./internal/service -run TestGetUserByID_Found
```

### 4. Bật chế độ verbose để xem chi tiết

Thêm cờ `-v` để xem từng test case:

```bash
go test ./internal/service -run TestGetUserByID -v
```

### 5. Gợi ý tích hợp với VS Code

- Cài extension **Go** (Go Team at Google).
- Mở file `*_test.go`, bạn sẽ thấy nút "run test" / "debug test" ngay bên cạnh từng hàm test.


