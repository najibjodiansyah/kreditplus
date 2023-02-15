create migrations 

migrate create -ext sql -dir app/config/databbase/migrations <nama_file>

run migrations 

migrate -databse "mysql://root@tcp(localhost:3360)/kreditplus" -path app/config/database/migrations up

rollback migrations 

migrate -databse "mysql://root@tcp(localhost:3360)/kreditplus" -path app/config/database/migrations down
