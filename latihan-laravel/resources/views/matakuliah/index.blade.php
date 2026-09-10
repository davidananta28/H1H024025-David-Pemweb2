@extends('layouts.app')

@section('judul', 'Daftar Matakuliah')

@section('konten')
    <h1 class="h3 mb-4">Daftar Matakuliah</h1>

    <form method="GET" action="{{ route('matakuliah.index') }}" class="mb-4">
        <input type="text" name="q" value="{{ $q }}" placeholder="Cari matakuliah">
        <button type="submit">Cari</button>
    </form>

    <table class="table table-bordered bg-white">
        <tr>
            <th>Kode</th>
            <th>Nama</th>
            <th>SKS</th>
            <th>Aksi</th>
        </tr>

        @forelse ($matakuliah as $m)
            <tr>
                <td>{{ $m['kode'] }}</td>
                <td>{{ $m['nama'] }}</td>
                <td>
                    <x-badge-sks :sks="$m['sks']" />
                </td>
                <td>
                    <a href="{{ route('matakuliah.show', $m['kode']) }}">
                        Detail
                    </a>
                </td>
            </tr>
        @empty
            <tr>
                <td colspan="4">Data tidak ditemukan.</td>
            </tr>
        @endforelse
    </table>
@endsection