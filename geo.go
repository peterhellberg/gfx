package gfx

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
func (gp GeoPoint) Vec(tileSize, zoom int) Vec {
	scale := MathPow(2, float64(zoom))
	fts := float64(tileSize)

	return V(
		((float64(gp.Lon)+180)/360)*scale*fts,
		(fts/2)-(fts*MathLog(MathTan((Pi/4)+((float64(gp.Lat)*Pi/180)/2)))/(2*Pi))*scale,
	)
}

// GeoTile for the GeoPoint at the given zoom level.
func (gp GeoPoint) GeoTile(zoom int) GeoTile {
	latRad := Degrees(gp.Lat).Radians()
	n := MathPow(2, float64(zoom))

	return GeoTile{
		Zoom: zoom,
		X:    int(n * (float64(gp.Lon) + 180) / 360),
		Y:    int((1.0 - MathLog(MathTan(latRad)+(1/MathCos(latRad)))/Pi) / 2.0 * n),
	}
}

// NewGeoPointFromTileNumbers creates a new GeoPoint based on the given tile numbers.
// https://wiki.openstreetmap.org/wiki/Slippy_map_tilenames#Tile_numbers_to_lon..2Flat.
func NewGeoPointFromTileNumbers(zoom, x, y int) GeoPoint {
	n := MathPow(2, float64(zoom))
	latRad := MathAtan(MathSinh(Pi * (1 - (2 * float64(y) / n))))

	return GP(latRad*180/Pi, (float64(x)/n*360)-180)
}

// GeoTile consists of a Zoom level, X and Y values.
type GeoTile struct {
	Zoom int
	X    int
	Y    int
}

// GT creates a new GeoTile.
func GT(zoom, x, y int) GeoTile {
	return GeoTile{Zoom: zoom, X: x, Y: y}
}

// GeoPoint for the GeoTile.
func (gt GeoTile) GeoPoint() GeoPoint {
	n := MathPow(2, float64(gt.Zoom))
	latRad := MathAtan(MathSinh(Pi * (1 - (2 * float64(gt.Y) / n))))

	return GP(latRad*180/Pi, (float64(gt.X)/n*360)-180)
}

// Rawurl formats a URL string with Zoom, X and Y.
func (gt GeoTile) Rawurl(format string) string {
	return Sprintf(format, gt.Zoom, gt.X, gt.Y)
}
