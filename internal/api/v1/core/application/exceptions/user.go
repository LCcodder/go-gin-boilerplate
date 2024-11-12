package exceptions

var ErrUserAlreadyExists = Error_{
	StatusCode: 400,
	Message:    "User already exists",
}
var ErrUserNotFound = Error_{
	StatusCode: 404,
	Message:    "User not found",
}
