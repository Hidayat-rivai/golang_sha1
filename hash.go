package main

import "fmt"
import "crypto/sha1"

func main(){
	var text = "Hidayat"
	var sha = sha1.New()
	sha.Write([]byte(text))
	var enkripsi = sha.Sum(nil)
	var akhir = fmt.Sprintf("%x",enkripsi)
	
	fmt.Println(akhir)
}