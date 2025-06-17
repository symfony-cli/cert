// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cert

import (
	"os"
)

var (
	FirefoxProfiles = []string{os.Getenv("HOME") + "/.mozilla/firefox/*",
		os.Getenv("HOME") + "/.mozilla/firefox-trunk/*"}
	NSSBrowsers = "Firefox and/or Chrome/Chromium"

	CertutilInstallHelp string
)

func (ca *CA) installPlatform() error {
	return nil
}

func (ca *CA) uninstallPlatform() error {
	return nil
}
