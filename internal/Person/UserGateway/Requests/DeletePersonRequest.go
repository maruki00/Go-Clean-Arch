package requests

type DeletePersonRequest struct {
	Id int `validate:"required"`
}
