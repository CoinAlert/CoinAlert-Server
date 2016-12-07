package main

// Save a device to the database.
func SaveDevice(d Device) error {
	sel := Device{Id: d.Id}
	_, err := db.Upsert(sel, d)
	return err
}
