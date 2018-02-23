// consignment-servce/datastore.go

package main

import (
	"gopkg.in/mgo.v2"
)

// CreateSession creates the main session to our mongodb instance
func CreateSession(host string) (*mgo.Session, error) {
	session, err := mgo.Dial(host)
	if err != nil {
		return nil, err
	}

	session.SetMode(mgo.Monotonic, true)

	return session, nil
}

// That's it, pretty straight forward. It takes a host string as an argument
// returns a session to our datastore and of course a potential error
// so that we can handle that on start-up.
