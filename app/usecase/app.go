package usecase

import (
	"context"
	"errors"
	"fmt"
	"kms/app/entity"
	"kms/app/errs"
	"kms/app/secure"
	"kms/app/service"
	"kms/internal/database/sqldb"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
)

// appAudit is the combination of a domain App and its audit data
type appAudit struct {
	App         *entity.App
	SimpleAudit *entity.SimpleAudit
}

// newAPIKeyResponse initializes an APIKeyResponse. The app.APIKey is
// decrypted and set to the Key field as part of initialization.
func newAPIKeyResponse(key entity.APIKey) service.APIKeyResponse {
	return service.APIKeyResponse{Key: key.Key(), DeactivationDate: key.DeactivationDate().String()}
}

// newAppResponse initializes an AppResponse
func newAppResponse(aa appAudit) *service.AppResponse {
	var keys []service.APIKeyResponse
	for _, key := range aa.App.APIKeys {
		akr := newAPIKeyResponse(key)
		keys = append(keys, akr)
	}
	return &service.AppResponse{
		ExternalID:          aa.App.ExternalID.String(),
		Name:                aa.App.Name,
		Description:         aa.App.Description,
		CreateAppExtlID:     aa.SimpleAudit.Create.App.ExternalID.String(),
		CreateUserFirstName: aa.SimpleAudit.Create.User.FirstName,
		CreateUserLastName:  aa.SimpleAudit.Create.User.LastName,
		CreateDateTime:      aa.SimpleAudit.Create.Moment.Format(time.RFC3339),
		UpdateAppExtlID:     aa.SimpleAudit.Update.App.ExternalID.String(),
		UpdateUserFirstName: aa.SimpleAudit.Update.User.FirstName,
		UpdateUserLastName:  aa.SimpleAudit.Update.User.LastName,
		UpdateDateTime:      aa.SimpleAudit.Update.Moment.Format(time.RFC3339),
		APIKeys:             keys,
	}
}

// AppUsecase is a service for creating an App
type AppUsecase struct {
	Datastorer      sqldb.Datastorer
	APIKeyGenerator entity.APIKeyGenerator
	EncryptionKey   *[32]byte
}

// Create is used to create an App
func (s *AppUsecase) Create(ctx context.Context, r *service.CreateAppRequest, adt entity.Audit) (ar *service.AppResponse, err error) {
	const op errs.Op = "service/AppService.Create"

	var (
		a  *entity.App
		aa appAudit
	)
	nap := newAppParams{
		Name:        r.Name,
		Description: r.Description,
		// when creating an app, the org the app belongs to must be
		// the same as the org which the user is transacting.
		Org:             adt.App.Org,
		ApiKeyGenerator: s.APIKeyGenerator,
		EncryptionKey:   s.EncryptionKey,
	}
	a, err = newApp(nap)
	if err != nil {
		return nil, errs.E(op, err)
	}
	aa = appAudit{
		App: a,
		SimpleAudit: &entity.SimpleAudit{
			Create: adt,
			Update: adt,
		},
	}

	// start db txn using pgxpool
	var tx pgx.Tx
	tx, err = s.Datastorer.BeginTx(ctx)
	if err != nil {
		return nil, errs.E(op, err)
	}
	// defer transaction rollback and handle error, if any
	defer func() {
		err = s.Datastorer.RollbackTx(ctx, tx, err)
	}()

	err = createAppTx(ctx, tx, aa)
	if err != nil {
		return nil, errs.E(op, err)
	}

	// commit db txn using pgxpool
	err = s.Datastorer.CommitTx(ctx, tx)
	if err != nil {
		return nil, errs.E(op, err)
	}

	return newAppResponse(appAudit{App: a, SimpleAudit: &entity.SimpleAudit{Create: adt, Update: adt}}), nil
}

type newAppParams struct {
	// name: app name
	Name string
	// description: app description
	Description string
	// org: the org the app belongs to
	Org *entity.Org
	// apiKeyGenerator: random string generator used to create API key for app
	ApiKeyGenerator entity.APIKeyGenerator
	// encryptionKey: encryption key used to encrypt the generated API key
	EncryptionKey *[32]byte
	// Provider is the OAuth2 provider
	Provider entity.Provider
	// ProviderClientID is the unique Client ID given by the Provider
	// which represents an application
	ProviderClientID string
}

