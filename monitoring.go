package main

import "fmt"

const NMAX = 999

type komponen struct {
	idkomponen           int
	nama, noseri, status string
}
type sensor struct {
	idkomponen  int
	suhu, beban float64
}
type tabpc [NMAX]komponen
type tabsensor [NMAX]sensor

func addkomponen(T *tabpc, S *tabsensor, nk *int) {
	var pilihan, i int
	fmt.Println("===TAMBAH KOMPONEN===")
	if *nk < NMAX {
		fmt.Print("Nama komponen: ")
		fmt.Scan(&T[*nk].nama)
		fmt.Print("No seri: ")
		fmt.Scan(&T[*nk].noseri)
		for i < *nk {
			if T[i].noseri == T[*nk].noseri {
				fmt.Println("Nomor seri sudah terdata")
				fmt.Print("No seri: ")
				fmt.Scan(&T[*nk].noseri)
				i = 0 //kalau ketemu riset i ke 0 agar ngecek dari awal lagi
			} else {
				i++ //kalau blm ketemu i bertambah 1
			}
		}
		fmt.Print("Suhu komponen (°C): ")
		fmt.Scan(&S[*nk].suhu)
		fmt.Print("Beban komponen (%): ")
		fmt.Scan(&S[*nk].beban)
		if S[*nk].suhu > 100 && S[*nk].beban > 100 {
			T[*nk].status = "Blue_Screen"
		} else if S[*nk].suhu > 85 {
			T[*nk].status = "Overheat"
		} else if S[*nk].beban > 90 {
			T[*nk].status = "Lag"
		} else {
			T[*nk].status = "Normal"
		}
		T[*nk].idkomponen = *nk + 1
		S[*nk].idkomponen = T[*nk].idkomponen
		*nk++
		fmt.Println("Komponen berhasil ditambahkan")
		showdata(*T, *S, *nk)
		fmt.Println("1. Tambah komponen")
		fmt.Println("2. Kembali")
		fmt.Print("Pilih menu (1-2): ")
		fmt.Scan(&pilihan)
		if pilihan == 1 {
			addkomponen(T, S, nk)
		} else if pilihan != 2 {
			for pilihan != 1 && pilihan != 2 {
				fmt.Println("Pilihan tidak valid, silakan pilih menu yang tersedia")
				fmt.Print("Pilih menu (1-2): ")
				fmt.Scan(&pilihan)
			}
		}
	} else {
		fmt.Println("data sudah penuh!")
	}
}
func updtkomponen(T *tabpc, S *tabsensor, nk int) {
	var carikomoponen, i, j, pilihan int
	var ketemu bool
	fmt.Println("===UPDATE KOMPONEN===")
	showdata(*T, *S, nk)
	fmt.Println("klik 0 untuk kembali ke menu utama")
	fmt.Print("Masukkan ID komponen yang ingin diupdate: ")
	fmt.Scan(&carikomoponen)
	if carikomoponen != 0 {
		for i = 0; i < nk; i++ {
			if T[i].idkomponen == carikomoponen {
				ketemu = true
				fmt.Print("Nama komponen: ")
				fmt.Scan(&T[i].nama)
				fmt.Print("No seri: ")
				fmt.Scan(&T[i].noseri)
				j = 0
				for j < nk {
					if T[j].noseri == T[i].noseri && j != i { // fungsi j != i agar skip yang dimana index dan value sama dengan dirinya sendiri/nomor seri boleh sama dengan data seblum diubah ketika update data
						fmt.Println("Nomor seri sudah terdata")
						fmt.Print("No seri: ")
						fmt.Scan(&T[i].noseri)
						j = 0 //kalau ketemu riset i ke 0 agar ngecek dari awal lagi
					} else {
						j++ //kalau blm ketemu i bertambah 1
					}
				}
				fmt.Print("Suhu komponen (°C): ")
				fmt.Scan(&S[i].suhu)
				fmt.Print("Beban komponen (%): ")
				fmt.Scan(&S[i].beban)
				fmt.Println("Komponen berhasil diupdate")
				showdata(*T, *S, nk)
				fmt.Println("1. Update komponen")
				fmt.Println("2. Kembali")
				fmt.Print("Pilih menu (1-2): ")
				fmt.Scan(&pilihan)
				if pilihan == 1 {
					updtkomponen(T, S, nk)
				} else if pilihan != 2 {
					for pilihan != 1 && pilihan != 2 {
						fmt.Println("Pilihan tidak valid, silakan pilih menu yang tersedia")
						fmt.Print("Pilih menu (1-2): ")
						fmt.Scan(&pilihan)
					}
				}
				if S[i].suhu > 100 && S[i].beban > 100 {
					T[i].status = "Blue_Screen"
				} else if S[i].suhu > 85 {
					T[i].status = "Overheat"
				} else if S[i].beban > 90 {
					T[i].status = "Lag"
				} else {
					T[i].status = "Normal"
				}
			}
		}
		if ketemu == false {
			fmt.Println("ID komponen tidak ditemukan")
			updtkomponen(T, S, nk)
		}
	}
}
func delkomponen(T *tabpc, S *tabsensor, nk *int) {
	var carikomoponen, i, j int
	var ketemu bool
	fmt.Println("===HAPUS KOMPONEN===")
	showdata(*T, *S, *nk)
	fmt.Println("klik 0 untuk kembali ke menu utama")
	fmt.Print("Masukkan ID komponen yang ingin dihapus: ")
	fmt.Scan(&carikomoponen)
	if carikomoponen != 0 {
		for i < *nk {
			if T[i].idkomponen == carikomoponen {
				ketemu = true
				j = i
				for j < *nk-1 {
					T[j] = T[j+1]
					S[j] = S[j+1]
					j++
				}
				*nk = *nk - 1
				fmt.Println("Komponen berhasil dihapus")
			}
			i++
		}
		if ketemu == false {
			fmt.Println("ID komponen tidak ditemukan")
			delkomponen(T, S, nk)
		}
	}
}
func statuskomponen(T *tabpc, S *tabsensor, nk int) {
	var i int
	var bermasalah bool
	fmt.Println("===STATUS KOMPONEN===")
	fmt.Println("\n-------------------------------------------------------------------------------------")
	fmt.Printf("%-5s%-20s%-14s%-13s%-16s%-22s\n", "ID", "Nama", "Suhu", "Beban", "Status", "NoSeri")
	fmt.Println("-------------------------------------------------------------------------------------")
	bermasalah = false
	for i = 0; i < nk; i++ {
		if T[i].status != "Normal" {
			bermasalah = true
			fmt.Printf("%-5d%-20s%-6.2f°C      %-6.2f%%      %-16s%-20s\n", T[i].idkomponen, T[i].nama, S[i].suhu, S[i].beban, T[i].status, T[i].noseri)
		}
	}
	if bermasalah == false {
		fmt.Println("Semua komponen dalam kondisi normal")
	}
	fmt.Println("-------------------------------------------------------------------------------------")
}
func searchname(T *tabpc, S *tabsensor, nk int) int {
	var mid, left, right, pass, i int
	var carinama string
	var temppc komponen
	var tempsen sensor
	fmt.Println("===CARI KOMPONEN BERDASARKAN NAMA===")
	fmt.Println("Data telah di sorting")
	pass = 1
	for pass < nk {
		temppc = T[pass]
		tempsen = S[pass]
		i = pass
		for i > 0 && temppc.nama < T[i-1].nama {
			T[i] = T[i-1]
			S[i] = S[i-1]
			i--
		}
		T[i] = temppc
		S[i] = tempsen
		pass++
	}
	fmt.Println("Masukkan nama komponen yang ingin dicari: ")
	fmt.Scan(&carinama)
	left = 0
	right = nk - 1
	for left <= right {
		mid = (left + right) / 2
		if T[mid].nama == carinama {
			return mid
		} else if T[mid].nama < carinama {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}
func searchstatus(T *tabpc, S *tabsensor, nk int) {
	var ketemu bool
	var caristatus string
	var i int
	fmt.Println("===CARI KOMPONEN BERDASARKAN STATUS===")
	fmt.Println("Masukkan status komponen yang ingin dicari: ")
	fmt.Scan(&caristatus)
	ketemu = false
	i = 0
	fmt.Println("\n-------------------------------------------------------------------------------------")
	fmt.Printf("%-5s%-20s%-14s%-13s%-16s%-22s\n", "ID", "Nama", "Suhu", "Beban", "Status", "NoSeri")
	fmt.Println("-------------------------------------------------------------------------------------")
	for i < nk {
		if T[i].status == caristatus {
			ketemu = true
			fmt.Printf("%-5d%-20s%-6.2f°C      %-6.2f%%      %-16s%-20s\n", T[i].idkomponen, T[i].nama, S[i].suhu, S[i].beban, T[i].status, T[i].noseri)
		}
		i++
	}
	fmt.Println("-------------------------------------------------------------------------------------")
	if ketemu == false {
		fmt.Println("Data tidak ditemukan.")
		fmt.Println("-------------------------------------------------------------------------------------")
	}
}
func sortingnoseridesc(T *tabpc, S *tabsensor, nk int) {
	var pass, idx, i int
	var temppc komponen
	var tempsen sensor
	pass = 1
	for pass <= nk-1 {
		idx = pass - 1
		i = pass
		for i < nk {
			if T[idx].noseri < T[i].noseri {
				idx = i
			}
			i++
		}
		temppc = T[pass-1]
		tempsen = S[pass-1]
		T[pass-1] = T[idx]
		S[pass-1] = S[idx]
		T[idx] = temppc
		S[idx] = tempsen
		pass++
	}
}
func sortingnoseriasc(T *tabpc, S *tabsensor, nk int) {
	var i, pass int
	var temppc komponen
	var tempsen sensor
	pass = 1
	for pass < nk {
		temppc = T[pass]
		tempsen = S[pass]
		i = pass
		for i > 0 && temppc.noseri < T[i-1].noseri {
			T[i] = T[i-1]
			S[i] = S[i-1]
			i--
		}
		T[i] = temppc
		S[i] = tempsen
		pass++
	}
}
func statistik(T tabpc, S tabsensor, nk int) {
	var totalsuhu, avgsuhu float64
	var totalbermasalah int
	var i int
	totalbermasalah = 0
	totalsuhu = 0
	avgsuhu = 0
	for i < nk {
		totalsuhu = totalsuhu + S[i].suhu
		if T[i].status != "Normal" {
			totalbermasalah++
		}
		i++
	}
	avgsuhu = totalsuhu / float64(nk)
	fmt.Println("===STATISTIK KOMPONEN===")
	fmt.Printf("Jumlah komponen yang bermasalah: %d komponen\n", totalbermasalah)
	fmt.Printf("Rata-rata suhu seluruh Komponen: %.2f°C\n", avgsuhu)
	fmt.Println("-----------------------------------------")
}
func menukomponen(T *tabpc, S *tabsensor, nk *int) {
	var pilih int
	var run bool
	run = true
	for run {
		fmt.Println("===MENU KOMPONEN===")
		showdata(*T, *S, *nk)
		fmt.Println("1. TAMBAH KOMPONEN")
		fmt.Println("2. UPDATE KOMPONEN")
		fmt.Println("3. HAPUS KOMPONEN")
		fmt.Println("4. KEMBALI")
		fmt.Print("PILIH MENU (1-4): ")
		fmt.Scan(&pilih)

		if pilih == 1 {
			addkomponen(T, S, nk)
		} else if pilih == 2 {
			updtkomponen(T, S, *nk)
		} else if pilih == 3 {
			delkomponen(T, S, nk)
		} else if pilih == 4 {
			run = false
		} else {
			fmt.Print("Pilihan tidak valid, silakan pilih menu yang tersedia")
		}
	}
}
func showdata(T tabpc, S tabsensor, nk int) {
	var i int
	fmt.Println("\n-------------------------------------------------------------------------------------")
	fmt.Printf("%-5s%-20s%-14s%-13s%-16s%-22s\n", "ID", "Nama", "Suhu", "Beban", "Status", "NoSeri")
	fmt.Println("-------------------------------------------------------------------------------------")
	for i = 0; i < nk; i++ {
		fmt.Printf("%-5d%-20s%-6.2f°C      %-6.2f%%      %-16s%-20s\n", T[i].idkomponen, T[i].nama, S[i].suhu, S[i].beban, T[i].status, T[i].noseri)
	}
	fmt.Println("-------------------------------------------------------------------------------------")
}
func main() {
	var nk, pilih, pilihpencarian int
	var Nnama int
	var run bool
	run = true
	var T tabpc = tabpc{{1, "NVIDIA_RTX_5090TI", "NV-RTX59TI", "Normal"}, {2, "AMD_RADEON_R9_370", "RD-AMDR937", "Normal"}, {3, "AMD_RADEON_R9_370X", "RD-AMDR937X", "Normal"}, {4, "NVIDIA_RTX_5090", "NV-RTX59", "Normal"}, {5, "NVIDIA_RTX_4060", "NV-RTX46", "Lag"}}
	var S tabsensor = tabsensor{{1, 68, 70}, {2, 72, 73}, {3, 76, 77}, {4, 78, 80}, {5, 78, 91}}
	nk = 5
	for run {
		fmt.Println("===MENU UTAMA===")
		fmt.Println("1. KOMPONEN")
		fmt.Println("2. CATATAN KOMPONEN")
		fmt.Println("3. CARI KOMPONEN")
		fmt.Println("4. TAMPILKAN SEMUA KOMPONEN BERDASARKAN NOSERI (DESC)")
		fmt.Println("5. TAMPILKAN SEMUA KOMPONEN BERDASARKAN NOSERI (ASCE)")
		fmt.Println("6. STATISTIK")
		fmt.Println("7. KELUAR")
		fmt.Print("PILIH MENU (1-7): ")
		fmt.Scan(&pilih)

		if pilih == 1 {
			menukomponen(&T, &S, &nk)
		} else if pilih == 2 {
			statuskomponen(&T, &S, nk)
		} else if pilih == 3 {
			fmt.Println("===MENU PENCARIAN===")
			fmt.Println("1. CARI NAMA KOMPONEN")
			fmt.Println("2. CARI STATUS KOMPONEN")
			fmt.Println("PILIH MENU (1-2): ")
			fmt.Scan(&pilihpencarian)
			if pilihpencarian == 1 {
				Nnama = searchname(&T, &S, nk)
				if Nnama == -1 {
					fmt.Println("Nama tidak ditemukan")
				} else {
					fmt.Println("\n-------------------------------------------------------------------------------------")
					fmt.Printf("%-5s%-20s%-14s%-13s%-16s%-22s\n", "ID", "Nama", "Suhu", "Beban", "Status", "NoSeri")
					fmt.Println("-------------------------------------------------------------------------------------")
					fmt.Printf("%-5d%-20s%-6.2f°C      %-6.2f%%      %-16s%-20s\n", T[Nnama].idkomponen, T[Nnama].nama, S[Nnama].suhu, S[Nnama].beban, T[Nnama].status, T[Nnama].noseri)
					fmt.Println("-------------------------------------------------------------------------------------")
				}
			} else if pilihpencarian == 2 {
				searchstatus(&T, &S, nk)
			} else {
				fmt.Println("Pilihan tidak valid, silakan pilih menu yang tersedia")
			}
		} else if pilih == 4 {
			sortingnoseridesc(&T, &S, nk)
			showdata(T, S, nk)
		} else if pilih == 5 {
			sortingnoseriasc(&T, &S, nk)
			showdata(T, S, nk)
		} else if pilih == 6 {
			statistik(T, S, nk)
		} else if pilih == 7 {
			fmt.Println("\nTERIMAKSIH\n")
			run = false
		} else {
			fmt.Println("Pilihan tidak valid, silakan pilih menu yang tersedia")
		}
	}
}
