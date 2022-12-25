package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

// type user struct {	结构体开头为小写，不是公有，修改为大写即可
type User struct {
	// username string
	// nickname string
	// sex      uint8
	// birthday time.Time
	Username string
	Nickname string
	Sex      uint8
	Birthday time.Time
}

func main() {
	// u := user{
	// 	sername: "坤坤",
	// 	nickname: "阿坤",
	// 	sex:      20,
	// 	birthday: time.Now(),
	// }
	u := User{
		Username: "坤坤",
		Nickname: "阿坤",
		Sex:      20,
		Birthday: time.Now(),
	}
	bs, err := json.Marshal(&u)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(string(bs))
}
