func (db *appdbimpl) CheckTripID(tripID int) (bool, error) {
	var ret int
	err := db.c.Get(&ret, "SELECT COUNT(*) FROM trips WHERE `id` = ?", tripID)
	if err != nil {
		return false, fmt.Errorf("select1 statement: %w", err)
	} else if ret > 0 {
		return true, nil
	}
	return false, nil
}