// newApp initializes an App with a single API Key
func newApp(nap newAppParams) (a *entity.App, err error) {
	const op errs.Op = "service/newApp"

	a = &entity.App{
		ID:               uuid.New(),
		ExternalID:       secure.NewID(),
		Org:              nap.Org,
		Name:             nap.Name,
		Description:      nap.Description,
		Provider:         nap.Provider,
		ProviderClientID: nap.ProviderClientID,
	}

	// create new API key
	keyDeactivation := time.Date(2099, 12, 31, 0, 0, 0, 0, time.UTC)
	var key entity.APIKey
	key, err = entity.NewAPIKey(nap.ApiKeyGenerator, nap.EncryptionKey, keyDeactivation)
	if err != nil {
		return nil, errs.E(op, err)
	}

	// add API key to app
	err = a.AddKey(key)
	if err != nil {
		return nil, errs.E(op, err)
	}

	return a, nil
}

// createAppTx creates the app in the database using a pgx.Tx. This is moved out of the
// app create handler function as it's also used when creating an org.
func createAppTx(ctx context.Context, tx pgx.Tx, aa appAudit) (err error) {
	const op errs.Op = "service/createAppTx"

	createAppParams := sqldb.CreateAppParams{
		AppID:                aa.App.ID,
		OrgID:                aa.App.Org.ID,
		AppExtlID:            aa.App.ExternalID.String(),
		AppName:              aa.App.Name,
		AppDescription:       aa.App.Description,
		AuthProviderID:       sqldb.NewNullInt32(int32(aa.App.Provider)),
		AuthProviderClientID: sqldb.NewNullString(aa.App.ProviderClientID),
		CreateAppID:          aa.SimpleAudit.Create.App.ID,
		CreateUserID:         aa.SimpleAudit.Create.User.NullUUID(),
		CreateTimestamp:      aa.SimpleAudit.Create.Moment,
		UpdateAppID:          aa.SimpleAudit.Update.App.ID,
		UpdateUserID:         aa.SimpleAudit.Update.User.NullUUID(),
		UpdateTimestamp:      aa.SimpleAudit.Update.Moment,
	}

	// create app database record using appstore
	var rowsAffected int64
	rowsAffected, err = datastore.New(tx).CreateApp(ctx, createAppParams)
	if err != nil {
		return errs.E(op, errs.Database, err)
	}

	if rowsAffected != 1 {
		return errs.E(op, errs.Database, fmt.Sprintf("rows affected should be 1, actual: %d", rowsAffected))
	}

	for _, key := range aa.App.APIKeys {

		createAppAPIKeyParams := datastore.CreateAppAPIKeyParams{
			ApiKey:          key.Ciphertext(),
			AppID:           aa.App.ID,
			DeactvDate:      key.DeactivationDate(),
			CreateAppID:     aa.SimpleAudit.Create.App.ID,
			CreateUserID:    aa.SimpleAudit.Create.User.NullUUID(),
			CreateTimestamp: aa.SimpleAudit.Create.Moment,
			UpdateAppID:     aa.SimpleAudit.Update.App.ID,
			UpdateUserID:    aa.SimpleAudit.Update.User.NullUUID(),
			UpdateTimestamp: aa.SimpleAudit.Update.Moment,
		}

		// create app API key database record using appstore
		var apiKeyRowsAffected int64
		apiKeyRowsAffected, err = datastore.New(tx).CreateAppAPIKey(ctx, createAppAPIKeyParams)
		if err != nil {
			return errs.E(op, errs.Database, err)
		}

		if apiKeyRowsAffected != 1 {
			return errs.E(op, errs.Database, fmt.Sprintf("rows affected should be 1, actual: %d", apiKeyRowsAffected))
		}
	}

	return nil
}

