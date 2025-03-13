package task

import (
	"fmt"
	"time"

	"go.uber.org/multierr"
)

type Task struct {
	id          string
	title       string
	description string
	status      Status
	createdAt   time.Time
	updatedAt   time.Time
}

type TaskFactoryConfig struct {
	MinCharactersForDescription int
	MaxCharactersForDescription int
	MinCharactersForTitle       int
	MaxCharactersForTitle       int
	DefaultStatus               Status
}

func (f TaskFactoryConfig) Validate() error {
	var err error

	if f.MinCharactersForDescription < 25 {
		err = multierr.Append(
			err,
			fmt.Errorf(
				"MinCharactersForDescription should be at least %d",
				f.MinCharactersForDescription,
			),
		)
	}
	if f.MaxCharactersForDescription > 250 {
		err = multierr.Append(
			err,
			fmt.Errorf(
				"MaxCharactersForDescription should be at most %d",
				f.MaxCharactersForDescription,
			),
		)
	}
	if f.MinCharactersForTitle < 5 {
		err = multierr.Append(
			err,
			fmt.Errorf(
				"MinCharactersForTitle should be at least %d",
				f.MinCharactersForTitle,
			),
		)
	}
	if f.MaxCharactersForTitle > 75 {
		err = multierr.Append(
			err,
			fmt.Errorf(
				"MaxCharactersForTitle should be at most %d",
				f.MaxCharactersForTitle,
			),
		)
	}
	if validateErr := f.DefaultStatus.Validate(); validateErr != nil {
		err = multierr.Append(err, validateErr)
	}
	return err
}

type TaskFactory struct {
	fc TaskFactoryConfig
}

func NewFactory(taskFactoryConfig TaskFactoryConfig) (TaskFactory, error) {
	if err := taskFactoryConfig.Validate(); err != nil {
		return TaskFactory{}, fmt.Errorf("invalid config passed to the task factory: %w", err)
	}
	return TaskFactory{fc: taskFactoryConfig}, nil
}

func MustNewFactory(taskFactoryConfig TaskFactoryConfig) TaskFactory {
	factory, err := NewFactory(taskFactoryConfig)
	if err != nil {
		panic(err)
	}
	return factory
}

func (factory TaskFactory) Config() TaskFactoryConfig {
	return factory.fc
}

func (factory TaskFactory) IsZero() bool {
	return factory == TaskFactory{}
}

type InvalidTitleError struct {
	title         string
	factoryConfig TaskFactoryConfig
}

func (err InvalidTitleError) Error() string {
	switch length := len(err.title); {
	case length < err.factoryConfig.MinCharactersForTitle:
		return fmt.Sprintf("title is shorter than the minimum allowed length, %d", length)
	case length > err.factoryConfig.MaxCharactersForTitle:
		return fmt.Sprintf("title is longer than the maximum allowed length, %d", length)
	default:
		return "Unexpected title error"
	}
}

type InvalidDescriptionError struct {
	description   string
	factoryConfig TaskFactoryConfig
}

func (err InvalidDescriptionError) Error() string {
	switch length := len(err.description); {
	case length < err.factoryConfig.MinCharactersForDescription:
		return fmt.Sprintf("description is shorter than the minimum allowed length, %d", length)
	case length > err.factoryConfig.MaxCharactersForDescription:
		return fmt.Sprintf("description is longer than the maximum allowed length, %d", length)
	default:
		return "Unexpected description error"
	}
}

func (factory TaskFactory) ValidateTask(task Task) error {
	titleLength := len(task.title)
	descriptionLength := len(task.description)

	if titleLength < factory.fc.MinCharactersForTitle || titleLength > factory.fc.MaxCharactersForTitle {
		return InvalidTitleError{
			title:         task.title,
			factoryConfig: factory.Config(),
		}
	}
	if descriptionLength < factory.fc.MinCharactersForDescription || descriptionLength > factory.fc.MaxCharactersForDescription {
		return InvalidDescriptionError{
			description:   task.description,
			factoryConfig: factory.Config(),
		}
	}
	return nil
}

func (factory TaskFactory) NewTask(title, description string) *Task {
	return &Task{
		id:          "",
		title:       title,
		description: description,
		status:      Pending,
		createdAt:   time.Now(),
		updatedAt:   time.Now(),
	}
}

func (task Task) String() string {
	return fmt.Sprintf("{\n\tid: %s\n\ttitle: %s\n\tdescription: %s\n\tstatus: %v\n\tcreatedAt: %v\n\tupdatedAt: %v\n}", task.id, task.title, task.description, task.status, task.createdAt, task.updatedAt)
}
