package forms

import (
	"net/http"
	"net/url"
)

type Form struct {
	url.Values
	Errors errors
}

func New(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

func (f *Form) HasValue(tagID string, request *http.Request) bool {
	x := request.Form.Get(tagID)
	if x == "" {
		f.Errors.AddError(tagID, "Field Empty!")
		return false
	}
	return true
}

func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}





































