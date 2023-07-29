park := requests.ParkCarInfo{
		Lat:          bot.GiverLat,
		Lon:          bot.GiverLon,
		ParkTime:     time.Now().Format(time.RFC3339),
		ParkAccuracy: 0.25,
	}

// Park the car in order to assign a parkid to the match
err = requests.ParkCar(bot.UserID, bot.Cid, park, logger, cfg)
if err != nil {
	logger.Error(bot.BotName, " has failed to park")
	wg.Done()
	return
}
logger.Println(bot.BotName, " has parked correctly")

// Unpark the car in order to be able to park next time as a giver
err = requests.UnparkCar(bot.UserID, bot.Cid, park, logger, cfg)
if err != nil {
	logger.Error(bot.BotName, " has failed to unpark")
	wg.Done()
	return
}
logger.Println(bot.BotName, " has unparked correctly")