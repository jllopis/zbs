// Copyright 2017 Joan Llopis. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package implementation

import (
	"errors"

	"github.com/jllopis/zbs/services"
	"github.com/jllopis/zbs/store"
	"golang.org/x/net/context"
)

// ServiceZbs proporciona la estructura para implementar el servicio Xxxx.
// Recibe una referencia para interactuar con el Storer.
type ServiceZbs struct {
	Store store.Storer
}

// FindJob devuelve el Job cuyo id corresponde con el solicitado
func (sz *ServiceZbs) FindJob(ctx context.Context, id *services.IdMessage) (*services.JobMsg, error) {
	if sz.Store == nil {
		return nil, errors.New("service store is nil")
	}

	return nil, errors.New("method not implemented")
	//	return sz.Store.FindJob(id.Id), nil
}

// ListJobs devuelve los Jobs resultantes de aplicar el filtro suministrado
func (sz *ServiceZbs) ListJobs(ctx context.Context, filter *services.JobFilterMsg) (*services.JobArrayMsg, error) {
	if sz.Store == nil {
		return nil, errors.New("service store is nil")
	}

	return sz.Store.ListJobs(filter)
}

// AddJob crea un nuevo registro
func (sz *ServiceZbs) AddJob(ctx context.Context, job *services.JobMsg) (*services.IdMessage, error) {
	if sz.Store == nil {
		return nil, errors.New("service store is nil")
	}
	return sz.Store.AddJob(job)
}

// UpdateJob actualiza el Job y lo devuelve actualizado
func (sz *ServiceZbs) UpdateJob(ctx context.Context, filter *services.JobMsg) (*services.JobMsg, error) {
	return nil, errors.New("method not implemented")
}

// DeleteJob elimina el registro y devuelve el registro eliminado
func (sz *ServiceZbs) DeleteJob(ctx context.Context, id *services.IdMessage) (*services.JobMsg, error) {
	if sz.Store == nil {
		return nil, errors.New("service store is nil")
	}
	return nil, errors.New("method not implemented")
	//	return sz.Store.DeleteJob(id.Id)
}
