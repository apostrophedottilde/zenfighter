# Notes

## To run the app
To deploy this app please run 'docker-compose up' in the go-assignment directory.
This will build and deploy the app alongside a mysql database container.
Environment variables in the docker-compose.yml specify all the db connection details needed for a development environment.

You can now access the app on localhost (or if you're on windows and using docker-toolbox you will have to get your vm's IP).
The app is exposed on port 8000.

Exposed endpoints are:

    - GET /knight
    - GET /knight/id
    - POST /knight
    - POST /fight?fighter1=X&fighter2=Y

I added the endpoint for fight because it didn't make sense to implement fighting fucntionality without allowing a trigger for it.

## To run the tests

- Please navigate to the 'RUN TESTS' directory.

- Run docker-compose build then docker-compose up.

- Navigate back to the 'go-assignment' directory.

- execute the following command: dbHost=localhost dbPort=3307 dbName=zenfightertest dbUser=newguy dbPass=password123 go test -p 1 ./domain ./providers/database ./adapters/http

I apologise that this is not an elligant solution at all, it was just the quickest at the time. 