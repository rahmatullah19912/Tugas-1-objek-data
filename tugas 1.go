package main

import "fmt"

type Identitas struct {
	Nik               int
	Nama              string
	Tempat_Tgl_Lahir  string
	Gol_Darah         string
	Alamat            string
	Jenis_Kelamin     string
	RT_RW             string
	Kel_Desa          string
	Kecamatan         string
	Agama             string
	Status_Perkawinan string
	Pekerjaan         string
	Kewarganegaraan   string
	Berlaku_Hingga    string
}

func main() {

	I := Identitas{
		Nik:               7371112007990007,
		Nama:              "MUH RAHMATULLAH",
		Tempat_Tgl_Lahir:  "UJUNG PANDANG, 20-07-1999",
		Jenis_Kelamin:     "LAKI-LAKI",
		Gol_Darah:         "-",
		Alamat:            "KOMP.YPPKG K.7 NO.20",
		RT_RW:             "007/001",
		Kel_Desa:          "PACCERAKANG",
		Kecamatan:         "BIRING KANAYA",
		Agama:             "ISLAM",
		Status_Perkawinan: "BELUM KAWIN",
		Pekerjaan:         "PELAJAR/MAHASISWA",
		Kewarganegaraan:   "WNI",
		Berlaku_Hingga:    "SEUMUR HIDUP",
	}
	fmt.Println("----------------------------------------")
	fmt.Println("PROVINSI SULAWESI SELATAN KABUPATEN MAKASSAR")
	fmt.Println("----------------------------------------")
	fmt.Println("Nik               :", I.Nik)
	fmt.Println("Nama              :", I.Nama)
	fmt.Println("Tempat_Tgl_Lahir  :", I.Tempat_Tgl_Lahir)
	fmt.Println("Jenis_kelamin     :", I.Jenis_Kelamin)
	fmt.Println("Gol_Darah         :", I.Gol_Darah)
	fmt.Println("Alamat            :", I.Alamat)
	fmt.Println("RT/RW             :", I.RT_RW)
	fmt.Println("Kel/Desa          :", I.Kel_Desa)
	fmt.Println("Kecamatan         :", I.Kecamatan)
	fmt.Println("Agama             :", I.Agama)
	fmt.Println("Status_perkawinan :", I.Status_Perkawinan)
	fmt.Println("Pekerjaan         :", I.Pekerjaan)
	fmt.Println("Kewarganegaraan   :", I.Kewarganegaraan)
	fmt.Println("Berlaku_hingga    :", I.Berlaku_Hingga)

}
