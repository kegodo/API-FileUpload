package main

import (
	"crypto/rand"
	"net/http"

	"FileUpload.Kevin.net/internal/data"
)

const randomStringSource = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789+_#$-!~"

func generateRandomString(length int64) string {
	s := make([]rune, length)
	r := []rune(randomStringSource)

	for i := range s {
		p, _ := rand.Prime(rand.Reader, len(r))
		x := p.Uint64()
		y := uint64(len(r))
		s[i] = r[x%y]
	}

	return string(s)

}

func (app *application) RandomStringHandler(w http.ResponseWriter, r *http.Request) {

	id, err := app.readIDParam(r)

	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	randString := generateRandomString(id)

	data := data.RandString{
		Data: randString,
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"data": data}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}

}
