package login

import (
	"errors"
	"testing"

	"github.com/cocus_challenger/pkg/api/entity"
	"github.com/cocus_challenger/pkg/auth"
	"github.com/cocus_challenger/pkg/storage"
)

func Test_manager_CreateUser(t *testing.T) {
	type args struct {
		l entity.LoginEntity
	}
	tests := []struct {
		name         string
		loginStorage storage.Login
		auth         auth.Auth
		args         args
		wantErr      bool
	}{
		{
			name: "Success",
			loginStorage: storage.LoginCustomMock{
				SaveMock: func(l entity.LoginEntity) error {
					return nil
				},
			},
			auth: auth.AuthCustomMock{
				GenerateHashPasswordMock: func(password string) (string, error) {
					return "hashedPassword", nil
				},
			},
			args: args{
				l: entity.LoginEntity{
					Username: "any_username",
					Password: "password",
				},
			},
			wantErr: false,
		},
		{
			name: "Error on generate hashedpassword",
			loginStorage: storage.LoginCustomMock{
				SaveMock: func(l entity.LoginEntity) error {
					return nil
				},
			},
			auth: auth.AuthCustomMock{
				GenerateHashPasswordMock: func(password string) (string, error) {
					return "", errPasswordHash
				},
			},
			wantErr: true,
		},
		{
			name: "Error on save new user",
			loginStorage: storage.LoginCustomMock{
				SaveMock: func(l entity.LoginEntity) error {
					return errors.New("some error")
				},
			},
			auth: auth.AuthCustomMock{
				GenerateHashPasswordMock: func(password string) (string, error) {
					return "hashedPassword", nil
				},
			},
			args: args{
				l: entity.LoginEntity{
					Username: "any_username",
					Password: "password",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := NewManager(tt.loginStorage, tt.auth)
			if err := m.CreateUser(tt.args.l); (err != nil) != tt.wantErr {
				t.Errorf("manager.CreateUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_manager_Login(t *testing.T) {
	type args struct {
		l entity.LoginEntity
	}
	tests := []struct {
		name         string
		loginStorage storage.Login
		args         args
		auth         auth.Auth
		wantErr      bool
	}{
		{
			name: "Success",
			loginStorage: storage.LoginCustomMock{
				LoginMock: func(l entity.LoginEntity) (entity.LoginEntity, error) {
					return entity.LoginEntity{
						Username: "any_username",
						Password: "hashedPawssord",
					}, nil
				},
			},
			args: args{
				l: entity.LoginEntity{
					Username: "any_username",
					Password: "password",
				},
			},
			auth: auth.AuthCustomMock{
				CheckPasswordHashMock: func(password, hash string) bool {
					return true
				},
			},
			wantErr: false,
		},
		{
			name: "Error to check hashedPassword",
			loginStorage: storage.LoginCustomMock{
				LoginMock: func(l entity.LoginEntity) (entity.LoginEntity, error) {
					return entity.LoginEntity{
						Username: "any_username",
						Password: "hashedPawssord",
					}, nil
				},
			},
			auth: auth.AuthCustomMock{
				CheckPasswordHashMock: func(password, hash string) bool {
					return false
				},
			},
			args: args{
				l: entity.LoginEntity{
					Username: "any_username",
					Password: "password",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := NewManager(tt.loginStorage, tt.auth)
			_, err := m.Login(tt.args.l)
			if (err != nil) != tt.wantErr {
				t.Errorf("manager.Login() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
