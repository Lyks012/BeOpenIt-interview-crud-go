# wait-for "${DATABASE_HOST}:${DATABASE_PORT}"


echo "Waiting for Mysql to start..."
ls -1l wait-for
./wait-for db:3306

echo "Starting the server..."
go run main.go