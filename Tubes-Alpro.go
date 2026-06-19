package main

import "fmt"

const NMAX = 100

type mahasiswa struct {
	nama              string
	nim               string
	nominalIuran      int
	totalIuran        int
	totalTunggakan    int
	tanggalPembayaran string
	bulanPembayaran   string
	tahunPembayaran   string
	statusPembayaran  string
}

type arrMahasiswa [NMAX]mahasiswa

type kas struct {
	saldo       int
	jumlahLunas int
}

func pilihMenu() {
	fmt.Println()
	fmt.Println("======== MENU SIKAS ========")
	fmt.Println("1. Tambah Data Mahasiswa")
	fmt.Println("2. Ubah Data Mahasiswa")
	fmt.Println("3. Hapus Data Mahasiswa")
	fmt.Println("4. Catat Transaksi")
	fmt.Println("5. Cari Mahasiswa yang Belum Membayar")
	fmt.Println("6. Urutkan Data Mahasiswa")
	fmt.Println("7. Tampilkan Statistik Kas")
	fmt.Println("8. Tampilkan Data Mahasiswa")
	fmt.Println("0. Keluar")
	fmt.Print("Pilih menu: ")
}

func tambahMahasiswa(M *arrMahasiswa, jumlahMahasiswa *int) {
	var n int
	fmt.Print("Jumlah mahasiswa yang ingin ditambahkan: ")
	fmt.Scan(&n)

	if n >= 1 && *jumlahMahasiswa+n <= 100 {

		jumlahLama := *jumlahMahasiswa

		*jumlahMahasiswa += n

		for i := 0; i < n; i++ {
			idx := jumlahLama + i
			fmt.Print("Masukkan nama mahasiswa (Gunakan underscore untuk spasi): ")
			fmt.Scan(&M[idx].nama)

			fmt.Print("Masukkan NIM mahasiswa: ")
			fmt.Scan(&M[idx].nim)

			fmt.Print("Masukkan nominal iuran: Rp. ")
			fmt.Scan(&M[idx].nominalIuran)
			fmt.Println()

			M[idx].tanggalPembayaran = "-"
			M[idx].bulanPembayaran = " "
			M[idx].tahunPembayaran = " "
			M[idx].statusPembayaran = "Belum Lunas"
			M[idx].totalIuran = 0
			M[idx].totalTunggakan = M[idx].nominalIuran
		}
		fmt.Println("\nData berhasil ditambahkan.")
		tampilkanData(M, *jumlahMahasiswa)
	} else {
		fmt.Println("Jumlah mahasiswa harus antara 1 dan 100.")
	}
}

func ubahMahasiswa(M *arrMahasiswa, n int) {
	var nim string
	var found bool
	var pilih int

	fmt.Print("Masukkan NIM mahasiswa yang ingin diubah datanya: ")
	fmt.Scan(&nim)

	// ubah data mahasiswa: sequential search
	for i := 0; i < n; i++ {
		if M[i].nim == nim {
			found = true

			fmt.Println("\nData ditemukan!")
			fmt.Println("1. Ubah Nama")
			fmt.Println("2. Ubah NIM")
			fmt.Println("3. Ubah Nominal Iuran")
			fmt.Print("Pilih data yang ingin diubah: ")
			fmt.Scan(&pilih)

			switch pilih {
			case 1:
				fmt.Print("Masukkan nama baru: ")
				fmt.Scan(&M[i].nama)

			case 2:
				fmt.Print("Masukkan NIM baru: ")
				fmt.Scan(&M[i].nim)

			case 3:
				fmt.Print("Masukkan nominal iuran baru: Rp. ")
				fmt.Scan(&M[i].nominalIuran)

				M[i].totalTunggakan = M[i].nominalIuran - M[i].totalIuran

			default:
				fmt.Println("Pilihan tidak valid.")
				return
			}

			fmt.Println("Data berhasil diubah.")
			tampilkanData(M, n)
			return
		}
	}

	if !found {
		fmt.Println("Data tidak ditemukan.")
	}
}

func hapusMahasiswa(M *arrMahasiswa, n *int) {
	var nim string
	var found bool
	var N int

	fmt.Print("Masukkan NIM mahasiswa yang ingin dihapus datanya: ")
	fmt.Scan(&nim)

	// hapus data mahasiswa: sequential search, geser array
	for i := 0; i < *n; i++ {
		if M[i].nim == nim {
			for j := i; j < *n-1; j++ {
				M[j] = M[j+1]
			}

			(*n)--
			fmt.Println("Data berhasil dihapus.")
			found = true
		}
	}
	if found {
		N = *n
		tampilkanData(M, N)
	} else {
		fmt.Println("Data tidak ditemukan.")
	}
	return
}

