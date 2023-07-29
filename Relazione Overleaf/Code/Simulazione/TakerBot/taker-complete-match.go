// Complete the match
err = requests.CompleteMatch(bot.UserID, bot.Cid, bot.MatchSchedule, logger, cfg)
if err != nil {
	logger.Error(bot.BotName, " has failed to complete the match")
	wg.Done()
	return
}
logger.Println(bot.BotName, " has completed the match")