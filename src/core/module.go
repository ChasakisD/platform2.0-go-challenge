package core

import (
	"gwi/assignment/core/data/database"
	"gwi/assignment/core/data/http"
)

func Initialize(connectionStr string, jwtSecret string) {
	http.InitializeJwt(jwtSecret)
	database.Initialize(connectionStr)
}
