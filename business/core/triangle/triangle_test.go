package triangle

import (
	"errors"
	"reflect"
	"testing"

	"github.com/cocus_challenger_refact/business/data/triangle"
)

func TestCore_Create(t *testing.T) {
	type fields struct {
		db triangle.TriangleInt
	}
	type args struct {
		t triangle.Triangle
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    triangle.Triangle
		wantErr bool
		err     error
	}{
		{
			name: "Success",
			fields: fields{
				db: triangle.TriangleIntCustomMock{
					SaveMock: func(t triangle.Triangle) error {
						return nil
					},
				},
			},
			args: args{
				triangle.Triangle{
					Id:    "1",
					Side1: 10,
					Side2: 10,
					Side3: 10,
				},
			},
			want: triangle.Triangle{
				Id:    "1",
				Side1: 10,
				Side2: 10,
				Side3: 10,
				Type:  "equilateral",
			},
			wantErr: false,
		},
		{
			name: "Error on save triangle",
			fields: fields{
				db: triangle.TriangleIntCustomMock{
					SaveMock: func(t triangle.Triangle) error {
						return errors.New("some error")
					},
				},
			},
			args: args{
				triangle.Triangle{
					Id:    "1",
					Side1: 10,
					Side2: 10,
					Side3: 10,
				},
			},
			wantErr: true,
		},
		{
			name: "Not a valid triangle",
			fields: fields{
				db: triangle.TriangleIntCustomMock{
					SaveMock: func(t triangle.Triangle) error {
						return nil
					},
				},
			},
			args: args{
				triangle.Triangle{
					Id:    "1",
					Side1: 5,
					Side2: 3,
					Side3: 8,
				},
			},
			wantErr: true,
			err:     triangle.ErrNotATriangle,
		},
		{
			name: "Isosceles triangle",
			fields: fields{
				db: triangle.TriangleIntCustomMock{
					SaveMock: func(t triangle.Triangle) error {
						return nil
					},
				},
			},
			args: args{
				triangle.Triangle{
					Id:    "1",
					Side1: 10,
					Side2: 10,
					Side3: 8,
				},
			},
			want: triangle.Triangle{
				Id:    "1",
				Side1: 10,
				Side2: 10,
				Side3: 8,
				Type:  "isosceles",
			},
			wantErr: false,
		},
		{
			name: "Equilateral triangle",
			fields: fields{
				db: triangle.TriangleIntCustomMock{
					SaveMock: func(t triangle.Triangle) error {
						return nil
					},
				},
			},
			args: args{
				triangle.Triangle{
					Id:    "1",
					Side1: 10,
					Side2: 10,
					Side3: 10,
				},
			},
			want: triangle.Triangle{
				Id:    "1",
				Side1: 10,
				Side2: 10,
				Side3: 10,
				Type:  "equilateral",
			},
			wantErr: false,
		},
		{
			name: "Scalene triangle",
			fields: fields{
				db: triangle.TriangleIntCustomMock{
					SaveMock: func(t triangle.Triangle) error {
						return nil
					},
				},
			},
			args: args{
				triangle.Triangle{
					Id:    "1",
					Side1: 10,
					Side2: 7,
					Side3: 5,
				},
			},
			want: triangle.Triangle{
				Id:    "1",
				Side1: 10,
				Side2: 7,
				Side3: 5,
				Type:  "scalene",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Core{
				db: tt.fields.db,
			}
			got, err := c.Create(tt.args.t)
			if (err != nil) != tt.wantErr {
				t.Errorf("Core.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.err != nil {
				if err != tt.err {
					t.Errorf("Core.Create() error = %v, wantErr %v", err, tt.err)
					return
				}
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Core.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCore_List(t *testing.T) {
	type fields struct {
		db triangle.TriangleInt
	}
	tests := []struct {
		name    string
		fields  fields
		want    triangle.Triangles
		wantErr bool
	}{
		{
			name: "Success",
			fields: fields{
				db: triangle.TriangleIntCustomMock{
					ListMock: func() (triangle.Triangles, error) {
						return []triangle.Triangle{{Id: "1"}}, nil
					},
				},
			},
			want:    triangle.Triangles{{Id: "1"}},
			wantErr: false,
		},
		{
			name: "Error on list all triangles from db",
			fields: fields{
				db: triangle.TriangleIntCustomMock{
					ListMock: func() (triangle.Triangles, error) {
						return []triangle.Triangle{{}}, errors.New("some error")
					},
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Core{
				db: tt.fields.db,
			}
			got, err := c.List()
			if (err != nil) != tt.wantErr {
				t.Errorf("Core.List() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Core.List() = %v, want %v", got, tt.want)
			}
		})
	}
}
