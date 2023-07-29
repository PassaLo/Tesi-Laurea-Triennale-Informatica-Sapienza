// Now that we have at least one bot ready start the simulation
	simulationID, err := requests.StartSimulation(csv[0][0], logger, cfg)
	if err != nil {
		logger.Error("failed to start simulation")
		return err
	}