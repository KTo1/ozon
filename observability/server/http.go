package server

import (
	"errors"
	"fmt"
	"github.com/KTo1/ozon/observability/model"
	"github.com/KTo1/ozon/observability/storage"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"time"
)

type FiberHandler struct {
	storage storage.NoteStorage
	tracer  trace.Tracer // для своих спанов
}

func NewFiberHandler(storage storage.NoteStorage, tracer trace.Tracer) FiberHandler {
	return FiberHandler{storage: storage, tracer: tracer}
}

func (f *FiberHandler) CreateNote(fiberctx *fiber.Ctx) error {
	ctx, span := f.tracer.Start(fiberctx.UserContext(), "My Really Create Note Span", trace.WithAttributes(attribute.String("My key", "my value")))
	defer span.End()

	input := struct {
		Title   string `json:"title"`
		Content string `json:"content"`
	}{}

	if err := fiberctx.BodyParser(&input); err != nil {
		return fmt.Errorf("body parsing error %w", err)
	}

	noteID := uuid.New()
	err := f.storage.Store(ctx, model.Note{
		ID:        noteID,
		Title:     input.Title,
		Content:   input.Content,
		CreatedAt: time.Now(),
	})

	if err != nil {
		time.Sleep(1 * time.Second)

		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return fiberctx.JSON(
		map[string]any{
			"note_id": noteID,
		})
}

func (f *FiberHandler) GetNote(fiberctx *fiber.Ctx) error {
	ctx, span := f.tracer.Start(fiberctx.UserContext(), "My Really Get Note Span")
	defer span.End()

	noteID, err := uuid.Parse(fiberctx.Query("note_id"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	// Зачастую наши функции делают несколько вызовов других. Для измерения времени каждой есть возможность разделить спан на события с помощью span.AddEvent(). В параметрах можно также задать атрибуты:
	span.AddEvent("call redis stogare", trace.WithAttributes(attribute.String("noteID", noteID.String())))

	note, err := f.storage.Get(ctx, noteID)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return fiber.NewError(fiber.StatusNotFound, err.Error())
		}

		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return fiberctx.JSON(note)
}
