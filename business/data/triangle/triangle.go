package triangle

import (
	"database/sql"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

type TriangleInt interface {
	Save(t Triangle) (Triangle, error)
	List() (Triangles, error)
}

type TrianglePostgres struct {
	Db *sql.DB
}

func (tp TrianglePostgres) Save(t Triangle) (Triangle, error) {

	tr := Triangle{
		Id:    uuid.New().String(),
		Side1: t.Side1,
		Side2: t.Side2,
		Side3: t.Side3,
		Type:  t.Type,
	}

	sqlStatement := `INSERT INTO triangle VALUES ($1, $2, $3, $4, $5)`
	_, err := tp.Db.Exec(sqlStatement, tr.Id, tr.Side1, tr.Side2, tr.Side3, tr.Type)
	if err != nil {
		log.Error().Err(err).Msgf("Error to insert an new triangle into db")
		return Triangle{}, err
	}

	return tr, nil
}

func (tp TrianglePostgres) List() (Triangles, error) {

	var ts []Triangle
	sqlStatement := `SELECT id, side1, side2, side3, type FROM triangle`
	rows, err := tp.Db.Query(sqlStatement)
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
