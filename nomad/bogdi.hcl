job "webserver" {
  datacenters = ["lux"]
  type = "service"

  group "webserver" {
    count = 3
    network {
      port "http" {
        to = 80
      }
    }

    service {
      name = "bogdi-blog-service"
      tags = [
        "traefik.enable=true",
        "traefik.http.routers.bogdi-blog.rule=Host(`bogdi.xyz`)",
        "traefik.http.routers.bogdi-blog.entrypoints=http"
      ]
      port = "http"
      check {
        name     = "alive"
        type     = "http"
        path     = "/"
        interval = "10s"
        timeout  = "2s"
      }
    }

    restart {
      attempts = 2
      interval = "30m"
      delay = "15s"
      mode = "fail"
    }

    task "bogdi-blog-container" {
      driver = "docker"
      config {
        image = "ghcr.io/gofort/bogdi:{ IMAGE_TAG }"
        ports = ["http"]
      }
    }
  }
}
