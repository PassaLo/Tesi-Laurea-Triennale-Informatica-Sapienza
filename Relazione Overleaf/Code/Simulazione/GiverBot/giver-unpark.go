// Unpark the car in order to be able to park for the next simulation
	err = requests.UnparkCar(bot.UserID, bot.Cid, park, logger, cfg)
	if err != nil {
		logger.Error(bot.BotName, " has failed to unpark")
		wg.Done()
		return
	}
	logger.Println(bot.BotName, " has unparked correctly")