package main

import "fmt"

func main() {
	var obj, num float32
	var err error
	fmt.Print("Вкажіть вартість товару: ")
	_, err = fmt.Scanln(&obj)
	if err != nil {
		fmt.Println("Вкажіть вартість товару!")
		return
	}
	fmt.Print("Вкажіть кількість одиниць товару: ")
	_, err = fmt.Scanln(&num)
	if err != nil {
		fmt.Println("Вкажіть кількість одиниць товару!")
		return
	}

	fmt.Printf("Загальна вартість покупки= %.2f\n", obj*num*1.07)

	//switch {
	//case year > age:
	//	fmt.Printf("Вам %d%s\n", year-age, " років(рік)")
	//case age > year:
	//	fmt.Println("Введений рік народження більше ніж поточний рік")
	//default:
	//	fmt.Println("Помилка вводу!")
	//}
}
