sudo yum install -y epel-release
sudo yum install -y nginx

sudo mkdir /etc/nginx/cert

openssl req -x509 -out /etc/nginx/cert/api.ioartigiano.com.crt -keyout /etc/nginx/cert/api.ioartigiano.com.key   -newkey rsa:2048 -nodes -sha256   -subj '/CN=api.ioartigiano.com' -extensions EXT -config <( printf "[dn]\nCN=api.ioartigiano.com\n[req]\ndistinguished_name = dn\n[EXT]\nsubjectAltName=DNS:api.ioartigiano.com\nkeyUsage=digitalSignature\nextendedKeyUsage=serverAuth")


sudo cat /opt/ioartigiano-be/provisioning/nginx.conf > /etc/nginx/nginx.conf


sudo systemctl enable nginx
sudo systemctl start nginx
+