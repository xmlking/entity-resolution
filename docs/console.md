

```cURL
# _cat for humans
GET _cat/plugins?v=true
GET _cat/health?v=true
GET _cat/nodes?v=true&h=ip,port,heapPercent,name
GET _cat/nodes?format=json&pretty
GET _cat/indices
GET _cat/indices?pretty

# _cluster
GET _cluster/stats
GET _cluster/health

# _search
GET _search
{
  "query": {
    "match_all": {}
  }
}

# demo
GET /kibana_sample_data_ecommerce

GET /kibana_sample_data_ecommerce/_settings

GET /kibana_sample_data_ecommerce/_mapping

GET kibana_sample_data_ecommerce/_search
{
  "query": {
    "match_all": {}
  }
}

```
