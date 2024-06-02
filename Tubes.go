package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
)

// Struct untuk pengguna
type Pengguna struct {
	nama, NIM, password, pin string
	point, highscore         int
}

// Struct untuk soal
type Soal struct {
	pertanyaan   string
	pilihan      [4]string
	kunciJawaban string
	benar, salah int
}

// Array
type DataPemain [1000]Pengguna
type DataSoal [1000]Soal

var soal DataSoal
var player DataPemain

//======================//
//		FUNC MAIN		//
//======================//

func main() {
	player[0] = Pengguna{
		nama:      "admin",
		NIM:       "admin",
		password:  "admin",
		pin:       "admin",
		highscore: -100,
	}
	var totalP int = 1
	var totalS int = 0
	menuAwal(&totalP, &totalS)
}

// Menu awal
func menuAwal(tP, tS *int) {
	var pilih int
	fmt.Println()
	fmt.Println("|====================|")
	fmt.Println("|        MENU        |")
	fmt.Println("| 1. Login           |")
	fmt.Println("| 2. Buat akun       |")
	fmt.Println("| 3. Lupa password   |")
	fmt.Println("| 4. Keluar          |")
	fmt.Println("|====================|")

	for {
		fmt.Print("[#] Pilihan (1/2/3/4): ")
		fmt.Scan(&pilih)
		if pilih == 1 || pilih == 2 || pilih == 3 || pilih == 4 {
			break
		}
	}

	switch pilih {
	case 1:
		loginNIM(&player, tP, tS)
	case 2:
		buatAkun(&player, tP)
		menuAwal(tP, tS)
	case 3:
		gantiPW(&player, *tP)
		menuAwal(tP, tS)
	case 4:
	}
}

// Membuat akun baru
func buatAkun(P *DataPemain, tP *int) {
	var i int
	var nim, pilih string
	i = *tP

	fmt.Print("> Masukkan nama: ")
	fmt.Scanln()
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	P[i].nama = scanner.Text()

	for {
		fmt.Print("> Masukkan NIM: ")
		fmt.Scan(&nim)

		if sequentialSearch(*P, *tP, nim) == -1 {
			P[i].NIM = nim
			break
		}
		fmt.Println("NIM sudah terdaftar")
	}

	fmt.Print("> Masukkan password: ")
	fmt.Scan(&P[i].password)

	for {
		fmt.Print("[#] Pin keamanan? (Y/N): ")
		fmt.Scan(&pilih)
		if pilih == "Y" || pilih == "y" {
			fmt.Print("> Buat pin: ")
			fmt.Scan(&P[i].pin)
			break
		} else if pilih == "N" || pilih == "n" {
			break
		}
	}

	fmt.Print("|Akun berhasil dibuat \n")
	*tP++
}

// Procedure login (NIM)
func loginNIM(P *DataPemain, tP, tS *int) {
	var found, nimLimit int
	var nim string
	var login bool = false

	for nimLimit != 3 && !login {
		fmt.Print("> NIM: ")
		fmt.Scan(&nim)
		found = sequentialSearch(*P, *tP, nim)
		if found == -1 {
			fmt.Println("|NIM tidak terdaftar")
			nimLimit++
		} else {
			login = true
			loginPW(P, found, tP, tS)
		}
	}
	if nimLimit == 3 {
		fmt.Println("|NIM salah 3 kali berturut")
		menuAwal(tP, tS)
	}
}

// Procedure login (password)
func loginPW(P *DataPemain, found int, tP, tS *int) {
	var pwLimit int
	var pw string
	var login bool = false

	for pwLimit != 3 && !login {
		fmt.Print("> Password: ")
		fmt.Scan(&pw)

		if P[found].NIM == "admin" && pw == P[found].password {
			menuAdmin(tP, tS)
			login = true
		} else if pw == P[found].password {
			fmt.Println("|Login Berhasil")
			login = true
			menuKuis(P, found, *tP, *tS)
		} else {
			fmt.Println("|Password salah")
			pwLimit++
		}
	}
	if pwLimit == 3 {
		fmt.Println("|Password salah 3 kali berturut")
		menuAwal(tP, tS)
	}
}

