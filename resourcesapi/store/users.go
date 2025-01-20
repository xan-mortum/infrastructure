package store

type Users struct {
}

func NewUsers() *Users {
	return &Users{}
}

func (u *Users) GetUsers() map[int]string {
	return map[int]string{
		1: "user1",
		2: "user2",
		3: "user3",
		4: "user4",
		5: "user5",
	}
}
