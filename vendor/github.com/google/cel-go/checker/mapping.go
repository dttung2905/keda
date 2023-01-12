// Copyright 2018 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package checker

import (
	exprpb "google.golang.org/genproto/googleapis/api/expr/v1alpha1"
)

type mapping struct {
	mapping map[string]*exprpb.Type
}

func newMapping() *mapping {
	return &mapping{
		mapping: make(map[string]*exprpb.Type),
	}
}

func (m *mapping) add(from *exprpb.Type, to *exprpb.Type) {
	m.mapping[typeKey(from)] = to
}

func (m *mapping) find(from *exprpb.Type) (*exprpb.Type, bool) {
	if r, found := m.mapping[typeKey(from)]; found {
		return r, found
	}
	return nil, false
}

func (m *mapping) copy() *mapping {
	c := newMapping()

	for k, v := range m.mapping {
		c.mapping[k] = v
	}
	return c
}
