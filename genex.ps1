go build ./cmd/crudgeon

if (Test-Path -Path ./output) {
    Remove-Item -Path ./output -Recurse -Force
}
./crudgeon -dir ./datafiles -path output -ns Name.Space