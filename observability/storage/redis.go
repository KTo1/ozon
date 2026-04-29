package storage

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/KTo1/ozon/observability/model"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

type NoteStorage struct {
	client redis.UniversalClient
}

func NewNoteStorage(client redis.UniversalClient) NoteStorage {
	return NoteStorage{client: client}
}

func (n *NoteStorage) Store(ctx context.Context, note model.Note) error {
	ctx, span := otel.Tracer("redis").Start(ctx, "redis.Set")
	defer span.End()

	data, err := json.Marshal(note)
	if err != nil {
		return fmt.Errorf("could not marshal note: %w", err)
	}

	span2 := trace.SpanFromContext(ctx)
	fmt.Printf("Redis before Store - TraceID: %s, SpanID: %s\n",
		span2.SpanContext().TraceID(),
		span2.SpanContext().SpanID())

	if err := n.client.Set(ctx, note.ID.String(), data, -1).Err(); err != nil {
		span2 := trace.SpanFromContext(ctx)
		fmt.Printf("Redis after Store - TraceID: %s, SpanID: %s\n",
			span2.SpanContext().TraceID(),
			span2.SpanContext().SpanID())

		return fmt.Errorf("could not store note: %w", err)
	}

	return nil
}

func (n *NoteStorage) Get(ctx context.Context, ID uuid.UUID) (*model.Note, error) {
	data, err := n.client.Get(ctx, ID.String()).Bytes()
	if err != nil {
		if err == redis.Nil {
			return nil, model.ErrNotFound
		}

		return nil, fmt.Errorf("could not get note: %w", err)
	}

	var note model.Note
	if err := json.Unmarshal(data, &note); err != nil {
		return nil, fmt.Errorf("could not unmarshal note: %w", err)
	}

	return &note, nil
}
