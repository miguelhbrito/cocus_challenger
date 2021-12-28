package login

import (
	"errors"
	"testing"

	"github.com/cocus_challenger_refact/business/auth"
	"github.com/cocus_challenger_refact/business/data/login"
)

func TestCore_CreateUser(t *testing.T) {
	type fields struct {
		db   login.LoginInt
		auth auth.Auth
	}
	type args struct {
		l login.Login
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Success",
			fields: fields{
				db: login.LoginIntCustomMock{
					SaveMock: func(l login.Login) error {
						return nil
					},
				},
				auth: auth.AuthCustomMock{
					GenerateHashPasswordMock: func(password string) (string, error) {
						return "hashedPassword", nil
					},
				},
			},
			args: args{
				l: login.Login{
					Username: "any_username",
					Password: "password",
				},
			},
			wantErr: false,
		},
		{
			name: "Error on generate hashedpassword",
			fields: fields{
				db: login.LoginIntCustomMock{
					SaveMock: func(l login.Login) error {
						return nil
					},
				},
				auth: auth.AuthCustomMock{
					GenerateHashPasswordMock: func(password string) (string, error) {
						return "", login.ErrPasswordHash
					},
				},
			},
			wantErr: true,
		},
		{
			name: "Error on save new user",
			fields: fields{
				db: login.LoginIntCustomMock{
					SaveMock: func(l login.Login) error {
						return errors.New("some error")
					},
				},
				auth: auth.AuthCustomMock{
					GenerateHashPasswordMock: func(password string) (string, error) {
						return "hashedPassword", nil
					},
				},
			},
			args: args{
				l: login.Login{
					Username: "any_username",
					Password: "password",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewCore(tt.fields.db, tt.fields.auth)
			if err := c.CreateUser(tt.args.l); (err != nil) != tt.wantErr {
				t.Errorf("Core.CreateUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCore_Login(t *testing.T) {
	type fields struct {
		db   login.LoginInt
		auth auth.Auth
	}
	type args struct {
		l login.Login
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Success",
			fields: fields{
				db: login.LoginIntCustomMock{
					LoginMock: func(l login.Login) (login.Login, error) {
						return login.Login{
							Username: "any_username",
							Password: "hashedPawssord",
						}, nil
					},
				},
				auth: auth.AuthCustomMock{
					CheckPasswordHashMock: func(password, hash string) bool {
						return true
					},
				},
			},
			args: args{
				l: login.Login{
					Username: "any_username",
					Password: "password",
				},
			},
			wantErr: false,
		},
		{
			name: "Error on get credentials from db",
			fields: fields{
				db: login.LoginIntCustomMock{
					LoginMock: func(l login.Login) (login.Login, error) {
						return login.Login{}, errors.New("some error")
					},
				},
				auth: auth.AuthCustomMock{
					CheckPasswordHashMock: func(password, hash string) bool {
						return true
					},
				},
			},
			args: args{
				l: login.Login{
					Username: "any_username",
					Password: "password",
				},
			},
			wantErr: true,
		},
		{
			name: "Error to check hashedPassword",
			fields: fields{
				db: login.LoginIntCustomMock{
					LoginMock: func(l login.Login) (login.Login, error) {
						return login.Login{
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
			},
			args: args{
				l: login.Login{
					Username: "any_username",
					Password: "password",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Core{
				db:   tt.fields.db,
				auth: tt.fields.auth,
			}
			_, err := c.Login(tt.args.l)
			if (err != nil) != tt.wantErr {
				t.Errorf("Core.Login() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
