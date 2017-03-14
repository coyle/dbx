// Copyright (C) 2017 Space Monkey, Inc.
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

package sqlgen

import (
	"testing"

	"gopkg.in/spacemonkeygo/dbx.v1/testutil"
)

func TestEqual(t *testing.T) {
	tw := testutil.Wrap(t)
	tw.Parallel()
	tw.Runp("fuzz identity", testEqualFuzzIdentity)
}

func testEqualFuzzIdentity(tw *testutil.T) {
	g := newGenerator(tw)

	for i := 0; i < 1000; i++ {
		sql := g.gen()

		if !sqlEqual(sql, sql) {
			tw.Logf("sql: %#v", sql)
			tw.Error()
		}
	}
}
