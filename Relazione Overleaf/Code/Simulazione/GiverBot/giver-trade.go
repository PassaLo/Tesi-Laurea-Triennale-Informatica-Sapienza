  csv := config.ReadCSV(logger, cfg.Host.CsvFile)
  config.SubstituteCSV(logger, csv, oldBot, newBot, cfg.Host.CsvFile)

  logger.Println(bot.BotName, " became a taker")
  logger.Println("Waiting for every match to end")
  wg.Done()