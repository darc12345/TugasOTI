package controller

import (
	"database/sql"
	"main/service"
)

type ControllerDB struct {
	serviceDB service.ServiceDB
}

func NewControllerDB(db *sql.DB) ControllerDB {
	return ControllerDB{serviceDB: service.NewServiceDB(db)}
}
