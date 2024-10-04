sudo mkdir -pv /opt/selectstar/
sudo touch /opt/selectstar/selectstar.pid

sudo chmod 775 /opt/selectstar/

sudo rm /usr/local/bin/selectstar-daemon
sudo rm /usr/local/bin/ss

go build -o selectstar-daemon
sudo cp selectstar-daemon /usr/local/bin/

cd cmd/ && go build -o ss
sudo cp ss /usr/local/bin/

echo "SelectStar has been installed successfully ..."
sudo ss server start
echo "and please open http://localhost:8091/"
echo "if you want to stop the server, run \"sudo ss server stop\""
echo "if you want to restart the app, run - \"sudo ss server start\""
