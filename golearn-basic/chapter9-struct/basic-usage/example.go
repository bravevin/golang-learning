package main

import "fmt"

type Qimiao struct {
	Name    string
	Sex     bool
	Age     int
	Hobbies []string
}

func main() {
	var qm Qimiao
	qm.Name = "qm"
	qm.Sex = false
	qm.Age = 24
	qm.Hobbies = []string{"玩游戏", "唱歌"}

	qm2 := Qimiao{
		"qm2", false, 48, []string{"sing", "play"},
	}
	qm3 := Qimiao{
		Name:    "qm3",
		Sex:     true,
		Age:     35,
		Hobbies: nil,
	}
	fmt.Println(qm, qm2, qm3)

	qm4 := new(Qimiao)
	qm4.Name = "QM4"
	a := qm4
	a.Age = 26
	fmt.Println(qm4, a)

	var qm5 *Qimiao
	qm5 = &qm2
	// qm5 := qm2

	qm5.Name = "qm5"
	(*qm5).Name = "newQm5"
	fmt.Println(qm2, qm5)
}
