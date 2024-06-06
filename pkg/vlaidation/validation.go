package validation

import (
	"errors"
	"exoplanet-service/internals"
)

func ValidateExoplanet(exoplanet internals.Exoplanet) error {
	if exoplanet.Name == "" {
		return errors.New("Name is required")
	}
	if exoplanet.Description == "" {
		return errors.New("Description is required")
	}
	if exoplanet.Distance <= 10 || exoplanet.Distance >= 1000 {
		return errors.New("Distance must be between 10 and 1000 light years")
	}
	if exoplanet.Radius <= 0.1 || exoplanet.Radius >= 10 {
		return errors.New("Radius must be between 0.1 and 10 Earth-radius units")
	}
	if exoplanet.Type == internals.Terrestrial {
		if exoplanet.Mass == nil {
			return errors.New("Mass is required for Terrestrial exoplanet")
		}
		if *exoplanet.Mass <= 0.1 || *exoplanet.Mass >= 10 {
			return errors.New("Mass must be between 0.1 and 10 Earth-mass units")
		}
	}
	return nil
}
