package main

import (
	"net/http"

	"FileUpload.Kevin.net/internal/data"
)

func (app *application) AboutMe(w http.ResponseWriter, r *http.Request) {

	data := data.About{
		Name:         "Kevin Godoy",
		Email:        "2018117874@ub.edu.bz",
		Occupation:   "Student",
		Organization: "University of Belize",
		Education:    "Associates in Information Technology",
		Address:      "Belmopan City, Belize",
	}

	err := app.writeJSON(w, http.StatusOK, envelope{"data": data}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}

}
