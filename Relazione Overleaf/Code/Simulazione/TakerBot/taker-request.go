// Send a taker request
	err := requests.PutTaker(bot.UserID, bot.Cid, bot.GiverSchedule, logger, cfg)
	if err != nil {
		logger.Error(bot.BotName, " has failed to send taker request")
		wg.Done()
		return
	}
	logger.Print(bot.BotName, " has correctly send the taker request")