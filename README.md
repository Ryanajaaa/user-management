
## 🚀 Cara Menjalankan

1. **Pastikan sudah install Go** (versi terbaru lebih baik).
2. **Clone repo ini:**
   ```sh
   git clone https://github.com/username/repo-name.git
   cd repo-name
   ```
3. **Jalankan aplikasi:**
   ```sh
   go run main.go
   ```
4. **API siap digunakan!** Secara default berjalan di `http://localhost:8080`

---

## 📂 Struktur Folder

```
/user-management
 ├── delivery/http        # Controller (handler API)
 ├── infrastructure       # Repository (penyimpanan data)
 ├── internal/domain      # Model data
 ├── internal/usecase     # Logika bisnis
 ├── main.go              # Entry point aplikasi
```

---

## 🔥 Daftar Endpoint

### 1️⃣ Register User
**Endpoint:**
```http
POST /api/register
```
**Request Body:**
```json
{
  "name": "John Doe",
  "email": "johndoe@example.com"
}
```
**Response:**
```json
{
  "message": "User registered",
  "user": {
    "id": 1,
    "name": "John Doe",
    "email": "johndoe@example.com"
  }
}
```

---

### 2️⃣ Login
**Endpoint:**
```http
POST /api/login
```
**Request Body:**
```json
{
  "email": "johndoe@example.com",
  "password": "password123"
}
```
**Response:**
```json
{
  "message": "Login successful",
  "token": "JWT-TOKEN-HERE"
}
```

---

### 3️⃣ Get Profile (Butuh Token)
**Endpoint:**
```http
GET /api/profile
```
**Header:**
```http
Authorization: Bearer JWT-TOKEN-HERE
```
**Response:**
```json
{
  "user": {
    "id": 1,
    "name": "John Doe",
    "email": "johndoe@example.com"
  }
}
```

---

### 4️⃣ Logout
**Endpoint:**
```http
POST /api/logout
```
**Header:**
```http
Authorization: Bearer JWT-TOKEN-HERE
```
**Response:**
```json
{
  "message": "Logout successful"
}
```
