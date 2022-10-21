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

## api/login (POST)
   INPUT
```json
   {"username": "test",	"password": "test"}
```
   OUTPUT
```json
   {"token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJRCI6MSwiZXhwIjoxNjY2MjEwNzQ1LCJ1c2VybmFtZSI6InRlc3QifQ.8PWstjzN8SD5lORNWAmAh1NBCGVL9GugE67rtzJN1NA"}
```

## /api/register (POST)
   INPUT
```json
  {"username": "test",	"password": "test","password_confirm": "test"}
```
   OUTPUT
```json
 {"message": "Registered!"}
```


 # SECURED ROUTES, BEARER NEEDED


## /api/user/event (POST)
   INPUT
```json
{"eventName": "newUnix", "eventTime": "1665997920","isSilent": "true","SnoozeDisabled": "false"}
```
   OUTPUT
```json
{"message": "New event Added!"}
```

## /api/user/event (DELETE)
   INPUT
```json
{"eventID": "1"}
```
   OUTPUT
```json
{"message": "Event Removed!"}
```

## /api/user/event (GET)

   OUTPUT
```json
[{"id":1,"eventname":"newUnix","eventtime":"1970-01-01T02:20:33Z","created_at":"2022-10-18T13:26:11.536961Z","userid":1,"issilent":false,"snoozedisabled":false},{"id":3,"eventname":"newUnix","eventtime":"1970-01-01T05:25:34Z","created_at":"2022-10-18T13:26:21.909119Z","userid":1,"issilent":true,"snoozedisabled":true}]
```



## /api/user/event (PATCH)
   INPUT
```json
{"eventID": "11", "eventTime": "1666095688"}
```
   OUTPUT
```json
{"message": "Event time updated!"}
```





