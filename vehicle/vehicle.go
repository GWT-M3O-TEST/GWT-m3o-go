package vehicle

import (
	"go.m3o.com/client"
)

type Vehicle interface {
	Lookup(*LookupRequest) (*LookupResponse, error)
}

func NewVehicleService(token string) *VehicleService {
	return &VehicleService{
		client: client.NewClient(&client.Options{
			Token: token,
		}),
	}
}

type VehicleService struct {
	client *client.Client
}

// Lookup a UK vehicle by it's registration number
func (t *VehicleService) Lookup(request *LookupRequest) (*LookupResponse, error) {

	rsp := &LookupResponse{}
	return rsp, t.client.Call("vehicle", "Lookup", request, rsp)

}

type LookupRequest struct {
	// the vehicle registration number
	Registration string `json:"registration,omitempty"`
}

type LookupResponse struct {
	// make of vehicle
	Make string `json:"make,omitempty"`
	// registration number
	Registration string `json:"registration,omitempty"`
	// date of last v5 issue
	LastV5Issued string `json:"last_v5_issued,omitempty"`
	// engine capacity
	EngineCapacity int32 `json:"engine_capacity,omitempty"`
	// month of first registration
	MonthOfFirstRegistration string `json:"month_of_first_registration,omitempty"`
	// tax status
	TaxStatus string `json:"tax_status,omitempty"`
	// wheel plan
	Wheelplan string `json:"wheelplan,omitempty"`
	// colour of vehicle
	Colour string `json:"colour,omitempty"`
	// mot status
	MotStatus string `json:"mot_status,omitempty"`
	// fuel type e.g petrol, diesel
	FuelType string `json:"fuel_type,omitempty"`
	// url of logo for the make
	LogoUrl string `json:"logo_url,omitempty"`
	// mot expiry
	MotExpiry string `json:"mot_expiry,omitempty"`
	// tax due data
	TaxDueDate string `json:"tax_due_date,omitempty"`
	// type approvale
	TypeApproval string `json:"type_approval,omitempty"`
	// year of manufacture
	YearOfManufacture int32 `json:"year_of_manufacture,omitempty"`
	// co2 emmissions
	Co2Emissions float64 `json:"co2_emissions,omitempty"`
}
