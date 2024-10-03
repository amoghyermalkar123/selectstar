sudo mkdir -pv /opt/selectstar/

sudo chmod 775 /opt/selectstar/

sudo rm /usr/local/bin/selectstar-daemon
sudo rm /usr/local/bin/selectstar

go build -o selectstar-daemon
sudo cp selectstar-daemon /usr/local/bin/

cd cmd/ && go build -o selectstar
sudo cp selectstar /usr/local/bin/

echo "SelectStar has been installed successfully ..."
sudo selectstar server start
echo "please open http://localhost:8091/"
echo "if you want to stop the server, run \"sudo selectstar server stop\""
echo "if you want to restart the server, run \"sudo selectstar server start\""
