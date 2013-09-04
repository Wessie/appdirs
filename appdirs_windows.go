package appdirs

import (
	"path/filepath"
	"strings"
	"syscall"
	"unsafe"
)

var (
	shell32, _        = syscall.LoadLibrary("shell32.dll")
	getFolderPathW, _ = syscall.GetProcAddress(shell32, "SHGetFolderPathW")
)

type Csidl uint

const (
	APPDATA        Csidl = 26
	COMMON_APPDATA       = 35
	LOCAL_APPDATA        = 28
)

func (app *App) UserData() string {
	var author string
	if author = app.Author; author == "" {
		author = app.Name
	}

	var csidl Csidl
	if app.Roaming {
		csidl = APPDATA
	} else {
		csidl = LOCAL_APPDATA
	}

	path, err := GetFolderPath(csidl)

	if err != nil {
		return ""
	}

	if path, err = filepath.Abs(path); err != nil {
		return ""
	}

	if app.Name != "" {
		return filepath.Join(path, author, app.Name)
	} else {
		return path
	}
}

func (app *App) SiteData() (path string) {
	path, err := GetFolderPath(COMMON_APPDATA)

	if err != nil {
		return ""
	}

	if path, err = filepath.Abs(path); err != nil {
		return ""
	}

	author := app.Author

	if author == "" {
		author = app.Name
	}

	if app.Name != "" {
		return filepath.Join(path, author, app.Name)
	} else {
		return path
	}
}

func GetFolderPath(csidl_const Csidl) (string, error) {
	var buf = strings.Repeat("0", 1024)
	cbuf, err := syscall.UTF16FromString(buf)
	if err != nil {
		return "", err
	}

	ret, _, callErr := syscall.Syscall6(
		uintptr(getFolderPathW),
		5,                    // The amount of arguments we have
		0,                    // Reserved argument, does nothing
		uintptr(csidl_const), // CSIDL value identifier
		0,                    // Access token, almost always 0
		0,                    // Flag to specify the path to be returned
		// null-terminated string to put the output in
		uintptr(unsafe.Pointer(&cbuf[0])),
		0, // Filler argument to syscall6, always 0
	)
	println(ret)

	if callErr != 0 && ret != 0 {
		return "", callErr
	}

	return syscall.UTF16ToString(cbuf), nil
}
