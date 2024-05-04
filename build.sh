EXE_NAME=ip2mmdb
OS=$(uname -s)
IP_LIST=china_ip_list.txt

go build -o bin/${EXE_NAME} main.go
go build -o bin/verify verify/verify_ip.go

# Get China IP list
curl -sSL https://ruleset.skk.moe/Clash/ip/china_ip.txt > bin/$IP_LIST
if [ $OS == "Darwin" ]; then
    sed -i '' '/^#/d' bin/$IP_LIST
else
    sed -i '/^#/d' bin/$IP_LIST
fi
curl -sSL https://raw.githubusercontent.com/gaoyifan/china-operator-ip/ip-lists/china6.txt >> bin/$IP_LIST

# Generate mmdb
./bin/${EXE_NAME} -s bin/$IP_LIST -d bin/Country.mmdb -t "GeoIP2-Country"

# Alternative mmdb for test: GeoIP2-CN
# curl -sSL "https://github.com/Hackl0us/GeoIP2-CN/raw/release/Country.mmdb" > bin/GeoIP2-CN.mmdb