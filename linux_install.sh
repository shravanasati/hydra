#!/bin/bash
chmod +x ./linux_install.sh

echo "Downloading hydra..."
curl -L "https://github.com/Shravan-1908/hydra/releases/latest/download/hydra-linux-amd64" -o hydra

echo "Placing hydra into PATH..."
mv ./hydra /bin


echo "hydra installation is completed!"
echo "You need to restart the shell to use hydra."
