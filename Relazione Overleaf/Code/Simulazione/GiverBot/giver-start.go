	// Step 4: Simulation
	// each bot is now independent of the others to complete the task
	// Catch SIGTERM or SIGINTERRUPT. By doing this we can't interrupt a not ended match
	cancelChan := make(chan os.Signal, 1)
	signal.Notify(cancelChan, syscall.SIGTERM, syscall.SIGINT)

	// Start the simulation for each bot
	for i := 0; i < len(botList); i++ {
		wg.Add(1)
		go simulate(botList[i], logger, cfg)
	}

	wg.Wait()
	// Check if any SIGTERM or SIGINTERRUPT signals were caught
	if len(cancelChan) >= cap(cancelChan) {
		sig := <-cancelChan
		logger.Println("Caught signal ", sig)
		break
	}
}