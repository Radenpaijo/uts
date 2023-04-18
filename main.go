package main

import (
	"fmt"
)

type Data struct {
	saldo         int
	nomerRekening int
	username      string
}

type Nasabah struct {
	data Data
	next *Nasabah
}

var rek int = 2334

func tambahNasabah(data *Nasabah, username string, saldo int) int {
	dummy := Nasabah{
		data: Data{
			saldo:         saldo,
			nomerRekening: rek,
			username:      username,
		},
	}

	if data.next == nil {
		data.next = &dummy
	} else {
		dummy.next = data.next
		data.next = &dummy
	}

	return rek
}

func login(data *Nasabah, rekening int, username string) *int {
	temp := data.next

	for temp != nil {
		if rekening == temp.data.nomerRekening && username == temp.data.username {
			return &temp.data.nomerRekening
		}
		temp = temp.next
	}

	return nil
}

func tarikSaldo(data *Nasabah, noRek int, tarik int) string {
	temp := data.next

	fmt.Println(noRek)
	fmt.Println(tarik)

	for temp != nil {
		if noRek == temp.data.nomerRekening {
			if tarik > temp.data.saldo {
				return "saldo anda tidak cukup"
			}
			temp.data.saldo = temp.data.saldo - tarik
			break
		}
		temp = temp.next
	}

	return "saldo anda telah di tarik"
}

func tambahSaldo(data *Nasabah, noRek int, saldo int) string {
	temp := data.next

	for temp != nil {
		if noRek == temp.data.nomerRekening {
			if saldo == 0 {
				return "Saldo tidak boleh 0"
			}

			temp.data.saldo += saldo
			break
		}
		temp = temp.next
	}

	return "Saldo anda sudah di tambahkan"
}

func cekSaldo(data *Nasabah, aw int) int {
	temp := data.next

	fmt.Println(aw)

	for temp != nil {
		if temp.data.nomerRekening == aw {
			return temp.data.saldo
		}
		temp = temp.next
	}

	return 0
}

func HapusAtm(data *Nasabah, token int) string {
	temp := data.next
	last := data

	for temp != nil {
		if temp.data.nomerRekening == token {
			if temp == data.next {
				data.next = temp.next
			} else {
				last.next = temp.next
			}
			break
		}
		last = temp
		temp = temp.next
	}

	return "Akun anda sudah di hapus"
}

func cekUser(data *Nasabah) {
	temp := data.next

	for temp != nil {
		fmt.Println("nama nasabah : ", temp.data.username)
		fmt.Println("rekening nasabah : ", temp.data.nomerRekening)
		temp = temp.next
	}
}

func main(){
	var menu int
	var username string
	var saldo int
	var tarik int
	var tambah int
	var nomer int
	var pilihAtm int
	var isi int
	data := Nasabah{}
	for {
		fmt.Println("MAIN MENU")
		fmt.Println("1. Daftar akun")
		fmt.Println("2. Masuk")
		fmt.Println("3. Cek semua nasabah")
		fmt.Println("4. Exit")
		fmt.Print("pilih : ")
	
		fmt.Scan(&menu)
		
		if menu == 1 {
	
			fmt.Print("Masukan nama anda : ")
			fmt.Scan(&username)
			fmt.Print("Masukan saldo anda : ")
			fmt.Scan(&saldo)
	
			rekening := tambahNasabah(&data, username, saldo)
			rek++

			fmt.Println("Akun anda sudah di buat, ini nomer rekening anda : ", rekening)

		}else if menu == 2 {

			fmt.Print("Masukan username : ")
			fmt.Scan(&username)
			fmt.Print("Masukan nomer Rekening : ")
			fmt.Scan(&nomer)

			token := login(&data, nomer, username)
			
			if token != nil {
				for {
					fmt.Println("1. tarik saldo")
					fmt.Println("2. tambah saldo")
					fmt.Println("3. cek saldo")
					fmt.Println("4. Hapus akun")
					fmt.Println("5. exit")
					fmt.Print("Pilih : ")
					fmt.Scan(&pilihAtm)
	
					if pilihAtm == 1 {

						fmt.Print("Masukan saldo yang mau anda tarik : ")
						fmt.Scan(&tarik)

						fmt.Println(tarikSaldo(&data, *token, tarik))

					}else if pilihAtm == 2 {
	
						fmt.Print("Masukan saldo yang mau anda tambah : ")
						fmt.Scan(&tambah)

						fmt.Println(tambahSaldo(&data, *token, tambah))
					}else if pilihAtm == 3 {

						isi = cekSaldo(&data, *token)
						fmt.Println("Saldo anda sebesar Rp.", isi)
	
					}else if pilihAtm == 4 {
						fmt.Println(HapusAtm(&data, *token))
					}else if pilihAtm == 5 {
						break
					}else {
						fmt.Println("Format tidak di temukan")
					}
				}

			}else {
				fmt.Println("Akun tidak di temukan")
			}
		}else if menu == 3 {
			cekUser(&data)
		}else if menu == 4 {
			break
		}else {
			fmt.Println("Format tidak di temukan")
		}
	}
}