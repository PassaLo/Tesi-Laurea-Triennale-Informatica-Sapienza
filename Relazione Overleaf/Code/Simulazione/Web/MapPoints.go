type MapPoints struct {
	Lat      float64 `json:"lat" db:"lat"`
	Lon      float64 `json:"lon" db:"lon"`
	GiverLat float64 `json:"giverlat" db:"giverlat"`
	GiverLon float64 `json:"giverlon" db:"giverlon"`
}