// Update is used to update an App. API Keys for an App cannot be updated.
func (s *AppUsecase) Update(ctx context.Context, r *entity.UpdateAppRequest, adt entity.Audit) (ar *entity.AppResponse, err error) {
	const op errs.Op = "service/AppService.Update"

	// start db txn using pgxpool
	var tx pgx.Tx
	tx, err = s.Datastorer.BeginTx(ctx)
	if err != nil {
		return nil, errs.E(op, err)
	}
	// defer transaction rollback and handle error, if any
	defer func() {
		err = s.Datastorer.RollbackTx(ctx, tx, err)
	}()

	// retrieve existing Org
	var aa appAudit
	aa, err = findAppByExternalIDWithAudit(ctx, tx, r.ExternalID)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, errs.E(op, errs.Validation, "No app exists for the given external ID")
		}
		return nil, errs.E(op, errs.Database, err)
	}
	// overwrite Update audit with the current audit
	aa.SimpleAudit.Update = adt

	// override fields with data from request
	aa.App.Name = r.Name
	aa.App.Description = r.Description

	updateAppParams := datastore.UpdateAppParams{
		AppName:         aa.App.Name,
		AppDescription:  aa.App.Description,
		UpdateAppID:     adt.App.ID,
		UpdateUserID:    adt.User.NullUUID(),
		UpdateTimestamp: adt.Moment,
		AppID:           aa.App.ID,
	}

	var rowsAffected int64
	rowsAffected, err = datastore.New(tx).UpdateApp(ctx, updateAppParams)
	if err != nil {
		return nil, errs.E(op, errs.Database, err)
	}

	if rowsAffected != 1 {
		return nil, errs.E(op, errs.Database, fmt.Sprintf("rows affected should be 1, actual: %d", rowsAffected))
	}

	// commit db txn using pgxpool
	err = s.Datastorer.CommitTx(ctx, tx)
	if err != nil {
		return nil, err
	}

	return newAppResponse(aa), nil
}

// Delete is used to delete an App
func (s *AppUsecase) Delete(ctx context.Context, extlID string) (dr entity.DeleteResponse, err error) {
	const op errs.Op = "service/AppService.Delete"

	// start db txn using pgxpool
	var tx pgx.Tx
	tx, err = s.Datastorer.BeginTx(ctx)
	if err != nil {
		return entity.DeleteResponse{}, errs.E(op, err)
	}
	// defer transaction rollback and handle error, if any
	defer func() {
		err = s.Datastorer.RollbackTx(ctx, tx, err)
	}()

	// retrieve existing App
	var a entity.App
	a, err = findAppByExternalID(ctx, tx, extlID)
	if err != nil {
		if err == pgx.ErrNoRows {
			return entity.DeleteResponse{}, errs.E(op, errs.Validation, "No app exists for the given external ID")
		}
		return entity.DeleteResponse{}, errs.E(op, errs.Database, err)
	}

	err = deleteAppTx(ctx, tx, a)
	if err != nil {
		return entity.DeleteResponse{}, errs.E(op, err)
	}

	// commit db txn using pgxpool
	err = s.Datastorer.CommitTx(ctx, tx)
	if err != nil {
		return entity.DeleteResponse{}, errs.E(op, err)
	}

	response := entity.DeleteResponse{
		ExternalID: extlID,
		Deleted:    true,
	}

	return response, nil
}

func deleteAppTx(ctx context.Context, tx pgx.Tx, a entity.App) (err error) {
	const op errs.Op = "service/deleteAppTx"

	// one-to-many API keys can be associated with an App. This will
	// delete them all.
	var apiKeysRowsAffected int64
	apiKeysRowsAffected, err = datastore.New(tx).DeleteAppAPIKeys(ctx, a.ID)
	if err != nil {
		return errs.E(op, errs.Database, err)
	}

	if apiKeysRowsAffected < 1 {
		return errs.E(op, errs.Database, fmt.Sprintf("rows affected should be at least 1, actual: %d", apiKeysRowsAffected))
	}

	var rowsAffected int64
	rowsAffected, err = datastore.New(tx).DeleteApp(ctx, a.ID)
	if err != nil {
		return errs.E(op, errs.Database, err)
	}

	if rowsAffected != 1 {
		return errs.E(op, errs.Database, fmt.Sprintf("rows affected should be 1, actual: %d", rowsAffected))
	}

	return nil
}

// FindByExternalID is used to find an App by its External ID
func (s *AppUsecase) FindByExternalID(ctx context.Context, extlID string) (ar *entity.AppResponse, err error) {
	const op errs.Op = "service/AppService.FindByExternalID"

	// start db txn using pgxpool
	var tx pgx.Tx
	tx, err = s.Datastorer.BeginTx(ctx)
	if err != nil {
		return nil, errs.E(op, err)
	}
	// defer transaction rollback and handle error, if any
	defer func() {
		err = s.Datastorer.RollbackTx(ctx, tx, err)
	}()

	var aa appAudit
	aa, err = findAppByExternalIDWithAudit(ctx, tx, extlID)
	if err != nil {
		return nil, errs.E(op, err)
	}

	return newAppResponse(aa), nil
}

