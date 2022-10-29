# Redis

My Hosted Redis https://app.redislabs.com/#/databases

## Operations

### Prerequisites

```shell
brew install redis
brew install redis-developer/tap/riot-file
brew install redis-developer/tap/riot-gen
```

### Setup

```shell
# start local redis
docker compose -f infra/redis.yml up
# stop local redis before restart again
docker compose -f infra/redis.yml down
# this will stop redis and remove all volumes
docker compose -f infra/redis.yml down -v 
# ssh to grafana
docker-compose exec grafana /bin/bash
```

#### Redis Insight

Redis db visualization dashboard

Download desktop Redis Insight App from [redis.com](https://redis.com/redis-enterprise/redis-insight/)
After opening app, connect using host:localhost, port:6379 and name:any_name_is_ok

#### Redis Grafana

For the first time, enable **Redis-Application** plugin at:

http://localhost:3000/plugins/redis-app/

### Schema

### Queries

Basics 

```shell
PING Marco!
MODULE LIST

# watch all command send to redis
MONITOR
# redis-cli MONITOR
```

```shell
KEYS *

HGETALL people:1750384707
HGETALL address:92a87e94-dbf0-44ff-a755-f7afc04116a8

# Delete all keys of the currently selected Redis database:
FLUSHDB

# Getting data from Streams
XRANGE events - +
```

<details>
  <summary>RedisGraph</summary>

```shell
# https://github.com/redis/jedis/blob/master/src/test/java/redis/clients/jedis/modules/graph/GraphAPITest.java
CREATE (:human{name:'danny',age:12})
CREATE (:person{name:'roi',age:32})
CREATE (:person{name:'amit',age:30})
MATCH (a:person), (b:person) WHERE (a.name = 'roi' AND b.name='amit')  CREATE (a)-[:knows]->(b)
MATCH (a:person) WHERE (a.name = 'roi') DELETE a
CREATE (:person{name:'roi',age:32})
MATCH (a:person), (b:person) WHERE (a.name = 'roi' AND b.name='amit')  CREATE (a)-[:knows]->(a)
MATCH (a:person) WHERE (a.name = 'roi') DELETE a
CREATE (:person{name:'roi',age:32})
CREATE (:person{name:'amit',age:30})
MATCH (a:person), (b:person) WHERE (a.name = 'roi' AND b.name='amit')  CREATE (a)-[:knows]->(a)
MATCH (a:person)-[e]->() WHERE (a.name = 'roi') DELETE e
CREATE (:person{name:'roi',age:32})
CREATE INDEX ON :person(age)
CREATE INDEX ON :person(age1)
DROP INDEX ON :person(age1)

# array support
CREATE (:person{name:'a',age:32,array:[0,1,2]})
WITH [0,1,2] as x return x
MATCH(n) return collect(n) as x
CREATE (:restaurant {location: point({latitude:30.27822306, longitude:-97.75134723})})
MATCH (restaurant) RETURN restaurant
UNWIND range(0,100) AS x WITH x AS x WHERE x = 100 RETURN x
MATCH (n:N {val:$val}) RETURN n.val

# Full-text search now included
# http://www.odbms.org/2020/04/introducing-redisgraph-2-0/
CALL db.idx.fulltext.createNodeIndex('Person','name')
CALL db.idx.fulltext.queryNodes('Person','Bob') YIELD node AS p RETURN p.name
# linked search
CALL db.idx.fulltext.queryNodes('Person','%yif%') YIELD node
MATCH p=(node)-[:CONNECTED*1..3]->(:Person {name:'Pieter'})
RETURN p.name, p.title, length(p) AS connectedness
ORDER BY connectedness ASC LIMIT 20

MATCH (me:Person)-->(friend:Person)-->(fof:Person)
WHERE me.name='Pieter'
RETURN friend.name, fof.name

MATCH (me:Person)-->(friend:Person)-[f:FRIEND]->(fof:Person)
WHERE me.name='Pieter'
RETURN friend, f, fof

# We can query for suggested introductions in the left-hand graph by following the “KNOWS” relationship:
MATCH (b:Person)-[:KNOWS]->(a:Person)-[:KNOWS]->(c:Person)
WHERE NOT (b)-[:KNOWS]-(c) AND b <> c
RETURN b, c

MATCH (b:Person)<-[:MANAGES]-(a:Person)-[:MANAGES]->(c:Person)
WHERE NOT (b)-[:PEER]-(c) AND b <> c
WITH b, c,
CASE b.level
WHEN 'Director' THEN 'mentor'
ELSE 'peer'
END AS role
MERGE (b)-[:PEER {role:role}]->(c)
```

```shell
# To create a node with multiple labels, you simply list all the labels separated by a colon:

GRAPH.QUERY g "CREATE (e:Employee:BoardMember {Name: 'Vincent Chan', Title: 'Web marketing lead'}) return e"

# To match a node with multiple labels (AND condition), you should use the same colon notation as well:

GRAPH.QUERY g "MATCH (e:Employee:BoardMember) return e"
```

To construct a `full-text index` on the title property of movies using the German language and using custom stopwords of all nodes with the label Movie:

More [detains](https://redis.com/blog/redisgraph-2-8-is-generally-available/)

```shell
GRAPH.QUERY DEMO_GRAPH "CALL db.idx.fulltext.createNodeIndex({ label: 'Movie', language: 'German', stopwords: ['a', 'ab'] }, 'title')"
```

RediSearch provides 3 additional field configuration options:

* weight – the importance of the text in the field
* nostem – skip stemming when indexing text
* phonetic – enable phonetic search on the text

To construct a full-text index on the title property with a phonetic search of all nodes with the label Movie:
```shell
GRAPH.QUERY DEMO_GRAPH "CALL db.idx.fulltext.createNodeIndex('Movie', {field: 'title', phonetic: 'dm:en'})"
```

</details>

### Data Loading

Sample Data https://github.com/redis-developer/redis-datasets/tree/master/user-database

```shell
redis-cli -h localhost -p 6379 < apps/entity-service/data/employee.redis
```

## Tools

## Reference
- Redis Developer Community [Projects](https://github.com/redis-developer)
- [RIOT: Redis Input/Output Tools](https://github.com/redis-developer/riot)
- [Terraform Provider Redis Cloud](https://github.com/RedisLabs/terraform-provider-rediscloud)
- Command-line utility for load generation and bechmarking NoSQL key-value databases[memtier_benchmark](https://github.com/RedisLabs/memtier_benchmark)

- **Redis Search**
    - Fast and Furious: Searching in a Distributed World with Highly Available, [Spring Data Redis](https://www.youtube.com/watch?v=QZdUXrzdxos)
    - Source Code on [GitHub](https://github.com/Redislabs-Solution-Architects/rediscogs)
    - [RediSearch](https://volkovlabs.com/i-taught-my-wife-how-to-use-redisearch-2-0-77d6f32660df) with **NPI** and **CommonWell Health Alliance** data, and  [Redis Application plug-in for Grafana](https://grafana.com/grafana/plugins/redis-datasource/)
    - [Real-time observability with Redis and Grafana](https://docs.google.com/presentation/d/1dt4lduof6qIZF1dJ8Sv4_sCjKYHBY_a5ODAVQSEANgE/edit#slide=id.g9bf045ab42_0_40)

- **Redis Graph**
    - RedisGraph [commands](https://oss.redislabs.com/redisgraph/commands/)
    - [RedisGraph bulk loader](https://github.com/RedisGraph/redisgraph-bulk-loader)
    - [search-graph-demo](https://github.com/stockholmux/conf19-search-graph-demo)
    - [Redis and Spring: Building High Performance RESTful APIs](https://github.com/wilvdb/redi2read/blob/main/src/main/kotlin/com/redislabs/edu/redi2read/services/RecommendationService.kt)
    - jedis [Graph SDK](https://github.com/redis/jedis/tree/master/src/main/java/redis/clients/jedis/graph)
      jedis [Graph SDK Examples](https://github.com/redis/jedis/blob/master/src/test/java/redis/clients/jedis/modules/graph/GraphAPITest.java)
    - Building a Recommendation Engine with RedisGraph[video](https://www.youtube.com/watch?v=ZLJ-3DJVufw)
    - Dungeons, Dragons, and Graph Databases - Guy Royse - NDC London 2022 - [Source Code](https://github.com/guyroyse/dnd-and-graph-databases) and [Video](https://www.youtube.com/watch?v=HqwY_TyxeJw)


- **Redis Spring**
    - [Spring Data Redis](https://docs.spring.io/spring-data/redis/docs/2.5.3/reference/html/#why-spring-redis)
    - [Accessing Data Reactively with Redis](https://spring.io/guides/gs/spring-data-reactive-redis/)
    - [Redis Reactive kotlin tests](https://github.com/spring-projects/spring-data-redis/blob/main/src/test/kotlin/org/springframework/data/redis/core/ReactiveHashOperationsExtensionsUnitTests.kt)
    
- **Redis Quarkus**
    - QUARKUS - USING THE [REDIS CLIENT](https://quarkus.io/guides/redis)

- **Example Projects**
    - [Microservices with Redis](https://github.com/redis-developer/redis-microservices-demo)
        - Has Streams, Graph and Search
        - Deploy the application on [Kubernetes](https://github.com/redis-developer/redis-microservices-demo/tree/master/kubernetes)

- **Redis Datasets**
   - Sample [Dataset](https://github.com/redis-developer/redis-datasets)
