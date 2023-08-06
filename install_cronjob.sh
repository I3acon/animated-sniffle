#!/usr/bin/env bash

read -p "Enter Name: " NAME
read -p "Enter NODE_URL: " NODE_URL
read -p "Enter Validator Address: " VAL_ADR
read -p "Enter Block Explorer: " BLOCK_EXPLORER
read -p "Enter MissedBlockThreshold: " THRESHOLD
read -p "Enter Frequency (ms): " FREQUENCY
read -p "Enter API_URL: " API_URL

echo '================================================='
echo -e "Name: \e[1m\e[32m$NAME\e[0m"
echo -e "NODE_URL: \e[1m\e[32m$NODE_URL\e[0m"
echo -e "Validator Address: \e[1m\e[32m$VAL_ADR\e[0m"
echo -e "Block Explorer: \e[1m\e[32m$BLOCK_EXPLORER\e[0m"
echo -e "MissedBlockThreshold: \e[1m\e[32m$THRESHOLD\e[0m"
echo -e "Frequency (ms): \e[1m\e[32m$FREQUENCY\e[0m"
echo -e "API_URL: \e[1m\e[32m$API_URL\e[0m"
echo '================================================='
sleep 3

echo -e "\e[1m\e[32m1. Installing $NAME Cronjob with $FREQUENCY ms frequency...\e[0m" && sleep 1
# install cosmos-exporter
wget -O cronjob https://github.com/I3acon/animated-sniffle/releases/download/Cronjob-v.1.1/cronjob
chmod +x cronjob

export USERNAME=$(whoami)
sudo -E bash -c "cat << EOF > /etc/systemd/system/$NAME.service
[Unit]
Description=$NAME Daemon
After=network-online.target

[Service]
User=$USERNAME
ExecStart=$HOME/cronjob $NODE_URL $VAL_ADR $BLOCK_EXPLORER $THRESHOLD $FREQUENCY $API_URL
Restart=always
RestartSec=3
LimitNOFILE=4096

[Install]
WantedBy=multi-user.target
EOF"

sudo systemctl daemon-reload
sudo systemctl enable $NAME
sudo systemctl start $NAME

echo -e "\e[1m\e[32mInstallation finished...\e[0m" && sleep 1
echo -e "\e[1m\e[32m2. Checking $NAME status...\e[0m" && sleep 1
sudo systemctl status $NAME

echo -e "\e[1m\e[32m3. Checking $NAME logs...\e[0m" && sleep 1
sudo journalctl -u $NAME -f