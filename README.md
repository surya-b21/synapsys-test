# SYNAPSYS TEST
This app was developed for synapsis test purposes. This app has seed and migration so you don't have to create and fill the database

## HOW TO RUN
Requirement
 - Docker Engine ( with docker compose )

Make sure you have create .env file. You can copy .env.example and fill variable
This is default .env file

    DB_HOST=mariadb
    DB_PORT=3306
    DB_DATABASE=synapsis
    DB_USERNAME=root
    DB_PASSWORD=th1s!sr0ot
    ENABLE_MIGRATION=true

For the first time you can enable migration. After that disable it.
You can follow this command to run app

    docker compose up -d
    docker compose exec app go run .
  You can login with this account
  

    email : admin@email.com
    password: admin12345

## ACCESS SWAGGER
After turn on docker container you can follow access [localhost:8080](http://localhost:8080/) to access swagger
