# Articles Backend - Golang

Backend ini dibuat dengan Golang menggunakan Gin + GORM.  
Repo ini bisa langsung dicloning dan dijalankan untuk development.

---

## Cara Menjalankan

1. **Clone repository**

```bash
git clone <repo-url>
cd <repo-folder>
```

2. **Install dependencies**

```bash
go mod tidy
```

3. **Buat file konfigurasi config.yaml**

```bash
service:
  port: ":9876"

database:
  dataSourceName: "username:password@tcp(127.0.0.1:3306)/article?charset=utf8mb4&parseTime=True&loc=Local"
```

4. **Siapkan database**

```bash
CREATE DATABASE article CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;
```

5. **Jalankan server**

```bash
go run cmd/main.go
```
