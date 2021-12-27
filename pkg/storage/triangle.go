package storage

import (
	dbconnect "github.com/cocus_challenger/db_connect"
	"github.com/cocus_challenger/pkg/api/entity"
	"github.com/rs/zerolog/log"
)

type TrianglePostgres struct{}

func NewTrianglePostgres() Triangle {
	return TrianglePostgres{}
}

func (tp TrianglePostgres) Save(t entity.Triangle) error {
	db := dbconnect.InitDB()
	defer db.Close()

	sqlStatement := `INSERT INTO triangle VALUES ($1, $2, $3, $4, $5)`

	_, err := db.Exec(sqlStatement, t.Id, t.Side1, t.Side2, t.Side3, t.Type)
	if err != nil {
		log.Error().Err(err).Msgf("Error to insert an new triangle into db")
		return err
	}

	return nil
}

func (tp TrianglePostgres) List() ([]entity.Triangle, error) {
	db := dbconnect.InitDB()
	defer db.Close()

	var ts []entity.Triangle

	sqlStatement := `SELECT id, side1, side2, side3, type FROM triangle`

	rows, err := db.Query(sqlStatement)
	if err != nil {
		log.Error().Err(err).Msg("Error to get all triangles from db")
		return nil, err
	}
	for rows.Next() {
		var t entity.Triangle
		err := rows.Scan(&t.Id, &t.Side1, &t.Side2, &t.Side3, &t.Type)
		if err != nil {
			log.Error().Err(err).Msg("Error to extract result from row")
		}
		ts = append(ts, t)
	}

	return ts, nil
}
