package addon

import (
	"context"

	"github.com/fsnotify/fsnotify"
	"github.com/maestre3d/gtavd/config"
	"github.com/maestre3d/gtavd/logging"
)

func Watch(ctx context.Context) {
	w, err := fsnotify.NewWatcher()
	if err != nil {
		logging.Error().Err(err).Msg("gtavd: failed to start addon-watcher module")
		return
	}
	defer func() {
		if err = w.Close(); err != nil {
			logging.Error().Err(err).Msg("gtavd-addon-watcher: failed to close os watcher")
		}
	}()
	for {
		if w != nil {
			modsPath := config.DefaultConfig.Addon.Dir
			removeWatcher(w, modsPath)
			if err = w.Add(modsPath); err != nil {
				logging.Error().
					Err(err).
					Str("path", modsPath).
					Msg("gtavd: failed to start addon file system watcher")
			} else {
				go listenFsEvents(w, modsPath)
			}
		}
		select {
		case <-config.OnConfigChange:
			continue
		case <-ctx.Done():
			return
		}
	}
}

func removeWatcher(w *fsnotify.Watcher, path string) {
	_ = w.Remove(path)
}

func listenFsEvents(w *fsnotify.Watcher, filePath string) {
	writeAddonsFromWatcher()
	logging.Info().
		Str("path", filePath).
		Msg("gtavd-addon-watcher: listening to addon directory")
	for e := range w.Events {
		logging.Debug().Str("event", e.String()).Msg("gtavd-addon-watcher: received fs event")
		writeAddonsFromWatcher()
	}
}

func writeAddonsFromWatcher() {
	if err := WriteAddons(); err != nil {
		logging.Error().Err(err).Msg("gtavd-addon-watcher: failed to write addons list")
	} else {
		logging.Info().Msg("gtavd-addon-watcher: dlclist.xml file was written")
	}
}
