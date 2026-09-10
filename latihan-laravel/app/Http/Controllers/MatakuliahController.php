<?php

namespace App\Http\Controllers;

use Illuminate\Http\Request;

class MatakuliahController extends Controller
{
    public function index(Request $request)
    {
        $matakuliah = [
            ['kode' => 'TK244001', 'nama' => 'Sistem Mikrokontroler', 'sks' => 3],
            ['kode' => 'TK244002', 'nama' => 'Pemrograman Web II', 'sks' => 2],
            ['kode' => 'TK244003', 'nama' => 'Jaringan Komputer', 'sks' => 3],
            ['kode' => 'TK244005', 'nama' => 'Basis Data', 'sks' => 3],
        ];

        $q = $request->query('q', '');

        if ($q) {
            $matakuliah = array_filter($matakuliah, function ($m) use ($q) {
                return stripos($m['kode'], $q) !== false ||
                    stripos($m['nama'], $q) !== false;
            });
        }

        return view('matakuliah.index', [
            'matakuliah' => $matakuliah,
            'q' => $q
        ]);
    }

    public function show($kode)
    {
        return view('matakuliah.show', [
            'kode' => $kode
        ]);
    }
}
