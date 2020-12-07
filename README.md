# Golang Microservices

## Curls

curl -v -d 'Bob' localhost:9006
curl -v localhost:9006/ -d '{"name": "tea", "description":"a nice cup of tea"}' | jq
curl -v localhost:9006/3 -XPUT -d '{"name": "tea", "description":"a nice cup of tea"}' | jq
curl localhost:9006 | jq (brew install jq)
