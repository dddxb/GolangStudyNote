package lib

import (
	"database/sql"
	"go.uber.org/zap"
)

var Db *sql.DB
var Logger *zap.Logger
var Sugar *zap.SugaredLogger