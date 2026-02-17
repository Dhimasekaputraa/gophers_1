package main

import (
	"fmt"

	"github.com/fatih/color"
)

func Hitung(harga float64, uang float64) {
	var kembalian float64
	var kurang float64

	kembalian = uang - harga
	kurang = harga - uang

	if uang == harga {
		color.Green("[SISTEM]: Transaksi berhasil : Uang anda pas")
	} else if uang > harga {
		color.Green("[SISTEM]: Transaksi berhasil : Kembalian anda %.2f", kembalian)
	} else {
		color.Red("[SISTEM]: Transaksi gagal : Uang lu kurang %.2f", kurang)
	}
}

func main() {
	var uang float64
	var harga float64

	fmt.Print("Harga barang : ")
	fmt.Scan(&harga)
	fmt.Print("uang Pembeli : ")
	fmt.Scan(&uang)

	Hitung(harga, uang)
}
