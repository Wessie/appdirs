// appdirs project doc.go

/*
This is a port of a python module used for finding what directory you 'should'
be using for saving your application data such as configuration, cache files or
other files.

The location of these directories is often hard to get right. The original python
module set out to change this into a simple API that returns you the exact
directory you need. This is a port of it to Go.

Depending on platform, this package exports you at the least 6 functions that
return various system directories. And one helper struct type that combines the
functions into methods for less arguments in your code.
*/
package appdirs

// UserDataDir returns the full path to the user-specific data directory.
//
//		"name" is the name of the application.
//		"author" is the name of the application author or company.
//		"version" is an optional version number.
//		"roaming" is windows specific, this uses the roaming appdata directory when set.
//
// This function uses XDG_DATA_HOME as defined by the XDG spec on *nix like systems.
//
// Examples of return values:
// 		Mac OS X: ~/Library/Application Support/<AppName>
// 		Unix: ~/.local/share/<AppName> # or in $XDG_DATA_HOME, if defined
// 		Win XP (not roaming): C:\Documents and Settings\<username>\Application Data\<AppAuthor>\<AppName>
// 		Win XP (roaming): C:\Documents and Settings\<username>\Local Settings\Application Data\<AppAuthor>\<AppName>
// 		Win 7 (not roaming): C:\Users\<username>\AppData\Local\<AppAuthor>\<AppName>
// 		Win 7 (roaming): C:\Users\<username>\AppData\Roaming\<AppAuthor>\<AppName>
func UserDataDir(name, author, version string, roaming bool) string {
	return userDataDir(name, author, version, roaming)
}

// SiteDataDir returns the full path to the user-shared data directory.
//
//		"name" is the name of the application.
//		"author" is the name of the application author or company.
//		"version" is an optional version number.
//
// This function uses XDG_DATA_DIRS[0] as by the XDG spec on *nix like systems.
//
// Examples of return values:
//		Mac OS X: /Library/Application Support/<AppName>
//		Unix: /usr/local/share/<AppName> or /usr/share/<AppName>
//		Win XP: C:\Documents and Settings\All Users\Application Data\<AppAuthor>\<AppName>
//		Vista: (Fail! "C:\ProgramData" is a hidden *system* directory on Vista.)
//		Win 7: C:\ProgramData\<AppAuthor>\<AppName> # Hidden, but writeable on Win 7.
//
// WARNING: Do not use this on Windows Vista, See the note above.
func SiteDataDir(name, author, version string) string {
	return siteDataDir(name, author, version)
}

// UserConfigDir returns the full path to the user-specific configuration directory
//
//		"name" is the name of the application.
//		"author" is the name of the application author or company.
//		"version" is an optional version number.
//		"roaming" is windows specific, this uses the roaming appdata directory when set.
//
// This function uses XDG_CONFIG_HOME as by the XDG spec on *nix like systems.
//
// Examples of return values:
//		Mac OS X: same as UserDataDir
//		Unix: ~/.config/<AppName> # or in $XDG_CONFIG_HOME, if defined
//		Win *: same as UserDataDir
func UserConfigDir(name, author, version string, roaming bool) string {
	return userConfigDir(name, author, version, roaming)
}

// SiteConfigDir returns the full path to the user-shared data directory.
//
//		"name" is the name of the application.
//		"author" is the name of the application author or company.
//		"version" is an optional version number.
//
// This function uses XDG_CONFIG_DIRS[0] as by the XDG spec on *nix like systems.
//
// Examples of return values:
//		Mac OS X: same as SiteDataDir
//		Unix: /etc/xdg/<AppName> or $XDG_CONFIG_DIRS[i]/<AppName> for each value in $XDG_CONFIG_DIRS
//		Win *: same as SiteDataDir
//		Vista: (Fail! "C:\ProgramData" is a hidden *system* directory on Vista.)
//
// WARNING: Do not use this on Windows Vista, see the note above.
func SiteConfigDir(name, author, version string) string {
	return siteConfigDir(name, author, version)
}

// UserCacheDir returns the full path to the user-specific cache directory.
//
//		"name" is the name of the application.
//		"author" is the name of the application author or company.
//		"version" is an optional version number.
//		"opinion" can be set to 'true' to append "Cache" to the base directory, or
//			'false' to omit it. (See the original python appdata documentation for
//			the reasons behind supporting this.)
//
// Examples of return values:
//		Mac OS X: ~/Library/Caches/<AppName>
//		Unix: ~/.cache/<AppName> (XDG default)
//		Win XP: C:\Documents and Settings\<username>\Local Settings\Application Data\<AppAuthor>\<AppName>\Cache
//		Vista: C:\Users\<username>\AppData\Local\<AppAuthor>\<AppName>\Cache
func UserCacheDir(name, author, version string, opinion bool) string {
	return userCacheDir(name, author, version, opinion)
}

// UserLogDir returns the full path to the user-specific log directory.
//
//		"name" is the name of the application.
//		"author" is the name of the application author or company.
//		"version" is an optional version number.
//		"opinion" can be set to 'true' to append "Logs" (windows) or "log" (unix) to
//			the base directory, or 'false' to omit it. (See the original python
//			appdata documentation for the reasons behind supporting this.)
//
// Examples of return values:
//		Mac OS X: ~/Library/Logs/<AppName>
//		Unix: ~/.cache/<AppName>/log # or under $XDG_CACHE_HOME if defined
//		Win XP: C:\Documents and Settings\<username>\Local Settings\Application Data\<AppAuthor>\<AppName>\Logs
//		Vista: C:\Users\<username>\AppData\Local\<AppAuthor>\<AppName>\Logs
func UserLogDir(name, author, version string, opinion bool) string {
	return userLogDir(name, author, version, opinion)
}
