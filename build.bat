cd %~dp0
cd mfp_web
cmd /c npm run build
cd ..
rmdir /S /Q bin
mkdir bin
copy /Y dictionary.txt bin\
set CGO_ENABLED=0
set GOARCH=amd64
set GOOS=linux
go build -trimpath -o bin/mfp-linux-amd64
set GOARCH=amd64
set GOOS=windows
go build -trimpath -o bin/mfp-windows-amd64.exe
pause
exit