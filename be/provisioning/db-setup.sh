
DATE() {
  date '+%Y-%m-%d %H:%M:%S'
}

echo "[$(DATE)] [Info] [System] Update package..."
yum update &> /dev/null

echo "[$(DATE)] [Info] [System] Install postgresql..."
sudo yum install -y postgresql-server postgresql-contrib &> /dev/null

sudo postgresql-setup initdb 

echo "[$(DATE)] [Info] [System] Start postgresql and enable on startup..."
sudo systemctl start postgresql &> /dev/null

sudo systemctl enable postgresql&> /dev/null


echo "[$(DATE)] [Info] [System] Setup config for vagrant user..."

sudo sed -i "s/#listen_address.*/listen_addresses '*'/" /var/lib/pgsql/data/postgresql.conf &> /dev/null

sudo cat > /var/lib/pgsql/data/pg_hba.conf <<EOF
# TYPE  DATABASE        USER            ADDRESS                 METHOD

# "local" is for Unix domain socket connections only
local   all             all                                     md5
# IPv4 local connections:
host    all             all             127.0.0.1/32            md5
# IPv6 local connections:
host    all             all             ::1/128                 md5
# Allow replication connections from localhost, by a user with the
# replication privilege.
#local   replication     postgres                                peer
#host    replication     postgres        127.0.0.1/32            ident
#host    replication     postgres        ::1/128                 ident
# Accept all IPv4 connections - FOR DEVELOPMENT ONLY!!!
host    all         all         0.0.0.0/0             md5

EOF

ALTER_POSTGRES_USER_SQL="ALTER USER postgres WITH ENCRYPTED PASSWORD 'postgres'"
CREATE_VAGRANT_ROLE="CREATE ROLE vagrant SUPERUSER LOGIN PASSWORD 'vagrant'"
CREATE_VAGRANT_DATABASE="CREATE DATABASE vagrant"

sudo -u postgres psql --command="$ALTER_POSTGRES_USER_SQL" &> /dev/null
sudo -u postgres psql --command="$CREATE_VAGRANT_ROLE" &> /dev/null
sudo -u postgres psql --command="$CREATE_VAGRANT_DATABASE" &> /dev/null

#sudo su postgres -c "psql -c \"CREATE ROLE vagrant SUPERUSER LOGIN PASSWORD 'vagrant'\" "
#sudo su postgres -c "psql -c \"CREATE DATABASE vagrant\" "



echo "[$(DATE)] [Info] [System] Restart PGSQL for config..."
sudo systemctl restart postgresql &> /dev/null

echo "[$(DATE)] [Info] [System] PSQL - Installation complete!"
echo "**********************************************************"
echo "*                                                        *"
echo "*     PGSQL VERSION:  $(psql --version)           *"
echo "*     Run: psql    #for discover db                      *"
echo "**********************************************************"