// Copyright 2018 The OpenPitrix Authors. All rights reserved.
// Use of this source code is governed by a Apache license
// that can be found in the LICENSE file.

package app

import (
	"context"

	"openpitrix.io/openpitrix/pkg/constants"
	"openpitrix.io/openpitrix/pkg/manager"
	"openpitrix.io/openpitrix/pkg/pb"
)

func NewAppManagerClient(ctx context.Context) (pb.AppManagerClient, error) {
	conn, err := manager.NewClient(ctx, constants.AppManagerHost, constants.AppManagerPort)
	if err != nil {
		return nil, err
	}
	return pb.NewAppManagerClient(conn), err
}
