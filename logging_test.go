package runscope

import (
	"fmt"
	"strings"
	"testing"
)

func TestDebugLog(t *testing.T) {
	output := &strings.Builder{}
	handler := func(level int, format string, args ...interface{}) {
		output.WriteString(fmt.Sprintf("[DEBUG] %s", fmt.Sprintf(format, args...)))
	}

	RegisterLogHandlers(handler, handler, handler)

	DebugF(1, "bucket %s uri %s", "foo", "http://exmaple.com")

	want := "[DEBUG] bucket foo uri http://exmaple.com"
	got := output.String()
	if want != got {
		t.Errorf("Want %s got %s", want, got)
	}
}

func TestDefaultHandler(t *testing.T) {
	DebugF(1, "bucket %s uri %s", "foo", "http://exmaple.com")
	DebugF(2, "bucket %s uri %s", "foo", "http://exmaple.com")
}
