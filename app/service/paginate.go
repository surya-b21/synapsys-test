package service

import "github.com/morkid/paginate"

// PG is global pagination variable
var PG *paginate.Pagination

// InitPG to init pagination
func InitPG() {
	pg := paginate.New()

	PG = pg
}
