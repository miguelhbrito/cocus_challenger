package triangle

type TriangleIntCustomMock struct {
	SaveMock func(t Triangle) error
	ListMock func() (Triangles, error)
}

func (tm TriangleIntCustomMock) Save(t Triangle) error {
	return tm.SaveMock(t)
}

func (tm TriangleIntCustomMock) List() (Triangles, error) {
	return tm.ListMock()
}
