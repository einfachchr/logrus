//
// Adds a hook to
//

package logrus

import (
	"fmt"
	"strings"

	"github.com/go-stack/stack"
)

type FilenameAndLineHook struct {
	tag string
}

func NewFilenameAndLineHook(tag string) *FilenameAndLineHook {
	return &FilenameAndLineHook{tag: tag}
}

func (f FilenameAndLineHook) Levels() []Level {
	return AllLevels
}

func (f FilenameAndLineHook) Fire(e *Entry) error {
	cs := stack.Trace()

	for _, c := range cs {
		filename := fmt.Sprintf("%+s", c)

		if strings.HasPrefix(filename, "github.com/Sirupsen/logrus") ||
			strings.HasPrefix(filename, "github.com/einfachchr") ||
			!strings.HasSuffix(filename, ".go") {
			continue
		}
		e.Data[f.tag] = fmt.Sprintf("%s:%s", filename, fmt.Sprintf("%d", c))
		break
	}

	return nil
}
