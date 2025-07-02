# se4458-go-auth-service

Go ile yazılmış, GORM tabanlı, MySQL destekli, JWT'siz (sadece temel auth) bir kullanıcı kimlik doğrulama servisidir. Swagger/OpenAPI ile dökümante edilmiştir.

## Özellikler
- Kullanıcı kayıt (register)
- Kullanıcı girişi (login)
- Şifreler bcrypt ile hashlenir
- MySQL desteği (GORM ile)
- Swagger UI ile API dökümantasyonu
- .env ile güvenli veritabanı bağlantısı

## Kurulum

### 1. Bağımlılıkları Yükle
```sh
go mod tidy
```

### 2. .env Dosyasını Oluştur
Proje köküne bir `.env` dosyası ekleyin:
```env
DB_USER=root
DB_PASS=parola
DB_HOST=127.0.0.1
DB_PORT=3306
DB_NAME=veritabani
```

### 3. MySQL'de Veritabanı Oluştur
```sql
CREATE DATABASE veritabani;
```
> Tabloyu elle oluşturmanıza gerek yok, uygulama ilk çalıştığında otomatik oluşur.

### 4. Swagger CLI Yükle (ilk kez kurulum için)
```sh
go install github.com/swaggo/swag/cmd/swag@latest
```

### 5. Swagger Dokümantasyonunu Oluştur
```sh
swag init -g cmd/main.go
```

### 6. Uygulamayı Başlat
```sh
go run cmd/main.go
```

## API Dökümantasyonu

Swagger UI'ya tarayıcınızdan erişin:
```
http://localhost:8080/swagger/index.html
```

## Temel Endpointler

- `POST /register` — Kullanıcı kaydı
- `POST /login` — Kullanıcı girişi