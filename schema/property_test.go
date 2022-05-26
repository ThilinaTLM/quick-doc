package schema

import "testing"
import "reflect"

func Test_Schema(t *testing.T) {
	got := FromObject(struct {
		Name string `json:"name" qd:"Person name"`
	}{
		Name: "test",
	}, Options{})
	want := []Property{
		{
			Name:        "name",
			Type:        PropType_STRING,
			Value:       "test",
			Description: "Person name",
		},
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
