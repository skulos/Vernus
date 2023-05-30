variables {
  version = "v0.0.0"
}

job "echo-b" {
  datacenters = ["phoenix"]
  namespace = "echo"

  group "echo" {
    count = 1
    task "server" {
      driver = "docker"

      config {
        image = "hashicorp/http-echo:latest"
        args  = [
          "-listen", ":18000",
          "-text", "${var.version}",
        ]
      }

      resources {
        network {
          mbits = 10
          port "http" {
            static = 18000
          }
        }
      }

      service {
        name = "echo-b"
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