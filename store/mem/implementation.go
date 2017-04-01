// Copyright 2017 Joan Llopis. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package mem

import (
	"errors"

	"github.com/jllopis/zbs/services"
)

func (m *MemStore) FindJob(*services.IdMessage) (*services.JobMsg, error) {
	return nil, errors.New("function not implemented")
}
func (m *MemStore) ListJobs(*services.JobFilterMsg) (*services.JobArrayMsg, error) {
	jobs := &services.JobArrayMsg{}
	for id, jobData := range m.Pool {
		jobs.Jobs = append(jobs.Jobs, &services.JobMsg{Id: id, JobData: jobData})
	}
	jobs.Total = int32(len(jobs.Jobs))
	return jobs, nil
	//	return nil, errors.New("function not implemented")
}
func (m *MemStore) AddJob(*services.JobMsg) (*services.IdMessage, error) {
	return nil, errors.New("function not implemented")
}
func (m *MemStore) UpdateJob(*services.JobMsg) (*services.JobMsg, error) {
	return nil, errors.New("function not implemented")
}
func (m *MemStore) DeleteJob(*services.IdMessage) (*services.JobMsg, error) {
	return nil, errors.New("function not implemented")
}
