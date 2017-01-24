#

## 1) consul-rundeck-resource-writer

Put the binary in your `$PATH`

```bash
CONSUL_HTTP_ADDR=consul.service.consul:8500 CONFIG_FILE=/etc/rundeck/consul-node-resources.yaml consul-rundeck-resource-writer
# or if you got consul running on 127.0.0.1:8500
CONFIG_FILE=/etc/rundeck/consul-node-resources.yaml consul-rundeck-resource-writer
```

`CONSUL_HTTP_ADDR` is optional and defaults to `127.0.0.1:8500`
`CONFIG_FILE` is optiona and defaults to `resources.yaml` in `$PWD`.

## 2) rundeck (debian)

In `/var/rundeck/projects/demo_project/etc/project.properties` put something similar to this

```
resources.source.2.config.file=/etc/rundeck/consul-node-resources.yaml
resources.source.2.config.format=resourceyaml
resources.source.2.config.generateFileAutomatically=true
resources.source.2.config.includeServerNode=true
resources.source.2.config.requireFileExists=false
resources.source.2.type=file
```

## 3) finish up

Reload Rundeck
