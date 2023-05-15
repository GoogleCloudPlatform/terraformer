package terraformutils

import (
	"reflect"
	"testing"
)

func TestEmptyWalkAndGet(t *testing.T) {
	structure := map[string]interface{}{}
	value := WalkAndGet("attr1", structure)

	if !reflect.DeepEqual(value, []interface{}{}) {
		t.Errorf("failed to get value %v", value)
	}
}

func TestEmptyNestedWalkAndGet(t *testing.T) {
	structure := map[string]map[string]interface{}{}
	value := WalkAndGet("attr1.attr2", structure)

	if !reflect.DeepEqual(value, []interface{}{}) {
		t.Errorf("failed to get value %v", value)
	}
}

func TestSimpleWalkAndGet(t *testing.T) {
	structure := map[string]interface{}{
		"attr1": "value",
	}
	value := WalkAndGet("attr1", structure)

	if !reflect.DeepEqual(value, []interface{}{"value"}) {
		t.Errorf("failed to get value %v", value)
	}
}

func TestSimpleArrayWalkAndGet(t *testing.T) {
	structure := map[string][]interface{}{
		"attr1": {"value"},
	}
	value := WalkAndGet("attr1", structure)

	if !reflect.DeepEqual(value, []interface{}{"value"}) {
		t.Errorf("failed to get value %v", value)
	}
}

func TestNestedWalkAndGet(t *testing.T) {
	structure := map[string]map[string]interface{}{
		"attr1": {
			"attr2": "value",
		},
	}
	value := WalkAndGet("attr1.attr2", structure)

	if !reflect.DeepEqual(value, []interface{}{"value"}) {
		t.Errorf("failed to get value %v", value)
	}
}

func TestNestedWalkWithDotInKeyAndGet(t *testing.T) {
	structure := map[string]map[string]interface{}{
		"attr1": {
			"attr2.attr3": "value",
		},
	}
	value := WalkAndGet("attr1.attr2.attr3", structure)

	if !reflect.DeepEqual(value, []interface{}{"value"}) {
		t.Errorf("failed to get value %v", value)
	}
}

func TestNestedArrayWalkAndGet(t *testing.T) {
	structure := mapI("attr1", []interface{}{
		mapI("attr2", "value1"),
		mapI("attr2", "value2")})
	value := WalkAndGet("attr1.attr2", structure)

	if !reflect.DeepEqual(value, []interface{}{"value1", "value2"}) {
		t.Errorf("failed to get value %v", value)
	}
}

func TestNonExistingWalkAndGet(t *testing.T) {
	structure := map[string]interface{}{
		"attr1": "test",
	}
	value := WalkAndGet("attr1.attr2", structure)

	if !reflect.DeepEqual(value, []interface{}{}) {
		t.Errorf("failed to get value %v", value)
	}
}

func TestSimpleWalkAndOverride(t *testing.T) {
	structure := map[string]interface{}{
		"attr1": "value",
	}
	WalkAndOverride("attr1", "value", "newValue", structure)

	if structure["attr1"] != "newValue" {
		t.Errorf("failed to set value")
	}
}

func TestSimpleArrayWalkAndOverride(t *testing.T) {
	structure := map[string][]interface{}{
		"attr1": {"value"},
	}
	WalkAndOverride("attr1", "value", "newValue", structure)

	if structure["attr1"][0] != "newValue" {
		t.Errorf("failed to set value")
	}
}

func TestSimpleWalkAndNotOverride(t *testing.T) {
	structure := map[string]interface{}{
		"attr1": "value",
	}
	WalkAndOverride("attr1", "differentValue", "newValue", structure)

	if structure["attr1"] != "value" {
		t.Errorf("failed to set value")
	}
}

func TestNonExistentWalkAndOverride(t *testing.T) {
	structure := map[string]interface{}{
		"attr1": "value",
	}
	WalkAndOverride("attr1.nonExistentAttr", "value", "newValue", structure)

	_, exists := structure["nonExistentAttr"]
	if exists {
		t.Errorf("failed to set value")
	}
}

func TestNestedWalkAndOverride(t *testing.T) {
	structure := map[string]map[string]interface{}{
		"attr1": {
			"attr2": "value",
		},
	}
	WalkAndOverride("attr1.attr2", "value", "newValue", structure)

	if structure["attr1"]["attr2"] != "newValue" {
		t.Errorf("failed to set value")
	}
}

func TestNestedArrayWalkAndOverride(t *testing.T) {
	structure := mapI("attr1", []interface{}{
		mapI("attr2", "value1"),
		mapI("attr2", "value2")})
	WalkAndOverride("attr1.attr2", "value2", "newValue", structure)

	if structure["attr1"].([]interface{})[0].(map[string]interface{})["attr2"] != "value1" || structure["attr1"].([]interface{})[1].(map[string]interface{})["attr2"] != "newValue" {
		t.Errorf("failed to set value")
	}
}

func TestNestedMapWalkAndOverride(t *testing.T) {
	structure := mapI("x", []interface{}{
		mapI("y", mapI("z", "42")),
	})
	WalkAndOverride("z.y", "z", "newValue", structure)

	expected := mapI("x", []interface{}{
		mapI("y", mapI("z", "42")),
	})
	if !reflect.DeepEqual(structure, expected) {
		t.Errorf("failed to set value")
	}
}

func TestEmptyWalkAndCheckField(t *testing.T) {
	structure := map[string]interface{}{}
	value := WalkAndCheckField("attr1", structure)

	if !reflect.DeepEqual(value, false) {
		t.Errorf("failed to get value %v", value)
	}
}

func TestSimpleWalkAndCheckField(t *testing.T) {
	structure := map[string]interface{}{
		"attr1": "value",
	}
	value := WalkAndCheckField("attr1", structure)

	if !reflect.DeepEqual(value, true) {
		t.Errorf("failed to get value %v", value)
	}
}
