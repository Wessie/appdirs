// A port of the excellent python module `appdirs`.
// See https://github.com/ActiveState/appdirs for the python version.
package appdirs

import (
	"os/user"
	"strings"
)

type App struct {
	Name    string
	Author  string
	Version string
	Roaming bool
	Opinion bool
}

// Return a new App struct, this uses the following defaults for the parameters
// not applicable in New.
//
// Roaming: false, Opinion: true
func New(name, author, version string) *App {
	return &App{
		Name:    name,
		Author:  author,
		Version: version,
		Roaming: false,
		Opinion: true,
	}
}

func (app *App) UserData() string {
	return UserDataDir(app.Name, app.Author, app.Version, app.Roaming)
}

func (app *App) SiteData() string {
	return SiteDataDir(app.Name, app.Author, app.Version)
}

func (app *App) SiteConfig() string {
	return SiteConfigDir(app.Name, app.Author, app.Version)
}

func (app *App) UserConfig() string {
	return UserConfigDir(app.Name, app.Author, app.Version, app.Roaming)
}

func (app *App) UserCache() string {
	return UserCacheDir(app.Name, app.Author, app.Version, app.Opinion)
}

func (app *App) UserLog() string {
	return UserLogDir(app.Name, app.Author, app.Version, app.Opinion)
}

func ExpandUser(path string) string {
	if u, err := user.Current(); err == nil {
		return strings.Replace(path, "~", u.HomeDir, -1)
	}
	return path
}
