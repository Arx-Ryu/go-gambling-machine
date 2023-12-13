package main

import (
	"fmt"
	"strings"
	"strconv"
)

func main() {
	var version int
	fmt.Println("Welcome to Gambling Machine!")
	fmt.Println(`Which version of Gambling Machine would you like to play?: 
	0 - Exit
	1 - Without Loop
	2 - With Loop`)
	fmt.Print("Version: ")
	fmt.Scanln(&version)

	if version == 1 {
		var temp int
		var num []int
		for i := 0; i < 3; i++ {
			fmt.Printf("Slot Number %d: ", i)
			fmt.Scanln(&temp)
			num = append(num, temp)
		}
		reply := gamblingMachine(num)
		fmt.Println(reply[0])
		fmt.Println(reply[1])
	} else if version == 2 {
		roll := 1
		balance := 5000
		var addict string
		for roll > 0 {
			var temp int
			var num []int
			for i := 0; i < 3; i++ {
				fmt.Printf("Slot Number %d: ", i)
				fmt.Scanln(&temp)
				num = append(num, temp)
			}
			slot := gamblingMachineV2(num)
			balance += slot
			if slot > 0 {
				fmt.Printf("You Win %d Dollars\n", slot)
				fmt.Printf("Your Total Credit Balance is %d Dollars\n", balance)
			} else {
				fmt.Printf("You Lose %d Dollars\n", (slot*-1))
				fmt.Printf("Your Total Credit Balance is %d Dollars\n", balance)
			}
			if balance <= 0 {
				fmt.Print("You are bankrupt!")
				roll = 0
			} else {
				fmt.Print("Roll Again? (y/n)")
				fmt.Scanln(&addict)
				if strings.Contains(addict, "n") {
					roll = 0
				}
			}		
		}
	} else if version == 0 {
		fmt.Print("Bye!")
	} else {
		//Error Handling
		fmt.Printf("Tidak ada Soal dengan Nomor %d", version)
	}
}

func gamblingMachine(num []int) []string {
	var slot int
	var reply []string
	var temp, tempB string
	balance := 5000
	if num[0] == num [1] && num[1] == num [2] {
		//Semua bilangan sama
		slot = (num[0] + num[1] + num[2]) * 200
		balance += slot
		temp = "You Win " + strconv.Itoa(slot) + " Dollars"
		tempB = "Your Total Credit Balance is " + strconv.Itoa(balance) + " Dollars"
		reply = append(reply, temp)
		reply = append(reply, tempB)
	} else if num[0] == num [1] || num[1] == num [2] || num[0] == num [2]{
		//1 Pasang bilangan sama
		slot = (num[0] + num[1] + num[2]) * 100
		balance += slot
		temp = "You Win " + strconv.Itoa(slot) + " Dollars"
		tempB = "Your Total Credit Balance is " + strconv.Itoa(balance) + " Dollars"
		reply = append(reply, temp)
		reply = append(reply, tempB)
	} else {
		//Tidak ada bilangan sama
		slot = (num[0] + num[1] + num[2]) * 50
		balance -= slot
		temp = "You Lose " + strconv.Itoa(slot) + " Dollars"
		tempB = "Your Total Credit Balance is " + strconv.Itoa(balance) + " Dollars"
		reply = append(reply, temp)
		reply = append(reply, tempB)
	}
	return reply
}

func gamblingMachineV2(num []int) int {
	var slot int
	if num[0] == num [1] && num[1] == num [2] {
		//Semua bilangan sama
		slot = (num[0] + num[1] + num[2]) * 200
	} else if num[0] == num [1] || num[1] == num [2] || num[0] == num [2]{
		//1 Pasang bilangan sama
		slot = (num[0] + num[1] + num[2]) * 100
	} else {
		//Tidak ada bilangan sama
		slot = (num[0] + num[1] + num[2]) * -50
	}
	return slot
}