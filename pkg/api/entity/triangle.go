package entity

import "github.com/cocus_challenger/pkg/api/response"

type Triangle struct {
	Id    string `json:"id"`
	Side1 int    `json:"side1"`
	Side2 int    `json:"side2"`
	Side3 int    `json:"side3"`
	Type  string `json:"type"`
}

type Triangles []Triangle

func (t Triangle) Response() response.Triangle {
	return response.Triangle{
		Id:    t.Id,
		Side1: t.Side1,
		Side2: t.Side2,
		Side3: t.Side3,
		Type:  t.Type,
	}
}

func (t Triangles) Response() []response.Triangle {
	resp := make([]response.Triangle, 0)
	for i := range t {
		resp = append(resp, t[i].Response())
	}
	return resp
}
