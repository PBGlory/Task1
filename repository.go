package taskService

import "gorm.io/gorm"

type TaskRepository interface {
	CreateTask(task Task) (Task, error)

	GetAllTasks() ([]Task, error)

	UpdateTaskById(id uint, task Task) (Task, error)

	DeleteTaskById(id uint) error
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *taskRepository {
	return &taskRepository{db: db}
}

func (repo *taskRepository) CreateTask(task Task) (Task, error) {
	result := repo.db.Create(&task)
	if result.Error != nil {
		return Task{}, result.Error
	}
	return task, nil
}

func (repo *taskRepository) GetAllTasks() ([]Task, error) {
	var tasks []Task
	err := repo.db.Find(&tasks).Error
	return tasks, err
}

func (repo *taskRepository) UpdateTaskById(id uint, task Task) (Task, error) {
	var existingTask Task

	if err := repo.db.First(&existingTask, id).Error; err != nil {
		return Task{}, err
	}

	result := repo.db.Model(&existingTask).Updates(task)
	if result.Error != nil {
		return Task{}, result.Error
	}
	return existingTask, nil
}

func (repo *taskRepository) DeleteTaskById(id uint) error {
	result := repo.db.Where("id = ?", id).Delete(&Task{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}
