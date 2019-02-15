package gfx

import "math"

// GeoPoint represents a geographic point with Lat/Lon.
type GeoPoint struct {
	Lon float64
	Lat float64
}

// GP creates a new GeoPoint
func GP(lat, lon float64) GeoPoint {
	return GeoPoint{Lon: lon, Lat: lat}
}

// Vec returns a vector for the geo point based on the given tileSize and zoom level.
func (gp GeoPoint) Vec(tileSize, zoom float64) Vec {
	scale := MathPow(2, zoom)

	return V(
		((gp.Lon+180)/360)*scale*tileSize,
		(tileSize/2)-(tileSize*math.Log(math.Tan((math.Pi/4)+
			((gp.Lat*math.Pi/180)/2)))/(2*math.Pi))*scale,
	)
}
