package errorz

var ErrAuthWrongCredentials = Error_{
	StatusCode: 401,
	Message:    "Wrong credentials",
}

var ErrAuthInvalidToken = Error_{
	StatusCode: 401,
	Message:    "Token is invalid",
}
