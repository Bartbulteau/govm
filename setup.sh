#!/bin/bash

clear
echo "Building amber and govm"
echo "Sudo password may be needed to install amber and govm in your bin directory..."
go build 
sudo cp govm /usr/local/bin/govm
sudo cp amber /usr/local/bin/amber
sudo cp -r _amber_modules /usr/local/bin/_amber_modules
echo "Done."
echo "Usage :"
echo "  amber [input] [executable name]"
echo "  govm [executable name]"