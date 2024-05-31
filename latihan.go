package main

import (
	"fmt"
	"math/rand"
)

type User struct {
	nama, NIM, password, pin string
	point, highscore         int
}

type Soal struct {
	pertanyaan   string
	pilihan      [4]string
	kunciJawaban string
	benar, salah int
}

type PlayerData [1000]User
type DataSoal [1000]Soal

var soal DataSoal
var player PlayerData
var totalP, totalS int = 1, 0

//======================//
//		FUNC MAIN		//
//======================//

func main() {
	player[0] = User{
		nama:      "admin",
		NIM:       "admin",
		password:  "admin",
		pin:       "admin",
		highscore: -100,
	}
	menuAwal()
}

//==================================//
//			  Menu Awal				//
//==================================//

func menuAwal() {
	var pilih int

	fmt.Println("|====================|")
	fmt.Println("|        MENU        |")
	fmt.Println("| 1. Login           |")
	fmt.Println("| 2. Buat akun       |")
	fmt.Println("| 3. Lupa password   |")
	fmt.Println("| 4. Keluar          |")
	fmt.Println("|====================|")

	fmt.Print("Pilihan (1/2/3/4): ")
	for {
		fmt.Scan(&pilih)
		if pilih == 1 || pilih == 2 || pilih == 3 || pilih == 4 {
			break
		}
	}

	switch pilih {
	case 1:
		loginNIM(&player)
	case 2:
		buatAkun(&player, &totalP)
		menuAwal()
	case 3:
		gantiPW(&player, totalP)
		menuAwal()
	case 4:
	}
}

//==================================//
//		Procedure Buat Akun			//
//==================================//

func buatAkun(P *PlayerData, totalP *int) {
	var i int
	var nim, pilih string
	i = *totalP

	fmt.Print("|Masukkan nama: ")
	fmt.Scan(&P[i].nama)

	for {
		fmt.Print("|Masukkan NIM: ")
		fmt.Scan(&nim)

		if searchPNIM(*P, *totalP, nim) == -1 {
			P[i].NIM = nim
			break
		}
		fmt.Println("NIM sudah terdaftar")
	}

	fmt.Print("|Masukkan password: ")
	fmt.Scan(&P[i].password)

	fmt.Print("Pin keamanan? (Y/N): ")
	for {
		fmt.Scan(&pilih)
		if pilih == "Y" || pilih == "y" {
			fmt.Scan(&P[i].pin)
			break
		} else if pilih == "N" || pilih == "n" {
			break
		}
	}

	fmt.Print("|Akun berhasil dibuat \n")
	*totalP++
}

//==================================//
//		Procedure Login(NIM)		//
//==================================//

func loginNIM(P *PlayerData) {
	var found, nimLimit int
	var nim string
	var login bool = false

	for nimLimit != 3 && !login {
		fmt.Print("|NIM: ")
		fmt.Scan(&nim)
		found = searchPNIM(*P, totalP, nim)
		if found == -1 {
			fmt.Println("|NIM tidak terdaftar")
			nimLimit++
		} else {
			login = true
			loginPW(P, found)
		}
	}
	if nimLimit == 3 {
		fmt.Println("|NIM salah 3 kali berturut")
		menuAwal()
	}
}

//==================================//
//		Procedure Login(Password)	//
//==================================//

func loginPW(P *PlayerData, found int) {
	var pwLimit int
	var pw string
	var login bool = false

	for pwLimit != 3 && !login {
		fmt.Print("|Password: ")
		fmt.Scan(&pw)
		//login = false

		if found == 0 && pw == P[found].password {
			menuAdmin()
			login = true
		} else if pw == P[found].password {
			fmt.Println("Login Berhasil")
			login = true
			menuKuis(P, found)
		} else {
			fmt.Println("Password salah")
			pwLimit++
		}
	}
	if pwLimit == 3 {
		fmt.Println("|Password salah 3 kali berturut")
		menuAwal()
	}
}

