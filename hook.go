/*
Package macos provides macOS Unified Loggging functionality to logrus via cgo.

Messages are logged as public.

See man os_log(3).
*/
package macos

// #include <ul.h>
import "C"

import (
	"fmt"
	"os"
	"unsafe"

	"github.com/sirupsen/logrus"
)

type level uint8

type UnifiedLogger struct {
	Subsystem string
	Category  string
	os_log_t  C.os_log_t
}

func New() *UnifiedLogger {
	l := &UnifiedLogger{}
	l.Finalize()
	return l
}

func (l *UnifiedLogger) Finalize() {
	if l.os_log_t != nil {
		return
	}

	if 0 < len(l.Subsystem) || 0 < len(l.Category) {
		s := C.CString(l.Subsystem)
		defer C.free(unsafe.Pointer(s))

		c := C.CString(l.Category)
		defer C.free(unsafe.Pointer(c))

		l.os_log_t = C.os_log_create(s, c)
	} else {
		l.os_log_t = C.os_log_default
	}
}

func (l *UnifiedLogger) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (l *UnifiedLogger) Fire(entry *logrus.Entry) error {
	line, err := entry.String()
	if err != nil {
		fmt.Fprintf(os.Stderr, "UnifiedLogger.Fire: Unable to read entry, %v", err)
		return err
	}

	go func() {
		cs := C.CString(line)
		defer C.free(unsafe.Pointer(cs))

		C.ul_log((C.uchar)(entry.Level), l.os_log_t, cs)
	}()

	return nil
}
