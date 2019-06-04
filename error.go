package music

var InvalidID = Error{
	Code:   "INVALID_ID",
	Detail: "The ID is not valid",
}

var UserNotFound = Error{
	Code:   "USER_NOT_FOUND",
	Detail: "Artist not found",
}

var ServerError = Error{
	Code:   "SERVER_ERROR",
	Detail: "Sorry, something went wrong",
}

type Error struct {
	Code   string      `json:"code"`
	Detail string      `json:"detail"`
	Meta   interface{} `json:"meta,omitempty"`
}

func (r Error) Error() string {
	return r.Code
}
