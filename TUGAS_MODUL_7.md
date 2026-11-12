# LAPORAN TUGAS PRAKTIKUM MODUL 7
## RESTful API CRUD dengan Fiber v3 dan GORM
**Mata Kuliah:** Pemrograman Web II  
**Program Studi:** S1 Teknik Komputer, Universitas Jenderal Soedirman  
**Nama:** David Ananta  
**NIM:** H1H024025  

---

## E. Penyelesaian Tugas Praktikum

### 1. Resource Matakuliah dengan CRUD Lengkap
Mengimplementasikan model, repository, handler, dan routing untuk resource `matakuliah` dengan arsitektur pemisahan lapisan (Clean Architecture/Layered Architecture):

- **Model:** `internal/model/matakuliah.go`
  Mendefinisikan struct `Matakuliah` dengan field `ID`, `Kode` (unique), `Nama`, `Sks`, `Semester`, `CreatedAt`, dan `UpdatedAt`, serta mapping tabel ke `matakuliahs`.
- **Repository:** `internal/repository/matakuliah_repository.go`
  - `AmbilMatakuliah(filter FilterMatakuliah)`: Mendukung pencarian (`cari`), filter semester, pengurutan, dan paginasi (`halaman`, `per_halaman`).
  - `AmbilMatakuliahPerId(id uint)`: Mengambil data matakuliah berdasarkan ID.
  - `SimpanMatakuliah(matakuliah *model.Matakuliah)`: Menyimpan data baru.
  - `PerbaruiMatakuliah(matakuliah *model.Matakuliah)`: Memperbarui data matakuliah.
  - `HapusMatakuliah(matakuliah *model.Matakuliah)`: Menghapus data matakuliah.
- **DTO:** `internal/handler/dto.go`
  - `PermintaanMatakuliah`: Menampung payload request JSON.
- **Handler:** `internal/handler/matakuliah_handler.go`
  - `DaftarMatakuliah`: Menangani `GET /api/matakuliah` dengan respons metadata paginasi.
  - `DetailMatakuliah`: Menangani `GET /api/matakuliah/:id`.
  - `BuatMatakuliah`: Menangani `POST /api/matakuliah` dengan validasi input (HTTP 201).
  - `PerbaruiMatakuliah`: Menangani `PUT /api/matakuliah/:id`.
  - `HapusMatakuliah`: Menangani `DELETE /api/matakuliah/:id`.
- **Router:** `internal/router/router.go`
  Mendaftarkan rute grup `/api/matakuliah`.

---

### 2. Endpoint GET /api/program-studi Beserta Jumlah Mahasiswa
- **Repository:** `internal/repository/program_studi_repository.go`
  Menggunakan kueri SQL agregasi dengan `LEFT JOIN` dan `COUNT` yang memperhitungkan soft delete (`AND mahasiswas.deleted_at IS NULL`):
  ```sql
  SELECT program_studis.id, program_studis.kode, program_studis.nama, program_studis.jenjang, 
         program_studis.created_at, program_studis.updated_at, 
         COUNT(mahasiswas.id) as jumlah_mahasiswa 
  FROM program_studis 
  LEFT JOIN mahasiswas ON mahasiswas.program_studi_id = program_studis.id AND mahasiswas.deleted_at IS NULL 
  GROUP BY program_studis.id ...
  ```
- **Handler & Router:** `internal/handler/program_studi_handler.go` & `internal/router/router.go`
  - Endpoint `GET /api/program-studi`: Mengembalikan seluruh daftar program studi beserta field `jumlah_mahasiswa`.
  - Endpoint `GET /api/program-studi/:id`: Mengembalikan detail satu program studi beserta `jumlah_mahasiswa`.

---

### 3. Penerapan Soft Delete pada Mahasiswa dan Endpoint Pemulihan Data
- **Model:** `internal/model/mahasiswa.go`
  Menambahkan field:
  ```go
  DeletedAt gorm.DeletedAt `gorm:"index" json:"dihapus_pada,omitempty"`
  ```
- **Mekanisme Soft Delete:**
  - Saat `DELETE /api/mahasiswa/:id` dipanggil, GORM secara otomatis menjalankan kueri:
    `UPDATE mahasiswas SET deleted_at = 'YYYY-MM-DD HH:MM:SS' WHERE id = ?`
  - Kueri umum (`Find`, `First`) otomatis mengecualikan data yang sudah dihapus (`WHERE deleted_at IS NULL`).
- **Endpoint Pemulihan (Restore):**
  - Mendaftarkan endpoint `POST /api/mahasiswa/:id/pulihkan` dan `POST /api/mahasiswa/:id/restore` (juga mendukung `PUT`).
  - Repository `Pulihkan(id uint)` menggunakan kueri unscoped:
    ```go
    config.DB.Unscoped().Model(&model.Mahasiswa{}).Where("id = ?", id).Update("deleted_at", nil)
    ```
  - Tambahan endpoint penunjang:
    `GET /api/mahasiswa/terhapus` atau `GET /api/mahasiswa?terhapus=true` untuk melihat daftar mahasiswa yang berada dalam status terhapus.

