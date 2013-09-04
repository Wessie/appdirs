// A port of the excellent python module `appdirs`.
// See https://github.com/ActiveState/appdirs for the python version.
package appdirs

type App struct {
	Name      string
	Author    string
	Version   string
	Roaming   bool
	MultiPath bool
	Opinion   bool
}

// Return a new App struct, this uses the following defaults for the parameters
// not applicable in New.
//
// Roaming: false, MultiPath: false, Opinion: true
func New(name, author, version string) *App {
	return &App{
		Name:      name,
		Author:    author,
		Version:   version,
		Roaming:   false,
		MultiPath: false,
		Opinion:   true,
	}
}