// FindAll is used to list all apps in the datastore
func (s *AppUsecase) FindAll(ctx context.Context) (sar []*entity.AppResponse, err error) {
	const op errs.Op = "service/AppService.FindAll"

	// start db txn using pgxpool
	var tx pgx.Tx
	tx, err = s.Datastorer.BeginTx(ctx)
	if err != nil {
		return nil, errs.E(op, err)
	}
	// defer transaction rollback and handle error, if any
	defer func() {
		err = s.Datastorer.RollbackTx(ctx, tx, err)
	}()

	var rows []datastore.FindAppsWithAuditRow
	rows, err = datastore.New(tx).FindAppsWithAudit(ctx)
	if err != nil {
		return nil, errs.E(op, errs.Database, err)
	}

	for _, row := range rows {
		a := &entity.App{
			ID:         row.AppID,
			ExternalID: secure.MustParseIdentifier(row.AppExtlID),
			Org: &entity.Org{
				ID:          row.OrgID,
				ExternalID:  secure.MustParseIdentifier(row.OrgExtlID),
				Name:        row.OrgName,
				Description: row.OrgDescription,
				Kind: &entity.OrgKind{
					ID:          row.OrgKindID,
					ExternalID:  row.OrgKindExtlID,
					Description: row.OrgKindDesc,
				},
			},
			Name:        row.AppName,
			Description: row.AppDescription,
			APIKeys:     nil,
		}

		sa := &entity.SimpleAudit{
			Create: entity.Audit{
				App: &entity.App{
					ID:          row.CreateAppID,
					ExternalID:  secure.MustParseIdentifier(row.CreateAppExtlID),
					Org:         &entity.Org{ID: row.CreateAppOrgID},
					Name:        row.CreateAppName,
					Description: row.CreateAppDescription,
					APIKeys:     nil,
				},
				User: &entity.User{
					ID:        row.CreateUserID.UUID,
					FirstName: row.CreateUserFirstName.String,
					LastName:  row.CreateUserLastName.String,
				},
				Moment: row.CreateTimestamp,
			},
			Update: entity.Audit{
				App: &entity.App{
					ID:          row.UpdateAppID,
					ExternalID:  secure.MustParseIdentifier(row.UpdateAppExtlID),
					Org:         &entity.Org{ID: row.UpdateAppOrgID},
					Name:        row.UpdateAppName,
					Description: row.UpdateAppDescription,
					APIKeys:     nil,
				},
				User: &entity.User{
					ID:        row.UpdateUserID.UUID,
					FirstName: row.UpdateUserFirstName.String,
					LastName:  row.UpdateUserLastName.String,
				},
				Moment: row.UpdateTimestamp,
			},
		}
		or := newAppResponse(appAudit{App: a, SimpleAudit: sa})

		sar = append(sar, or)
	}

	return sar, nil
}

func findAppByID(ctx context.Context, dbtx datastore.DBTX, id uuid.UUID) (entity.App, error) {
	const op errs.Op = "service/findAppByID"

	row, err := datastore.New(dbtx).FindAppByID(ctx, id)
	if err != nil {
		return entity.App{}, errs.E(op, errs.Database, err)
	}

	a := entity.App{
		ID:         row.AppID,
		ExternalID: secure.MustParseIdentifier(row.AppExtlID),
		Org: &entity.Org{
			ID:          row.OrgID,
			ExternalID:  secure.MustParseIdentifier(row.OrgExtlID),
			Name:        row.OrgName,
			Description: row.OrgDescription,
			Kind: &entity.OrgKind{
				ID:          row.OrgKindID,
				ExternalID:  row.OrgKindExtlID,
				Description: row.OrgKindDesc,
			},
		},
		Name:        row.AppName,
		Description: row.AppDescription,
		APIKeys:     nil,
	}

	return a, nil
}

func findAppByExternalID(ctx context.Context, dbtx datastore.DBTX, extlID string) (entity.App, error) {
	const op errs.Op = "service/findAppByExternalID"

	row, err := datastore.New(dbtx).FindAppByExternalID(ctx, extlID)
	if err != nil {
		return entity.App{}, errs.E(op, errs.Database, err)
	}

	a := entity.App{
		ID:         row.AppID,
		ExternalID: secure.MustParseIdentifier(row.AppExtlID),
		Org: &entity.Org{
			ID:          row.OrgID,
			ExternalID:  secure.MustParseIdentifier(row.OrgExtlID),
			Name:        row.OrgName,
			Description: row.OrgDescription,
			Kind: &entity.OrgKind{
				ID:          row.OrgKindID,
				ExternalID:  row.OrgKindExtlID,
				Description: row.OrgKindDesc,
			},
		},
		Name:        row.AppName,
		Description: row.AppDescription,
		APIKeys:     nil,
	}

	return a, nil
}

