// Catch SIGTERM or SIGINTERRUPT. We're using a second one to 
// stop the simulation IF a match has not been found yet
	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, syscall.SIGTERM, syscall.SIGINT)

	// We'll need to memorize some taker information
	var takerInfo *requests.TakerInfo
	logger.Print(bot.BotName, " is waiting for matches...")

	for { // Waiting for a match
		Wait(cfg.Time.Sleep)

		// Check if any SIGTERM or SIGINTERRUPT signals were caught
		if len(stopChan) >= cap(stopChan) {
			sig := <-stopChan
			logger.Println("Caught signal ", sig)
			wg.Done()
			return
		}
		// Check for matches
		takerInfo,err = requests.GetMatchTakerInfo(bot.UserID, bot.Cid, logger, cfg)
		if err != nil {
			logger.Error(bot.BotName, " has failed to get match taker info")
			wg.Done()
			return
		}

		// If the giver request get matched, retrieve taker info
// and exit the waiting loop
		if !takerInfo.MatchSchedule.IsZero() {
			bot.MatchSchedule = takerInfo.MatchSchedule.Format(time.RFC3339)
			logger.Println(bot.BotName, " has found a match")
			break
		}
	}
// Save the Uid and Cid of the giver
	var oldBot []string
	oldBot = append(oldBot, bot.UserID.String())
	oldBot = append(oldBot, bot.Cid.String())

	// Save the Uid and Cid of the taker
	var newBot []string
	newBot = append(newBot, takerInfo.TakerID.String())
	newBot = append(newBot, takerInfo.TakerCid.String())

	park := requests.ParkCarInfo{
		Lat:          bot.ParkLat,
		Lon:          bot.ParkLon,
		ParkTime:     time.Now().Format(time.RFC3339),
		ParkAccuracy: 0.25,
	}