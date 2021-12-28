package triangle

import (
	dbconnect "github.com/cocus_challenger_refact/platform/db_connect"
	"github.com/rs/zerolog/log"
)

type TriangleInt interface {
	Save(t Triangle) error
	List() (Triangles, error)
}

type TrianglePostgres struct{}

func (tp TrianglePostgres) Save(t Triangle) error {

	db := dbconnect.InitDB()
	defer db.Close()

	tr := Triangle{
		Id:    t.Id,
		Side1: t.Side1,
		Side2: t.Side2,
		Side3: t.Side3,
		Type:  t.Type,
	}

	sqlStatement := `INSERT INTO triangle VALUES ($1, $2, $3, $4, $5)`
	_, err := db.Exec(sqlStatement, tr.Id, tr.Side1, tr.Side2, tr.Side3, tr.Type)
	if err != nil {
		log.Error().Err(err).Msgf("Error to insert an new triangle into db")
		return err
	}

	return nil
}

func (tp TrianglePostgres) List() (Triangles, error) {
	db := dbconnect.InitDB()
	defer db.Close()

	var ts []Triangle
	sqlStatement := `SELECT id, side1, side2, side3, type FROM triangle`
	rows, err := db.Query(sqlStatement)
	if err != nil {
		log.Error().Err(err).Msg("Error to get all triangles from db")
		return nil, err
	}

	for rows.Next() {
		var t Triangle
		err := rows.Scan(&t.Id, &t.Side1, &t.Side2, &t.Side3, &t.Type)
		if err != nil {
			log.Error().Err(err).Msg("Error to extract result from row")
		}
		ts = append(ts, t)
	}

	return ts, nil
}
