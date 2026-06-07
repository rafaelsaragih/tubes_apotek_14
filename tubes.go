package main
import "fmt"

type obat struct {
	stok int
	nama, gejala, kadaluarsa string
}

type arrObat [1000]obat

func main() {
	var data arrObat
	var n int
	var pilih int

	for {

		fmt.Println("===== APOTEK SMART =====")
		fmt.Println("1. Tambah Obat")
		fmt.Println("2. Ubah Obat")
		fmt.Println("3. Hapus Obat")
		fmt.Println("4. Tampilkan Data")
		fmt.Println("5. Sequential Search Nama")
		fmt.Println("6. Sequential Search Gejala")
		fmt.Println("7. Binary Search Nama")
		fmt.Println("8. Selection Sort Expired")
		fmt.Println("9. Insertion Sort Expired")
		fmt.Println("10. Statistik")
		fmt.Println("0. Keluar")
		fmt.Print("Pilih: ")

		fmt.Scan(&pilih)

		if pilih == 1 {
			tambahObat(&data, &n)
		} else if 
	}
}

func tambahObat(A *arrObat, n *int) {

	fmt.Print("Nama: ")
	fmt.Scan(&(*A)[*n].nama)

	fmt.Print("Gejala: ")
	fmt.Scan(&(*A)[*n].gejala)

	fmt.Print("Stok: ")
	fmt.Scan(&(*A)[*n].stok)

	fmt.Print("Kadaluarsa (YYYY-MM-DD): ")
	fmt.Scan(&(*A)[*n].kadaluarsa)

	(*n)++
	fmt.Println("Data berhasil ditambahkan.")
}

func ubahObat(A *arrObat, n *int) {
	var ubah string
	var i int

	fmt.Print("Masukkan nama obat yang ingin diubah: ")
	fmt.Scan(&ubah)

	for i = 0; i < *n; i++ {
		if (*A)[i].nama == ubah {

			fmt.Println("Data ditemukan!")
			fmt.Println("Masukkan data baru")

			fmt.Print("Nama: ")
			fmt.Scan(&(*A)[i].nama)

			fmt.Print("Gejala: ")
			fmt.Scan(&(*A)[i].gejala)

			fmt.Print("Stok: ")
			fmt.Scan(&(*A)[i].stok)

			fmt.Print("Kadaluarsa: ")
			fmt.Scan(&(*A)[i].kadaluarsa)

			fmt.Println("Data berhasil diubah.")
			return
		}
	}

	fmt.Println("Data tidak ditemukan.")
}


func selectionSortExpired(A *arrObat, n int) {
	var i, j, min int

	for i = 0; i < n-1; i++ {

		min = i

		for j = i + 1; j < n; j++ {
			if (*A)[j].kadaluarsa < (*A)[min].kadaluarsa {
				min = j
			}
		}

		(*A)[i], (*A)[min] = (*A)[min], (*A)[i]
	}
}

func insertionSortExpired(A *arrObat, n int) {
	var pass, i int
	var temp obat

	for pass = 1; pass < n; pass++ {

		temp = (*A)[pass]
		i = pass - 1

		for i >= 0 && temp.kadaluarsa < (*A)[i].kadaluarsa {
			(*A)[i+1] = (*A)[i]
			i--
		}

		(*A)[i+1] = temp
	}
}

func statistik(data *arrObat, n int) {
	var i int
	var jumlah int

	jumlah = 0

	for i = 0; i < n; i++ {

		if (*data)[i].stok <= 5 {

			fmt.Println((*data)[i].nama, "-", (*data)[i].stok)
			jumlah++
		}
	}

	fmt.Println("Jumlah obat hampir habis :", jumlah)
}
