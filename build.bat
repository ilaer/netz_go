go build    -ldflags="-s -w -H=windowsgui" -o netz.exe
upx netz.exe
"C:\Program Files\WinRAR\WinRAR.exe" a -r -ep1 -idq -inul -y "netz.zip" "netz.exe"
timeout /nobreak /t 15