package main

import "fmt"
import "math"
import "os"

var kembali string

type hitung interface {
luas() float64
keliling() float64
}
	type segitiga struct {
	alas, tinggi  float64
	}

		func (S segitiga) luas() float64 {  
	return 0.5 * S.alas * S.tinggi
	}
	
		func (S segitiga) keliling() float64 {
	return S.alas * 3
		}

	type persegi struct {
	sisi float64
	}

		func (p persegi) luas() float64 {
		return math.Pow(p.sisi, 2)
		}

		func (p persegi) keliling() float64 {
		return p.sisi * 4
		}

func back(){
  fmt.Println("========================")
  fmt.Print("Kembali ke menu awal? [Y/N]")
  fmt.Scan(&kembali)
  if kembali == "Y" {
    main()
  } else if kembali == "N" {
     fmt.Println("Thanks")
     os.Exit(1)
  }else{
    fmt.Print("Anda salah Input !!!")
  }
}


func main() {
var menu int
var side, sides, large float64
var bangunDatar hitung
  	fmt.Printf("============================\n")
  	fmt.Printf("Program untuk menghitung :\n")
 	fmt.Printf("============================\n")
	fmt.Printf("1. Persegi\n")
 	fmt.Printf("2. Segitiga\n")
 	fmt.Printf("3. Exit\n")
 	fmt.Printf("===========================\n") 
  fmt.Scan(&menu)
  if menu == 1{
    fmt.Println("=====Persegi======")
    fmt.Println("Sisi persegi: ")
    fmt.Scan(&side)
    bangunDatar = persegi{side}
	fmt.Println("luas :", bangunDatar.luas())
	fmt.Println("keliling :", bangunDatar.keliling())
	back()
  } else if menu == 2{
  	fmt.Println("Alas Segitiga : ")
    fmt.Scan(&sides)
    fmt.Println("Luas Segitiga : ")
    fmt.Scan(&large)
    bangunDatar = segitiga{sides,large}
	fmt.Println("=====Segitiga=====")
	fmt.Println("luas :", bangunDatar.luas())
	fmt.Println("keliling :", bangunDatar.keliling())
	back()
  }else if menu == 3 {
     fmt.Println("Anda Keluar")
     os.Exit(1)
  } else {
    fmt.Println("Tidak Tersedia")
  }  

}
