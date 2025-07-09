package main

import (
	"log"
	"nil_dereference_analysis/common"
)

func process(user *common.User) {
	if user == nil {
		log.Println("user is nil")
	}
	copiedUser := user.Copy(nil)
	handleUser(copiedUser)
}

func handleUser(user *common.User) {
	// some logic with user
}
