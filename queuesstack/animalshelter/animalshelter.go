package animalshelter

import (
	"container/list"
	"errors"
)

const (
	Dog = iota
	Cat
)

type Animal struct {
	Name string
	Type int
}
type AnimalShelter struct {
	animal *list.List
}

func NewAnimalShelter() *AnimalShelter {
	return &AnimalShelter{
		animal: list.New(),
	}
}
func (as *AnimalShelter) Enqueue(animal Animal) error {
	if as.animal != nil {
		as.animal.PushBack(animal)
		return nil
	}
	return errors.New("queue not initialized")
}
func (as *AnimalShelter) DequeAny() (Animal, error) {
	animal, err := Animal{}, errors.New("animal not found")
	element := as.animal.Back()
	if element != nil {
		item := element.Value.(Animal)
		as.animal.Remove(element)
		return item, nil
	}
	return animal, err
}
func FindIf(l *list.List, op func(e *list.Element) bool) (Animal, bool) {
	var item Animal
	var found bool = false
	for e := l.Back(); e != nil; e = e.Prev() {
		if e != nil {
			found = op(e)
			if found {
				item = e.Value.(Animal)
				return item, found
			}
		}
	}
	return item, found
}
func (as *AnimalShelter) DequeDog() (Animal, error) {
	item, found := FindIf(as.animal, func(e *list.Element) bool {
		var item Animal
		item = e.Value.(Animal)
		if item.Type == Dog {
			return true
		}
		return false
	})
	if found {
		return item, nil
	}
	return Animal{}, errors.New("animal not found")
}
func (as *AnimalShelter) DequeCat() (Animal, error) {
	var item Animal
	var found bool
	for e := as.animal.Back(); e != nil; e = e.Prev() {
		if e != nil {
			item = e.Value.(Animal)
			if item.Type == Cat {
				found = true
				break
			}
		}
	}
	if found {
		return item, nil
	}
	return Animal{}, errors.New("animal not found")
}