---

### 4. Perbandingan Jumlah Baris Kode (Laravel vs Fiber) dan Analisis Perbedaan

#### Tabel Perbandingan Baris Kode (LoC - Lines of Code)

| Komponen Arsitektur | Laravel 11/13 (PHP) | Fiber v3 + GORM (Go) | Keterangan |
| :--- | :---: | :---: | :--- |
| **Model** | 45 baris (`Mahasiswa.php`) | 29 baris (`mahasiswa.go`) | Model Go lebih ringkas karena hanya struct + tag. |
| **Data Validation / DTO** | 72 baris (2 Form Requests) | 11 baris (`dto.go`) | Laravel memisahkan Store & Update request class. |
| **Transformasi / Resource** | 73 baris (`MahasiswaResource.php`) | 0 baris (Otomatis via JSON tags) | Go menggunakan struct tag `json:"..."`. |
| **Repository (Akses DB)** | 0 baris (Active Record bawaan) | 89 baris (`mahasiswa_repository.go`) | Fiber menerapkan pemisahan lapisan repository. |
| **Controller / Handler** | 95 baris (`MahasiswaController.php`) | 230 baris (`mahasiswa_handler.go`) | Handler Go mengelola type casting, pointer, dan error handling secara manual. |
| **Routing** | 1 baris (`Route::apiResource(...)`) | 12 baris (Registrasi manual per rute) | Laravel menyediakan resource macro 1 baris. |
| **Total LoC (Mahasiswa)** | **~286 baris** | **~371 baris** | **Fiber membutuhkan ~30% baris kode lebih banyak.** |

#### Analisis Penyebab Perbedaan

1. **Active Record (Laravel) vs Pemisahan Lapisan / Repository Pattern (Fiber/Go):**
   - Eloquent pada Laravel menggunakan pola *Active Record*, di mana class Model sekaligus berfungsi sebagai query builder (`Mahasiswa::create(...)`, `$mahasiswa->delete()`). Ini memangkas kebutuhan membuat lapisan repository terpisah untuk CRUD dasar.
   - Pada implementasi Fiber (Modul 7), arsitektur dibagi menjadi *Handler -> Repository -> Model*, sehingga kueri database dipisahkan ke dalam fungsi-fungsi repository tersendiri.

2. **Static Typing & Explicit Error Handling (Go) vs Dynamic Reflection (PHP/Laravel):**
   - Di Go, setiap kemungkinan error harus ditangani secara eksplisit (`if err != nil`). Konversi string parameter HTTP ke integer harus dilakukan manual (`strconv.Atoi`) disertai pengecekan validitas.
   - Penanganan field opsional atau nilai `false` memerlukan manipulasi pointer (`*bool`, dereferencing).
   - Laravel memanfaatkan *Route Model Binding* yang secara otomatis mencocokkan ID di URL ke instans Model di basis data dan melempar `404 Not Found` secara otomatis tanpa kode tambahan.

3. **Routing Abstraction:**
   - Laravel memiliki helper `Route::apiResource()`, yang dalam 1 baris langsung mendaftarkan 5 endpoint standar RESTful (`index`, `store`, `show`, `update`, `destroy`).
   - Fiber mengharuskan setiap rute dan metode HTTP (`Get`, `Post`, `Put`, `Delete`) didaftarkan secara manual pada router group.

4. **Karakteristik & Trade-off Performa:**
   - Meskipun Fiber (Go) membutuhkan baris kode yang lebih banyak dan lebih verbose, kode Go memberikan keuntungan besar dalam hal:
     - **Kecepatan eksekusi (throughput requests per second)** yang jauh melampaui PHP/Laravel.
     - **Konsumsi memori (RAM) yang jauh lebih hemat**.
     - **Static type safety** saat kompilasi (mencegah bug runtime).
     - **Single binary executable** mandiri tanpa dependensi web server eksternal (seperti Nginx/PHP-FPM).

---

## F. Jawaban Pertanyaan Pembahasan

### 1. Mengapa field Aktif pada DTO menggunakan pointer (`*bool`) sedangkan field lain tidak?
**Jawaban:**
Di bahasa Go, setiap tipe data primitif memiliki *zero value* bawaan saat variabel tidak diisi:
- Nilai zero value untuk `bool` adalah `false`.
- Nilai zero value untuk `string` adalah `""`.
- Nilai zero value untuk `int` adalah `0`.

Jika field `Aktif` didefinisikan sebagai `bool` biasa (non-pointer), ketika klien mengirim JSON tanpa menyertakan field `aktif` (misalnya pada pembaruan parsial `PUT` atau saat field bersifat opsional), deserializer JSON di Go akan otomatis mengisinya dengan `false`. Akibatnya, server tidak dapat membedakan apakah klien memang sengaja menonaktifkan mahasiswa (`"aktif": false`) atau klien tidak mengirimkan field tersebut sama sekali.

