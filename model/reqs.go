/*
	Copyright 2015 Lachlan Coote
	
	Licensed under the Apache License, Version 2.0 (the "License");
	you may not use this file except in compliance with the License.
	You may obtain a copy of the License at
	
		http://www.apache.org/licenses/LICENSE-2.0
		
	Unless required by applicable law or agreed to in writing, software
	distributed under the License is distributed on an "AS IS" BASIS,
	WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
	See the License for the specific language governing permissions and
	limitations under the License.
*/


package model


// An abstraction of an expression of some service that is required by a spec or some component thereof. 
type Requirement interface {
	
}

// A requirement for a specific type of service.
type RequiredService struct {
	ServiceId string
}

// A requirement for a specific identified component that may provide one or more services.
type RequiredComponent struct {
	ComponentId string
}

// A constructor method for required services.
func NewRequiredService(serviceId string) *RequiredService {
	req := new(RequiredService)
	req.ServiceId = serviceId
	return req
}

// A constructor method for required components.
func NewRequiredComponent(componentId string) *RequiredComponent {
	req := new(RequiredComponent)
	req.ComponentId = componentId
	return req
}

