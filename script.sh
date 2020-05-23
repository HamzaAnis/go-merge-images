rm go-merge-images
go build
mv go-merge-images ./test
cd test
./go-merge-images
