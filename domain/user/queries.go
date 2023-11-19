package user

func GetUser(id int64) User {
	for _, user := range data {
		if user.ID == id {
			return user
		}
	}

	return User{}
}