// Mengubah password player
func gantiPW(P *DataPemain, tP int) {
	var found, nimLimit int
	var nim, pin string
	var cek bool
	insertionSortNIM(P, tP)
	found = binarySearch(*P, tP, nim)

	for nimLimit != 3 && !cek {
		fmt.Print("> NIM: ")
		fmt.Scan(&nim)
		found = binarySearch(*P, tP, nim)
		if found == -1 {
			fmt.Println("|NIM tidak terdaftar")
			nimLimit++
		} else if P[found].pin == "" {
			fmt.Println("|Akun tidak memiliki pin")
			nimLimit++
		} else {
			fmt.Println("> Pin: ")
			fmt.Scan(&pin)
			if pin == P[found].pin {
				fmt.Print("|Silahkan ubah password: ")
				fmt.Scan(&P[found].password)
				fmt.Println(P[found].password)
			} else {
				fmt.Println("|Pin salah")
			}
		}
	}
	if nimLimit == 3 {
		fmt.Println("|NIM salah 3 kali berturut")
	}
}

// Menu admin
func menuAdmin(tP, tS *int) {
	var pilih int
	fmt.Println()
	fmt.Println("|====================|")
	fmt.Println("|     MENU ADMIN     |")
	fmt.Println("| 1. Tambah soal     |")
	fmt.Println("| 2. Ubah soal       |")
	fmt.Println("| 3. Hapus soal      |")
	fmt.Println("| 4. Daftar soal     |")
	fmt.Println("| 5. Info Soal       |")
	fmt.Println("| 6. Kembali         |")
	fmt.Println("|====================|")

	for {
		fmt.Print("[#] Pilihan (1/2/3/4/5/6): ")
		fmt.Scan(&pilih)
		if pilih == 1 || pilih == 2 || pilih == 3 || pilih == 4 || pilih == 5 || pilih == 6 {
			break
		}
	}
	switch pilih {
	case 1:
		tambahSoal(&soal, tS, *tS)
		menuAdmin(tP, tS)
	case 2:
		ubahSoal(&soal)
		menuAdmin(tP, tS)
	case 3:
		hapusSoal(&soal, tS)
		menuAdmin(tP, tS)
	case 4:
		listSoal(*tS)
		menuAdmin(tP, tS)
	case 5:
		infoSoal(soal, *tS)
		menuAdmin(tP, tS)
	case 6:
		menuAwal(tP, tS)
	}
}

// Menambah soal
func tambahSoal(S *DataSoal, tS *int, no int) {
	var i, bSoal int
	var exit string

	fmt.Print("> Masukkan banyak soal yang ingin ditambah: ")
	fmt.Scan(&bSoal)

	for i = no; i < no+bSoal; i++ {
		fmt.Print("[?] Masukkan pertanyaan: ")
		fmt.Scanln()
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		S[i].pertanyaan = scanner.Text()
		fmt.Println("[&] Masukkan pilihan jawaban:")
		fmt.Print("> a. ")
		pilihan := bufio.NewScanner(os.Stdin)
		pilihan.Scan()
		S[i].pilihan[0] = pilihan.Text()
		fmt.Print("> b. ")
		pilihan.Scan()
		S[i].pilihan[1] = pilihan.Text()
		fmt.Print("> c. ")
		pilihan.Scan()
		S[i].pilihan[2] = pilihan.Text()
		fmt.Print("> d. ")
		pilihan.Scan()
		S[i].pilihan[3] = pilihan.Text()

		for {
			fmt.Print("> Masukkan kunci jawaban: ")
			fmt.Scan(&S[i].kunciJawaban)
			exit = S[i].kunciJawaban
			if exit == "a" || exit == "b" || exit == "c" || exit == "d" {
				break
			}
		}

		fmt.Println("|Pertanyaan berhasil dibuat")
		*tS++
	}
}

