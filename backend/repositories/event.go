package repositories

import (
	"context"
	"fmt"

	"github.com/sahildhargave/ticket-project-v1/models"
	"gorm.io/gorm"
)

type EventRepository struct {
	db *gorm.DB
}

func (r *EventRepository) GetMany(ctx context.Context) ([]*models.Event, error) {
	events := []*models.Event{}

	res := r.db.Model(&models.Event{}).Find(&events)

	if res.Error != nil {
		return nil, fmt.Errorf("Something went wrong")
	}

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

func NewEventRepository(db *gorm.DB) models.EventRepository {
	return &EventRepository{
		db: db,
	}
}
