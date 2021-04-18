# Elasticsearch Data Modeling

## Schema

<p align="center">
  <img src="images/schema.drawio.svg" width="60%">
</p>

## Tips

- Problem with Elasticsearch is, it does not support ACID transactions. Changes to individual documents are *ACIDic*, but not changes involving multiple documents.
- use **bulk api** to update multiple types of documents in one transaction.  
- Elasticsearch provides two concepts that can help with denormalization:
  - Nested Document /Query
  - Parent & Child Relationship.
- Defines a parent/child relationship for documents in the same index.
- Use [join](https://www.elastic.co/guide/en/elasticsearch/reference/current/parent-join.html) field type for **associate** and **disassociate** related docs
- As we wanted all of the properties to be searchable — each property was configured as `index: not_analyzed`.
- We set **docvalues** and **_all** to **false**.

## Reference

- [Elasticsearch as a Graph Database](https://medium.com/@imriqwe/elasticsearch-as-a-graph-database-bc0eee7f7622)
- [Handling Relationships Using Nested Objects in Elasticsearch](https://qbox.io/blog/handling-relationships-using-nested-objects-elasticsearch)
- [Handling Data Denormalization Issues in Elasticsearch](https://qbox.io/blog/handling-data-denormalisation-issues-in-elasticsearch/)
