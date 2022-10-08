package data

import "FileUpload.Kevin.net/internal/validator"

type aboutMe struct {
	Name         string `json:"name"`
	Email        string `json:"email"`
	Occupation   string `json:"occupation"`
	Organization string `json:"organization"`
	Education    string `json:"education"`
	Address      string `json:"address"`
}

func ValidateEntires(v *validator.Validator, bio *aboutMe) {
	v.Check(bio.Name != "", "name", "must be provided")
	v.Check(len(bio.Name) <= 200, "name", "must not be more than 200 bytes long")

	v.Check(bio.Email != "", "Email", "must be provided")
	v.Check(len(bio.Email) <= 200, "Email", "must not be more than 200 bytes long")

	v.Check(bio.Organization != "", "Organization", "must be provided")
	v.Check(len(bio.Organization) <= 200, "Organization", "must not be more than 200 bytes long")

	v.Check(bio.Occupation != "", "Occupation", "must be provided")
	v.Check(len(bio.Occupation) <= 200, "Occupation", "must not be more than 200 bytes long")

	v.Check(bio.Education != "", "Education", "must be provided")
	v.Check(len(bio.Education) <= 200, "Education", "must not be more than 200 bytes long")

	v.Check(bio.Address != "", "Address", "must be provided")
	v.Check(len(bio.Address) <= 200, "Address", "must not be more than 200 bytes long")

}
