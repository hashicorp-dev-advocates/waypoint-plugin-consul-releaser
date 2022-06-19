package builder

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/hashicorp/waypoint-plugin-sdk/component"
	"github.com/hashicorp/waypoint-plugin-sdk/terminal"
)

type Builder struct {
	config BuildConfig
}

// Implement Configurable
func (b *Builder) Config() (interface{}, error) {
	return &b.config, nil
}

// Implement ConfigurableNotify
func (b *Builder) ConfigSet(config interface{}) error {
	_, ok := config.(*BuildConfig)
	if !ok {
		// The Waypoint SDK should ensure this never gets hit
		return fmt.Errorf("Expected *BuildConfig as parameter")
	}

	return nil
}

// Implement Builder
func (b *Builder) BuildFunc() interface{} {
	// return a function which will be called by Waypoint
	return b.build
}

func (b *Builder) build(ctx context.Context, ui terminal.UI, ji *component.JobInfo) (*Binary, error) {

	u := ui.Status()
	defer u.Close()
	u.Update("Building application")

	// set the name
	b.config.Name = ji.App

	releaserURL := os.Getenv("CONSUL_RELEASER_URL")
	jsonData, err := json.Marshal(b.config)
	if err != nil {
		u.Step(terminal.StatusError, "JSON marshal failed")

		return nil, err
	}

	resp, err := http.Post(releaserURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		u.Step(terminal.StatusError, "POST to CONSUL_RELEASER_URL failed")

		return nil, err
	}

	bconv, err := io.ReadAll(resp.Body)
	if err != nil {
		u.Step(terminal.StatusError, "Could not read JSON")

		return nil, err
	}

	u.Step(terminal.StatusOK, string(bconv))

	return &Binary{
		Location: "<placeholder>",
	}, nil
}
