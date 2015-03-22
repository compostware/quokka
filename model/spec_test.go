package model

import (
    "fmt"
    "reflect"
    "testing"
)


const (
	serv1, serv2 = "leia", "donut"
	frag = "yodel"
)

var (
	req1, req2 = NewRequiredService(serv1), NewRequiredService(serv2)
	reqs = []Requirement{ req1, req2 }
)


func TestEmptySpec(t *testing.T) {
	spec := NewSpec()
	if spec.Fragment() != "" {
		t.Errorf("Unexpected fragment value %s", spec.Fragment())		
	} else if len(spec.Requires()) != 0 {
		t.Errorf("Unexpected requirements %s", spec.Requires())		
	}
}

func TestPopulatedSpec(t *testing.T) {
	spec := NewSpec()
	spec.AddRequirement(req1)
	spec.AddRequirement(req2)
	spec.SetFragment(frag)
	
	if spec.Fragment() != frag {
		t.Errorf("Unexpected fragment value %s", spec.Fragment())		
	} else if !reflect.DeepEqual(spec.Requires(), reqs) {
		t.Errorf("Unexpected requirements %s", spec.Requires())		
	}
}

// Just validating my understanding and use of "append()".
func TestSpecWithReqsOutstripCapacity(t *testing.T) {
	spec := NewSpec()
	
	reqCap := defaultReqCapacity + 1
	reqs := make ([]Requirement, reqCap, reqCap)
	
	for i := 0; i < reqCap; i++ {
		req := NewRequiredService(fmt.Sprintf("a%n", i))
		reqs[i] = req
		spec.AddRequirement(req)
	}
	
	if !reflect.DeepEqual(spec.Requires(), reqs) {
		t.Errorf("Unexpected requirements %s", spec.Requires())		
	}
}

