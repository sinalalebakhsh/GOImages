package helpers

import "log"

func ErrorCheck(err error) {
	if err != nil {
		log.Fatal(err)
	}
}