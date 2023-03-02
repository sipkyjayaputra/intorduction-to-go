package main

import (
	"log"
	"sync"
)

func main() {
	// how to define a variable
	// var a string    //define
	// a = "air putih" //assign
	// log.Println(a)

	// var b int64 = 2 // define and assign
	// log.Println(b)

	// c := 2 // assign
	// log.Println(c)

	// // how to define a model / data structure
	// type User struct {
	// 	Name string
	// 	Age  int64
	// }
	// var user User
	// user = User{
	// 	Name: "Paijo",
	// 	Age:  20,
	// }
	// // log.Println(user)

	// var user2 User = User{
	// 	Name: "Paimin",
	// 	Age:  21,
	// }
	// // log.Println(user2)

	// user3 := User{
	// 	Name: "Sipky",
	// 	Age:  22,
	// }
	// log.Println(user3)

	// how to define an array
	// arrayString := []string{"data1", "data2"}
	// string1 := arrayString[1]
	// log.Println(string1)

	// arrayUser := []User{user, user2, user3}
	// user1 := arrayUser[2]
	// log.Println(user1)

	// how to build enum in golang
	// type Status string
	// const (
	// 	available    Status = "Available"
	// 	notAvailable Status = "Not Available"
	// )
	// log.Println(available)

	// how to define a function
	// total := Tambah(10, 20)
	// fmt.Println("print disini: ", total)

	// contract function
	// bisa membuat function menjadi private
	// kucing1 := Kucing{
	// 	Kaki: 4,
	// 	Mata: 2,
	// 	Nama: "Helli",
	// }
	// kucing1.Minum()

	// looping in go
	// for loop
	// for i := 0; i < 100; i++ {
	// 	log.Println(i)
	// }
	// arrayInt := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// for j := 0; j < len(arrayInt); j++ {
	// 	log.Println(arrayInt[j])
	// }
	// for _, val := range arrayInt {
	// 	log.Println(" value: ", val)
	// }
	// while do or do while
	// loopIdx := 0
	// loop := true
	// for loop {
	// 	log.Println(loopIdx)
	// 	loopIdx++
	// 	if loopIdx == 10 {
	// 		loop = !loop
	// 	}
	// }

	// conditional in go
	// if else
	// minum := "sirup"
	// if minum == "kopi" {
	// 	log.Println("begadang")
	// } else {
	// 	log.Println("tidur")
	// }
	// switch case
	// switch minum {
	// case "kopi":
	// 	log.Println("begadang")
	// case "sirup":
	// 	log.Println("seger")
	// default:
	// 	log.Println("tidur")
	// }

	// private and public
	// log.Println(helper.Angka1)
	// log.Println(helper.tambah(2, 5))

	// how to use go routine
	// memungkinkan menjalankan proses sekaligus dalam waktu bersamaan
	angkaChan := make(chan int)
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go cetak(angkaChan, &wg)
		angkaChan <- i
	}
	wg.Wait()

	// how to use channel
	// channel a
	// orang 1 dan orang 2
	// orang 1 -> aku cinta kamu -> orang 2
	// kamu cinta aku, cinta kamu aku, aku kamu cinta

	// combine goroutine with channel

	// gin
	// gorm
	// mysql
}

func cetak(c chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	angka := <-c
	log.Println("angka: ", angka)
}
