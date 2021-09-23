package main

import "fmt"

type Hero struct {
	On bool
	Ammo,
	Power int
}

func (h *Hero) Shoot() bool {
	if h.On == false {
		return false
	} else {
		if h.Ammo > 0 {
			h.Ammo--
			return true
		} else {
			return false
		}
	}
}
func (h *Hero) RideBike() bool {
	if h.On == false {
		return false
	} else {
		if h.Power > 0 {
			h.Power--
			return true
		} else {
			return false
		}
	}
}
func main() {

	testStruct := &Hero{}
	/*
	 * Экземпляр созданной вами структуры необходимо передать в качестве
	 * аргумента функции testStruct, которая выполнит проверку соблюдения
	 * всех условий задания/
	 */
	fmt.Println(testStruct)
}
