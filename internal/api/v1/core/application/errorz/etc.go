package errorz

var ErrDatabaseError = Error_{
	StatusCode: 503,
	Message:    "Database internal error",
}

var ErrServiceUnavailable = Error_{
	StatusCode: 503,
	Message:    "Service unavailable",
}
