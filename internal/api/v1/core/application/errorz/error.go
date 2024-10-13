package errorz

type Error_ struct {
	StatusCode uint   `json:"status_code"`
	Message    string `json:"message"`
}
