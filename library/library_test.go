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

package library

import (
    "testing"
    "github.com/compostware/quokka/model"
)

const (
	compIdA = "a"
	compIdB = "b"
	compFragA = "a-diddly-a-a"
	compFragB = "b-lli-b-ly"
)

func TestPublishNewComponent(t *testing.T) {
	lib := buildTargetLibWithCompAPublished()
	
	result, err := lib.Retrieve(compIdA)
	
	if err != nil {
		t.Errorf("Unexpected error received %s", err)
	} else if result == nil {
		t.Error("Publish failed, nil returned by retrieve")
	} else if result.Id != compIdA || result.Fragment != compFragA {
		t.Errorf("Publish failed, expected (%s,%s) but got (%s, %s)", 
			compIdA, compFragA,
			result.Id, result.Fragment)
	}
}

func TestRePublishComponent(t *testing.T) {
	lib := buildTargetLibWithCompAPublished()
	
	compA2 := model.NewComponent(compIdA, compFragB)
	err := lib.Publish(compA2)
	
	result, err := lib.Retrieve(compIdA)
	
	if err != nil {
		t.Errorf("Unexpected error received %s", err)
	} else if result == nil {
		t.Error("Re-Publish failed, nil returned by retrieve")
	} else if result.Id != compIdA || result.Fragment != compFragB {
		t.Errorf("Re-Publish failed, expected (%s,%s) but got (%s, %s)", 
			compIdA, compFragB,
			result.Id, result.Fragment)
	}
}

func TestRetrieveNonExistentComponent(t *testing.T) {
	lib := buildTargetLibWithCompAPublished()
	
	result, err := lib.Retrieve(compIdB)
	
	if err == nil {
		t.Error("Error expected, none received")
	} else if result != nil {
		t.Errorf("Retrieve failed, expected nil but got (%s,%s)",  
			result.Id, result.Fragment)
	}
}

func buildTargetLibWithCompAPublished() LibraryClient {
	compA := model.NewComponent(compIdA, compFragA)
	
	lib := NewLibraryClient()
	lib.Publish(compA)
	
	return lib
}

