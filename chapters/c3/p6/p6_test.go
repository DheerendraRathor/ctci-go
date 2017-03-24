package p6

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "time"
)

func TestAnimalShelter_Enqueue(t *testing.T) {
    shelter := AnimalShelter{}

    shelter.Enqueue(Dog)
    shelter.Enqueue(Cat)
    shelter.Enqueue(Cat)
    shelter.Enqueue(Cat)

    assert.Equal(t, 1, shelter.dogQueue.Length())
    assert.Equal(t, 3, shelter.catQueue.Length())

    defer func() {
        if r := recover(); r == nil {
            assert.Fail(t, "Expected panic but got nil")
        }
    }()
    shelter.Enqueue(3)
}

func TestAnimalShelter_Dequeue(t *testing.T) {
    shelter := AnimalShelter{}

    shelter.Enqueue(Dog)
    time1 := time.Now().UTC()
    shelter.Enqueue(Cat)
    time2 := time.Now().UTC()
    shelter.Enqueue(Cat)
    time3 := time.Now().UTC()
    shelter.Enqueue(Dog)
    time4 := time.Now().UTC()

    animal, _ := shelter.DequeueAny()
    assert.Equal(t, Dog, animal.animalType)
    assert.True(t, time1.After(animal.arrivalTime))

    animal, _ = shelter.DequeueCat()
    assert.Equal(t, Cat, animal.animalType)
    assert.True(t, time2.After(animal.arrivalTime))
    assert.True(t, time1.Before(animal.arrivalTime))

    animal, _ = shelter.DequeueAny()
    assert.Equal(t, Cat, animal.animalType)
    assert.True(t, time3.After(animal.arrivalTime))
    assert.True(t, time2.Before(animal.arrivalTime))

    _, err := shelter.DequeueCat()
    assert.NotNil(t, err)

    animal, _ = shelter.DequeueAny()
    assert.Equal(t, Dog, animal.animalType)
    assert.True(t, time4.After(animal.arrivalTime))
    assert.True(t, time3.Before(animal.arrivalTime))

    animal, err = shelter.DequeueAny()
    assert.NotNil(t, err)

    animal, err = shelter.DequeueDog()
    assert.NotNil(t, err)

    shelter.Enqueue(Dog)
    animal, err = shelter.DequeueDog()
    assert.Nil(t, err)
    assert.Equal(t, Dog, animal.animalType)

    shelter.Enqueue(Cat)
    animal, err = shelter.DequeueAny()
    assert.Nil(t, err)
    assert.Equal(t, Cat, animal.animalType)
}
