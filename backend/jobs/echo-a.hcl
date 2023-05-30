variables {
  version = "v0.0.0"
}

job "echo-a" {
  datacenters = ["phoenix"]
  namespace = "echo"

  group "echo" {
    count = 1
    task "server" {
      driver = "docker"

      config {
        image = "hashicorp/http-echo:latest"
        args  = [
          "-listen", ":19000",
          "-text", "${var.version}",
        ]
      }

      resources {
        network {
          mbits = 10
          port "http" {
            static = 19000
          }
        }
      }

      service {
        name = "echo-a"
        port = "http"

        check {
          type     = "http"
          path     = "/health"
          interval = "2s"
          timeout  = "2s"
        }
      }
    }
  }
}