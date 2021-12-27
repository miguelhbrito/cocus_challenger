package storage

import (
	"database/sql"

	dbconnect "github.com/cocus_challenger/db_connect"
	"github.com/cocus_challenger/pkg/api/entity"
	"github.com/rs/zerolog/log"
)

type LoginPostgres struct{}

func NewLoginPostgres() Login {
	return LoginPostgres{}
}

func (lp LoginPostgres) Save(t entity.LoginEntity) error {
	db := dbconnect.InitDB()
	defer db.Close()
	sqlStatement := `INSERT INTO login VALUES ($1, $2)`
	_, err := db.Exec(sqlStatement, t.Username, t.Password)
	if err != nil {
		log.Error().Err(err).Msgf("Error to insert a new user into db")
		return err
	}
	return nil
}

func (lp LoginPostgres) Login(l entity.LoginEntity) (entity.LoginEntity, error) {
	db := dbconnect.InitDB()
	defer db.Close()
	var lr entity.LoginEntity
	sqlStatement := `SELECT username, password FROM login WHERE username = $1`
	result := db.QueryRow(sqlStatement, l.Username)
	err := result.Scan(&lr.Username, &lr.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Error().Err(err).Msgf("Not found user with username %s", l.Username)
			return entity.LoginEntity{}, err
		}
		log.Error().Err(err).Msgf("Error to get credentials from db, with id %s", l.Username)
		return entity.LoginEntity{}, err
	}
	return lr, nil
}
