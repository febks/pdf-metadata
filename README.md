go run main.go list infile/gojek.pdf > log/system.log 2>&1

go run main.go update input.pdf output.pdf Title="dummy" Author="dummy" Creator="dummy" > log/system.log 2>&1

go run main.go remove input.pdf clean.pdf > log/system.log 2>&1