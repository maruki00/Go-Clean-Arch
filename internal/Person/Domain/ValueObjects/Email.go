package valueobject

type Email struct {
	email string
}

func NewEmail(email string) *Email {
	if email == "" {
		panic("email should not be empty")
	}

	return &Email{
		email: email,
	}
}

func (obj *Email) String() string {
	return obj.email
}
