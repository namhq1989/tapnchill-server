package domain

type Plan string

const (
	PlanUnknown Plan = ""
	PlanFree    Plan = "free"
	PlanPro     Plan = "pro"
)

func (p Plan) IsValid() bool {
	return p != PlanUnknown
}

func (p Plan) String() string {
	return string(p)
}

func ToPlan(value string) Plan {
	switch value {
	case PlanFree.String():
		return PlanFree
	case PlanPro.String():
		return PlanPro
	default:
		return PlanUnknown
	}
}
