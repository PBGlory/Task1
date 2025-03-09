package taskService

type TaskService struct {
	repo TaskRepository
}

func NewTaskService(repo TaskRepository) *TaskService {
	return &TaskService{repo: repo}
}

func (service *TaskService) CreateTask(task Task) (Task, error) {
	return service.repo.CreateTask(task)
}

func (service *TaskService) GetAllTasks() ([]Task, error) {
	return service.repo.GetAllTasks()
}

func (service *TaskService) UpdateTaskByID(id uint, task Task) (Task, error) {
	return service.repo.UpdateTaskById(id, task)
}

func (service *TaskService) DeleteTaskById(id uint) error {
	return service.repo.DeleteTaskById(id)
}
