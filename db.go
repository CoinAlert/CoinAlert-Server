package main

import (
	"gopkg.in/mgo.v2"
)

// Save a device to the database.
func SaveDevice(s *mgo.Session, d Device) error {
	session := s.Copy()
	defer session.Close()
	c := session.DB(database).C(devices_collection)

	sel := Device{Id: d.Id}

	_, err := c.Upsert(sel, d)
	return err
}
