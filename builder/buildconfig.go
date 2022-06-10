package builder

type BuildConfig struct {
	Name     string `hcl:"name,optional"`
	Releaser struct {
		PluginName string `hcl:"plugin_name,optional"`
		Config     struct {
			ConsulService string `hcl:"consul_service, optional"`
		} `hcl:"config,optional"`
	} `hcl:"releaser,optional"`
	Runtime struct {
		PluginName string `hcl:"plugin_name,optional"`
		Config     struct {
			Deployment string `hcl:"deployment,optional"`
			Namespace  string `hcl:"namespace,optional"`
		} `hcl:"config,optional"`
	} `hcl:"runtime,optional"`
	Strategy struct {
		PluginName string `hcl:"plugin_name,optional"`
		Config     struct {
			Interval       string `hcl:"interval,optional"`
			InitialTraffic int    `hcl:"initial_traffic,optional"`
			TrafficStep    int    `hcl:"traffic_step,optional"`
			MaxTraffic     int    `hcl:"max_traffic,optional"`
			ErrorThreshold int    `hcl:"error_threshold,optional"`
		} `hcl:"config,optional"`
	} `hcl:"strategy,optional"`
	Monitor struct {
		PluginName string `hcl:"plugin_name,optional"`
		Config     struct {
			Address string `hcl:"address,optional"`
			Queries []struct {
				Name   string `hcl:"name",optional`
				Preset string `hcl:"preset,optional"`
				Min    int    `hcl:"min,optional"`
				Max    int    `hcl:"max,omitempty,optional"`
			} `hcl:"queries,optional"`
		} `hcl:"config,optional"`
	} `hcl:"monitor,optional"`
}