//==================================//
//	   Procedure Ubah Password		//
//==================================//

func gantiPW(P *PlayerData, totalP int) {
	var found int
	var nim, pin string
	fmt.Print("|NIM: ")
	fmt.Scan(&nim)
	found = searchPNIM(*P, totalP, nim)
	if found == -1 {
		fmt.Println("|NIM tidak terdaftar")
		gantiPW(P, totalP)
	} else if P[found].pin == "" {
		fmt.Println("Akun tidak memiliki pin")
	} else {
		fmt.Println("Masukkan pin")
		fmt.Scan(&pin)
		if pin == P[found].pin {
			fmt.Print("Silahkan ubah password: ")
			fmt.Scan(&P[found].password)
			fmt.Println(P[found].password)
		} else {
			fmt.Println("Pin salah")
		}
	}
}

//==================================//
//			  Menu Admin			//
//==================================//

func menuAdmin() {
	var pilih int

	fmt.Println("|====================|")
	fmt.Println("|     MENU ADMIN     |")
	fmt.Println("| 1. Tambah soal     |")
	fmt.Println("| 2. Ubah soal       |")
	fmt.Println("| 3. Hapus soal      |")
	fmt.Println("| 4. Daftar soal     |")
	fmt.Println("| 5. Info Soal       |")
	fmt.Println("| 6. Kembali         |")
	fmt.Println("|====================|")

	fmt.Print("Pilihan (1/2/3/4/5/6): ")
	for {
		fmt.Scan(&pilih)
		if pilih == 1 || pilih == 2 || pilih == 3 || pilih == 4 || pilih == 5 || pilih == 6 {
			break
		}
	}
	switch pilih {
	case 1:
		tambahSoal(&soal, &totalS, totalS)
		menuAdmin()
	case 2:
		ubahSoal(&soal, totalS)
		menuAdmin()
	case 3:
		hapusSoal(&soal, &totalS)
		menuAdmin()
	case 4:
		listSoal(totalS)
		menuAdmin()
	case 5:
		infoSoal(soal, totalS)
		menuAdmin()
	case 6:
		menuAwal()
	}
}

//==================================//
//		Procedure Tambah Soal		//
//==================================//

func tambahSoal(S *DataSoal, totalS *int, no int) {
	var i, bSoal int
	var exit string

	fmt.Print("Masukkan banyak soal yang ingin ditambah: ")
	fmt.Scan(&bSoal)

	for i = no; i < no+bSoal; i++ {
		fmt.Print("|Masukkan pertanyaan: ")
		fmt.Scan(&S[i].pertanyaan)
		fmt.Println("|Masukkan pilihan jawaban:")
		fmt.Print("|a. ")
		fmt.Scan(&S[i].pilihan[0])
		fmt.Print("|b. ")
		fmt.Scan(&S[i].pilihan[1])
		fmt.Print("|c. ")
		fmt.Scan(&S[i].pilihan[2])
		fmt.Print("|d. ")
		fmt.Scan(&S[i].pilihan[3])
		fmt.Print("|Masukkan kunci jawaban: ")

		for {
			fmt.Scan(&S[i].kunciJawaban)
			exit = S[i].kunciJawaban
			if exit == "a" || exit == "b" || exit == "c" || exit == "d" {
				break
			}
		}

		fmt.Print("|Pertanyaan berhasil dibuat \n")
		*totalS++
	}
}

//==================================//
//		 Procedure Ubah Soal		//
//==================================//

func ubahSoal(S *DataSoal, totalS int) {
	var no int

	fmt.Println("|==================")
	fmt.Print("|Nomor soal yang mau diubah: ")
	fmt.Scan(&no)
	tambahSoal(S, &totalS, no-1)
}

//==================================//
//		Procedure Hapus Soal		//
//==================================//

