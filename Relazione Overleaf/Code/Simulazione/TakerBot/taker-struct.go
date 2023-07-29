type Bot struct {
	BotName       string
	UserID        uuid.UUID
	Cid           uuid.UUID
	GiverSchedule time.Time
	MatchSchedule string
	GiverLat      float64
	GiverLon      float64
	StartingLat   float64
	StartingLon   float64
	CurrentTripID int
	SimulationID  int
}