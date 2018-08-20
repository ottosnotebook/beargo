package cmd

import (
	"os"
	"path/filepath"

	"github.com/gohugoio/hugo/deps"
	"github.com/gohugoio/hugo/hugofs"
	"github.com/gohugoio/hugo/hugolib"
	"github.com/spf13/jwalterweatherman"
	"github.com/spf13/viper"
)

// Hugo ...
type Hugo struct {
	root    string
	cfg     *viper.Viper
	sites   *hugolib.HugoSites
	rebuild chan bool
}

// NewHugoSite ...
func NewHugoSite(root string) (*Hugo, error) {
	v, err := loadHugoConfig(root)
	if err != nil {
		return nil, err
	}

	f := hugofs.NewDefault(v)

	siteOpts := deps.DepsCfg{
		Logger: jwalterweatherman.NewNotepad(jwalterweatherman.LevelError, jwalterweatherman.LevelError, os.Stdout, os.Stdout, "", 0),
		Fs:     f,
		Cfg:    v,
	}

	sites, err := hugolib.NewHugoSites(siteOpts)
	if err != nil {
		return nil, err
	}

	return &Hugo{
		root:    root,
		cfg:     v,
		sites:   sites,
		rebuild: make(chan bool),
	}, nil
}

func loadHugoConfig(root string) (*viper.Viper, error) {

	if !filepath.IsAbs(root) {
		cwd, err := os.Getwd()
		if err != nil {
			return nil, err
		}
		root = filepath.Join(cwd, root)
	}

	cfgOpts := hugolib.ConfigSourceDescriptor{
		Fs: hugofs.Os,
		/* Path: c.h.source, */
		WorkingDir: root,
		Filename:   filepath.Join(root, "config.yaml"),
	}

	config, _, err := hugolib.LoadConfig(cfgOpts)

	config.Set("workingDir", root)

	return config, err
}

// Build ...
func (h *Hugo) Build() error {

	err := copyChildren(filepath.Join(h.root, "static"), filepath.Join(h.root, "public"))
	if err != nil {
		return err
	}

	// you need to copy the static site
	return h.sites.Build(hugolib.BuildCfg{})
}
