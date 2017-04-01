// Copyright 2017 Joan Llopis. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package store

import "github.com/jllopis/zbs/services"

// Storer es la interfaz que deben implementar los diferentes backends que realizen
// la persistencia del microservicio
type Storer interface {
	// Conecta con el Store
	Dial(options Options) error
	// Status devuelve el estado de la conexión
	Status() (int, string)
	// Close cierra la conexión con el Store
	Close() error
	Job
}

type Job interface {
	FindJob(*services.IdMessage) (*services.JobMsg, error)
	ListJobs(*services.JobFilterMsg) (*services.JobArrayMsg, error)
	AddJob(*services.JobMsg) (*services.IdMessage, error)
	UpdateJob(*services.JobMsg) (*services.JobMsg, error)
	DeleteJob(*services.IdMessage) (*services.JobMsg, error)
}

const (
	// DISCONNECTED indicates that there is no connection with the Storer
	DISCONNECTED = iota
	// CONNECTED indicate that the connection with the Storer is up and running
	CONNECTED
)

var (
	// StatusStr is a string representation of the status of the connections with the Storer
	StatusStr = []string{"Disconnected", "Connected"}
)

// Options is a map to hold the database connection options
type Options map[string]interface{}
