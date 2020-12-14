# Golang Microservices

## Curls

curl -v -d 'Bob' localhost:9006
curl -v localhost:9006/ -d '{"name": "tea", "Price": 1.23, "description":"a nice cup of tea", "SKU": "abc-def-ghi"}' | jq
curl -v localhost:9006/1 -XPUT -d '{"name": "tea", "description":"a nice cup of tea"}' | jq
curl localhost:9006 | jq (brew install jq)
