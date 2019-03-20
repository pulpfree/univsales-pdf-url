package validate

import (
	"errors"
	"time"

	"github.com/pulpfree/gdps-fs-dwnld/model"
)

// Time form constant
const (
	timeShortForm  = "20060102"
	timeRecordForm = "2006-01-02"
)

// Date function
func Date(dateInput string) (time.Time, error) {

	date, err := time.Parse(timeRecordForm, dateInput)
	if err != nil {
		return date, err
	}

	// Ensure date is not future dated
	today := time.Now()
	if today.Year() < date.Year() || today.YearDay() < date.YearDay() {
		return date, errors.New("Invalid date. Date cannot be future, must be less than current date")
	}

	return date, err
}

// RequestInput function
func RequestInput(r *model.RequestInput) (res *model.Request, err error) {

	if r == nil {
		err = errors.New("Missing RequestInput values")
		return nil, err
	}

	res = new(model.Request)

	res.Date, err = Date(r.Date)
	if err != nil {
		return res, err
	}

	// Little to validate with stationID
	res.StationID = r.StationID

	return res, nil
}
