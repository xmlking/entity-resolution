# Entity Resolution

entity resolution

## Overview

Schema

<p align="center">
  <img src="docs/images/schema.drawio.svg" width="60%">
</p>
<!-- 
![My Diagram](docs/images/schema.drawio.svg) -->

### TODO

- [x] Start/Stop ES with docker-compose
- [ ] Generate test data
- [ ] Batch load test data with [benthos](https://www.benthos.dev/)
- [ ] 

## Getting Started

### Start

```bash
docker compose up
# docker compose up elasticsearch
# open kibana UI
open http://localhost:5601/app/dev_tools#/console
```

### Stop

```bash
docker compose down
```

## Reference

- [zentity](https://zentity.io/)
- [real-time-entity-resolution](https://www.slideshare.net/o19s/real-time-entity-resolution-with-elasticsearch-haystack-2018)
- [zentity-docs](https://zentity.io/docs/basic-usage/)
- [Elasticsearch as a Graph Database](https://medium.com/@imriqwe/elasticsearch-as-a-graph-database-bc0eee7f7622)