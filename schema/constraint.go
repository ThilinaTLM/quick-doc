package schema

// ConstraintType Constraint types
type ConstraintType string

const (
	ConType_REQUIRED ConstraintType = "REQUIRED"
	ConType_MAX      ConstraintType = "MAX"
	ConType_MIN      ConstraintType = "MIN"
	ConType_BETWEEN  ConstraintType = "BETWEEN"
)

type Constraint struct {
	Type ConstraintType `json:"type"`
	Min  int            `json:"min,omitempty"`
	Max  int            `json:"max,omitempty"`
}
