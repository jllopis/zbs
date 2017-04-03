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
}

func (m *MemStore) AddJob(job *services.JobMsg) (*services.IdMessage, error) {
	if _, ok := m.Pool[job.Id]; ok {
		return nil, errors.New("record exists")
	}

	nextID := int64(len(m.Pool) + 1)
	m.Pool[nextID] = job.JobData
	if err := m.fsync(); err != nil {
		return nil, err
	}
	return &services.IdMessage{Id: nextID}, nil
}

func (m *MemStore) UpdateJob(*services.JobMsg) (*services.JobMsg, error) {
	return nil, errors.New("function not implemented")
}

func (m *MemStore) DeleteJob(*services.IdMessage) (*services.JobMsg, error) {
	return nil, errors.New("function not implemented")
}
