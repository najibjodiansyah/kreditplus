create migrations 

migrate create -ext sql -dir app/config/database/migrations <nama_file>

run migrations 

migrate -database "mysql://root:@tcp(localhost:3306)/kreditplus" -path app/config/database/migrations up

rollback migrations 

migrate -databsae "mysql://root:@tcp(localhost:3306)/kreditplus" -path app/config/database/migrations down

https://www.simulasikredit.com/simulasi_kredit_mobil_simpel.php

challange ur self make unit and integration test using flag golang

monitoring using open telemetry and jaeger
