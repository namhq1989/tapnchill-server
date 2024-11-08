package domain

type TaskStatus string

const (
	TaskStatusUnknown TaskStatus = ""
	TaskStatusTodo    TaskStatus = "todo"
	TaskStatusDone    TaskStatus = "done"
)

func (t TaskStatus) IsValid() bool {
	return t != TaskStatusUnknown
}

func (t TaskStatus) String() string {
	return string(t)
}

func ToTaskStatus(value string) TaskStatus {
	switch value {
	case TaskStatusTodo.String():
		return TaskStatusTodo
	case TaskStatusDone.String():
		return TaskStatusDone
	default:
		return TaskStatusUnknown
	}
}
