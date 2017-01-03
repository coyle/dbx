// Copyright (C) 2016 Space Monkey, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ast

import "text/scanner"

type Select struct {
	Pos        scanner.Position
	FuncSuffix string
	Limit      *Limit
	Fields     *FieldRefs
	Joins      []*Join
	Where      []*Where
	OrderBy    *OrderBy
}

type Limit struct {
	Pos    scanner.Position
	Amount int
}

type Join struct {
	Pos   scanner.Position
	Left  *FieldRef
	Right *FieldRef
	Type  JoinType
}

type JoinType int

const (
	LeftJoin JoinType = iota
)

type Where struct {
	Pos   scanner.Position
	Left  *FieldRef
	Op    Operator
	Right *FieldRef
}

type OrderBy struct {
	Pos        scanner.Position
	Fields     *FieldRefs
	Descending bool
}
