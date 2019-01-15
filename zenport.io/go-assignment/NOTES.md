# Notes
To deploy this app please run 'docker-compose up' in the go-assignment directory.
This will build and deploy the app alongside a mysql database container.
Environment variables in the docker-compose.yml specify all the db connection details needed. Feel free to change them but make sure you change them in the sevice description too.

Once docker-compose up has completed there will be a database connection available on localhost.

To run the tests please run the following command:
- dbHost=localhost dbPort=3306 dbName=zenfightertest dbUser=newguy dbPass=password123 go test -p 1 ./domain ./providers/database ./adapters/http
