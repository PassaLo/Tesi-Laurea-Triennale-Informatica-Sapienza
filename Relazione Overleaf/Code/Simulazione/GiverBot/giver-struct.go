type Bot struct {
	BotName       string
	UserID        uuid.UUID
	Cid           uuid.UUID
	GiverSchedule time.Time
	MatchSchedule string
	ParkLat       float64
	ParkLon       float64
}