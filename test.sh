cd delegates
go test -coverprofile=coverage.out
sleep 15
cd ..
cd signins
go test -coverprofile=coverage.out
sleep 15
cd ..
cd handlers
go test -coverprofile=coverage.out
sleep 15
cd ..