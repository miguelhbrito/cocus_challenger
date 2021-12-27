package triangle

import (
	"errors"
	"reflect"
	"testing"

	"github.com/cocus_challenger/pkg/api/entity"
	"github.com/cocus_challenger/pkg/storage"
)

func Test_manager_Create(t *testing.T) {
	type args struct {
		t entity.Triangle
	}
	tests := []struct {
		name            string
		m               manager
		storageTriangle storage.Triangle
		args            args
		want            entity.Triangle
		wantErr         bool
		err             error
	}{
		{
			name: "Success",
			storageTriangle: storage.TriangleCustomMock{
				SaveMock: func(t entity.Triangle) error {
					return nil
				},
			},
			args: args{
				entity.Triangle{
					Id:    "1",
					Side1: 10,
					Side2: 10,
					Side3: 10,
				},
			},
			want: entity.Triangle{
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
			storageTriangle: storage.TriangleCustomMock{
				SaveMock: func(t entity.Triangle) error {
					return errors.New("some error")
				},
			},
			args: args{
				entity.Triangle{
					Id:    "1",
					Side1: 10,
					Side2: 10,
					Side3: 10,
					Type:  "equilateral",
				},
			},
			wantErr: true,
		},

		{
			name: "Not a valid triangle",
			storageTriangle: storage.TriangleCustomMock{
				SaveMock: func(t entity.Triangle) error {
					return nil
				},
			},
			args: args{
				entity.Triangle{
					Id:    "1",
					Side1: 5,
					Side2: 3,
					Side3: 8,
				},
			},
			wantErr: true,
			err:     errNotATriangle,
		},
		{
			name: "Isosceles triangle",
			storageTriangle: storage.TriangleCustomMock{
				SaveMock: func(t entity.Triangle) error {
					return nil
				},
			},
			args: args{
				entity.Triangle{
					Id:    "1",
					Side1: 10,
					Side2: 10,
					Side3: 8,
				},
			},
			want: entity.Triangle{
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
			storageTriangle: storage.TriangleCustomMock{
				SaveMock: func(t entity.Triangle) error {
					return nil
				},
			},
			args: args{
				entity.Triangle{
					Id:    "1",
					Side1: 10,
					Side2: 10,
					Side3: 10,
				},
			},
			want: entity.Triangle{
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
			storageTriangle: storage.TriangleCustomMock{
				SaveMock: func(t entity.Triangle) error {
					return nil
				},
			},
			args: args{
				entity.Triangle{
					Id:    "1",
					Side1: 10,
					Side2: 7,
					Side3: 5,
				},
			},
			want: entity.Triangle{
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
			m := NewManager(tt.storageTriangle)
			got, err := m.Create(tt.args.t)
			if (err != nil) != tt.wantErr {
				t.Errorf("manager.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.err != nil {
				if err != tt.err {
					t.Errorf("manager.Create() error = %v, wantErr %v", err, tt.err)
					return
				}
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("manager.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_manager_List(t *testing.T) {
	tests := []struct {
		name            string
		storageTriangle storage.Triangle
		m               manager
		want            entity.Triangles
		wantErr         bool
	}{
		{
			name: "Success",
			storageTriangle: storage.TriangleCustomMock{
				ListMock: func() ([]entity.Triangle, error) {
					return []entity.Triangle{{Id: "1"}}, nil
				},
			},
			want:    []entity.Triangle{{Id: "1"}},
			wantErr: false,
		},
		{
			name: "Error on list all triangles from db",
			storageTriangle: storage.TriangleCustomMock{
				ListMock: func() ([]entity.Triangle, error) {
					return []entity.Triangle{{}}, errors.New("some error")
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := NewManager(tt.storageTriangle)
			got, err := m.List()
			if (err != nil) != tt.wantErr {
				t.Errorf("manager.List() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("manager.List() = %v, want %v", got, tt.want)
			}
		})
	}
}
