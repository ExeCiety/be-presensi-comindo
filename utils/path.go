package utils

var RootPath *string

func GetRootPath() string {
	return *RootPath
}

func SetRootPath(cwd string) {
	RootPath = &cwd
}
