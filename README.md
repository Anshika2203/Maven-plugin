# Maven-plugin

A [Drone](https://drone.io) plugin to build Java applications using [Apache Maven](https://maven.apache.org).

## Plugin Image

The plugin `anshikaanand/maven-plugin` is available for the following architectures:

| OS            | Tag                                |
| ------------- | ---------------------------------- |
| linux/amd64   | `linux-amd64`                      |
| linux/arm64   | `linux-arm64`                      |
| windows/amd64 | `windows-amd64`                    |

## Examples

```
# Plugin YAML
- step:
    type: Plugin
    name: maven-plugin-arm64
    identifier: maven-plugin-arm64
    spec:
        connectorRef: harness-docker-connector
        image: anshikaanand/maven-plugin:linux-arm64
       

- step:
    type: Plugin
    name: maven-plugin-amd64
    identifier: maven-plugin-amd64
    spec:
        connectorRef: harness-docker-connector
        image: anshikaanand/maven-plugin:linux-amd64
        