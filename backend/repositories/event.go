package repositories

import (
	"context"
	"time"

	"github.com/sahildhargave/ticket-project-v1/models"
)

type EventRepository struct {
	db any
}

func (r *EventRepository) GetMany(ctx context.Context) ([]*models.Event, error) {
	events := []*models.Event{}

	events = append(events, &models.Event{
		ID:        "123424324343",
		Name:      "Example Event",
		Location:  "India,ls",
		Date:      time.Now(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	return events, nil
}

func (r *EventRepository) GetOne(ctx context.Context, eventId uint) (*models.Event, error) {
	return nil, nil
}

func (r *EventRepository) CreateOne(ctx context.Context, event *models.Event) (*models.Event, error) { // Changed to *models.Event
	// Your implementation here
	return nil, nil
}

// DeleteOne implements models.EventRepository.
func (r *EventRepository) DeleteOne(ctx context.Context, eventId uint) error {
	panic("unimplemented")
}

// UpdateOne implements models.EventRepository.
func (r *EventRepository) UpdateOne(ctx context.Context, eventId uint, updateData map[string]interface{}) (*models.Event, error) {
	panic("unimplemented")
}

func NewEventRepository(db any) models.EventRepository {
	return &EventRepository{
		db: db,
	}
}
