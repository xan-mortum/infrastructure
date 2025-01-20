package store

type Books struct {
}

func NewBooks() *Books {
	return &Books{}
}

func (b *Books) GetBooks() map[int]string {
	return map[int]string{
		1: "book1",
		2: "book2",
		3: "book3",
		4: "book4",
		5: "book5",
	}
}
