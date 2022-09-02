package advanced

import (
	"encoding/json"
	"fmt"
	"log"
)

type person struct {
	First string
}

// Marshal func
func marshal() {
	p1 := person{
		First: "A",
	}
	p2 := person{
		First: "B",
	}

	xp := []person{p1, p2}

	bs, err := json.Marshal(xp) // dịch struct thành json
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(string(bs)) // [{"First":"A"},{"First":"B"}]

	xp2 := []person{}

	err = json.Unmarshal(bs, &xp2) // dịch json thành struct
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(xp2) // [{A} {B}]
}