// findAppByExternalIDWithAudit retrieves App data from the datastore
// given a unique external ID, which is then hydrated into an App
// and audit struct.
func findAppByExternalIDWithAudit(ctx context.Context, dbtx datastore.DBTX, extlID string) (appAudit, error) {
	const op errs.Op = "service/findAppByExternalIDWithAudit"

	var (
		row datastore.FindAppByExternalIDWithAuditRow
		err error
	)

	row, err = datastore.New(dbtx).FindAppByExternalIDWithAudit(ctx, extlID)
	if err != nil {
		return appAudit{}, errs.E(op, errs.Database, err)
	}

	a := &entity.App{
		ID:         row.AppID,
		ExternalID: secure.MustParseIdentifier(row.AppExtlID),
		Org: &entity.Org{
			ID:          row.OrgID,
			ExternalID:  secure.MustParseIdentifier(row.OrgExtlID),
			Name:        row.OrgName,
			Description: row.OrgDescription,
			Kind: &entity.OrgKind{
				ID:          row.OrgKindID,
				ExternalID:  row.OrgKindExtlID,
				Description: row.OrgKindDesc,
			},
		},
		Name:        row.AppName,
		Description: row.AppDescription,
		APIKeys:     nil,
	}

	sa := &entity.SimpleAudit{
		Create: entity.Audit{
			App: &entity.App{
				ID:          row.CreateAppID,
				ExternalID:  secure.MustParseIdentifier(row.CreateAppExtlID),
				Org:         &entity.Org{ID: row.CreateAppOrgID},
				Name:        row.CreateAppName,
				Description: row.CreateAppDescription,
				APIKeys:     nil,
			},
			User: &entity.User{
				ID:        row.CreateUserID.UUID,
				FirstName: row.CreateUserFirstName.String,
				LastName:  row.CreateUserLastName.String,
			},
			Moment: row.CreateTimestamp,
		},
		Update: entity.Audit{
			App: &entity.App{
				ID:          row.UpdateAppID,
				ExternalID:  secure.MustParseIdentifier(row.UpdateAppExtlID),
				Org:         &entity.Org{ID: row.UpdateAppOrgID},
				Name:        row.UpdateAppName,
				Description: row.UpdateAppDescription,
				APIKeys:     nil,
			},
			User: &entity.User{
				ID:        row.UpdateUserID.UUID,
				FirstName: row.UpdateUserFirstName.String,
				LastName:  row.UpdateUserLastName.String,
			},
			Moment: row.UpdateTimestamp,
		},
	}

	return appAudit{App: a, SimpleAudit: sa}, nil
}

func findAppByProviderClientID(ctx context.Context, tx pgx.Tx, id string) (*entity.App, error) {
	const op errs.Op = "service/findAppByProviderClientID"

	row, err := datastore.New(tx).FindAppByProviderClientID(ctx, entity.NewNullString(id))
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errs.E(op, errs.NotExist, fmt.Sprintf("no app registered for provider client ID: %s", id))
		} else {
			return nil, errs.E(op, errs.Database, err)
		}
	}

	a := entity.App{
		ID:         row.AppID,
		ExternalID: secure.MustParseIdentifier(row.AppExtlID),
		Org: &entity.Org{
			ID:          row.OrgID,
			ExternalID:  secure.MustParseIdentifier(row.OrgExtlID),
			Name:        row.OrgName,
			Description: row.OrgDescription,
			Kind: &entity.OrgKind{
				ID:          row.OrgKindID,
				ExternalID:  row.OrgKindExtlID,
				Description: row.OrgKindDesc,
			},
		},
		Name:        row.AppName,
		Description: row.AppDescription,
		APIKeys:     nil,
	}

	return &a, nil
}

// FindAppByName finds an App in the database given an org and app name.
func FindAppByName(ctx context.Context, tx datastore.DBTX, o *entity.Org, name string) (*entity.App, error) {
	const op errs.Op = "service/FindAppByName"

	findAppByNameParams := datastore.FindAppByNameParams{
		OrgID:   o.ID,
		AppName: name,
	}

	dbAppRow, err := datastore.New(tx).FindAppByName(ctx, findAppByNameParams)
	if err != nil {
		return nil, errs.E(op, errs.Database, err)
	}

	a := &entity.App{
		ID:          dbAppRow.AppID,
		ExternalID:  secure.MustParseIdentifier(dbAppRow.AppExtlID),
		Org:         o,
		Name:        dbAppRow.AppName,
		Description: dbAppRow.AppDescription,
		APIKeys:     nil,
	}

	return a, nil
}