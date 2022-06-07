package users

type publicUser struct {
	ID      int64  `json:"id"`
	Created string `json:"created"`
	Status  string `json:"status"`
}

type privateUser struct {
	ID        int64  `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Created   string `json:"created"`
	Status    string `json:"status"`
}

// Marshal is a method that returns a public or private user
func (user *User) Marshal(isPublic bool) interface{} {
	if isPublic {
		return publicUser{
			ID:      user.ID,
			Created: user.Created,
			Status:  user.Status,
		}
	}
	return privateUser{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Created:   user.Created,
		Status:    user.Status,
	}
}

// Marshal is a method that returns a public or private users
func (users Users) Marshal(isPublic bool) []interface{} {
	result := make([]interface{}, len(users))
	for index, user := range users {
		result[index] = user.Marshal(isPublic)
	}
	return result
}
