if (Test-Path -Path ./output) {
    Remove-Item -Path ./output -Recurse -Force
}
go run ./cmd/crudgeon -dir ./datafiles -path output -ns Name.Space