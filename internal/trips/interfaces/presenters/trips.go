package presenters

type AddTripRequest struct {
	Name string `json:"name"`
}

type AddTripResponse struct {
	Name string `json:"name"`
}