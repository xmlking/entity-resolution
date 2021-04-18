# Entity Resolution

entity resolution

## Overview

Schema

<p align="center">
  <img src="docs/images/er.drawio.svg" width="60%">
</p>

<!-- ![overview](https://s3-eu-west-1.amazonaws.com/graphaware/assets/graphAidedSearchIntro2.png) -->
<!-- 
![My Diagram](docs/images/schema.drawio.svg) -->

Links: [Elasticsearch](docs/elasticsearch.md), [Data Modeling](docs/modeling.md), [Kibana Console](/docs/console.md)

### TODO

- [x] Start/Stop ES with docker-compose
- [ ] Generate test data
- [ ] Batch load test data with [benthos](https://www.benthos.dev/)

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
- [GraphAware Graph-Aided Search](https://github.com/graphaware/graph-aided-search)
- [Graph-Powered Search: Neo4j & Elasticsearch](https://graphaware.com/assets/graphpoweredsearch-neo4j-elasticsearch.pdf)
