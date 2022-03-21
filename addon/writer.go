package addon

import (
	"encoding/xml"
	"io/ioutil"
	"path"
	"sort"

	"github.com/maestre3d/gtavd/xmlutil"

	"github.com/maestre3d/gtavd/config"
	"github.com/maestre3d/gtavd/logging"
)

const (
	dlcListFileName = "dlclist.xml"
)

var (
	// hashSet used to avoid duplicated on-the-go
	dlcItems = map[string]struct{}{
		"platform:/dlcPacks/mpBeach/":        {},
		"platform:/dlcPacks/mpBusiness/":     {},
		"platform:/dlcPacks/mpChristmas/":    {},
		"platform:/dlcPacks/mpValentines/":   {},
		"platform:/dlcPacks/mpBusiness2/":    {},
		"platform:/dlcPacks/mpHipster/":      {},
		"platform:/dlcPacks/mpIndependence/": {},
		"platform:/dlcPacks/mpPilot/":        {},
		"platform:/dlcPacks/spUpgrade/":      {},
		"platform:/dlcPacks/mpLTS/":          {},
	}
)

type paths struct {
	Items []string `xml:"Item"`
}

type sMandatoryPacksData struct {
	XMLName xml.Name `xml:"SMandatoryPacksData"`
	Paths   paths    `xml:"Paths"`
}

func WriteAddons() error {
	if err := loadDlcItems(config.DefaultConfig.Addon.Dir); err != nil {
		return err
	}
	if err := loadDlcItems(config.DefaultConfig.Addon.ModDir); err != nil {
		return err
	}

	logging.Debug().Msgf("gtavd-addon-watcher: found %d addon directories", len(dlcItems))
	return writeAddonsFile()
}

func loadDlcItems(dlcPath string) error {
	readDir, err := ioutil.ReadDir(dlcPath)
	if err != nil {
		return err
	}

	for _, d := range readDir {
		if !d.IsDir() {
			continue
		}
		if _, ok := config.DefaultConfig.Addon.Blacklist[d.Name()]; !ok {
			key := "dlcpacks:/" + d.Name() + "/"
			dlcItems[key] = struct{}{}
		}
	}
	return nil
}

func newDlcList() sMandatoryPacksData {
	items := make([]string, 0, len(dlcItems))
	for i := range dlcItems {
		items = append(items, i)
	}
	sort.Slice(items, func(i, j int) bool {
		return items[j] < items[i]
	})
	return sMandatoryPacksData{
		Paths: paths{
			Items: items,
		},
	}
}

func writeAddonsFile() error {
	data := newDlcList()
	fileName := path.Join(config.DefaultConfig.DaemonPath, dlcListFileName)
	return xmlutil.WriteFile(fileName, data)
}
