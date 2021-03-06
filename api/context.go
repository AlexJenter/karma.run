// Copyright 2017 karma.run AG. All rights reserved.
// Use of this source code is governed by an AGPL license that can be found in the LICENSE file.
package api

type contextKeyCodecT struct{}
type contextKeyDatabaseT struct{}
type contextKeyUserIdT struct{}

var (
	ContextKeyCodec    = contextKeyCodecT{}
	ContextKeyDatabase = contextKeyDatabaseT{}
	ContextKeyUserId   = contextKeyUserIdT{}
)
