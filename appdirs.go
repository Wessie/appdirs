// A port of the excellent python module `appdirs`.
// See https://github.com/ActiveState/appdirs for the python version.
package appdirs

type App struct {
	Name      string
	Author    string
	Version   string
	Roaming   bool
	MultiPath bool
}
