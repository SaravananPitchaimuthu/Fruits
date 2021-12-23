package domain

type FruitRepositoryStub struct {
	fruits []Fruit
}

func (f FruitRepositoryStub) FindAll() ([]Fruit, error) {
	return f.fruits, nil
}

func NewFruitRepositoryStub() FruitRepositoryStub {
	fruits := []Fruit{
		{"1001", "Apple", "20", "300", "0"},
		{"1001", "Orange", "30", "500", "1"},
	}
	return FruitRepositoryStub{fruits}
}
