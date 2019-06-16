package loc

var (
	Eurasia      = makeLoc("Eurasia", 45.4507, 68.8319, Earth)
	Europe       = makeLoc("Europe", 54.5260, 15.2551, Eurasia)
	Asia         = makeLoc("Asia", 34.0479, 100.6197, Eurasia)
	Americas     = makeLoc("Americas", 54.5260, -105.2551, Earth)
	NorthAmerica = makeLoc("North America", 54.5260, -105.2551, Americas)
	SouthAmerica = makeLoc("South America", -8.7832, -55.4915, Americas)
	Australia    = makeLoc("Australia", 25.2744, 133.7751, Earth)
	Africa       = makeLoc("Africa", 8.7832, 34.5085, Earth)
)
