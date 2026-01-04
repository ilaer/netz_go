set GOARCH=386
set CGO_ENABLED=1
del netz32.zip
timeout /nobreak /t 3
go build  -o netz32.exe    main.go
"C:\Program Files\WinRAR\WinRAR.exe" a -r -ep1 -idq -inul -y "netz32.zip" "netz32.exe" 
timeout /nobreak /t 15
pause