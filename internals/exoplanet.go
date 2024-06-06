package internals

type ExoplanetType string

const (
	GasGiant    ExoplanetType = "GasGiant"
	Terrestrial ExoplanetType = "Terrestrial"
)

type Exoplanet struct {
	ID          string        `json:"id"`
	Name        string        `json:"name"`
	Description string        `json:"description"`
	Distance    int           `json:"distance"`       // in light years
	Radius      float64       `json:"radius"`         // in Earth-radius units
	Mass        *float64      `json:"mass,omitempty"` // in Earth-mass units (only for Terrestrial)
	Type        ExoplanetType `json:"type"`
}

var exoplanets = map[string]Exoplanet{}
