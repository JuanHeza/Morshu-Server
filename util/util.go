package util

import "log"

func IsOk(err error) bool {
	if err != nil {
		log.Println("There Was an Error", err)
		return false
	}
	return true
}
