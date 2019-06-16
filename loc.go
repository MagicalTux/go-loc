package loc

type Location struct {
	Name     string
	Lat, Lon float64
	Parent   *Location
}
