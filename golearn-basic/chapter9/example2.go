package main

import "fmt"

type Qimiao2 struct {
	Name    string
	Sex     bool
	Age     int
	Hobbies []string
	MyHome  Home
	Home
}

type Home struct {
	address string
}

func (h *Home) Open() {
	fmt.Println("open " + h.address)
}

func (q *Qimiao2) Song(song string) (retstr string) {
	retstr = "I am the return value"
	fmt.Println(q.Name + "唱了一首" + song)
	return
}

func main() {
	var qm Qimiao2
	qm.Name = "Vincent"
	qm.MyHome.address = "Peking"
	qm.MyHome.Open()
	qm.address = "Shanghai"
	qm.Open()
	ret1 := qm.Song("qing hua ci")
	fmt.Println(ret1)
}
