package error

const (
	UnknownError = iota + 1
	RequestError
	DatabaseError
	UnauthorizedError
	InternalServerError
)

const (
	InternalUserNotFound = "internal user not found"
)
