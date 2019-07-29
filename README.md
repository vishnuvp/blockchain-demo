# Blockchain Demo

Largely or completely influenced by [Code your own blockchain in Go](https://medium.com/@mycoralhealth/code-your-own-blockchain-in-less-than-200-lines-of-go-e296282bcffc)

Display the blockchain in JSON format
```
curl -X GET \
   http://localhost:18080 \
   -H 'Content-Type: application/json' \
```
* Create a new block

```
curl -X POST \
   http://localhost:18080 \
   -H 'Content-Type: application/json' \
   -d '{"Data": 5}'
```