package client

import (
	"context"

	"habits/api"
	"templates/internal/habit"
)

// CreateHabit calls the CreateHabit endpoint of the habits service.
func (hc *HabitsClient) CreateHabit(ctx context.Context, h habit.Habit) error {
	freq := int32(h.WeeklyFrequency)

	_, err := hc.cli.CreateHabit(ctx, &api.CreateHabitRequest{
		Name:            string(h.Name),
		WeeklyFrequency: &freq,
	})
	return err
}
