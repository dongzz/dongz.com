package main

import (
	"errors"
	"fmt"
	"os"
	"time"
)

/**
单向环形链表
*/
type circelQueue struct {
	id       int
	name     string
	nextNode *circelQueue
}

func newCircelQueue(id int, name string) (q *circelQueue) {
	q = &circelQueue{id: id, name: name}
	q.nextNode = q
	return
}

//新增节点
func (q *circelQueue) insertNode(id int, name string) (err error) {
	if q.hasNode(id) {
		err = errors.New("ID already exist ")
	}
	newNode := circelQueue{id: id, name: name, nextNode: q.nextNode}
	q.nextNode = &newNode
	return
}

//删除节点
func deleteNode(q *circelQueue, id int) (newQueue *circelQueue, err error) {
	node, err := q.getNode(id)
	if err != nil {
		return
	}

	temp := q
	for temp.nextNode != node {
		temp = temp.nextNode
	}

	if q == node {
		newQueue = node.nextNode
		temp.nextNode = newQueue
		q = nil
	} else {
		temp.nextNode = node.nextNode
	}

	node = nil
	return
}

//获取节点
func (q *circelQueue) getNode(id int) (node *circelQueue, err error) {
	temp := q
	for temp.id != id {
		temp = temp.nextNode
		if temp == q {
			err = errors.New("ID don't exist ")
			return
		}
	}
	node = temp
	return
}

//是否存在节点
func (q *circelQueue) hasNode(id int) (flag bool) {
	temp := q
	flag = false
	for temp.id != id {
		temp = temp.nextNode
		if temp.id == q.id {
			return
		}
	}
	flag = true
	return
}

func (q *circelQueue) show() {
	temp := q
	fmt.Printf("head -> id: %v,name: %v \n", q.id, q.name)
	for temp.nextNode != q {
		temp = temp.nextNode
		fmt.Printf("nextNode -> id: %v,name: %v \n", temp.id, temp.name)

	}
}

func main() {
	//创建一个头节点
	head := newCircelQueue(1, "1号")

	head.insertNode(2, "2号")
	head.insertNode(3, "3号")
	head.insertNode(4, "4号")

	head, _ = deleteNode(head, 1)
	head.show()

	time.Sleep(4)
	os.Exit(0)
}
