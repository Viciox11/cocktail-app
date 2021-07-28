
DATE() {
  date '+%Y-%m-%d %H:%M:%S'
}

echo "[$(DATE)] [Info] [System] Update package..."
yum update &> /dev/null
systemctl stop postgresql.service
sudo setsebool -P httpd_can_network_connect 1

echo "[$(DATE)] [Info] [System] Install httpd and php for psql..."
sudo yum install -y httpd &> /dev/null

sudo yum install -y php php-pgsql &> /dev/null
echo "[$(DATE)] [Info] [Systemi] Download adminer..."


wget https://github.com/vrana/adminer/releases/download/v4.7.1/adminer-4.7.1.php -O /var/www/html/index.php &> /dev/null
sudo sed -i "s/80/8080/g" /etc/httpd/conf/httpd.conf

echo "[$(DATE)] [Info] [System] Restart HTTPD for config..."
systemctl start httpd.service &> /dev/null
systemctl enable httpd.service &> /dev/null
systemctl start postgresql.service &> /dev/null

echo "[$(DATE)] [Info] [System] ADMINER - Installation complete!"
echo "**********************************************************"
echo "*                                                        *"
echo "*     Adminer installed                                  *"
echo "*     Go to: http://192.168.12.34/index.php              *"
echo "**********************************************************"
