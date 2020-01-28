package disaster


import "github.com/FundStation/models"

type DisasterService interface {
	ViewDisaster() ([]models.Disaster, error)
	ViewEvent() ([]models.Event, error)
}
