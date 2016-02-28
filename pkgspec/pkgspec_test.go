// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pkgspec

import "testing"

func TestParse(t *testing.T) {
	list := []struct {
		Spec string
		Pkg  *Pkg
		Err  error
	}{
		{"abc/def", nil, nil},
		{"", nil, ErrEmptyPath},
		{"::", nil, ErrEmptyPath},
		{"::foo", nil, ErrEmptyPath},
		{"abc/def::", nil, ErrEmptyOrigin},
		{"abc/def::foo/bar/vendor/abc/def", nil, nil},
		{"abc/def::foo/bar/vendor/abc/def@", nil, nil},
		{"abc/def::foo/bar/vendor/abc/def@v1.2.3", nil, nil},
		{"abc/def@", nil, nil},
		{"abc/def@v1.2.3", nil, nil},
	}

	for _, item := range list {
		pkg, err := Parse(item.Spec)
		if err != nil && item.Err != nil {
			if err != item.Err {
				t.Errorf("For %q, got error %q but expected error %q", item.Spec, err, item.Err)
				continue
			}
			continue
		}
		if err == nil && item.Err != nil {
			t.Errorf("For %q, got nil error but expected error %q, %#v", item.Spec, item.Err, pkg)
			continue
		}
		if pkg == nil {
			t.Errorf("For %q, got nil pkg", item.Spec)
			continue
		}
		str := pkg.String()
		if str != item.Spec {
			t.Errorf("For %q, round tripped to %q", item.Spec, str)
			continue
		}
	}
}
