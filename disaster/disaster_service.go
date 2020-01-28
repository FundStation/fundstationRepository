package disaster


import "github.com/FundStation2/models"

type DisasterService interface {
	ViewDisaster() ([]models.Disaster, error)
	ViewEvent() ([]models.Event, error)
}