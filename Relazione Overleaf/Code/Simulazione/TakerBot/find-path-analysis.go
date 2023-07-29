var path geo.PointSet
	var pathLength []int

	for _, route := range resp.Routes {
		for _, leg := range route.Legs {
			// leg.Steps =The entire trip
			for i, step := range leg.Steps {
				// step = one segment of the entire road
				for _, p := range step.Geometry.PointSet {
					// We memorize the length of this segment as  
          // the number of coordinate sets in it
					// We limit it at 15 to avoid a long waiting
					if step.Geometry.PointSet.Length() >= 15 {
						pathLength = append(pathLength, 15)
					} else {
						pathLength = append(pathLength, step.Geometry.PointSet.Length())
					}
					// Every coordinate set of this segment
					point := p
					path.Push(&point)
					// We only use the first coordinate set of each segment because otherwise we
					// would have way too many coordinates to follow and memorize for each trip.
					break
				}
				// Sometimes the last segment is composed of more than a single coordinates 
				// set. This check is needed to get the real last coordinate
				if i == len(leg.Steps)-1 {
					lastPoint := step.Geometry.PointSet[len(step.Geometry.PointSet)-1]
					if lastPoint != *path.Last() {
						path.Push(&lastPoint)
					}
				}
			}
		}
	}
	return path, pathLength, nil
}