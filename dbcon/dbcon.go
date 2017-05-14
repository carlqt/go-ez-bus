package dbcon

import (
	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
)

var (
	//DBcon is the connection handler
	DBcon  *sqlx.DB
	SDBcon *squirrel.StatementBuilderType
)
