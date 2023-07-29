// RandomCoordinates Generate Random coordinates in an area to park
func RandomCoordinates(bot *types.Bot, logger logrus.FieldLogger, 
cfg config.BotConfiguration) error {
	/*
    .
    .
    .
    Section to generate a set or random coordinates inside the 
    requested area
    .
    .
    .
    */
	  // In this next section we make the request to Open Street Map to find 
    // the nearest road to the generated coordinates where the car can park

		// Construct the custom client.
		client := overpass.NewWithSettings("https://overpass-api.de/api/interpreter", 
cfg.Bot.Quantity, http.DefaultClient)

		// Make the query. 
		query := fmt.Sprint("[out:json];(way(around:50," + fmt.Sprintf("%f", cx+x) + "," 
+ fmt.Sprintf("%f", cy+y) + ")
[highway~'^(primary|secondary|tertiary|residential)$'][name];);out geom;")
		result, err := client.Query(query)
		if err != nil {
			logger.Error(bot.BotName, " has failed to send API request to open street map")
			return err
		}

	// If the request didn't find any available roads we'll go back to generate 
    // a new set of  random coordinates
		if len(result.Ways) > 0 {
			// The request could return 0 or many roads, but we just need one. 
			for _, c := range result.Ways {
        // For each road we get many coordinates set but we only need one 
        // (in our case, we use the center point of the road)
				i := len(c.Geometry) / 2
				bot.ParkLat = c.Geometry[i].Lat
				bot.ParkLon = c.Geometry[i].Lon
				return nil
			}
		}
		time.Sleep(time.Duration(2) * time.Second)
	}
}