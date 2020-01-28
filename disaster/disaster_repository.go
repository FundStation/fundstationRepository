package disaster


import "github.com/FundStation/models"

type DisasterRepository interface {
	SelectDisaster() ([]models.Disaster, error)
	SelectEvent() ([]models.Event, error)
}
