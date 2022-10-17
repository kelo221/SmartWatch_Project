## ENDPOINTS

* /api/login POST
   * {"username": "test",	"password": "test"}
   * {"token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJRCI6MSwiZXhwIjoxNjY2MjEwNzQ1LCJ1c2VybmFtZSI6InRlc3QifQ.8PWstjzN8SD5lORNWAmAh1NBCGVL9GugE67rtzJN1NA"}

* /api/register
   * {"username": "test",	"password": "test","password_confirm": "test"}
   * {"message": "Registered!"}

* /api/user/event **BEARER NEEDED**
   * {"eventName": "newUnix", "eventTime": "1665997920","isSilent": "true","SnoozeDisabled": "false"}
   * {"message": "New event Added!"}

* /api/user/event **BEARER NEEDED**
   * {"eventName": "newUnix", "eventTime": "1665997920","isSilent": "true","SnoozeDisabled": "false"}
   * {"message": "New event Added!"}

* /api/user/event **BEARER NEEDED**
   * {"eventID": "1"}
   * {"message": "Event Removed!"}
   
 * /api/user/events **BEARER NEEDED**
   * ARRAY OF EVENTS


## INIT Database

1. psql.exe -U postgres
2. CREATE DATABASE smartwatch_project; \q
3. psql.exe -U postgres -d smartwatch_project -f db/schema.sql

* OPTIONAL
    * (for backend developement generate ORM: sqlboiler psql) 


