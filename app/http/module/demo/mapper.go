package demo

import "github.com/godemo/coredemo/app/provider/demoProvider"

func StudentsToUserDTOs(students []demoProvider.Student) []UserDTO {
	ret := []UserDTO{}
	for _, student := range students {
		t := UserDTO{
			ID:   student.ID,
			Name: student.Name,
		}
		ret = append(ret, t)
	}
	return ret
}

func UserModelsToUserDTOs(models []UserModel) []UserDTO {
	ret := []UserDTO{}
	for _, user := range models {
		t := UserDTO{
			ID:   user.UserId,
			Name: user.Name,
		}
		ret = append(ret, t)
	}
	return ret
}
