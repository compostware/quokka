package model


// An abstraction of an expression of some service that is required by a spec or some component thereof. 
type Requirement interface {
	
}

// A requirement for a specific identified service.
type RequiredService struct {
	ServiceId string
}

// A constructor method for required services.
func NewRequiredService(serviceId string) (req *RequiredService) {
	req = new(RequiredService)
	req.ServiceId = serviceId
	return
}

