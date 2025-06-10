package dao

import (
	"utes-x-api/model"
	sqlboiler "utes-x-api/sqlboiler/entity"
)

func toUserModel(userDto sqlboiler.User) (*model.User, error) {
	user := model.User{
		ID:        model.MustParseID(userDto.ID),
		Username:  userDto.Username,
		Email:     userDto.Email,
		CreatedAt: userDto.CreatedAt,
		UpdatedAt: userDto.UpdatedAt.Ptr(),
	}
	return &user, nil
}

func toUserModelSlice(usersDto sqlboiler.UserSlice) ([]model.User, error) {
	users := make([]model.User, len(usersDto))
	for i, userDto := range usersDto {
		user, err := toUserModel(*userDto)
		if err != nil {
			return nil, err
		}
		users[i] = *user
	}
	return users, nil
}
