package utils

import "os"

func ExitErrorApplyingMigrations() {
	os.Exit(131)
}

func ExitErrorReadingMigrations() {
	os.Exit(130)
}

func ExitErrorConnectingToDatabase() {
	os.Exit(129)
}

func ExitErrorAccessingFileSystem() {
	os.Exit(128)
}

func ExitErrorInvalidArguments() {
	os.Exit(127)
}

func ExitSuccessfullyWithIncompleteMigrations() {
	os.Exit(1)
}

func ExitSuccessfully() {
	os.Exit(0)
}
