package shared

import "github.com/namhq1989/tapnchill-server/pkg/task/domain"

type Service struct {
	taskRepository domain.TaskRepository
	goalRepository domain.GoalRepository
}

func NewService(taskRepository domain.TaskRepository, goalRepository domain.GoalRepository) Service {
	return Service{
		taskRepository: taskRepository,
		goalRepository: goalRepository,
	}
}
