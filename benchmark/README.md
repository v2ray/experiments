# 性能测试脚本

## 准备工作

```bash
# Update to latest binaries
apt-get update
apt-get -y upgrade
apt-get -y install git

# Install Go
curl -O https://storage.googleapis.com/golang/go1.7.4.linux-amd64.tar.gz
tar -C /usr/local -xzf go1.7.4.linux-amd64.tar.gz

# Download this repo
mkdir -p ./git/src/github.com/v2ray
cd ./git/src/github.com/v2ray
git clone https://github.com/v2ray/experiments.git
cd experiments/benchmark/testcases
source ./env.sh

# Build benchmark tools
./build.sh
```

## 单项测试

### 直连
```bash
./bare.sh
```

### V2Ray

```bash
# Install
curl -O https://install.direct/go.sh
chmod +x go.sh
./go.sh

# V2Ray direct: Dokodemo-door <-> Freedom
./v2ray_bare.sh

# V2Ray VMess
./v2ray_vmess.sh
```

### Shadowsocks Libev

```bash
# Install (Debian Jessie)
sh -c 'printf "deb http://httpredir.debian.org/debian jessie-backports main" > /etc/apt/sources.list.d/jessie-backports.list'
apt-get update
apt-get install shadowsocks-libev

# Test
./ss_libev.sh
```

### ShadowsocksR Python

```bash
# Install
mkdir -p $GOPATH/src/github.com/shadowsocksr
pushd $GOPATH/src/github.com/shadowsocksr
git clone https://github.com/shadowsocksr/shadowsocksr.git
cd shadowsocksr
bash initcfg.sh
popd

# Test
./ssr_py.sh
```
