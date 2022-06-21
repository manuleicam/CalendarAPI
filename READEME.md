# Interview API

This golang API works inside a docker container that can be built using the dockerfile in this folder

This API will start with the example cases already on the database.

### Commands to start API:
1. docker build -t calendarapi . 
2. docker run -p 80:80 calendarapi 

After that the docker will be running on the port 80 in the machine IP.

# File Struct
+ main.go starts the API

+ In my project i have 2 resources "USERS" and "TIME SLOTS"

5 main folders:

    + All 5 folders are splited inside between resources.

+   businessLayer:
        + Files that redirect requests base in the method used to call the endpoint.
        + Files that read the request and write the response everytime someone reachs an endpoint.

+   documentation:
        + Files that explain how the project work.

+   model:
        + Files that work as my databases. Both the files have a "MAP" and only functions inside this files can access the respective "MAP".

+   module:
        + Files where all my resources' structs are coded.

+   routes:
        + File where all the endpoints are. This file redirects the requests to other functions based in the endpoint called.


# Times Slots Endpoints
### delete timeslot based on userid and timeslotid
#### "/users/{userId}/time-slots/{timeSlotId}"

Example: 

    +   DELETE  localhost/users/1/time-slots/2    -> Deletes the time slot with id = 2 that belongs to the user with id = 1
            Response Body:
                Plain text: "Time slot deleted"

### get and create user timeslot by user id
#### "/time-slots/users/{id}"

Example: 

    +   GET  localhost/time-slots/users/1  -> Gets the time slots for the user with id = 1
            Response Body:
                JSON    {
                            "userID": 1,
                            "days": [
                                {
                                    "id": 0,
                                    "day": "2022-06-07T00:00:00Z",
                                    "hourBeg": 9,
                                    "hourEnd": 10
                                }
                            ]
                        }
    
    +   POST  localhost/time-slots/users/1  -> Creates the time slots in the request body for the user with id = 1
            Request Body:
                JSON    {
                            "days": [
                                {
                                    "day": "2022-06-07T00:00:00Z",
                                    "hourBeg": 9,
                                    "hourEnd": 10
                                }
                            ]
                        }

            Response Body:
                Plain text: "Time Slot Created"        

### get the match between users IDs (user id in query parameters)
#### "/time-slots/users?id=..."

Example: 

    +   GET  localhost/time-slots/users?id=1,2  -> Gets the time slots that match between the users with id 1 and 2
            Response Body:
                JSON    [
                            {
                                "id": 0,
                                "day": "2022-06-07T00:00:00Z",
                                "hourBeg": 9,
                                "hourEnd": 10
                            }
                        ]

# Users Endpoints
### Gets users, delete all users, creates new user
#### "/users"

Example:

    +   GET     localhost/users     -> Returns all users in JSON
        Response Body:
            JSON    [{
                        "name": "NAME",
                        "position": "POSITION"
                    }]
            
    +   POST    localhost/users     -> Creates a new user and returns the new user ID
        Request Body:
            JSON    {
                        "name": "NEW_NAME",
                        "position": "NEW_POSITION"
                    } 
        Response Body:
            Plain text: "New Id: 1"

    +   DELETE  localhost/users     -> Deletes all the users
        Response Body:
            Plain text: "Users Removed"

### get, update and delete user info by user ID
#### "/users/{id}"

Example: 

    +   GET     localhost/users/1     -> Returns info about user with id = 1
        Response Body:
            JSON    [{
                        "name": "NAME",
                        "position": "POSITION"
                    }]
            
    +   PUT    localhost/users/1     -> updates the user info with id = 1 to the info in the request body
        Request Body:
            JSON    {
                        "name": "NEW_NAME",
                        "position": "NEW_POSITION"
                    } 
        Response Body:
            Plain text: "ID: 1 Updated"

    +   DELETE  localhost/users/1     -> Deletes the user with id = 1
        Response Body:
            Plain text: "ID: 1 Deleted"

