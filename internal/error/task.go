package apperrors

import "errors"

var Task = struct {
	InvalidGoalID                 error
	GoalNotFound                  error
	GoalDeleteErrorTasksRemaining error

	InvalidTaskID error
	TaskNotFound  error
}{
	InvalidGoalID:                 errors.New("task_invalid_goal_id"),
	GoalNotFound:                  errors.New("task_goal_not_found"),
	GoalDeleteErrorTasksRemaining: errors.New("task_goal_delete_error_tasks_remaining"),

	InvalidTaskID: errors.New("task_invalid_task_id"),
	TaskNotFound:  errors.New("task_not_found"),
}
