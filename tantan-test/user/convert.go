package user

func (u *User) Convert() interface{} {
	return &struct {
		User
		Type string `json:"type"`
	}{
		*u,
		"user",
	}
}

func (l *UserList) Convert() interface{} {
	converter := []interface{}{}
	for _, user := range *l {
		converter = append(converter, user.Convert())
	}

	return converter
}
