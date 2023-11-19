package status

var Success = Status{
	Code:    200,
	Message: "Success",
}

var UnknownError = Status{
	Code:    500,
	Message: "Unknown Error",
}
