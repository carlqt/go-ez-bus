package dbcon

import (
	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
)

var (
	//DBX is the connection handler for sqlx
	DBX    *sqlx.DB
	SDBcon *squirrel.StatementBuilderType
)
