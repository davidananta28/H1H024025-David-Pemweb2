<?php

namespace App\Http\Controllers;

use Illuminate\Http\Request;

class MahasiswaController extends Controller
{
    public function index()
    {
        $daftarMahasiswa = [
            ['nim' => 'H1H024001', 'nama' => 'Bahtiar', 'angkatan' => 2020],
            ['nim' => 'H1H024002', 'nama' => 'Wisnu', 'angkatan' => 2020],
            ['nim' => 'H1H024003', 'nama' => 'Pandu', 'angkatan' => 2020],
        ];

        return view('mahasiswa.index', ['daftarMahasiswa' => $daftarMahasiswa]);
    }
    public function show(string $nim)
    {
        return view('mahasiswa.show', ['nim' => $nim]);
    }
    public function cari(Request $request)
    {
        $kataKunci = $request->query('q', '');

        return response()->json([
            'kata_kunci' => $kataKunci,
            'metode' => $request->method(),
            'path' => $request->path(),
        ]);
    }
}
