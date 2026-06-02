package response

// Error Responses for swagger annotation

type BadReqErrorResponse struct {
	Code  int    `json:"code" example:"400"`
	Error string `json:"error" example:"bad request: incorrect format"`
}

type StatusUnauthorizedErrorResponse struct {
	Code  int    `json:"code" example:"401"`
	Error string `json:"error" example:"unauthorized"`
}

type StatusForbiddenErrorResponse struct {
	Code  int    `json:"code" example:"403"`
	Error string `json:"error" example:"forbidden"`
}

type NotFoundErrorResponse struct {
	Code  int    `json:"code" example:"404"`
	Error string `json:"error" example:"user not found"`
}

type StatusConflictErrorResponse struct {
	Code  int    `json:"code" example:"409"`
	Error string `json:"error" example:"user already exists"`
}

type StatusUnprocessableEntityErrorResponse struct {
	Code  int    `json:"code" example:"422"`
	Error string `json:"error" example:"unprocessable entity"`
}

type InternalErrorResponse struct {
	Code  int    `json:"code" example:"500"`
	Error string `json:"error" example:"internal server error"`
}
