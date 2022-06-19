# Consul Release Controller Waypoint Plugin
Waypoint plugin for Consul Release Controller

## Building the plugin

To build the plugin for local development

```
make build
```

then to install locally for testing

```
make install
```

## Building an On Demand Runner

An On Demand Runner image for Waypoint can be built using the foll

## Configuring Waypoint

The plugin requires the location of the release controller, this can be set using environment variables
for the runner profile.

```
-env-var="RELEASE_CONTROLLER_ADDR=http://127.0.0.1:8080" \
```

An example configuration can be seen below

```
waypoint runner profile set \
  -name=release-controller \
  -plugin-type=release-controller \
  -env-var="RELEASE_CONTROLLER_ADDR=http://10.5.0.12:8080" \
  -env-var="WAYPOINT_SERVER_ADDR=10.5.0.12:9701" \
  -env-var="WAYPOINT_SERVER_TLS=true" \
  -env-var="WAYPOINT_SERVER_TLS_SKIP_VERIFY=true" \
  -oci-url="nicholasjackson/consul-releaser-odr:0.1.0" 
```

```
waypoint project apply -runner-profile=release-controller hashicraft
```