// Mengubah soal
func ubahSoal(S *DataSoal) {
	var no int
	var exit string

	fmt.Println("|==================")
	fmt.Print("> Nomor soal yang mau diubah: ")
	fmt.Scan(&no)

	fmt.Print("> Masukkan pertanyaan: ")
	fmt.Scanln()
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	S[no-1].pertanyaan = scanner.Text()
	fmt.Println("[&] Masukkan pilihan jawaban:")
	fmt.Print("> a. ")
	pilihan := bufio.NewScanner(os.Stdin)
	pilihan.Scan()
	S[no-1].pilihan[0] = pilihan.Text()
	fmt.Print("> b. ")
	pilihan.Scan()
	S[no-1].pilihan[1] = pilihan.Text()
	fmt.Print("> c. ")
	pilihan.Scan()
	S[no-1].pilihan[2] = pilihan.Text()
	fmt.Print("> d. ")
	pilihan.Scan()
	S[no-1].pilihan[3] = pilihan.Text()

	for {
		fmt.Print("> Masukkan kunci jawaban: ")
		fmt.Scan(&S[no-1].kunciJawaban)
		exit = S[no-1].kunciJawaban
		if exit == "a" || exit == "b" || exit == "c" || exit == "d" {
			break
		}
	}
	fmt.Println("|Soal berhasil diubah \n")
}

// Menghapus Soal
func hapusSoal(S *DataSoal, tS *int) {
	var i, no int

	fmt.Print("> Nomor soal yang mau dihapus: ")
	fmt.Scan(&no)
	for i = no - 1; i < *tS; i++ {
		S[i].pertanyaan = S[i+1].pertanyaan
		S[i].pilihan[0] = S[i+1].pilihan[0]
		S[i].pilihan[1] = S[i+1].pilihan[1]
		S[i].pilihan[2] = S[i+1].pilihan[2]
		S[i].pilihan[3] = S[i+1].pilihan[3]
		S[i].kunciJawaban = S[i+1].kunciJawaban
	}
	fmt.Printf("[x] Soal no %d berhasil dihapus \n", no)
	*tS--
}

// Menampilkan daftar pertanyaan dan jawabannya
func listSoal(tS int) {
	for i := 0; i < tS; i++ {
		fmt.Print(i+1, ". ", soal[i].pertanyaan, "\n")
		fmt.Print("a. ", soal[i].pilihan[0], "\n")
		fmt.Print("b. ", soal[i].pilihan[1], "\n")
		fmt.Print("c. ", soal[i].pilihan[2], "\n")
		fmt.Print("d. ", soal[i].pilihan[3], "\n")
		fmt.Println("[*] Jawabannya adalah", soal[i].kunciJawaban)
		fmt.Println()
	}
	fmt.Print("Enter untuk kembali... ")
	fmt.Scanln()
	fmt.Scanln()
}

