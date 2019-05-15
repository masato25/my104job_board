package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
	"github.com/gofrs/uuid"
)

type Job struct {
	ID          uuid.UUID `json:"id" db:"id"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
	CompanyName string    `json:"company_name" db:"company_name"`
	JobName     string    `json:"job_name" db:"job_name"`
	Area        string    `json:"area" db:"area"`
	Welfare     string    `json:"welfare" db:"welfare"`
	J           string    `json:"j" db:"j"`
	C           string    `json:"c" db:"c"`
	JobCat      string    `json:"job_cat" db:"job_cat"`
	SalLow      int       `json:"sal_low" db:"sal_low"`
	SalHigh     int       `json:"sal_high" db:"sal_high"`
	DESCRIPTION string    `json:"description" db:"description"`
	OtherDes    string    `json:"other_des" db:"other_des"`
	Profile     string    `json:"profile" db:"profile"`
	Link        string    `json:"link" db:"link"`
	Manager     bool      `json:"manager" db:"manager"`
	NeedOnBt    bool      `json:"need_on_bt" db:"need_on_bt"`
}

// String is not required by pop and may be deleted
func (j Job) String() string {
	jj, _ := json.Marshal(j)
	return string(jj)
}

// Jobs is not required by pop and may be deleted
type Jobs []Job

// String is not required by pop and may be deleted
func (j Jobs) String() string {
	jj, _ := json.Marshal(j)
	return string(jj)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (j *Job) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Field: j.CompanyName, Name: "CompanyName"},
		&validators.StringIsPresent{Field: j.JobName, Name: "JobName"},
		&validators.StringIsPresent{Field: j.Area, Name: "Area"},
		&validators.StringIsPresent{Field: j.Welfare, Name: "Welfare"},
		&validators.StringIsPresent{Field: j.J, Name: "J"},
		&validators.StringIsPresent{Field: j.C, Name: "C"},
		&validators.StringIsPresent{Field: j.JobCat, Name: "JobCat"},
		&validators.IntIsPresent{Field: j.SalLow, Name: "SalLow"},
		&validators.IntIsPresent{Field: j.SalHigh, Name: "SalHigh"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (j *Job) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (j *Job) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
