package chapter6_struct

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Flying struct {
	Name string
}

func (f *Flying) Fly() {
	fmt.Println(f.Name, "can fly")
}

type Walkable struct {
	Name string
}

func (w *Walkable) Walk() {
	fmt.Println(w.Name, "can walk")
}

type Human struct {
	Walkable
}

type Bird struct {
	Flying
	Walkable
}

func TestInherit(t *testing.T) {
	human := new(Human)
	human.Name = "human"
	human.Walk()

	bird := new(Bird)
	bird.Walkable.Name = "bird"
	bird.Walk()
	bird.Flying.Name = "bird"
	bird.Fly()
}

type Wheel struct {
	Size int
}

type Engine struct {
	Power  int
	Energy int
}

type Car struct {
	Wheel
	Engine
}

func TestCar(t *testing.T) {
	car := new(Car)
	car.Size = 4
	car.Power = 200
	fmt.Println(car)

	car1 := Car{
		Wheel: Wheel{
			Size: 4,
		},
		Engine: Engine{
			Power:  200,
			Energy: 10,
		},
	}

	fmt.Printf("%+v \n", car1)

	json, _ := json.Marshal(car1)
	fmt.Printf("%s", json)
}
