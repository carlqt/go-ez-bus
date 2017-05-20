package env

import (
	"net/http"

	"github.com/Masterminds/squirrel"
	"github.com/carlqt/ez-bus/config"
	"github.com/jmoiron/sqlx"
)

var (
	//DBX is the connection handler for sqlx
	DBX        *sqlx.DB
	SDBcon     *squirrel.StatementBuilderType
	Conf       *config.Config
	HttpClient *http.Client
)
