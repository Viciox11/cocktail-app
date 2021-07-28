#!/bin/bash
# This is the entry point for configuring the system.
#####################################################
DATE() {
  date '+%Y-%m-%d %H:%M:%S'
}

echo "[$(DATE)] [Info] [System] Install basic tools..."
sudo yum install -y git vim wget net-tools &> /dev/null

echo "[$(DATE)] [Info] [System] Install gcc compiler..."
sudo yum install -y gcc &> /dev/null

echo "[$(DATE)] [Info] [System] Import go centos repo..."
wget https://dl.google.com/go/go1.13.3.linux-amd64.tar.gz &> /dev/null

echo "[$(DATE)] [Info] [System] Verify TAR sha256sum..."
sha256sum go1.13.3.linux-amd64.tar.gz &> /dev/null

echo "[$(DATE)] [Info] [System] Extract golang tar..."
sudo tar -C /usr/local -xzf go1.13.3.linux-amd64.tar.gz &> /dev/null

echo "[$(DATE)] [Info] [System] Set go ENV PATH..."

echo "export PATH=$PATH:/usr/local/go/bin:~/go/bin" >> /home/vagrant/.bash_profile

echo "export SERVER_HOST=0.0.0.0" >>/home/vagrant/.bash_profile
echo "export ENV=development" >>/home/vagrant/.bash_profile
echo "export SERVER_PORT=5000" >>/home/vagrant/.bash_profile
echo "export PSQL_HOST=localhost" >>/home/vagrant/.bash_profile
echo "export PSQL_DATABASE=vagrant" >>/home/vagrant/.bash_profile
echo "export PSQL_USER=vagrant" >>/home/vagrant/.bash_profile
echo "export PSQL_PSW=vagrant" >>/home/vagrant/.bash_profile
echo "export PSQL_PORT=5432" >>/home/vagrant/.bash_profile

source /home/vagrant/.bash_profile

echo "[$(DATE)] [Info] [System] Install air for hot reloading..."
curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b /home/vagrant/go/bin &> /dev/null

echo 'alias air="~/go/bin/air"' >> /home/vagrant/.bashrc
source /home/vagrant/.bashrc

echo "[$(DATE)] [Info] [System] Creating resources directory..."
mkdir -p /resources/images
chown -R vagrant:vagrant /resources/images

chown -R vagrant:vagrant /home/vagrant/go

sudo setenforce 0

echo "[$(DATE)] [Info] [System] GO - Installation complete!"
echo "**********************************************************"
echo "*                                                        *"
echo "*     GO VERSION:  $(go version)       *"
echo "*     Run: cd /opt/ioartigiano-be/cmd/server             *"
echo "*     Run: go run main.go # or air for hot reloading     *"
echo "**********************************************************"

