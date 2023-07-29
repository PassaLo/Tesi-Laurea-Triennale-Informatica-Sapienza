for {
		var botList []types.Bot //  available bot in runtime
		csv = config.ReadCSV(logger, cfg.Host.CsvFile)

		//  Load all the data relative to the bot from the csv file
		//  it stops when we reach the desired number of bots
		for i, record := range csv {
			if len(botList) >= cfg.Bot.Quantity {
				break
			}

			var bot types.Bot
			uid, _ := uuid.FromString(record[0])
			name, err := requests.GetUserNickname(uid, logger, cfg)
			if err != nil {
				logger.Warn("bot not retrieved")
			} else {
				bot.BotName = name
				logger.Print(bot.BotName, "'s info correctly retrieved")
				bot.UserID = uid
				bot.Cid, _ = uuid.FromString(record[1])

				// In this case the already created car is not valid, 
        // so we must create a new one
				if !requests.ValidCar(uid, bot.Cid, logger, cfg) {
					cid, reqErr := requests.CreateCar(bot.UserID, i, logger, cfg)
					bot.Cid, _ = uuid.FromString(cid)
					if reqErr != nil {
						logger.Error(bot.BotName, "'s car not created")
						continue
					} else {
						csv[i][1] = bot.Cid.String()
						botList = append(botList, bot)
						logger.Print(bot.BotName, "'s car correctly created")
					}
				} else {
					botList = append(botList, bot)
					logger.Print(bot.BotName, "'s car correctly retrieved")
				}
			}

			// Generate random coordinates to park the car at a random parking spot
			err = requests.RandomCoordinates(&bot, logger, cfg)
			if err != nil {
				logger.Error(bot.BotName, " failed to generate coordinates for the park")
				continue
			}
			// The newly created car must be parked since the bot is a giver
			park := requests.ParkCarInfo{
				Lat:          bot.ParkLat,
				Lon:          bot.ParkLon,
				ParkTime:     time.Now().Format(time.RFC3339),
				ParkAccuracy: 0.25,
			}

			err = requests.ParkCar(bot.UserID, bot.Cid, park, logger, cfg)
			if err != nil {
				logger.Error(bot.BotName, " has failed to park")
				continue
			}
			logger.Print(bot.BotName, "'s car correctly parked")
		}