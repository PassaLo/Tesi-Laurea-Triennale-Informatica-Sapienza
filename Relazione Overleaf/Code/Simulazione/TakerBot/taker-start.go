		// Step 5: Freedom: the bot is ready to simulate
		// each bot is now independent of the others to complete the task

		// The api we use to get the road only accepts 2 request at the time, 
    // so we need a semaphore to limit them
		sem := make(chan bool, 2)

		// Start the simulation for each bot
		for i := 0; i < min; i++ {
			wg.Add(1)
			go simulate(botList[i], logger, cfg, sem)
		}
		wg.Wait()

		// Check if any SIGTERM or SIGINTERRUPT signals were caught
		if len(cancelChan) >= cap(cancelChan) {
			sig := <-cancelChan
			logger.Println("Caught signal ", sig, ". Interrupting Simulation")
			// Stop the simulation by putting an EndTime in the db
			err = requests.StopSimulation(botList[0].UserID, simulationID, logger, cfg)
			if err != nil {
				logger.Error("failed to stop the simulation")
			}
			break
		}
  }