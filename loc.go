package loc

import "github.com/MagicalTux/go-cheapruler"

type Location struct {
	Name     string
	Lat, Lon float64
	Parent   *Location
	Depth    int
}

func (l *Location) Distance(l2 *Location) float64 {
	// compute distance with ruler
	cr := cheapruler.New(l.Lon, cheapruler.Kilometers)
	return cr.Distance([]float64{l.Lat, l.Lon}, []float64{l2.Lat, l2.Lon})
}

func makeLoc(name string, lat, lon float64, parent *Location) *Location {
	return &Location{
		Name:   name,
		Lat:    lat,
		Lon:    lon,
		Parent: parent,
		Depth:  parent.Depth + 1,
	}
}
