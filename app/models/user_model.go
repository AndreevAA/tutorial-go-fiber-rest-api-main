package models

import (
	_ "database/sql/driver"
	_ "encoding/json"
	_ "errors"
	"github.com/google/uuid"
	_ "time"
)

// User struct to describe user object.
type User struct {
	Id       uuid.UUID `db:"id" json:"id" validate:"required,uuid"`
	Login    string    `db:"login" json:"login" validate:"required,lte=255"`
	Password string    `db:"password" json:"password" validate:"required,lte=255"`
	AcsLevel int       `db:"lvl_access" json:"lvl_access" validate:"required,len=1"`
}