// Menampilkan Soal dengan jumlah benar dan salah terbanyak
func infoSoal(S DataSoal, n int) {
	var i, j int
	fmt.Println("|Soal Termudah: ")
	selectionSortB(&S, n)
	for i = 0; i < 5; i++ {
		fmt.Printf("|%d. %s\n", i+1, S[i].pertanyaan)
		for j = 0; j < S[i].benar; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
	fmt.Println("|================|")
	fmt.Println("Soal Tersusah: ")
	selectionSortS(&S, n)
	for i = 0; i < 5; i++ {
		fmt.Printf("|%d. %s\n", i+1, S[i].pertanyaan)
		for j = 0; j < S[i].salah; j++ {
			fmt.Print("x")
		}
		fmt.Println()
	}

	fmt.Print("Enter untuk kembali... ")
	fmt.Scanln()
	fmt.Scanln()
}

// Menu
func menuKuis(P *DataPemain, found int, tP, tS int) {
	var pilih int
	fmt.Println()
	fmt.Println("|====================|")
	fmt.Println("|     MENU KUIS      |")
	fmt.Println("| 1. Mulai kuis      |")
	fmt.Println("| 2. Highscore       |")
	fmt.Println("| 3. Peringkat       |")
	fmt.Println("| 4. Kembali         |")
	fmt.Println("|====================|")

	for {
		fmt.Print("[#] Pilihan (1/2/3/4): ")
		fmt.Scan(&pilih)
		if pilih == 1 || pilih == 2 || pilih == 3 || pilih == 4 {
			break
		}
	}
	switch pilih {
	case 1:
		mulaiKuis(P, &soal, tS, found)
		menuKuis(P, found, tP, tS)
	case 2:
		highscore(*P, found)
		menuKuis(P, found, tP, tS)
	case 3:
		peringkat(*P)
		menuKuis(P, found, tP, tS)
	case 4:
		menuAwal(&tP, &tS)
	}
}

// Memulai Kuis
func mulaiKuis(P *DataPemain, S *DataSoal, tS, found int) {
	var i int
	var poin, poinAkhir float64
	var jawab string
	P[found].point = 0

	randomize(S, tS)

	if tS > 10 {
		tS = 10
	}

	if tS == 0 {
		fmt.Println("|Belum ada pertanyaan")
	} else {
		poin = 100 / float64(tS)
		for i = 0; i < tS; i++ {
			fmt.Printf("%d. %s?\n", i+1, S[i].pertanyaan)
			fmt.Println("a.", S[i].pilihan[0])
			fmt.Println("b.", S[i].pilihan[1])
			fmt.Println("c.", S[i].pilihan[2])
			fmt.Println("d.", S[i].pilihan[3])
			fmt.Print("> Jawab: ")
			fmt.Scan(&jawab)

			if jawab == S[i].kunciJawaban {
				fmt.Println("|Jawaban Benar|")
				fmt.Printf("Point +%.3g\n\n", poin)
				poinAkhir += poin
				S[i].benar++
			} else {
				fmt.Println("|Jawaban Salah|")
				fmt.Println("Jawaban :", S[i].kunciJawaban)
				fmt.Println("Point +0\n")
				S[i].salah++
			}
		}
		P[found].point = int(poinAkhir)
		fmt.Println("Score akhir:", P[found].point)
		fmt.Printf("Selamat, anda mendapatkan Rp%d\n", 100000000/P[found].point)
	}

	if P[found].highscore < P[found].point {
		updateHigh(P, found)
	}

	fmt.Print("Enter untuk kembali... ")
	fmt.Scanln()
	fmt.Scanln()
}

// Menampilkan Highscore pribadi
func highscore(P DataPemain, found int) {
	fmt.Println("|Nama:", P[found].nama)
	fmt.Println("|NIM:", P[found].NIM)
	fmt.Println("|Highscore anda:", P[found].highscore)

	fmt.Print("Enter untuk kembali... ")
	fmt.Scanln()
	fmt.Scanln()
}

// Menampilkan peringkat 1 sampai 3 Player dengan highscore tertinggi
func peringkat(P DataPemain) {
	var i int
	insertionSortH(&P, 5)
	for i = 0; i < 3; i++ {
		if i == 0 {
			fmt.Print("[1] ")
		} else if i == 1 {
			fmt.Print("[2] ")
		} else if i == 2 {
			fmt.Print(".3. ")
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
func sequentialSearch(P DataPemain, n int, x string) int {
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

// Mencari NIM Player
func binarySearch(P DataPemain, n int, x string) int {
	var l, r, m, found int
	l = 1
	r = n - 1
	found = -1
	for l <= r && found == -1 {
		m = (l + r) / 2
		if x < P[m].NIM {
			r = m - 1
		} else if x > P[m].NIM {
			l = m + 1
		} else {
			found = m
		}
	}
	return found
}

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
//			Insertion Sort 				//
//======================================//

// Mengurutkan Player berdasarkan Highscore (descending)
func insertionSortH(P *DataPemain, n int) {
	var pass, i int
	var temp Pengguna
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

// Mengurutkan Player berdasarkan NIM (ascending)
func insertionSortNIM(P *DataPemain, n int) {
	var pass, i int
	var temp Pengguna
	pass = 2

	for pass < n {
		i = pass
		temp = P[pass]

		for i > 0 && temp.NIM < P[i-1].NIM {
			P[i] = P[i-1]
			i--
		}
		P[i] = temp
		pass++
	}
}

// Mengupdate Highscore Player
func updateHigh(P *DataPemain, found int) {
	P[found].highscore = P[found].point
	fmt.Println("Highscore terupdate")
}

// Mengacak soal kuis
func randomize(S *DataSoal, tS int) {
	for i := 0; i < tS; i++ {
		j := rand.Intn(i + 1)
		S[i], S[j] = S[j], S[i]
	}
}
