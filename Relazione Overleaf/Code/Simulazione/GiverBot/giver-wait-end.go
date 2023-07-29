// Waiting for the match to end
for { 
    Wait(cfg.Time.Sleep)
    takerInfo,err = requests.GetMatchTakerInfo(bot.UserID, bot.Cid, logger, cfg)
    if err != nil {
	     logger.Error(bot.BotName, " has failed to get old match info")
	     wg.Done()
	     return
    }
    // Check if the old match ended
    if takerInfo.MatchSchedule.IsZero() {
	     logger.Println(bot.BotName, " previous match ended")
	     break
    }
}