package response

type Triangle struct {
	Id    string `json:"id"`
	Side1 int    `json:"side1"`
	Side2 int    `json:"side2"`
	Side3 int    `json:"side3"`
	Type  string `json:"type"`
}
