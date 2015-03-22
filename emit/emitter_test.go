package emit

import (
    "bytes"
    "testing"
	"github.com/compostware/quokka/library"
	"github.com/compostware/quokka/model"
)

const (
	compId1 = "comp-a"
	compFragment1 = "comp-a-fragment-a\n"
	
	compId2 = "comp-b"
	compFragment2 = "comp-b-fragment-b"
	
	compId3 = "comp-c"
	
	fragment = "Output line 1\noutput line 2"
)

var (
	lib = library.NewInMemoryLibrary()
	target = NewEmitter(lib)
	buf bytes.Buffer 
)

func TestEmitEmpty(t *testing.T) {
	clearAll()
	
	r := model.NewRealization()
	
	err := target.Emit(r, &buf)
	
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	} else if buf.String() != "" {
		t.Errorf("Expected empty output, got: %s", buf.String())
	}
}

func TestEmitWithSimpleFragment(t *testing.T) {
	clearAll()
	
	r := model.NewRealization()
	r.SetFragment(fragment)
	
	err := target.Emit(r, &buf)
	
	if err != nil {
		t.Errorf("Unexpected error: '%s'", err)
	} else if expected := fragment + "\n"; buf.String() != expected {
		t.Errorf("Expected output '%s', got: '%s'", expected, buf.String())
	}
}

func TestEmitWithComponent(t *testing.T) {
	clearAll()
	
	r := model.NewRealization()
	r.AddComponent(model.NewComponentReference(compId1))
	
	err := target.Emit(r, &buf)
	
	if err != nil {
		t.Errorf("Unexpected error: '%s'", err)
	} else if buf.String() != compFragment1 {
		t.Errorf("Expected output '%s', got: '%s'", compFragment1, buf.String())
	}
}

func TestEmitWithComponentWithoutNewline(t *testing.T) {
	clearAll()
	
	r := model.NewRealization()
	r.AddComponent(model.NewComponentReference(compId2))
	
	err := target.Emit(r, &buf)
	
	if err != nil {
		t.Errorf("Unexpected error: '%s'", err)
	} else if expected := compFragment2 + "\n"; buf.String() != expected {
		t.Errorf("Expected output '%s', got: '%s'", expected, buf.String())
	}
}

func TestEmitWithComponentWithBadComponentId(t *testing.T) {
	clearAll()
	
	r := model.NewRealization()
	r.AddComponent(model.NewComponentReference(compId3))
	
	err := target.Emit(r, &buf)
	
	if err == nil {
		t.Errorf("Expected an error, but got null")
	}
}

func TestEmitWithFullComposition(t *testing.T) {
	clearAll()
	
	r := model.NewRealization()
	r.AddComponent(model.NewComponentReference(compId1))
	r.AddComponent(model.NewComponentReference(compId2))
	r.SetFragment(fragment)
	
	err := target.Emit(r, &buf)
	
	expected := compFragment1 + 
		compFragment2 + "\n" +
		fragment + "\n";
	
	if err != nil {
		t.Errorf("Unexpected error: '%s'", err)
	} else if buf.String() != expected {
		t.Errorf("Expected output '%s', got: '%s'", expected, buf.String())
	}
}

// resets the shared test state, should be called at the beginning of each test
func clearAll() {
	buf.Reset()
	lib.Clear()	
	registerComponent(compId1, compFragment1)
	registerComponent(compId2, compFragment2)
}

func registerComponent(id, fragment string) {
	lib.Publish(model.NewComponent(id, fragment))
}

