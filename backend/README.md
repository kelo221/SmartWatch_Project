# STARTING THE SERVER
## REQUIREMENTS
* PostgreSQL
* Golang 1.19.2 or newer


## INIT Database

```bash
psql.exe -U postgres
CREATE DATABASE smartwatch_project; \q
psql.exe -U postgres -d smartwatch_project -f db/schema.sql
```
* OPTIONAL
    * (for backend developement generate ORM: sqlboiler psql) 

After the database has been initialized you may start the backend.

```bash
cd backend
go run .
```


# ROUTES

## JSON

### api/login (POST)
   INPUT
```json
   {"username": "test",	"password": "test"}
```
   OUTPUT
```json
   {"token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJRCI6MSwiZXhwIjoxNjY2MjEwNzQ1LCJ1c2VybmFtZSI6InRlc3QifQ.8PWstjzN8SD5lORNWAmAh1NBCGVL9GugE67rtzJN1NA"}
```

### /api/register (POST)
   INPUT
```json
  {"username": "test",	"password": "test","password_confirm": "test"}
```
   OUTPUT
```json
 {"message": "Registered!"}
```


 # SECURED ROUTES, BEARER NEEDED

## JSON

### /api/user/event (GET)

   OUTPUT
```json
[{"id":1,"eventname":"newUnix","eventtime":"1970-01-01T02:20:33Z","created_at":"2022-10-18T13:26:11.536961Z","userid":1,"issilent":false,"snoozedisabled":false},{"id":3,"eventname":"newUnix","eventtime":"1970-01-01T05:25:34Z","created_at":"2022-10-18T13:26:21.909119Z","userid":1,"issilent":true,"snoozedisabled":true}]
```
### /api/user/event (POST)
   INPUT
```json
{"eventName": "newUnix", "eventTime": "1665997920","isSilent": "true","SnoozeDisabled": "false"}
```
   OUTPUT
```json
{"message": "New event Added!"}
```

### /api/user/eventDate (DELETE)
   INPUT
```json
{
	"eventDate": "1970-01-01 05:25:34"
}
```
   OUTPUT
```json
{"message": "Event Removed!"}
```

### /api/user/event (PATCH)
   INPUT
```json
{
	"eventName": "newUnixMod",
	"eventTimeOld": "1970-01-01 00:02:03",
	"eventTimeNewUnix": "1000",
	"isSilent": "true",
	"SnoozeDisabled": "true"
}
```
   OUTPUT
```json
{"message": "Event time updated!"}
```

<br/>

## PROTOBUFFER

[Decode results](https://protobuf-decoder.netlify.app/)

### /api/proto/event/ (GET)

   OUTPUT
```
2220080112076e6577556e69781a0308f141220c08f3d2ba9a0610e8bf85800228012225080312076e6577556e69781a0408ce9801220c08fdd2ba9a0610989cc0b103280130013801
```

### /api/proto/debug/ (GET)
Returns an example protobuffer object.

   OUTPUT
```
0806120b6a2e4576656e746e616d651a02080122020801280b30013801
```


### /api/proto/event/:protobufferHex (POST)
   INPUT
```
http://localhost:8000/api/proto/event/0806120b6a2e4576656e746e616d651a02080122020801280b30013801
```
   OUTPUT
```json
{"message": "New event Added!"}
```

### /api/proto/event/:protobufferHex (DELETE)
   INPUT
```
http://localhost:8000/api/proto/event/0806120b6a2e4576656e746e616d651a02080122020801280b30013801
```
   OUTPUT
```json
{"message": "Event Removed!"}
```


### /api/proto/event/:protobufferHex (PATCH)
   INPUT
```
http://localhost:8000/api/proto/event/0806120b6a2e4576656e746e616d651a02080122020801280b30013801
```
   OUTPUT
```json
{"message": "Event time updated!"}
```

