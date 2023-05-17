package main

import "fmt"

//type Room uint8

// 비트에 의미 부여
const (
	MasterRoom uint8 = 1 << iota
	LivingRoom
	BathRoom
	SmallRoom
)

// bit 플래그 사용시 아래 3개 함수는 자주 사용하는 패턴이다.
func SetLight(rooms, room uint8) uint8 {
	// BitFlag 켜지는 효과
	return rooms | room
}

func ResetLight(rooms, room uint8) uint8 {
	// Bit Clear - 해당 bit 값을 0으로 변경
	return rooms &^ room
}

// isLightOn 없이 사용하려면 플래그 개수 만큼의 boolean 변수를 사용해야 한다.
func IsLightOn(rooms, room uint8) bool {
	return rooms&room == room
}

//

func TurnLights(rooms uint8) {
	if IsLightOn(rooms, MasterRoom) {
		fmt.Println("Turn on MasterRoom Light")
	}
	if IsLightOn(rooms, BathRoom) {
		fmt.Println("Turn on BathRoom Light")
	}
	if IsLightOn(rooms, SmallRoom) {
		fmt.Println("Turn on Kitchen Light")
	}
	if IsLightOn(rooms, LivingRoom) {
		fmt.Println("Turn on LivingRoom Light")
	}
}

func main() {
	var rooms uint8 = 0
	rooms = SetLight(rooms, MasterRoom)
	rooms = SetLight(rooms, BathRoom)
	rooms = SetLight(rooms, SmallRoom)

	rooms = ResetLight(rooms, SmallRoom)

	TurnLights(rooms)
}
