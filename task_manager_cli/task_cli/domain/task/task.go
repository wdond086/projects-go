package task

import (
	"fmt"
	"time"

	"go.uber.org/multierr"
)

type Task struct {
	id          string
	description string
	status      Status
	createdAt   time.Time
	updatedAt   time.Time
}

type TaskFactoryConfig struct {
	MinCharactersForDescription int
	MaxCharactersForDescription int
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
