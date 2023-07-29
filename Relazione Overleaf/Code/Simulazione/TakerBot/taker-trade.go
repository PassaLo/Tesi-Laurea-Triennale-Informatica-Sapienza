	// Save the Uid and Cid of the taker
	var oldBot []string
	oldBot = append(oldBot, bot.UserID.String())
	oldBot = append(oldBot, bot.Cid.String())

	// Save the Uid, the Cid and the park coordinates of the giver
	var newBot []string
	newBot = append(newBot, giverInfo.GiverID.String())
	newBot = append(newBot, giverInfo.GiverCid.String())
	newBot = append(newBot, fmt.Sprintf("%f", bot.GiverLat))
	newBot = append(newBot, fmt.Sprintf("%f", bot.GiverLon))

	// Now we replace the taker with the giver in the taker's csv file. 
  // By doing this a giver becomes a taker and will
	// start from where it left as a giver.
	// (The old giver does the same thing for the taker, 
  // allowing it to become a giver)
	csv := config.ReadCSV(logger, cfg.Host.CsvFile)
	config.SubstituteCSV(logger, csv, oldBot, newBot, cfg.Host.CsvFile)

	logger.Println(bot.BotName, " became a giver")
	logger.Println("Waiting for every match to end")
	wg.Done()
}