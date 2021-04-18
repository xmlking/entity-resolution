# Entity Resolution

entity resolution

## Start

```bash
docker compose up
# docker compose up elasticsearch
# open kibana UI
open http://localhost:5601/app/dev_tools#/console
```

## Test

```bash
curl -XGET 'http://localhost:9200'
http GET :9200
# het health
http GET :9200/_cluster/health
# list plugins
http GET :9200/_cat/plugins
```
