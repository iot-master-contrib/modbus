
# 整体编译
go env -w GOPROXY=https://goproxy.cn,direct
go env -w GOPRIVATE=*.gitlab.com,*.gitee.com
go env -w GOSUMDB=off

app="modbus"
version="1.0.0"

read -t 5 -p "please input version(default:$version)" ver
if [ -n "${ver}" ];then
	version=$ver
fi


pkg="github.com/zgwit/iot-master/v3/pkg/build"
gitHash=$(git show -s --format=%H)
buildTime=$(date -d today +"%Y-%m-%d %H:%M:%S")

# -w -s
ldflags="-X '${pkg}.Version=$version' \
-X '${pkg}.GitHash=$gitHash' \
-X '${pkg}.Build=$buildTime'"

#export CGO_ENABLED=1
#CC=x86_64-w64-mingw32-gcc

export GOARCH=amd64

export GOOS=windows
go build -ldflags "$ldflags" -o modbus.exe main.go

export GOOS=linux
go build -ldflags "$ldflags" -o modbus main.go