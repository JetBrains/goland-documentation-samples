package common

import "log"

func logUserEvent(msg string, user *User) {
	log.Printf("user event: %s: name=%s, age=%d", msg, user.Name, user.Age)
}
