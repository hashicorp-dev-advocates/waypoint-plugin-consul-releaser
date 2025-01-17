project = "test"

app "payments" {
  build {
    use "consul-release-controller" {
      releaser {
        plugin_name = "consul"

        config {
          consul_service = "payments"
        }
      }

      runtime {
        plugin_name = "nomad"

        config {
          deployment = "payments-deployment"
        }
      }

      strategy {
        plugin_name = "canary"

        config {
          initial_delay   = "30s"
          interval        = "30s"
          initial_traffic = 10
          traffic_step    = 20
          max_traffic     = 100
          error_threshold = 5
        }
      }

      monitor {
        plugin_name = "prometheus"

        config {
          address = "http://localhost:9090"

          query {
            name   = "request-success"
            preset = "envoy-request-success"
            min    = 99
          }

          query {
            name   = "request-duration"
            preset = "envoy-request-duration"
            min    = 20
            max    = 200
          }
        }
      } 
    }
  }

  deploy {
    use "consul-release-controller" { }
  }
}
