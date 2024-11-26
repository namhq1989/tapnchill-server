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

const (
	FreePlanMaxHabits int64 = 5
	ProPlanMaxHabits  int64 = 20

	FreePlanMaxGoals int64 = 5
	ProPlanMaxGoals  int64 = 20

	FreePlanMaxTaskPerGoal int64 = 20
	ProPlanMaxTaskPerGoal  int64 = 50

	FreePlanMaxNotes int64 = 50
	ProPlanMaxNotes  int64 = 10000
)
