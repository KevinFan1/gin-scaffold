package vo

type ApiError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (a *ApiError) Error() string {
	return a.Message
}
