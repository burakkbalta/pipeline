package pipeline

import (
	"reflect"
	"testing"
)

type Add struct{}
type Square struct{}

func (s Square) Execute(context *int) {
	*context = (*context) * (*context)
}

func (a Add) Execute(context *int) {
	*context = (*context) + (*context)
}

func TestPipeline_SquareAdd(t *testing.T) {
	context := 3
	expected := 18 // (3*3) - (9+9)

	pipeline := CreatePipeline[int]()
	pipeline.Next(Square{}).Next(Add{})
	pipeline.Execute(&context)

	if !reflect.DeepEqual(context, expected) {
		t.Errorf("Expected %v, got %v", expected, context)
	}
}

func TestPipeline_AddSquare(t *testing.T) {
	context := 3
	expected := 36 // (3+3) - (6*6)

	pipeline := CreatePipeline[int]()
	pipeline.Next(Add{}).Next(Square{})
	pipeline.Execute(&context)

	if !reflect.DeepEqual(context, expected) {
		t.Errorf("Expected %v, got %v", expected, context)
	}
}

func TestPipeline_MultipleAdd(t *testing.T) {
	context := 3
	expected := 48 // (3+3) - (6+6) - (12+12) - (24+24)

	pipeline := CreatePipeline[int]()
	pipeline.Next(Add{}).Next(Add{}).Next(Add{}).Next(Add{})
	pipeline.Execute(&context)

	if !reflect.DeepEqual(context, expected) {
		t.Errorf("Expected %v, got %v", expected, context)
	}
}

func TestPipeline_MultipleSquare(t *testing.T) {
	context := 3
	expected := 6561 // (3*3) - (9*9) - (81-81)

	pipeline := CreatePipeline[int]()
	pipeline.Next(Square{}).Next(Square{}).Next(Square{})
	pipeline.Execute(&context)

	if !reflect.DeepEqual(context, expected) {
		t.Errorf("Expected %v, got %v", expected, context)
	}
}

func TestPipeline_MixedAddSquare(t *testing.T) {
	context := 3
	expected := 2592 // (3+3) - (6*6) - (36*36) - (1296+1296)

	pipeline := CreatePipeline[int]()
	pipeline.Next(Add{}).Next(Square{}).Next(Square{}).Next(Add{})
	pipeline.Execute(&context)

	if !reflect.DeepEqual(context, expected) {
		t.Errorf("Expected %v, got %v", expected, context)
	}
}
