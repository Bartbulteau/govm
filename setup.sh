#!/bin/bash

clear
echo ""

echo "
    ======================================
          Amber and GoVM version 0.0.2
    ======================================
    Copyright (c) 2018 Barthelemy Bulteau
"
echo ""

echo "Building amber and govm..."
go build src/vm.go
echo ""
echo "Sudo password may be needed to install amber and govm in your bin directory..."
mv vm govm
sudo cp govm /usr/local/bin/govm
sudo cp src/amber /usr/local/bin/amber
sudo cp -r src/_amber_modules /usr/local/bin/_amber_modules
echo "Done."
echo ""
echo "Usage :"
echo "  amber [input] [executable name]"
echo "  govm [executable name]"
echo ""