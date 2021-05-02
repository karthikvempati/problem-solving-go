package main

import (
	"fmt"

	"github.com/karthikvempati/data-structures-go/queue"
)

func main() {
	//Successful Orders
	CheckOrders([]int{1, 2, 3}, []int{12, 4, 67}, []int{1, 12, 2, 4, 3, 67})

	//Failed Orders
	CheckOrders([]int{1, 2, 3}, []int{12, 4, 67}, []int{1, 4, 2, 12, 3, 67})
}

func CheckOrders(dineIn []int, takeOut []int, served []int) {
	if len(dineIn) == 0 || len(takeOut) = 0{
		fmt.Println("Orders served successfully")
		return
	}

	//Check if dineIn + takeOut Length is equal to served

	dineInQueue := CreateQueue(dineIn)
	takeOutQueue := CreateQueue(takeOut)
	for _, val := range served {
		dineInCurrent := dineInQueue.Peek()
		takeOutCurrent := takeOutQueue.Peek()
		if dineInCurrent == val {
			_ = dineInQueue.DeQueue()
			continue
		} else if takeOutCurrent == val {
			_ = takeOutQueue.DeQueue()
			continue
		} else {
			fmt.Println("Error: Orders served incorrectly")
			return
		}
	}

	fmt.Println("Orders served successfully")
}

func CreateQueue(orders []int) queue.Queue {
	q := queue.Queue{}
	for _, value := range orders {
		q.EnQueue(value)
	}
	return q
}
