<?php

namespace Database\Seeders;

use App\Models\User;
use Illuminate\Database\Console\Seeds\WithoutModelEvents;
use Illuminate\Database\Seeder;
use App\Models\Mahasiswa;
use App\Models\Matakuliah;

class DatabaseSeeder extends Seeder
{
    use WithoutModelEvents;

    /**
     * Seed the application's database.
     */
    public function run(): void
    {
        $this->call(ProgramStudiSeeder::class);
        Mahasiswa::factory()->count(30)->create();
        $this->call(MatakuliahSeeder::class);

        $matakuliahIds = Matakuliah::pluck('id');

        Mahasiswa::all()->each(function (Mahasiswa $mahasiswa) use ($matakuliahIds) {
            $diambil = $matakuliahIds->random(min(3, $matakuliahIds->count()));
            foreach ($diambil as $id) {
                $mahasiswa->matakuliah()->attach($id, [
                    'nilai' => fake()->randomFloat(2, 60, 100),
                ]);
            }
        });
    }
}
