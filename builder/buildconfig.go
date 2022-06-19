package builder

type BuildConfig struct {
	Name     string    `json:"name"`
	Releaser *Releaser `hcl:"releaser,block" json:"releaser"`
	Runtime  *Runtime  `hcl:"runtime,block" json:"runtime"`
	Strategy *Strategy `hcl:"strategy,block" json:"strategy"`
	Monitor  *Monitor  `hcl:"monitor,block" json:"monitor"`
}

type Releaser struct {
	PluginName string          `hcl:"plugin_name" json:"plugin_name"`
	Config     *ReleaserConfig `hcl:"config,block" json:"config"`
}

type ReleaserConfig struct {
	ConsulService string `hcl:"consul_service" json:"consul_service"`
}

type Runtime struct {
	PluginName string         `hcl:"plugin_name" json:"plugin_name"`
	Config     *RuntimeConfig `hcl:"config,block" json:"config"`
}

type RuntimeConfig struct {
	Deployment string `hcl:"deployment" json:"deployment"`
	Namespace  string `hcl:"namespace,optional" json:"namespace,omitempty"`
}

type Strategy struct {
	PluginName string          `hcl:"plugin_name" json:"plugin_name"`
	Config     *StrategyConfig `hcl:"config,block" json:"config"`
}

type StrategyConfig struct {
	Interval       string `hcl:"interval,optional" json:"interval"`
	InitialDelay   string `hcl:"initial_delay,optional" json:"initial_delay"`
	InitialTraffic int    `hcl:"initial_traffic,optional" json:"initial_traffic"`
	TrafficStep    int    `hcl:"traffic_step,optional" json:"traffic_step"`
	MaxTraffic     int    `hcl:"max_traffic,optional" json:"max_traffic"`
	ErrorThreshold int    `hcl:"error_threshold,optional" json:"error_threshold"`
}

type Monitor struct {
	PluginName string         `hcl:"plugin_name" json:"plugin_name"`
	Config     *MonitorConfig `hcl:"config,block" json:"config"`
}

type MonitorConfig struct {
	Address string  `hcl:"address" json:"address"`
	Queries []Query `hcl:"query,block" json:"query"`
}

type Query struct {
	Name   string `hcl:"name" json:"name"`
	Preset string `hcl:"preset" json:"preset"`
	Min    int    `hcl:"min" json:"min"`
	Max    int    `hcl:"max,optional" json:"max"`
}
