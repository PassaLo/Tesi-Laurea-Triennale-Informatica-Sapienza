func (db *appdbimpl) CheckTripUser(userID uuid.UUID, tripID int) (bool, error) {
	var parkingUser string
	err := db.c.Get(&parkingUser, "SELECT parkedby FROM parks WHERE `parkid` = ?", tripID)
	if err != nil {
		return false, fmt.Errorf("select1 statement: %w", err)
	}
	if userID.String() != parkingUser {
		return false, nil
	}
	return true, nil
}