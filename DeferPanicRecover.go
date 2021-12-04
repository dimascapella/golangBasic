package main

func main() {
	runApp()

	run(0)

	runApps(true)
}

// Defer Function Merupakan fungsi yang terjadwalkan dan akan tetap dieksekusi walapun terjadi error
func log() {
	println("Selesai Memanggil Function")
}

func runApp() {
	defer log()
	println("Run App")
}

// Pannic Function Merupakan fungsi yang digunakan untuk menghentikan program dan dipanggil waktu terjadi error, Defer akan tetap berjalan walaupun terjadi error

func debugging() {
	println("End Program")
	message := recover()
	println("Error Message ,", message)
}

func run(number int) {
	defer debugging()
	if number == 0 {
		panic("Tidak Boleh 0")
	} else {
		println(number)
	}
}

// Recover Function Digunakan Untuk Menangkap Data Panic Function Sehingga Program Akan Tetap Berjalan
func debug() {
	message := recover()
	println("Error Message ,", message)
}

func runApps(error bool) {
	defer debug()
	if error {
		panic("Error App")
	}
}
