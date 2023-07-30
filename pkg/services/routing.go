package services

import "math"

type Routing struct{}

func (r *Routing) Distance(lat1, lon1, lat2, lon2 float64) float64 {
	const earthRadiusKm = 6371.0
	dLat := r.DegToRad(lat2 - lat1)
	dLon := r.DegToRad(lon2 - lon1)
	a := math.Sin(dLat/2)*math.Sin(dLat/2) + math.Cos(r.DegToRad(lat1))*math.Cos(r.DegToRad(lat2))*math.Sin(dLon/2)*math.Sin(dLon/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	return earthRadiusKm * c
}

func (r *Routing) DegToRad(deg float64) float64 {
	return deg * (math.Pi / 180.0)
}
