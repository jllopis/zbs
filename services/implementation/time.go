// Copyright 2017 Joan Llopis. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package implementation

import (
	"time"

	"github.com/jllopis/zbs/services"

	google_protobuf1 "github.com/golang/protobuf/ptypes/empty"
	"golang.org/x/net/context"
)

type ServerTime struct{}

func (st *ServerTime) GetServerTime(context.Context, *google_protobuf1.Empty) (*services.ServerTimeMessage, error) {
	return &services.ServerTimeMessage{Value: time.Now().UTC().UnixNano()}, nil
}