func hapusSoal(S *DataSoal, totalS *int) {
	var i, no int

	fmt.Print("|Nomor soal yang mau dihapus: ")
	fmt.Scan(&no)
	for i = no - 1; i < *totalS; i++ {
		S[i].pertanyaan = S[i+1].pertanyaan
		S[i].pilihan[0] = S[i+1].pilihan[0]
		S[i].pilihan[1] = S[i+1].pilihan[1]
		S[i].pilihan[2] = S[i+1].pilihan[2]
		S[i].pilihan[3] = S[i+1].pilihan[3]
		S[i].kunciJawaban = S[i+1].kunciJawaban
	}
	fmt.Printf("Soal no %d berhasil dihapus", no)
	*totalS--
}

//==================================//
//		Procedure Daftar Soal		//
//==================================//

func listSoal(totalS int) {
	for i := 0; i < totalS; i++ {
		fmt.Print(i+1, ") ", soal[i].pertanyaan, "\n")
		fmt.Print("a.", soal[i].pilihan[0], "                ")
		fmt.Print("b.", soal[i].pilihan[1], "\n")
		fmt.Print("c.", soal[i].pilihan[2], "                ")
		fmt.Print("d.", soal[i].pilihan[3], "\n")
		fmt.Println("Jawabannya adalah", soal[i].kunciJawaban)
		fmt.Println()
	}
	fmt.Print("Enter untuk kembali... ")
	fmt.Scanln()
}

//==================================//
//		Procedure Info Soal			//
//==================================//

func infoSoal(S DataSoal, n int) {
	var i int
	fmt.Println("|Soal Termudah: ")
	selectionSortB(&S, n)
	for i = 0; i < 5; i++ {
		fmt.Printf("|%d. %s\nJumlah benar = %d\n", i+1, S[i].pertanyaan, S[i].benar)
	}
	fmt.Println("|Soal Tersusah: ")
	selectionSortS(&S, n)
	for i = 0; i < 5; i++ {
		fmt.Printf("|%d. %s\nJumlah salah = %d\n", i+1, S[i].pertanyaan, S[i].salah)
	}

	fmt.Print("Enter untuk kembali... ")
	fmt.Scanln()
}

//==================================//
//		      Menu Kuis				//
//==================================//

func menuKuis(P *PlayerData, found int) {
	var pilih int
	fmt.Println("|====================|")
	fmt.Println("|     MENU KUIS      |")
	fmt.Println("| 1. Mulai kuis      |")
	fmt.Println("| 2. Highscore       |")
	fmt.Println("| 3. Peringkat       |")
	fmt.Println("| 4. Kembali         |")
	fmt.Println("|====================|")

	fmt.Print("Pilihan (1/2/3/4): ")
	for {
		fmt.Scan(&pilih)
		if pilih == 1 || pilih == 2 || pilih == 3 || pilih == 4 {
			break
		}
	}
	switch pilih {
	case 1:
		mulaiKuis(P, &soal, totalS, found)
		menuKuis(P, found)
	case 2:
		highscore(*P, found)
		menuKuis(P, found)
	case 3:
		peringkat(*P)
		menuKuis(P, found)
	case 4:
		menuAwal()
	}
}

//==================================//
//		Procedure Mulai Kuis		//
//==================================//

