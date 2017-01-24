#

## 1) consul

In file `/etc/consul/watch_nodes-rundeck-nodes.json` put something similar to this, assuming the binary of `consul-rundeck-resource-writer` exist in `/opt/`

```json
{
  "watches": [
    {
      "type": "nodes",
      "handler": "CONFIG_FILE=/var/rundeck/projects/demo_project/etc/resources.yaml /opt/consul-rundeck-resource-writer"
    }
  ]
}
```

## 2) rundeck (debian)

In `/var/rundeck/projects/demo_project/etc/project.properties` put something similar to this

```
resources.source.2.config.file=/var/rundeck/projects/demo_project/etc/resources.yaml
resources.source.2.config.format=resourceyaml
resources.source.2.config.generateFileAutomatically=true
resources.source.2.config.includeServerNode=true
resources.source.2.config.requireFileExists=false
resources.source.2.type=file
```

## 3) finish up

Reload Rundeck
