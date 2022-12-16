package main

type UserApi struct {
	storage []*User
}

type User struct {
	Id       string
	Email    string
	FullName string
}

type UpdateUserRequest struct {
	Id       string
	FullName *string
	Email    *string
}

type Error string

func (e Error) Error() string { return string(e) }

var (
	UserNotFound      Error = "not_found"
	ErrNotImplemented Error = "not_implemented"
)

func (api *UserApi) Update(request UpdateUserRequest) (*User, error) {
	//TODO: Implement here
	return nil, ErrNotImplemented
}

type apiTestCase struct {
	users  []*User
	input  UpdateUserRequest
	output *User
	err    error
}

func pointer(v string) *string {
	return &v
}
