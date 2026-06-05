# House Honter Backend

Repository pembelajaran **Golang** yang berisi materi dasar bahasa Go sekaligus contoh implementasi sederhana pola **Repository – Service** untuk fitur *Category*.

- **Module:** `househonterbackend/v2`
- **Go version:** 1.25.6
- **Dependency utama:** [`testify`](https://github.com/stretchr/testify) untuk unit test & mocking

---

## 🚀 Menjalankan Project

```bash
# Mengunduh dependency
go mod tidy

# Menjalankan file tertentu
go run helper/hellow_world.go

# Menjalankan seluruh test
go test ./...

# Menjalankan test pada package tertentu
go test ./service/...
```

---

## 📂 Struktur Project

### Fitur Utama (Repository – Service Pattern)

| Folder / File | Keterangan |
| --- | --- |
| `Entity/category.go` | Definisi struct `Category` (Id, Name). |
| `repository/category_repository.go` | Interface `CategoryRepository` (kontrak akses data). |
| `repository/category_repository_mock.go` | Mock repository berbasis `testify/mock` untuk testing. |
| `service/category_service.go` | Business logic `CategoryService` (mengambil category by id). |
| `service/category_server_test.go` | Unit test untuk `CategoryService`. |

Alur: `Service` bergantung pada interface `Repository`, sehingga mudah diuji dengan mock tanpa database nyata.

### Materi Dasar Golang

| Folder | Topik |
| --- | --- |
| `variable/` | Deklarasi variabel |
| `type_declaration/` | Type declaration / type alias |
| `number/` | Tipe data angka |
| `boolean/` | Tipe data boolean |
| `string/` | Manipulasi string |
| `array/` | Array |
| `slice/` | Slice |
| `struct/` | Struct |
| `index/` | Indexing |
| `matematika/` | Operasi matematika |
| `operasi_boolean/` | Operasi boolean / logika |
| `perbandingan/` | Operator perbandingan |
| `Macth/` | Pencocokan (switch / match) |
| `flowandflow/` | Control flow (if / for) |
| `errors/` | Error handling |
| `helper/` | Fungsi helper + contoh test |

### Package Bawaan Go

| Folder | Package |
| --- | --- |
| `package_format/` | `fmt` |
| `package_strconv/` | `strconv` |
| `package_string/` | `strings` |

---

## 🧪 Testing

Project ini menggunakan `testify` untuk assertion dan mocking. Contoh file test:

- `helper/hellow_world_test.go`
- `service/category_server_test.go`

```bash
go test -v ./...
```

---

## 📌 Catatan

Repository ini bersifat **pembelajaran**, sehingga sebagian folder berisi eksperimen materi Go satu per satu, sementara fitur *Category* menjadi contoh penerapan arsitektur yang lebih terstruktur.
