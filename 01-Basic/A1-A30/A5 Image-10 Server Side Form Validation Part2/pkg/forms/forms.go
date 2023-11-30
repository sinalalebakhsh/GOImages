package forms

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"github.com/asaskevich/govalidator"
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

func (f *Form) HasRequired(tagIDs ...string) {
	for _, tagID := range tagIDs {
		value := f.Get(tagID)
		if strings.TrimSpace(value) == "" {
			f.Errors.AddError(tagID, "This field cant's be blank!")
		}
	}
}

func (f *Form) HasValue(tagID string, request *http.Request) bool {
	x := request.Form.Get(tagID)
	return x != ""
	// if x == "" {
	// 	f.Errors.AddError(tagID, "Field Empty!")
	// 	return false
	// }
}

func (f *Form) MinLength(tagID string, length int, request *http.Request) bool {
	x := request.Form.Get(tagID)
	if len(x) < length {
		f.Errors.AddError(tagID, fmt.Sprintf("This field must be %d characters long or more", length))
		return false
	}
	return true
}

func (f *Form) IsEmail(tagId string) {
	if !govalidator.IsEmail(f.Get(tagId)) {
		f.Errors.AddError(tagId, "Invalid Email")
	}
}

func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}





































