// Follow said route
	for i, point := range path {
		// Save in the db the coordinates we're reaching
		err = requests.SaveCoordinates(bot.UserID, logger, cfg,   bot.CurrentTripID, 
        point.Y(), point.X())
		if err != nil {
			logger.Error(bot.BotName, " has failed to update taker position")
		}

		takerPos := requests.TakerPosition{
			ParkLat:  point.Y(),
			ParkLon:  point.X(),
			Schedule: bot.MatchSchedule,
			Eta:      path.Length() - i,
		}

		// Calculate the time the taker will need to reach those coordinates
		randomNum := cfg.Time.MinTime
		if (cfg.Time.MaxTime - cfg.Time.MinTime) > 0 {
			bg := big.NewInt(cfg.Time.MaxTime - cfg.Time.MinTime)
			n, err := rand.Int(rand.Reader, bg)
			if err != nil {
				logger.Error(bot.BotName," has failed to generate a random wait number")
			}
			randomNum = n.Int64() + cfg.Time.MinTime
		}
		// The time needed is defined by a random time with a max and a min 
    // written in the config file * the length of this road segment
		Wait(int(randomNum) * pathLength[i])
		// Update position
		err = requests.UpdateTakerPosition(bot.UserID,bot.Cid,logger,cfg, takerPos)
		if err != nil {
			logger.Error(bot.BotName, " has failed to update taker position")
		}
	}