package main

import "fmt"

type Step interface {
	Execute(Runtime, TaskReporter) error
	Description() string
}

type StopStep struct {
	services []string
}

func NewStopStep(config StepConfig) Step {
	return &StopStep{
		services: config.Stop,
	}
}

func (s *StopStep) Execute(runtime Runtime, reporter TaskReporter) error {
	return runtime.StopServices(s.services...)
}

func (s *StopStep) Description() string {
	return fmt.Sprintf("<Stop: %v>", s.services)
}

type StartStep struct {
	services []string
}

func NewStartStep(config StepConfig) Step {
	return &StartStep{
		services: config.Start,
	}
}

func (s *StartStep) Execute(runtime Runtime, reporter TaskReporter) error {
	return runtime.StartServices(s.services...)
}

func (s *StartStep) Description() string {
	return fmt.Sprintf("<Start: %v>", s.services)
}

type AssertStep struct {
}

func NewAssertStep(name string) Step {
	return &AssertStep{}
}

func (s *AssertStep) Execute(runtime Runtime, reporter TaskReporter) error {
	return nil
}

func (s *AssertStep) Description() string {
	return fmt.Sprintf("<Assert:>")
}