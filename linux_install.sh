#!/bin/bash

echo "Downloading hydra..."
curl -L "https://github.com/Shravan-1908/hydra/releases/latest/download/hydra-linux-amd64" -o hydra

echo "Adding hydra into PATH..."

mkdir -p ~/.hydra

chmod u+x ./hydra

mv ./hydra ~/.hydra
echo "export PATH=$PATH:~/.hydra" >> ~/.bashrc

echo "hydra installation is completed!"
echo "You need to restart the shell to use hydra."
