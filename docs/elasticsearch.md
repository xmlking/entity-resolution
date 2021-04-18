# Elasticsearch

## Verify

```bash
curl -XGET 'http://localhost:9200'
http GET :9200
# get health
http GET :9200/_cluster/health
# list plugins
http GET :9200/_cat/plugins
# list indices
http GET  /_cat/indices
```

## Console

### Index

```bash
GET /_cat/health?v=true
# create
PUT /my-index-000001
# get
GET /my-index-000001
# list
GET /_cat/indices
# Retrieve settings
GET /my-index-000001/_settings
```

### Mapping

```bash
# Retrieve
GET /my-index-000001/_mapping
```

### Ingest

```bash
PUT /my-index-000001/_doc/1
{
  "name": "John Doe"
}

GET /my-index-000001/_doc/1
```

### Mapping

```bash
```

### Cleanup

```bash
# cleanup 
DELETE /my-index-000001
```
