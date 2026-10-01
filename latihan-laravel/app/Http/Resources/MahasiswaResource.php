<?php

namespace App\Http\Resources;

use Illuminate\Http\Request;
use Illuminate\Http\Resources\Json\JsonResource;

class MahasiswaResource extends JsonResource
{
    public function toArray(Request $request): array
    {
        $fields = $request->query('fields');

        if ($fields) {
            $fields = array_map('trim', explode(',', $fields));

            $allowedFields = [
                'id',
                'nim',
                'nama',
                'email',
                'angkatan',
                'ipk',
                'aktif',
                'program_studi',
                'dibuat_pada',
            ];

            $fields = array_intersect($fields, $allowedFields);

            $data = [];

            foreach ($fields as $field) {
                if ($field === 'program_studi') {
                    $data['program_studi'] = $this->whenLoaded(
                        'programStudi',
                        function () {
                            return [
                                'id' => $this->programStudi->id,
                                'kode' => $this->programStudi->kode,
                                'nama' => $this->programStudi->nama,
                            ];
                        }
                    );
                } elseif ($field === 'dibuat_pada') {
                    $data['dibuat_pada'] = $this->created_at?->toIso8601String();
                } else {
                    $data[$field] = $this->{$field};
                }
            }

            return $data;
        }

        return [
            'id' => $this->id,
            'nim' => $this->nim,
            'nama' => $this->nama,
            'email' => $this->email,
            'angkatan' => $this->angkatan,
            'ipk' => (float) $this->ipk,
            'aktif' => $this->aktif,
            'program_studi' => $this->whenLoaded('programStudi', function () {
                return [
                    'id' => $this->programStudi->id,
                    'kode' => $this->programStudi->kode,
                    'nama' => $this->programStudi->nama,
                ];
            }),
            'dibuat_pada' => $this->created_at?->toIso8601String(),
        ];
    }
}
