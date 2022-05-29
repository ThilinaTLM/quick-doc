package schema

import (
	"reflect"
	"testing"
)

type UserAccount1 struct {
	Name string   `json:"name"`
	Age  int      `json:"age,omitempty"`
	Log  LogEntry `json:"log"`
}

type UserAccount2 struct {
	Name string    `json:"name"`
	Age  int       `json:"age,omitempty"`
	Log  *LogEntry `json:"log"`
}

type UserAccount3 struct {
	Name string      `json:"name"`
	Age  int         `json:"age,omitempty"`
	Log  interface{} `json:"log"`
}

type LogEntry struct {
	Date  string `json:"date"`
	Valid bool   `json:"valid"`
}

func Test_TypeString(t *testing.T) {
	sb := NewBuilderDefault()

	got, err := sb.GetSchema("Test User")

	want := &Property{
		Type:  PropType_STRING,
		Value: "Test User",
	}

	if err != nil {
		t.Errorf("error while generating schema, %e", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("not match got=%v; want=%v", got, want)
	}
}

func Test_TypeStringPointer(t *testing.T) {
	sb := NewBuilderDefault()

	str := "Test User"
	got, err := sb.GetSchema(&str)

	want := &Property{
		Type:  PropType_STRING,
		Value: "Test User",
	}

	if err != nil {
		t.Errorf("error while generating schema, %e", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("not match got=%v; want=%v", got, want)
	}
}

func Test_TypeBoolean(t *testing.T) {
	sb := NewBuilderDefault()

	got, err := sb.GetSchema(true)

	want := &Property{
		Type:  PropType_BOOLEAN,
		Value: "true",
	}

	if err != nil {
		t.Errorf("error while generating schema, %e", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("not match got=%v; want=%v", got, want)
	}
}

func Test_TypeInteger(t *testing.T) {
	sb := NewBuilderDefault()

	got, err := sb.GetSchema(10)

	want := &Property{
		Type:  PropType_INTEGER,
		Value: "10",
	}

	if err != nil {
		t.Errorf("error while generating schema, %e", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("not match got=%v; want=%v", got, want)
	}
}

func Test_TypeStruct_1(t *testing.T) {
	sb := NewBuilderDefault()

	got, err := sb.GetSchema(UserAccount1{
		Name: "Test User",
		Age:  22,
		Log: LogEntry{
			Date:  "2022-01-21",
			Valid: false,
		},
	})

	want := &Property{
		Type: PropType_OBJECT,
		Properties: []Property{
			{
				Type:  PropType_STRING,
				Name:  "name",
				Value: "Test User",
			},
			{
				Type:  PropType_INTEGER,
				Name:  "age",
				Value: "22",
			},
			{
				Type: PropType_OBJECT,
				Name: "log",
				Properties: []Property{
					{
						Type:  PropType_STRING,
						Name:  "date",
						Value: "2022-01-21",
					},
					{
						Type:  PropType_BOOLEAN,
						Name:  "valid",
						Value: "false",
					},
				},
			},
		},
	}

	if err != nil {
		t.Errorf("error while generating schema, %e", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("not match got=%v; want=%v", got, want)
	}
}

func Test_TypeStruct_2(t *testing.T) {
	sb := NewBuilderDefault()

	got, err := sb.GetSchema(UserAccount2{
		Name: "Test User",
		Age:  22,
		Log:  nil,
	})

	want := &Property{
		Type: PropType_OBJECT,
		Properties: []Property{
			{
				Type:  PropType_STRING,
				Name:  "name",
				Value: "Test User",
			},
			{
				Type:  PropType_INTEGER,
				Name:  "age",
				Value: "22",
			},
			{
				Type: PropType_OBJECT, // nil object
				Name: "log",
				Properties: []Property{
					{
						Type:  PropType_STRING,
						Name:  "date",
						Value: "nil",
					},
					{
						Type:  PropType_BOOLEAN,
						Name:  "valid",
						Value: "nil",
					},
				},
			},
		},
	}

	if err != nil {
		t.Errorf("error while generating schema, %e", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("not match \ngot =%v\nwant=%v", got, want)
	}
}

func Test_TypeStruct_3(t *testing.T) {
	sb := NewBuilderDefault()

	got, err := sb.GetSchema(UserAccount3{
		Name: "Test User",
		Age:  22,
		Log: &LogEntry{
			Date:  "2022-02-01",
			Valid: false,
		},
	})

	want := &Property{
		Type: PropType_OBJECT,
		Properties: []Property{
			{
				Type:  PropType_STRING,
				Name:  "name",
				Value: "Test User",
			},
			{
				Type:  PropType_INTEGER,
				Name:  "age",
				Value: "22",
			},
			{
				Type: PropType_OBJECT, // nil object
				Name: "log",
				Properties: []Property{
					{
						Type:  PropType_STRING,
						Name:  "date",
						Value: "2022-02-01",
					},
					{
						Type:  PropType_BOOLEAN,
						Name:  "valid",
						Value: "false",
					},
				},
			},
		},
	}

	if err != nil {
		t.Errorf("error while generating schema, %e", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("not match \ngot =%v\nwant=%v", got, want)
	}
}

func Test_TypeStruct_4(t *testing.T) {
	sb := NewBuilderDefault()

	got, err := sb.GetSchema(UserAccount3{
		Name: "Test User",
		Age:  22,
		Log:  nil,
	})

	want := &Property{
		Type: PropType_OBJECT,
		Properties: []Property{
			{
				Type:  PropType_STRING,
				Name:  "name",
				Value: "Test User",
			},
			{
				Type:  PropType_INTEGER,
				Name:  "age",
				Value: "22",
			},
		},
	}

	if err != nil {
		t.Errorf("error while generating schema, %e", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("not match \ngot =%v\nwant=%v", got, want)
	}
}

func Test_TypeSlice(t *testing.T) {
	sb := NewBuilderDefault()

	got, err := sb.GetSchema([]int{1, 2, 3})

	want := &Property{
		Type: PropType_ARRAY,
		Properties: []Property{
			{
				Type:  PropType_INTEGER,
				Value: "1",
			},
			{
				Type:  PropType_INTEGER,
				Value: "2",
			},
			{
				Type:  PropType_INTEGER,
				Value: "3",
			},
		},
	}

	if err != nil {
		t.Errorf("error while generating schema, %e", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("not match \ngot =%v\nwant=%v", got, want)
	}
}
