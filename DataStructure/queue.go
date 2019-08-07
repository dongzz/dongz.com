package main

import (
	"errors"
	"fmt"
)

/**
队列
先进先出
*/
type queue struct {
	data int
	next *queue
}

func newQueue() (q queue) {
	q = queue{}
	return
}

func (q *queue) addQueue(data int) {
	temp := q
	for temp.next != nil {
		temp = temp.next
	}
	temp.next = &queue{data: data}
}

func (q *queue) minusQueue() (data int, err error) {
	if q.next == nil {
		err = errors.New("queue is empty")
		return
	}
	data = q.next.data
	q.next = q.next.next
	return
}

func (q *queue) show() {
	temp := q
	for temp.next != nil {
		fmt.Print(temp.next.data, "\t")
		temp = temp.next
	}
	fmt.Println()
}

//数组模拟队列
type arrQueue struct {
	data [10]int
	font int
	rear int
}

func newArrQueue() (q arrQueue) {
	q.font, q.rear = 0, 0
	return
}

func (q *arrQueue) addArrQueue(data int) (err error) {
	if (q.rear-q.font+1)%10 == 0 {
		err = errors.New("queue is full")
		return
	}
	q.data[q.rear] = data
	q.rear++
	q.rear %= len(q.data)
	return
}

func (q *arrQueue) minusArrQueue() (data int, err error) {
	if q.rear == q.font {
		err = errors.New("queue is empty")
		return
	}
	data = q.data[q.font]
	q.data[q.font] = 0
	q.font++
	q.font %= len(q.data)
	return
}

func main() {
	firstNode := newQueue()

	arrQueue := newArrQueue()

	var check int
	flag := true

	for flag {
		fmt.Println("--------   1，add   ---------")
		fmt.Println("--------   2，minus   ---------")
		fmt.Println("--------   3，exit   ---------")
		fmt.Println("--------   pick    ---------")

		fmt.Scanln(&check)

		switch check {
		case 1:
			fmt.Println("give me you number")
			var num int
			fmt.Scanln(&num)
			err := arrQueue.addArrQueue(num)
			if err != nil {
				fmt.Println(err)
			}
			firstNode.addQueue(num)
			firstNode.show()
		case 2:
			_, err := arrQueue.minusArrQueue()
			if err != nil {
				fmt.Println("array ", err)
			}
			_, err = firstNode.minusQueue()
			if err != nil {
				fmt.Println("node ", err)
			}
		case 3:
			fmt.Println("exit system")
			flag = false
		default:
			fmt.Println("pick again")
			continue
		}
		fmt.Println(arrQueue)
		fmt.Println("-------------")
		firstNode.show()
	}
}
