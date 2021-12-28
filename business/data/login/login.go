package login

import (
	"database/sql"

	dbconnect "github.com/cocus_challenger_refact/platform/db_connect"
	"github.com/rs/zerolog/log"
)

type LoginInt interface {
	Save(l Login) error
	Login(l Login) (Login, error)
}

type LoginPostgres struct{}

func (lp LoginPostgres) Save(le Login) error {
	db := dbconnect.InitDB()
	defer db.Close()

	l := Login{
		Username: le.Username,
		Password: le.Password,
	}

	sqlStatement := `INSERT INTO login VALUES ($1, $2)`
	_, err := db.Exec(sqlStatement, l.Username, l.Password)
	if err != nil {
		log.Error().Err(err).Msgf("Error to insert a new user into db")
		return err
	}

	return nil
}

func (lp LoginPostgres) Login(le Login) (Login, error) {
	db := dbconnect.InitDB()
	defer db.Close()

	l := Login{
		Username: le.Username,
		Password: le.Password,
	}

	var lr Login
	sqlStatement := `SELECT username, password FROM login WHERE username = $1`
	result := db.QueryRow(sqlStatement, l.Username)
	err := result.Scan(&lr.Username, &lr.Password)

	if err != nil {
		if err == sql.ErrNoRows {
			log.Error().Err(err).Msgf("Not found user with username %s", l.Username)
			return Login{}, err
		}
		log.Error().Err(err).Msgf("Error to get credentials from db, with id %s", l.Username)
		return Login{}, err
	}

	return lr, nil
}
