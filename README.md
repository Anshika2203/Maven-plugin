# Maven-plugin

A [Drone](https://drone.io) plugin to build Java applications using [Apache Maven](https://maven.apache.org).

## Plugin Image

The plugin `anshikaanand/maven-plugin` is available for the following architectures:

| OS            | Tag                                |
| ------------- | ---------------------------------- |
| linux/amd64   | `linux-amd64`                      |
| linux/arm64   | `linux-arm64`                      |
| windows/amd64 | `windows-amd64`                    |


## Parameters

| Parameter                                                                        | Comments                                                                                                                                  |
|:---------------------------------------------------------------------------------|-------------------------------------------------------------------------------------------------------------------------------------------|
| context_dir <span style="font-size: 10px"><br/>`optional`</span>                 | The context directory within the source repository where `pom.xml` is found to execute the maven goals. Defaults to Drone workspace root. |
| goals <span style="font-size: 10px"><br/>`optional`</span>                       | An array of maven goals to run.Defaults: `-DskipTests clean install`.                                                                     |
| maven_modules <span style="font-size: 10px"><br/>`optional`</span>               | An array of maven modules to be built incase of a multi module maven project.                                                             |
| maven_mirror_url <span style="font-size: 10px"><br/>`optional`</span>            | The Maven repository mirror url.                                                                                                          |
| server_user <span style="font-size: 10px"><br/>`optional`</span>                 | The username for the maven repository manager server.                                                                                     |
| server_password <span style="font-size: 10px"><br/>`optional`</span>             | The password for the maven repository manager server.                                                                                     |
| proxy_user <span style="font-size: 10px"><br/>`optional`</span>                  | The username for the proxy server.                                                                                                        |
| proxy_password <span style="font-size: 10px"><br/>`optional`</span>              | The password for the proxy server.                                                                                                        |
| proxy_port <span style="font-size: 10px"><br/>`optional`</span>                  | Port number for the proxy server.                                                                                                         |
| proxy_host <span style="font-size: 10px"><br/>`optional`</span>                  | Proxy server Host.                                                                                                                        |
| proxy_non_proxy_hosts <span style="font-size: 10px"><br/>`optional`</span>       | An array of non proxy server hosts.                                                                                                       |
| proxy_protocol <span style="font-size: 10px"><br/>`optional`</span>              | Protocol for the proxy ie http or https.                                                                                                  |


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
        