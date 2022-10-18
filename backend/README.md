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


# ENDPOINTS

* /api/login POST
   * {"username": "test",	"password": "test"}
   * {"token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJRCI6MSwiZXhwIjoxNjY2MjEwNzQ1LCJ1c2VybmFtZSI6InRlc3QifQ.8PWstjzN8SD5lORNWAmAh1NBCGVL9GugE67rtzJN1NA"}

* /api/register POST
   * {"username": "test",	"password": "test","password_confirm": "test"}
   * {"message": "Registered!"}

  ## SECURED ROUTER, BEARER NEEDED

* /api/user/event **POST**
   * {"eventName": "newUnix", "eventTime": "1665997920","isSilent": "true","SnoozeDisabled": "false"}
   * {"message": "New event Added!"}

* /api/user/event **DELETE**
   * {"eventID": "1"}
   * {"message": "Event Removed!"}
   
* /api/user/event **GET**
   * ARRAY OF EVENTS
   
* /api/user/event **DELETE**
   * {"eventID": "11", "eventTime": "1666095688"}
   * {"message": "Event time updated!"}
   






