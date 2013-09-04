package appdirs

import (
	"path/filepath"
)

func UserDataDir(name, author, version string, roaming bool) (path string) {
	path = ExpandUser("~/Library/Application Support")

	if name != "" {
		path = filepath.Join(path, name)
	}

	if name != "" && version != "" {
		path = filepath.Join(path, version)
	}

	return path
}

func SiteDataDir(name, author, version string) (path string) {
	path = ExpandUser("/Library/Application Support")

	if name != "" {
		path = filepath.Join(path, name)
	}

	if name != "" && version != "" {
		path = filepath.Join(path, version)
	}

	return path
}

func UserConfigDir(name, author, version string, roaming bool) (path string) {
	return UserDataDir(name, author, version, roaming)
}

func SiteConfigDir(name, author, version string) (path string) {
	return SiteDataDir(name, author, version, false)
}

func UserCacheDir(name, author, version string, opinion bool) (path string) {
	path = ExpandUser("~/Library/Caches")

	if name != "" {
		path = filepath.Join(path, name)
	}

	if name != "" && version != "" {
		path = filepath.Join(path, version)
	}

	return path
}

func UserLogDir(name, author, version string, opinion bool) (path string) {
	path = ExpandUser("~/Library/Logs")

	path = filepath.Join(path, name)

	if name != "" && version != "" {
		path = filepath.Join(path, version)
	}
}
