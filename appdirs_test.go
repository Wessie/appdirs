package appdirs

import (
	"testing"
)

func TestUserData(t *testing.T) {
	app := New("Testing", "Tester", "1.0.0")

	if r := app.UserData(); r == "" {
		t.Error("Received empty string.")
	} else {
		t.Logf("UserData: %s", r)
	}
}

func TestSiteData(t *testing.T) {
	app := New("Testing", "Tester", "1.0.0")

	if r := app.SiteData(); r == "" {
		t.Error("Received empty string.")
	} else {
		t.Logf("SiteData: %s", r)
	}
}

func TestSiteConfig(t *testing.T) {
	app := New("Testing", "Tester", "1.0.0")

	if r := app.SiteConfig(); r == "" {
		t.Error("Received empty string.")
	} else {
		t.Logf("SiteConfig: %s", r)
	}
}

func TestUserCache(t *testing.T) {
	app := New("Testing", "Tester", "1.0.0")

	if r := app.UserCache(); r == "" {
		t.Error("Received empty string.")
	} else {
		t.Logf("UserCache: %s", r)
	}
}

func TestUserConfig(t *testing.T) {
	app := New("Testing", "Tester", "1.0.0")

	if r := app.UserConfig(); r == "" {
		t.Error("Received empty string.")
	} else {
		t.Logf("UserConfig: %s", r)
	}
}

func TestUserLog(t *testing.T) {
	app := New("Testing", "Tester", "1.0.0")

	if r := app.UserLog(); r == "" {
		t.Error("Received empty string.")
	} else {
		t.Logf("UserLog: %s", r)
	}
}
