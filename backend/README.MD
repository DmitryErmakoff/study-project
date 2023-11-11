#sneaker service 

#REST API

GET /sneakers -- list of sneakers -- 200, 404, 500 \
GET /sneakers/:id -- sneaker by id -- 200, 404, 500 \
POST /sneakers/:id -- create sneaker -- 204, 4xx, Header Location: url \
PUT /sneakers/:id -- fully update sneaker -- 204/200, 404, 400, 500 \
PATCH /sneakers/:id -- partially update sneaker -- 204/200, 404, 400, 500 \
DELETE /sneakers/:id -- delete sneaker by id -- 204, 404, 400 \
