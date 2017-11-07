env GOOS=windows GOARCH=386 go build
mv badfood.exe badfood-win386.exe
env GOOS=windows GOARCH=amd64 go build
mv badfood.exe badfood-win64.exe
env GOOS=linux GOARCH=amd64 go build
mv badfood badfood-linux64
env GOOS=darwin GOARCH=386 go build
mv badfood badfood-darwin386
env GOOS=darwin GOARCH=amd64 go build
mv badfood badfood-darwinAmd64
