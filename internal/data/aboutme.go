package data

import "FileUpload.Kevin.net/internal/validator"

type About struct {
	Name         string   `json:"name"`
	Email        string   `json:"email"`
	Occupation   string   `json:"occupation"`
	Organization string   `json:"organization"`
	Education	 string   `json:"education"`
	Address		 string	  `json:"address"`
}

err := app.readJSON(w, r, &input)
if err != nil {
	app.badrequestResponse(w, r, err)
	return
}

v := validator.New()
	//use the check method to execute our validation checks
	v.Check(input.Name != "", "name", "must be provided")
	v.Check(len(input.Name) <= 200, "name", "must not be more than 200 bytes long")

	v.Check(input.Email != "", "Email", "must be provided")
	v.Check(len(input.Email) <= 200, "Email", "must not be more than 200 bytes long")

	v.Check(input.Organization != "", "Organization", "must be provided")
	v.Check(len(input.Organization) <= 200, "Organization", "must not be more than 200 bytes long")
	
	v.Check(input.Occupation != "", "Occupation", "must be provided")
	v.Check(len(input.Occupation) <= 200, "Occupation", "must not be more than 200 bytes long")
	
	v.Check(input.Education != "", "Education", "must be provided")
	v.Check(len(input.Education) <= 200, "Education", "must not be more than 200 bytes long")
	

}
	fmt.Fprintf(w, "%+v\n", input)

func (app *application) AboutMe(w http.ResponseWriter, r *http.Request){
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	About := data.About{
		Name:			"Kevin Godoy",
		Email: 			"2018117874@ub.edu.bz",
		Occupation: 	"Student",
		Organization:   "University of Belize",
		Education: 		"Associates in Information Technology"
		Address:		"Belmopan City, Belize"
	}
}
