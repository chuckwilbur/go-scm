// Copyright 2018 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package stash

import (
	"context"

	"github.com/drone/go-scm/scm"
)

type tokenService struct {
}

func (t *tokenService) Refresh(context.Context, *scm.Token) (bool, error) {
	// this function is a no-op because Bitbucket Server
	// does not implement refresh tokens.
	return false, nil
}
