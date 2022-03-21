package global

import (
	"os"
	"sync"

	"github.com/maestre3d/gtavd/fsutil"

	"github.com/rs/zerolog/log"
)

var (
	DefaultPath     string
	defaultPathOnce = sync.Once{}
)

func init() {
	defaultPathOnce.Do(func() {
		userPath, _ := os.UserHomeDir()
		DefaultPath = userPath + "/.gtavd"
		if err := fsutil.TryCreatePath(DefaultPath); err != nil {
			log.Err(err).Msg("gtavd: failed to create gtavd path")
		}
	})
}
