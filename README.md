# blog

# REST API

GET /notes -- list of notes -- 200, 500
POST /notes -- create note -- 201, 500

# Init

To get the default configuration file, run with the init command.
For example go run main.go init
When using Docker, set the db database port in the config file.
In the docker-compose file, when mapping the directory, set the user to the current system user.

# Start 

To get the default configuration file, run with the start command.
For example go run main.go start