Dengan menggunakan pointer `*bool`:
- Jika klien **tidak mengirim** field `aktif`, nilai pointer akan bernilai `nil`.
- Jika klien **mengirimkan** `"aktif": true`, pointer akan menunjuk ke nilai `true`.
- Jika klien **mengirimkan** `"aktif": false`, pointer akan menunjuk ke nilai `false`.
Hal ini memungkinkan handler membedakan antara nilai yang tidak diubah (`nil`) dan nilai yang sengaja diisi `false`.

### 2. Jelaskan perbedaan Save, Update, dan Updates pada GORM.
**Jawaban:**
- **`Save`**:
  Berfungsi menyimpan seluruh field dari struct model.
  - Jika struct belum memiliki primary key (atau bernilai 0), `Save` akan menjalankan kueri `INSERT`.
  - Jika struct sudah memiliki primary key, `Save` akan menjalankan kueri `UPDATE` yang memperbarui **seluruh kolom** pada tabel, termasuk kolom yang memiliki zero value.
- **`Update`**:
  Berfungsi memperbarui **satu kolom tunggal** tertentu pada baris yang dipilih. Method ini tetap memperbarui kolom tersebut meskipun nilai yang dimasukkan adalah zero value (misalnya: `db.Model(&m).Update("aktif", false)`).
- **`Updates`**:
  Berfungsi memperbarui **banyak kolom sekaligus**.
  - Jika dipanggil dengan parameter berupa `struct`, GORM secara default hanya akan memperbarui kolom yang nilainya **bukan zero value** (kolom bernilai `0`, `""`, atau `false` akan diabaikan).
  - Jika dipanggil dengan parameter berupa `map[string]interface{}`, GORM akan memperbarui seluruh kolom yang tercantum dalam map, termasuk yang bernilai zero value.

### 3. Apa risiko penggunaan AutoMigrate pada lingkungan produksi?
**Jawaban:**
Method `AutoMigrate` pada GORM dirancang untuk kemudahan dan kenyamanan pada tahap pengembangan (*development*), namun berisiko tinggi jika digunakan langsung pada lingkungan produksi (*production*):
1. **Tidak Menghapus Kolom Usang (*No Column Dropping*):**
   `AutoMigrate` hanya membuat tabel baru, menambah kolom baru, atau menambah indeks. AutoMigrate tidak akan pernah menghapus kolom yang sudah tidak dipakai atau mengubah tipe data kolom secara otomatis untuk mencegah kehilangan data. Akibatnya, skema basis data produksi bisa menjadi kotor dan tidak sinkron dengan codebase.
2. **Tidak Ada Riwayat Versi & Rollback:**
   Berbeda dengan alat migrasi berbasis berkas (seperti Laravel Migration atau golang-migrate), `AutoMigrate` tidak memiliki versioning ataupun mekanisme `down()` / rollback jika terjadi kegagalan deployment.
3. **Risiko Concurrency Lock / Deadlock pada Multi-Instance:**
   Pada lingkungan produksi modern dengan banyak replika container/pod yang menyala bersamaan, eksekusi `AutoMigrate` secara simultan saat aplikasi boot dapat memicu *table locking* atau *deadlock* pada DDL database.

### 4. Bandingkan mekanisme eager loading pada GORM dengan `with()` pada Eloquent.
**Jawaban:**
- **Mekanisme pada GORM (`Preload`):**
  - Menggunakan method `.Preload("Relasi")` (contoh: `config.DB.Preload("ProgramStudi").Find(&daftar)`).
  - GORM akan mengeksekusi kueri utama untuk mengambil data parent, lalu mengumpulkan foreign key, dan mengeksekusi kueri kedua secara otomatis (misalnya `SELECT * FROM program_studis WHERE id IN (...)`) kemudian memetakan hasilnya ke struct relasi Go.
  - GORM juga mendukung eager loading berbasis join tunggal melalui `.Joins("ProgramStudi")` khusus untuk relasi 1-to-1 (belongs to).
- **Mekanisme pada Eloquent (`with()`):**
  - Menggunakan method `::with('relasi')` (contoh: `Mahasiswa::with('programStudi')->get()`).
  - Eloquent menjalankan kueri parent terlebih dahulu, kemudian menjalankan kueri `WHERE IN` untuk model terkait, lalu mencocokkannya ke dalam properti model PHP di memori.
- **Perbandingan:**
  - Keduanya sama-sama memecahkan masalah **N+1 Query Problem** dengan menjalankan kueri agregasi berbasis `IN (...)`.
  - Pada Eloquent, relasi didefinisikan secara dinamis melalui method PHP (`belongsTo`, `hasMany`).
  - Pada GORM, relasi didefinisikan secara deklaratif menggunakan tipe struct pointer dan GORM struct tag (`gorm:"foreignKey:ProgramStudiID"`), yang diperiksa saat inisialisasi skema.
