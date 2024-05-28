build:
	go build -o ./bin/mac_arcloud-effects-cli app/main.go
	GOOS=windows GOARCH=amd64 go build -o ./bin/win_arcloud-effects-cli.exe app/main.go
