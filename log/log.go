package log

import (
	"fmt"

	"github.com/ttacon/chalk"
)

// Infof print info message to stdout
func Infof(msg string, args ...interface{}) {
	fmt.Println(chalk.Green.Color(fmt.Sprintf(msg, args...)))
}

// Warnf print warn message to stdout
func Warnf(msg string, args ...interface{}) {
	fmt.Println(chalk.Cyan.Color(fmt.Sprintf(msg, args...)))
}

// Errorf print error message to stdout
func Errorf(msg string, args ...interface{}) {
	fmt.Println(chalk.Red.Color(fmt.Sprintf(msg, args...)))
}

// ErrorOrInfof print the message to stdout according the error value
func ErrorOrInfof(err error,
	infof func(),
	errorf func(),
) {
	if err != nil {
		errorf()
		return
	}

	infof()
}
