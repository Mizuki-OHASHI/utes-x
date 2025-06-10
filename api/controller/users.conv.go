package controller

import "utes-x-api/model"

func toUserResponse(user model.User) User {
	return User{
		Id:        user.ID.String(),
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: nil,
	}
}

func toUserResponseSlice(users []model.User) []User {
	res := make([]User, len(users))
	for i, user := range users {
		res[i] = toUserResponse(user)
	}
	return res
}
