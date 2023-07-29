var giverInfo *requests.GiverInfo
	logger.Print(bot.BotName, " is waiting for matches...")

	for { // Waiting for the match to start
		Wait(cfg.Time.Sleep)
		giverInfo,err = requests.GetMatchGiverInfo(bot.UserID, bot.Cid, logger, cfg)
		if err != nil {
			logger.Error(bot.BotName, " has failed to get match giver info")
			wg.Done()
			return
		}
		// If the taker request get matched, retrieve giver info and exit the 
    // waiting loop
		if giverInfo.GiverLat != 0 && giverInfo.GiverLon != 0 {
			bot.CurrentTripID = giverInfo.GiverParkID
			bot.GiverLat = giverInfo.GiverLat
			bot.GiverLon = giverInfo.GiverLon
			bot.MatchSchedule = giverInfo.MatchSchedule.Format(time.RFC3339)
			break
		}
	}