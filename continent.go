package loc

var (
	Eurasia      = &Location{Name: "Eurasia", Lat: 45.4507, Lon: 68.8319, Parent: Earth, Depth: Earth.Depth + 1}
	Europe       = &Location{Name: "Europe", Lat: 54.5260, Lon: 15.2551, Parent: Eurasia, Depth: Eurasia.Depth + 1}
	Asia         = &Location{Name: "Asia", Lat: 34.0479, Lon: 100.6197, Parent: Eurasia, Depth: Eurasia.Depth + 1}
	Americas     = &Location{Name: "Americas", Lat: 54.5260, Lon: 105.2551, Parent: Earth, Depth: Earth.Depth + 1}
	NorthAmerica = &Location{Name: "North America", Lat: 54.5260, Lon: 105.2551, Parent: Americas, Depth: Americas.Depth + 1}
	SouthAmerica = &Location{Name: "South America", Lat: 8.7832, Lon: 55.4915, Parent: Americas, Depth: Americas.Depth + 1}
	Australia    = &Location{Name: "Australia", Lat: 25.2744, Lon: 133.7751, Parent: Earth, Depth: Earth.Depth + 1}
	Africa       = &Location{Name: "Africa", Lat: 8.7832, Lon: 34.5085, Parent: Earth, Depth: Earth.Depth + 1}
)
