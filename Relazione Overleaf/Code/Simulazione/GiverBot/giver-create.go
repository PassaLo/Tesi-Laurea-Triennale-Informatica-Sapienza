// Step 3: Create and Load Bots info and writes the csv file
	var csv [][]string //  total bot
	csv = config.ReadCSV(logger, cfg.Host.CsvFile)

	// Create some bots if there aren't enough in the csv
	if len(csv) < cfg.Bot.Quantity {
		for i := len(csv); i < cfg.Bot.Quantity; i++ {
			// The newly created bot
			var bot types.Bot
			// The record to add in the csv file
			var record []string
			bot.BotName = fmt.Sprintf("%s%d", cfg.Bot.Name, len(csv)+1)
			// Create the bot
			uid, reqErr := requests.CreateUser(bot.BotName, logger, cfg)
			bot.UserID, _ = uuid.FromString(uid)
			if reqErr != nil {
				logger.Error(bot.BotName, " not created")
				continue
			} else {
				logger.Print(bot.BotName, " correctly created")
			}

			record = append(record, bot.UserID.String())

			// Create the car
			cid, reqErr := requests.CreateCar(bot.UserID, len(csv)+1, logger, cfg)
			bot.Cid, _ = uuid.FromString(cid)
			if reqErr != nil {
				logger.Error(bot.BotName, "'s car not created")
				continue
			}
			logger.Print(bot.BotName, "'s car correctly created")

			record = append(record, bot.Cid.String())
			csv = append(csv, record)
		}
		// Rewrite the csv file with the new bots added
		config.WriteCSV(logger, csv, cfg.Host.CsvFile)
	}