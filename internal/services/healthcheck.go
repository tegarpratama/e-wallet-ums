package services

import "ewallet-ums/internal/interfaces"

type Healthcheck struct {
	HealthcheckReposity interfaces.IHealthcheckRepo
}

func (s *Healthcheck) HealthcheckServices() (string, error) {
	return "it's works", nil
}
