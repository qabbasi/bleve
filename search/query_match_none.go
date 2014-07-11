//  Copyright (c) 2014 Couchbase, Inc.
//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file
//  except in compliance with the License. You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
//  Unless required by applicable law or agreed to in writing, software distributed under the
//  License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
//  either express or implied. See the License for the specific language governing permissions
//  and limitations under the License.
package search

import (
	"github.com/couchbaselabs/bleve/index"
)

type MatchNoneQuery struct {
	BoostVal float64 `json:"boost,omitempty"`
	Explain  bool    `json:"explain,omitempty"`
}

func (q *MatchNoneQuery) Boost() float64 {
	return q.BoostVal
}

func (q *MatchNoneQuery) Searcher(index index.Index) (Searcher, error) {
	return NewMatchNoneSearcher(index, q)
}

func (q *MatchNoneQuery) Validate() error {
	return nil
}
