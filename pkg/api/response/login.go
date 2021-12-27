package response

type Token struct {
	Token   string `json:"token"`
	ExpTime int64  `json:"expTime"`
}
