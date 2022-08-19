package main

import (
	"database/sql"
	"net/http"
)

func tryLogin(user userData, db *sql.DB, c echo.Context) error {
	var response Response
	token := checkUser(user, db)
	if token == "-1" {
		response.Message = "failed to login! \n"
		response.Data = "there are several possible issues: database issues, or the data is just wrong"
		response.httpstatus = http.StatusOK
		response.Mail = user.Email
	} else {
		response.Data = token
		response.Mail = user.Email
		response.Message = " successfully loggeed in!"
		response.httpstatus = http.StatusOK
	}
	return c.JSON(response.httpstatus, response)
}
