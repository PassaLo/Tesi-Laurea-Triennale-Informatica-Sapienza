	// For now, we're using this time for the schedule, 
  // but we could also change it for a custom one
	bot.GiverSchedule = time.Now().Add(time.Minute)
	// Send a giver request
	err := requests.PutGiver(bot.UserID, bot.Cid, bot.GiverSchedule, logger, cfg)
	if err != nil {
		logger.Error(bot.BotName, " has failed to send giver request")
		Wait(cfg.Time.Sleep)
		wg.Done()
		return
	}
	logger.Print(bot.BotName, " has correctly send the giver request")