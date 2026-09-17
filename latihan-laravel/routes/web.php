<?php

use Illuminate\Support\Facades\Route;
use App\Http\Controllers\MahasiswaWebController;

Route::get('/', function () {
    return view('welcome');
});

Route::get('/salam', function () {
    return "Selamat datang di Pemrograman Web II";
});

Route::get('/mahasiswa/{nim}', function (string $nim) {
    return 'Data mahasiswa dengan NIM ' . $nim;
});

Route::get('/matakuliah/{kode?}', function (?string $kode = null) {
    if ($kode === null) {
        return 'Menampilkan seluruh matakuliah';
    }
    return 'Menampilkan matakuliah kode ' . $kode;
});

Route::get('/semester/{angka}', function (int $angka) {
    return 'Semester ke ' . $angka;
})->whereNumber('angka');

use App\Http\Controllers\MahasiswaController;

Route::get('/data-mahasiswa', [MahasiswaController::class, 'index'])->name('mahasiswa.index');
Route::get('/data-mahasiswa/{nim}', [MahasiswaController::class, 'show'])->name('mahasiswa.show');
Route::get('/cari-mahasiswa', [MahasiswaController::class, 'cari']);

use App\Http\Controllers\MatakuliahController;

Route::get('/data-matakuliah', [MatakuliahController::class, 'index'])->name('matakuliah.index');

Route::get('/data-matakuliah/{kode}', [MatakuliahController::class, 'show'])->name('matakuliah.show');

Route::get('/mahasiswa-data', [
    MahasiswaWebController::class,
    'index'
])->name('mahasiswa.data');
