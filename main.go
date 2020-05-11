package main

import "fmt"

func main()  {
	//duration := []int{1, 1, 5, 1, 1}
	//var persons []person
	var jmlLoket int

	fmt.Print("Masukkan jumlah loket:")
	_, err := fmt.Scanf("%d", &jmlLoket)
	if err != nil {
		fmt.Println("Error Input :", err.Error())
	}
}