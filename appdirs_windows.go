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

func UserDataDir(name, author, version string, roaming bool) (path string) {
	if author == "" {
		author = name
	}

	var csidl Csidl
	if roaming {
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

	if name != "" {
		path = filepath.Join(path, author, name)
	}

	if name != "" && version != "" {
		path = filepath.Join(path, version)
	}

	return path
}

func SiteDataDir(name, author, version string) (path string) {
	path, err := GetFolderPath(COMMON_APPDATA)

	if err != nil {
		return ""
	}

	if path, err = filepath.Abs(path); err != nil {
		return ""
	}

	if author == "" {
		author = name
	}

	if name != "" {
		path = filepath.Join(path, author, name)
	}

	if name != "" && version != "" {
		path = filepath.Join(path, version)
	}

	return path
}

func UserConfigDir(name, author, version string, roaming bool) string {
	return UserDataDir(name, author, version, roaming)
}

func SiteConfigDir(name, author, version string) (path string) {
	return SiteDataDir(name, author, version)
}

func UserCacheDir(name, author, version string, opinion bool) (path string) {
	if author == "" {
		author = name
	}

	path, err := GetFolderPath(LOCAL_APPDATA)

	if err != nil {
		return ""
	}

	if path, err = filepath.Abs(path); err != nil {
		return ""
	}

	if name != "" {
		path = filepath.Join(path, author, name)
		if opinion {
			path = filepath.Join(path, "Cache")
		}
	}

	if name != "" && version != "" {
		path = filepath.Join(path, version)
	}

	return path
}

func UserLogDir(name, author, version string, opinion bool) (path string) {
	path = UserDataDir(name, author, version, false)

	if opinion {
		path = filepath.Join(path, "Logs")
	}

	return path
}

// A helper function to receive a CSIDL folder from windows. This is exported
// for package users for if they will want to receive a different CSIDL folder
// than the ones we support.
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

	if callErr != 0 && ret != 0 {
		return "", callErr
	}

	return syscall.UTF16ToString(cbuf), nil
}
