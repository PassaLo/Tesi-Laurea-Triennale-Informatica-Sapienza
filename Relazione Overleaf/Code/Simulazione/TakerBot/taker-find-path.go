// Now we save the simulationID in the trip tied to this match 
// (we got it in the GetMatchGiverInfo)
	err = requests.UpdateTripSimulation(bot.UserID, bot.SimulationID,    
    bot.CurrentTripID, logger, cfg)
	if err != nil {
		logger.Error(bot.BotName," failed to add simulationId to the trip")
	}

	// Start notifying the user with his position
	logger.Print(bot.BotName, " is updating taker position...")
	// As stated before the api only allows 2 requests at the time
	sem <- true
	// Find a route to follow to get to the giver
	path, pathLength, err := requests.FindPath(bot.GiverLat, bot.GiverLon, 
     bot.StartingLat, bot.StartingLon)
	if err != nil {
		logger.Error(bot.BotName, " has failed to find a path")
		wg.Done()
		return
	}
	Wait(1000)
	func() { <-sem }()