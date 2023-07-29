// Step 4: Search for giver requests and assign schedule
		// search for giver request until there are one or more valid giver requests
		var giverList []requests.GiverList

		logger.Print("searching giver..")
		for {
			// If we get a SIGKILL signal it's necessary to stop the simulation
			if len(cancelChan) >= cap(cancelChan) {
				sig := <-cancelChan
				logger.Println("Caught signal ", sig, ". Interrupting Simulation")
				// Stop the simulation by putting an EndTime in the db
				err = requests.StopSimulation(botList[0].UserID, simulationID, logger, cfg)
				if err != nil {
					logger.Error("failed to stop the simulation")
				}
				return nil
			}
			// Waiting the giver request
			Wait(cfg.Time.Sleep)
			giverList, err = requests.GetGiverList(botList[0].UserID, botList[0].Cid, logger, cfg)
			if err != nil {
				logger.Error("failed to get giver list")
			}

			if len(giverList) > 0 {
				var validTime = time.Now()

				// Take the valid schedule for bot to complete the match
				_, min, sec := validTime.Clock()
				if min%cfg.Time.Schedule != 0 || sec != 0 {
					validTime = validTime.Add(time.Duration(-(min%cfg.Time.Schedule)*60-sec) * 
                    time.Second).Add(time.Duration(cfg.Time.Schedule) * time.Minute)
				}

				// Filter the request with a valid schedule
				n := 0
				for _, giver := range giverList {
					if giver.Schedule.Format(time.RFC3339) == validTime.Format(time.RFC3339) {
						giverList[n] = giver
						n++
					}
				}

				giverList = giverList[:n]

				if n > 0 {
					break
				}
			}
		}
		min := min(len(giverList), len(botList))
		logger.Print("found ", len(giverList), " givers, ", min, " requests are going to be fulfilled")

		// Assign giver schedule to the bot
		for i := 0; i < min; i++ {
			botList[i].GiverSchedule = *giverList[i].Schedule
		}
