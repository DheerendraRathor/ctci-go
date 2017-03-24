package p6

import (
    "time"
    "github.com/DheerendraRathor/ctci-go/collections"
)

type AnimalType uint

const (
    Dog AnimalType = iota
    Cat
)

type Animal struct {
    arrivalTime time.Time
    animalType AnimalType
}

type AnimalShelter struct {
    dogQueue collections.Queue
    catQueue collections.Queue
}

func (shelter *AnimalShelter) Enqueue(animalType AnimalType) {
    animal := Animal{arrivalTime: time.Now().UTC(), animalType: animalType}
    switch animalType {
    case Cat:
        shelter.catQueue.Add(animal)
    case Dog:
        shelter.dogQueue.Add(animal)
    default:
        panic("Unknown Animal")
    }
}

func (shelter *AnimalShelter) DequeueAny() (Animal, *collections.DataStructureEmptyError) {
    cat, _ := shelter.catQueue.Peek()
    dog, _ := shelter.dogQueue.Peek()

    if shelter.catQueue.IsEmpty() {
        if shelter.dogQueue.IsEmpty() {
            return Animal{}, &collections.DataStructureEmptyError{}
        }

        shelter.dogQueue.Remove()
        return dog.(Animal), nil
    }

    if shelter.dogQueue.IsEmpty() {     // Cat queue is not empty but dog queue is. Purrrr
        shelter.catQueue.Remove()
        return cat.(Animal), nil
    }

    catAnimal := cat.(Animal)
    dogAnimal := dog.(Animal)

    if catAnimal.arrivalTime.Before(dogAnimal.arrivalTime) {
        shelter.catQueue.Remove()
        return catAnimal, nil
    }

    shelter.dogQueue.Remove()
    return dogAnimal, nil
}

func (shelter *AnimalShelter) DequeueDog() (Animal, *collections.DataStructureEmptyError) {

    if shelter.dogQueue.IsEmpty() {
        return Animal{}, &collections.DataStructureEmptyError{}
    }

    dog, _ := shelter.dogQueue.Remove()
    dogAnimal := dog.(Animal)

    return dogAnimal, nil
}

func (shelter *AnimalShelter) DequeueCat() (Animal, *collections.DataStructureEmptyError) {

    if shelter.catQueue.IsEmpty() {
        return Animal{}, &collections.DataStructureEmptyError{}
    }

    cat, _ := shelter.catQueue.Remove()
    catAnimal := cat.(Animal)

    return catAnimal, nil
}
