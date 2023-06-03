package crud

import (
	"context"
	"encoding/json"
	"github.com/goadify/goadify/modules/crud/helpers"
	"github.com/goadify/goadify/modules/crud/hydrator"
	"github.com/goadify/goadify/modules/crud/models"
	"github.com/goadify/openapi/crud/go/gen"
	"github.com/pkg/errors"
)

type httpHandler struct {
	isDevMode bool
	em        *entityMaster
}

var (
	ErrEntityNotRegistered              = errors.New("entity not registered")
	ErrRepositoryNotSupportsCreate      = errors.New("repository not supports create actions")
	ErrRepositoryNotSupportsRead        = errors.New("repository not supports read actions")
	ErrRepositoryNotSupportsUpdate      = errors.New("repository not supports update actions")
	ErrRepositoryNotSupportsDelete      = errors.New("repository not supports delete actions")
	ErrAccessDenied                     = errors.New("access denied")
	ErrIdentifierNotConvertableToString = errors.New("identifier not convertable to string")
	ErrUnexpectedConvertingRecord       = errors.New("something went wrong while converting record")
)

func (h *httpHandler) checkAccess(ctx context.Context, repository Repository) (bool, error) {
	accessRules, err := repository.AccessRules(ctx)
	if err != nil {
		return false, err
	}

	if !helpers.InSlice(accessRules, AccessReadRule) {
		return false, nil
	}

	return true, nil
}

func (h *httpHandler) convertRecord(record Record) (*models.IdentifiedRecord, error) {
	id, ok := helpers.ConvertToString(record.ID())

	if !ok {
		return nil, ErrIdentifierNotConvertableToString
	}

	return &models.IdentifiedRecord{
		ID:   id,
		Data: record,
	}, nil
}

func (h *httpHandler) convertRecords(records []Record) ([]models.IdentifiedRecord, error) {
	res := make([]models.IdentifiedRecord, len(records))
	for i := 0; i < len(records); i++ {
		model, err := h.convertRecord(records[i])
		if err != nil {
			return nil, err
		}

		if model == nil {
			return nil, ErrUnexpectedConvertingRecord
		}
		res[i] = *model
	}

	return res, nil
}

func (h *httpHandler) GetEntitiesMappings(_ context.Context) ([]gen.EntityMapping, error) {
	ems := h.em.EntityMappings()

	return hydrator.EntityMappings(ems), nil
}

func (h *httpHandler) CreateRecord(ctx context.Context, req *gen.Record, params gen.CreateRecordParams) (*gen.IdentifiedRecord, error) {
	repository, ok := h.em.Repository(params.Name)
	if !ok {
		return nil, errors.Wrap(ErrEntityNotRegistered, params.Name)
	}

	allowed, err := h.checkAccess(ctx, repository)
	if err != nil {
		return nil, err
	}

	if !allowed {
		return nil, ErrAccessDenied
	}

	creatableRepository, ok := repository.(RepositoryCreatable)
	if !ok {
		return nil, errors.Wrap(ErrRepositoryNotSupportsCreate, params.Name)
	}

	record := repository.NewRecord()

	err = json.Unmarshal(req.Data, &record)
	if err != nil {
		return nil, err
	}

	err = creatableRepository.Create(ctx, record)
	if err != nil {
		return nil, err
	}

	convertedRecord, err := h.convertRecord(record)
	if err != nil {
		return nil, err
	}

	result, err := hydrator.Record(*convertedRecord)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (h *httpHandler) GetRecordById(ctx context.Context, params gen.GetRecordByIdParams) (*gen.IdentifiedRecord, error) {
	repository, ok := h.em.Repository(params.Name)
	if !ok {
		return nil, errors.Wrap(ErrEntityNotRegistered, params.Name)
	}

	allowed, err := h.checkAccess(ctx, repository)
	if err != nil {
		return nil, err
	}

	if !allowed {
		return nil, ErrAccessDenied
	}

	readableRepos, ok := repository.(RepositoryReadable)
	if !ok {
		return nil, errors.Wrap(ErrRepositoryNotSupportsRead, params.Name)
	}

	record, err := readableRepos.GetByID(ctx, params.ID)
	if err != nil {
		return nil, err
	}

	convertedRecord, err := h.convertRecord(record)
	if err != nil {
		return nil, err
	}

	result, err := hydrator.Record(*convertedRecord)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (h *httpHandler) GetRecords(ctx context.Context, params gen.GetRecordsParams) (*gen.RecordsResponse, error) {
	repository, ok := h.em.Repository(params.Name)
	if !ok {
		return nil, errors.Wrap(ErrEntityNotRegistered, params.Name)
	}

	allowed, err := h.checkAccess(ctx, repository)
	if err != nil {
		return nil, err
	}

	if !allowed {
		return nil, ErrAccessDenied
	}

	readableRepos, ok := repository.(RepositoryReadable)
	if !ok {
		return nil, errors.Wrap(ErrRepositoryNotSupportsRead, params.Name)
	}

	records, totalCount, err := readableRepos.GetList(
		ctx,
		int32(params.Page.Or(1)),
		int32(params.Page.Or(20)),
	)

	if err != nil {
		return nil, err
	}

	convertedRecords, err := h.convertRecords(records)
	if err != nil {
		return nil, err
	}

	items, err := hydrator.Records(convertedRecords)
	if err != nil {
		return nil, err
	}

	return &gen.RecordsResponse{
		TotalCount: totalCount,
		Items:      items,
	}, nil
}

func (h *httpHandler) UpdateRecordById(ctx context.Context, req *gen.Record, params gen.UpdateRecordByIdParams) (*gen.IdentifiedRecord, error) {
	repository, ok := h.em.Repository(params.Name)
	if !ok {
		return nil, errors.Wrap(ErrEntityNotRegistered, params.Name)
	}

	allowed, err := h.checkAccess(ctx, repository)
	if err != nil {
		return nil, err
	}

	if !allowed {
		return nil, ErrAccessDenied
	}

	updatableRepos, ok := repository.(RepositoryUpdatable)
	if !ok {
		return nil, errors.Wrap(ErrRepositoryNotSupportsUpdate, params.Name)
	}

	record := repository.NewRecord()

	err = json.Unmarshal(req.Data, &record)
	if err != nil {
		return nil, err
	}

	err = updatableRepos.Update(ctx, record, params.ID)
	if err != nil {
		return nil, err
	}

	convertedRecord, err := h.convertRecord(record)
	if err != nil {
		return nil, err
	}

	result, err := hydrator.Record(*convertedRecord)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (h *httpHandler) DeleteRecordById(ctx context.Context, _ *gen.Record, params gen.DeleteRecordByIdParams) error {
	repository, ok := h.em.Repository(params.Name)
	if !ok {
		return errors.Wrap(ErrEntityNotRegistered, params.Name)
	}

	allowed, err := h.checkAccess(ctx, repository)
	if err != nil {
		return err
	}

	if !allowed {
		return ErrAccessDenied
	}

	deletableRepos, ok := repository.(RepositoryDeletable)
	if !ok {
		return errors.Wrap(ErrRepositoryNotSupportsDelete, params.Name)
	}

	err = deletableRepos.Delete(ctx, params.ID)
	if err != nil {
		return err
	}

	return nil
}

func (h *httpHandler) NewError(_ context.Context, err error) *gen.ErrorStatusCode {
	return hydrator.Error(err, hydrator.ErrorDisplayInternalMessages(h.isDevMode))
}

func newHttpHandler(entityMaster *entityMaster, isDevMode bool) *httpHandler {
	return &httpHandler{
		em:        entityMaster,
		isDevMode: isDevMode,
	}
}
