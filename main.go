package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	getActRoomID(3, "2123213")
}

func getActRoomID(actType int8, actID string) (roomID string, err error) {
	roomID = fmt.Sprintf("%d%04d%02d", time.Now().Unix(), rand.Intn(10000), actType)
	fmt.Println(roomID)
	return
}
