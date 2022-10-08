package main

import (
	"fmt"
	"net/http"

	"FileUpload.Kevin.net/internal/data"
	"FileUpload.Kevin.net/internal/validator"
)

func (app *application) AboutMe(w http.ResponseWriter, r *http.Request) {
	//our target decode destination
	var input struct {
		Name         string `json:"name"`
		Email        string `json:"email"`
		Occupation   string `json:"occupation"`
		Organization string `json:"organization"`
		Education    string `json:"education"`
		Address      string `json:"address"`
	}
	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}
	//copying the values
	jsondata := &data.About{
		Name:         input.Name,
		Email:        input.Email,
		Occupation:   input.Occupation,
		Organization: input.Organization,
		Education:    input.Education,
		Address:      input.Address,
	}

	//initialize a new validator instance
	v := validator.New()

	//check the map to determine if there were any validation errors
	if data.EntriesValidation(v, jsondata); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}
	//Display the request
	fmt.Fprintf(w, "%+v\n", input)
}
