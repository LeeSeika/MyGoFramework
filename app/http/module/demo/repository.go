package demo

type Repository struct {
}

func NewRepository() *Repository {
	return &Repository{}
}

func (r *Repository) GetUserIds() []int {
	return []int{1, 2}
}

func (r *Repository) GetUsersByIds([]int) []UserModel {
	model1 := UserModel{UserId: 1, Name: "Ana", Age: 25}
	model2 := UserModel{UserId: 2, Name: "Bob", Age: 24}
	return []UserModel{model1, model2}
}
