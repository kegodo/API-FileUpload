// Filename cmd/api/handlers.go
package main

import (
	"net/http"
)

func (app *application) RandomStringHandler(w http.ResponseWriter, r *http.Request) {
	userInput, err := app.readIDParam(r)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	if userInput > 9999 {
		app.methodNotAllowedResponse(w, r, err)
		return
	}

	randomString := app.generateRandomString(int(userInput))

	err = app.writeJSON(w, http.StatusOK, envelope{"Random String": randomString}, nil)

	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
