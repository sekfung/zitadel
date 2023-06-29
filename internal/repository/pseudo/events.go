package pseudo

import (
	"context"
	"time"

	"github.com/zitadel/zitadel/internal/eventstore"
)

const (
	eventTypePrefix    = eventstore.EventType("pseudo.")
	ScheduledEventType = eventTypePrefix + "timestamp"
)

var _ eventstore.Event = (*ScheduledEvent)(nil)

type ScheduledEvent struct {
	*eventstore.BaseEvent `json:"-"`
	// TODO: `json:"-"`
	Timestamp time.Time `json:"timestamp"`
	// TODO: `json:"-"`
	InstanceIDs []string `json:"instanceIDs"`
	// TODO: `json:"-"`
	TriggeringEvent eventstore.Event `json:"triggeringEvent"`
}

func NewScheduledEvent(
	ctx context.Context,
	timestamp time.Time,
	triggeringEvent eventstore.Event,
	instanceIDs ...string,
) *ScheduledEvent {
	return &ScheduledEvent{
		BaseEvent: eventstore.NewBaseEventForPush(
			ctx,
			&NewAggregate().Aggregate,
			ScheduledEventType,
		),
		Timestamp:       timestamp,
		InstanceIDs:     instanceIDs,
		TriggeringEvent: triggeringEvent,
	}
}