func catatTransaksi(M *arrMahasiswa, n int, K *kas) {
	var nim string
	var pembayaran int
	var found bool

	fmt.Print("Masukkan NIM mahasiswa: ")
	fmt.Scan(&nim)

	// catat transaksi: sequential search
	for i := 0; i < n; i++ {
		if M[i].nim == nim {
			fmt.Print("Masukkan jumlah pembayaran: Rp. ")
			fmt.Scan(&pembayaran)

			fmt.Print("Masukkan tanggal pembayaran: ")
			fmt.Scan(&M[i].tanggalPembayaran)

			fmt.Print("Masukkan bulan pembayaran: ")
			fmt.Scan(&M[i].bulanPembayaran)

			fmt.Print("Masukkan tahun pembayaran: ")
			fmt.Scan(&M[i].tahunPembayaran)

			// update total iuran
			M[i].totalIuran = M[i].totalIuran + pembayaran
			K.saldo = K.saldo + pembayaran

			// update tunggakan
			M[i].totalTunggakan = M[i].nominalIuran - M[i].totalIuran
			M[i].statusPembayaran = "Belum Lunas"

			// update status pembayaran jika sudah lunas
			if M[i].totalIuran >= M[i].nominalIuran && M[i].statusPembayaran != "Lunas" {
				M[i].statusPembayaran = "Lunas"
				K.jumlahLunas++
				M[i].totalTunggakan = 0
			}

			fmt.Println("Transaksi berhasil dicatat.")
			tampilkanData(M, n)
			found = true
			return
		}
	}
	if !found {
		fmt.Println("Data tidak ditemukan.")
	}
}

