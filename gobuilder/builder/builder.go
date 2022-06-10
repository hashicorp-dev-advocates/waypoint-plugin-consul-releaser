package builder

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/hashicorp/waypoint-plugin-sdk/terminal"
)

//type BuildConfig struct {
//OutputName string `hcl:"output_name,optional"`
//Source     string `hcl:"source,optional"`
//}

type Builder struct {
	config BuildConfig
}

// Implement Configurable
func (b *Builder) Config() (interface{}, error) {
	return &b.config, nil
}

// Implement ConfigurableNotify
func (b *Builder) ConfigSet(config interface{}) error {
	c, ok := config.(*BuildConfig)
	if !ok {
		// The Waypoint SDK should ensure this never gets hit
		return fmt.Errorf("Expected *BuildConfig as parameter")
	}

	_, err := os.Stat(c.Name)
	if err != nil {
		return fmt.Errorf("App name does not exist.")
	}

	return nil
}

// Implement Builder
func (b *Builder) BuildFunc() interface{} {
	// return a function which will be called by Waypoint
	return b.build
}

// A BuildFunc does not have a strict signature, you can define the parameters
// you need based on the Available parameters that the Waypoint SDK provides.
// Waypoint will automatically inject parameters as specified
// in the signature at run time.
//
// Available input parameters:
// - context.Context
// - *component.Source
// - *component.JobInfo
// - *component.DeploymentConfig
// - hclog.Logger
// - terminal.UI
// - *component.LabelSet
//
// The output parameters for BuildFunc must be a Struct which can
// be serialzied to Protocol Buffers binary format and an error.
// This Output Value will be made available for other functions
// as an input parameter.
// If an error is returned, Waypoint stops the execution flow and
// returns an error to the user.
func (b *Builder) build(ctx context.Context, ui terminal.UI) (*Binary, error) {
	u := ui.Status()
	defer u.Close()
	u.Update("Building application")

	if b.config.Name == "" {
		b.config.Name = "app"
	}

	releaser_url := os.Getenv("CONSUL_RELEASER_URL")
	json_data, err := json.Marshal(b.config)
	if err != nil {
		u.Step(terminal.StatusError, "JSON marshal failed")

		return nil, err
	}

	resp, err := http.Post(releaser_url, "application/json", bytes.NewBuffer(json_data))
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
