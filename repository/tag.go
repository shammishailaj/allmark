// Copyright 2013 Andreas Koch. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package repository

import (
	"fmt"
	"strings"
)

type Tag struct {
	Name string
}

func NewTag(name string) (*Tag, error) {

	normalized := normalizeTagName(name)
	if normalized == "" {
		return nil, fmt.Errorf("Cannot create a tag from an empty string")
	}

	return &Tag{
		Name: normalized,
	}, nil
}

func normalizeTagName(name string) string {
	return strings.ToLower(strings.TrimSpace(name))
}