func mulaiKuis(P *PlayerData, S *DataSoal, totalS, found int) {
	var i, poin int
	var jawab string
	P[found].point = 0

	randomize(S)

	if totalS > 10 {
		totalS = 10
	}

	if totalS == 0 {
		fmt.Println("Belum ada pertanyaan")
	} else {
		poin = 100 / totalS
		for i = 0; i < totalS; i++ {
			fmt.Printf("%d. %s?\n", i+1, S[i].pertanyaan)
			fmt.Println("a.", S[i].pilihan[0])
			fmt.Println("b.", S[i].pilihan[1])
			fmt.Println("c.", S[i].pilihan[2])
			fmt.Println("d.", S[i].pilihan[3])
			fmt.Print("Jawab: ")
			fmt.Scan(&jawab)

			if jawab == S[i].kunciJawaban {
				fmt.Println("Jawaban Benar")
				fmt.Printf("Point +%d\n\n", poin)
				P[found].point += poin
				S[i].benar++
			} else {
				fmt.Println("Jawaban Salah")
				fmt.Println("Point +0\n")
				S[i].salah++
			}
		}
		fmt.Println("Score akhir:", P[found].point)
	}

	if P[found].highscore < P[found].point {
		updateHigh(P, found)
	}

	fmt.Print("Enter untuk kembali... ")
	fmt.Scanln()
	fmt.Scanln()
}

//==================================//
//		 Procedure Highscore 		//
//==================================//

func highscore(P PlayerData, found int) {
	fmt.Println("Nama:", P[found].nama)
	fmt.Println("NIM:", P[found].NIM)
	fmt.Println("Highscore anda:", P[found].highscore)

	fmt.Print("Enter untuk kembali... ")
	fmt.Scanln()
	fmt.Scanln()
}

//==================================//
//		Procedure Peringkat 		//
//==================================//

func peringkat(P PlayerData) {
	var i int
	insertionSort(&P, 5)
	for i = 0; i < 3; i++ {
		if i == 0 {
			fmt.Print("[*1*]")
		} else if i == 1 {
			fmt.Print("[2]")
		} else if i == 2 {
			fmt.Print(".3.")
		} else {
			fmt.Print(i + 1)
		}
		fmt.Println(P[i].nama, P[i].highscore)
	}
	fmt.Print("Enter untuk kembali... ")
	fmt.Scanln()
	fmt.Scanln()

}

//======================================//
//		    Sequential Search			//
//======================================//

// Mencari NIM Player

func searchPNIM(P PlayerData, n int, x string) int {
	var i, found int
	i = 0
	found = -1
	for i <= n && found == -1 {
		if P[i].NIM == x {
			found = i
		}
		i++
	}
	return found
}

//======================================//
//		      Binary Search				//
//======================================//

//======================================//
//		Selection Sort (Descending)		//
//======================================//

// Mengurutkan Soal berdasarkan Jumlah Benar atau Salah

func selectionSortB(S *DataSoal, n int) {
	var pass, idx, i int
	var temp Soal
	pass = 1

	for pass < n {
		idx = pass - 1
		i = pass
		for i < n {
			if S[idx].benar < S[i].benar {
				idx = i
			}
			i++
		}
		temp = S[pass-1]
		S[pass-1] = S[idx]
		S[idx] = temp
		pass++
	}
}

func selectionSortS(S *DataSoal, n int) {
	var pass, idx, i int
	var temp Soal
	pass = 1

	for pass < n {
		idx = pass - 1
		i = pass
		for i < n {
			if S[idx].salah < S[i].salah {
				idx = i
			}
			i++
		}
		temp = S[pass-1]
		S[pass-1] = S[idx]
		S[idx] = temp
		pass++
	}
}

//======================================//
//		Insertion Sort (Descending)		//
//======================================//

// Mengurutkan Player berdasarkan Highscore

func insertionSort(P *PlayerData, n int) {
	var pass, i int
	var temp User
	pass = 1

	for pass < n {
		i = pass
		temp = P[pass]

		for i > 0 && temp.highscore > P[i-1].highscore {
			P[i] = P[i-1]
			i--
		}
		P[i] = temp
		pass++
	}
}

// Mengupdate Highscore Player

func updateHigh(P *PlayerData, found int) {
	P[found].highscore = P[found].point
	fmt.Println("Highscore terupdate")
}

func randomize(S *DataSoal) {
	fmt.Println(len(S))
	for i := 0; i < totalS; i++ {
		j := rand.Intn(i + 1)
		S[i], S[j] = S[j], S[i]
	}
}
