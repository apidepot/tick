// Copyright (c) 2016-2023 The tick developers. All rights reserved.
// Project site: https://github.com/apidepot/tick
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package tick

import "testing"

func assertString(t *testing.T, label string, got, want string) {
	if got != want {
		t.Errorf("\t got = %s %s\n\t\t\twant = %s", got, label, want)
	}
}
