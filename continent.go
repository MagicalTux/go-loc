package loc

var (
	Eurasia      = &Location{Name: "Eurasia", Lat: 45.4507, Lon: 68.8319, Parent: Earth}
	Europe       = &Location{Name: "Europe", Lat: 54.5260, Lon: 15.2551, Parent: Eurasia}
	Asia         = &Location{Name: "Asia", Lat: 34.0479, Lon: 100.6197, Parent: Eurasia}
	Americas     = &Location{Name: "Americas", Lat: 54.5260, Lon: 105.2551, Parent: Earth}
	NorthAmerica = &Location{Name: "North America", Lat: 54.5260, Lon: 105.2551, Parent: Americas}
	SouthAmerica = &Location{Name: "South America", Lat: 8.7832, Lon: 55.4915, Parent: Americas}
	Australia    = &Location{Name: "Australia", Lat: 25.2744, Lon: 133.7751, Parent: Earth}
	Africa       = &Location{Name: "Africa", Lat: 8.7832, Lon: 34.5085, Parent: Earth}
)
