package dbcon

import (
	"database/sql"

	"github.com/Masterminds/squirrel"
)

var (
	//DBcon is the connection handler
	DBcon  *sql.DB
	SDBcon *squirrel.StatementBuilderType
)
