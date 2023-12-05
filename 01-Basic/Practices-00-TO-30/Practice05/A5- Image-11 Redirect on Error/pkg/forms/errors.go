package forms

type errors map[string][]string

func (e errors) AddError(tagID, message string) {
	e[tagID] = append(e[tagID], message)
}

func (e errors) GetError(tagID string) string {
	es := e[tagID]
	if len(es) == 0 {
		return ""
	} else {
		return es[0]
	}
}




