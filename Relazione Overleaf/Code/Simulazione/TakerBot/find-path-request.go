// We use an external api to get the road to travel to get to the giver
	client := osrm.NewFromURL("https://router.project-osrm.org")

	ctx, cancelFn := context.WithTimeout(context.Background(), time.Second)
	defer cancelFn()

	resp, err := client.Route(ctx, osrm.RouteRequest{
		Profile: "car",
		Coordinates: osrm.NewGeometryFromPointSet(geo.PointSet{
			{takerLon, takerLat},
			{giverLon, giverLat},
		}),
		Steps:       osrm.StepsTrue,
		Annotations: osrm.AnnotationsFalse,
		Overview:    osrm.OverviewFalse,
		Geometries:  osrm.GeometriesPolyline6,
	})
	if err != nil {
		log.Printf("route failed: %v", err)
		return nil, nil, err
	}