func cariBelumBayar(M *arrMahasiswa, n int) {
	var pilih int
	var found bool
	var nim string
	var left, right, mid int
	var temp mahasiswa
	var pass, j, idx int

	fmt.Println("\n===== PENCARIAN MAHASISWA =====")
	fmt.Println("1. Tampilkan daftar mahasiswa belum lunas")
	fmt.Println("2. Cari mahasiswa berdasarkan NIM")
	fmt.Print("Pilih: ")
	fmt.Scan(&pilih)

	switch pilih {

	case 1:
		found = false
		fmt.Println("\nDaftar mahasiswa yang belum lunas:")
		fmt.Println("Nama         NIM             Total Tunggakan")
		fmt.Println("--------------------------------------------")
		// tampilkan mahasiswa yang belum lunas: sequential search
		for i := 0; i < n; i++ {
			if M[i].statusPembayaran != "Lunas" {
				fmt.Printf("%-12s %-15s Rp. %-10d\n", M[i].nama, M[i].nim, M[i].totalTunggakan)
				found = true
			}
		}

		if !found {
			fmt.Println("Semua mahasiswa sudah lunas.")
		}

	case 2:
		// urutkan data mahasiswa berdasarkan NIM: selection sort
		for pass = 1; pass < n; pass++ {
			idx = pass - 1

			for j = pass - 1; j < n; j++ {
				if M[j].nim < M[idx].nim {
					idx = j
				}
			}

			temp = M[idx]
			M[idx] = M[pass-1]
			M[pass-1] = temp
		}

		fmt.Print("Masukkan NIM mahasiswa: ")
		fmt.Scan(&nim)

		// cari mahasiswa berdasarkan NIM: binary search
		left = 0
		right = n - 1
		found = false

		for left <= right && !found {
			mid = (left + right) / 2

			if M[mid].nim == nim {
				found = true

				if M[mid].statusPembayaran != "Lunas" {
					fmt.Println("\nMahasiswa ditemukan:")
					fmt.Println("Nama      :", M[mid].nama)
					fmt.Println("NIM       :", M[mid].nim)
					fmt.Println("Tunggakan :", M[mid].totalTunggakan)
					fmt.Println("Status    :", M[mid].statusPembayaran)
				} else {
					fmt.Println("Mahasiswa sudah lunas.")
				}

			} else if M[mid].nim < nim {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}

		if !found {
			fmt.Println("Mahasiswa tidak ditemukan.")
		}

	default:
		fmt.Println("Pilihan tidak tersedia.")
	}
}

func insertionSortNamaAscending(M *arrMahasiswa, n int) {
	var pass, i int
	var temp mahasiswa

	// urutkan data mahasiswa berdasarkan nama (ascending): insertion sort
	for pass = 1; pass < n; pass++ {
		temp = M[pass]
		i = pass
		for i > 0 && temp.nama < M[i-1].nama {
			M[i] = M[i-1]
			i--
		}
		M[i] = temp
	}
	fmt.Println("Data berhasil diurutkan berdasarkan nama (A-Z).")
	tampilkanData(M, n)
}

func insertionSortNamaDescending(M *arrMahasiswa, n int) {
	var pass, i int
	var temp mahasiswa

	// urutkan data mahasiswa berdasarkan nama (descending): insertion sort
	for pass = 1; pass < n; pass++ {
		temp = M[pass]
		i = pass
		for i > 0 && temp.nama > M[i-1].nama {
			M[i] = M[i-1]
			i--
		}
		M[i] = temp
	}
	fmt.Println("Data berhasil diurutkan berdasarkan nama (Z-A).")
	tampilkanData(M, n)
}

func sortTunggakanAscending(M *arrMahasiswa, n int) {
	var pass, j, idx int
	var temp mahasiswa

	// urutkan data mahasiswa berdasarkan tunggakan (ascending): selection sort
	for pass = 1; pass < n; pass++ {
		idx = pass - 1
		for j = pass - 1; j < n; j++ {
			if M[j].totalTunggakan < M[idx].totalTunggakan {
				idx = j
			}
		}
		temp = M[idx]
		M[idx] = M[pass-1]
		M[pass-1] = temp
	}
	fmt.Println("Data diurutkan berdasarkan tunggakan (↑)")
	tampilkanData(M, n)
}

func sortTunggakanDescending(M *arrMahasiswa, n int) {
	var pass, j, idx int
	var temp mahasiswa

	// urutkan data mahasiswa berdasarkan tunggakan (descending): selection sort
	for pass = 1; pass < n; pass++ {
		idx = pass - 1
		for j = pass - 1; j < n; j++ {
			if M[j].totalTunggakan > M[idx].totalTunggakan {
				idx = j
			}
		}
		temp = M[idx]
		M[idx] = M[pass-1]
		M[pass-1] = temp
	}
	fmt.Println("Data diurutkan berdasarkan tunggakan (↓)")
	tampilkanData(M, n)
}

func statistik(K kas) {
	fmt.Println("\n===== STATISTIK KAS =====")
	fmt.Println("Total saldo kas: Rp. ", K.saldo)
	fmt.Println("Jumlah mahasiswa lunas:", K.jumlahLunas)
}

func tampilkanData(M *arrMahasiswa, n int) {
	var i int

	if n == 0 {
		fmt.Println("Belum ada data mahasiswa.")
		return
	}

	fmt.Println("\n====================================== DATA MAHASISWA ======================================")

	fmt.Printf("%-12s %-15s %-10s %-10s %-12s %-16s %-15s\n",
		"Nama", "NIM", "Nominal", "Total", "Tunggakan", "Tanggal Bayar", "Status")

	for i = 0; i < n; i++ {
		fmt.Printf("%-12s %-15s %-10d %-10d %-12d %-2s %-2s %-10s %-15s\n",
			M[i].nama,
			M[i].nim,
			M[i].nominalIuran,
			M[i].totalIuran,
			M[i].totalTunggakan,
			M[i].tanggalPembayaran,
			M[i].bulanPembayaran,
			M[i].tahunPembayaran,
			M[i].statusPembayaran,
		)
	}
}

func main() {
	var M arrMahasiswa
	var K kas
	var jumlahMahasiswa int
	var menu, pilih int
	var selesai bool

	fmt.Println("Selamat datang di Sistem Informasi Kas (SIKAS)")

	selesai = false

	for !selesai {
		pilihMenu()
		fmt.Scan(&menu)

		switch menu {
		case 0:
			fmt.Println("Program selesai.")
			selesai = true

		case 1:
			tambahMahasiswa(&M, &jumlahMahasiswa)

		case 2:
			ubahMahasiswa(&M, jumlahMahasiswa)

		case 3:
			hapusMahasiswa(&M, &jumlahMahasiswa)

		case 4:
			catatTransaksi(&M, jumlahMahasiswa, &K)

		case 5:
			cariBelumBayar(&M, jumlahMahasiswa)

		case 6:
			fmt.Println("1. Urutkan berdasarkan nama (A-Z)")
			fmt.Println("2. Urutkan berdasarkan nama (Z-A)")
			fmt.Println("3. Urutkan berdasarkan tunggakan (↑)")
			fmt.Println("4. Urutkan berdasarkan tunggakan (↓)")
			fmt.Print("Pilih: ")
			fmt.Scan(&pilih)

			if pilih == 1 {
				insertionSortNamaAscending(&M, jumlahMahasiswa)
			} else if pilih == 2 {
				insertionSortNamaDescending(&M, jumlahMahasiswa)
			} else if pilih == 3 {
				sortTunggakanAscending(&M, jumlahMahasiswa)
			} else if pilih == 4 {
				sortTunggakanDescending(&M, jumlahMahasiswa)
			} else {
				fmt.Println("Pilihan tidak tersedia.")
			}

		case 7:
			statistik(K)

		case 8:
			tampilkanData(&M, jumlahMahasiswa)

		default:
			fmt.Println("Menu tidak tersedia.")
		}
	}
}
