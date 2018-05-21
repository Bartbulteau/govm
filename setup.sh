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

echo "Building GoVM..."
go build src/main.go
if [ $? -eq 0 ]; then
    echo "govm build succeeded"
else
    echo "Fail : unable to build govm
Do you have Go installed ?
Check by typing : go version"
    exit 0
fi

echo ""
echo "Sudo password may be needed to install amber and govm in your bin directory..."
mv main govm
if [ $? -eq 0 ]; then
    echo "Success : renaming files"
else
    echo "Fail : can't rename vm to govm"
    exit 0
fi
sudo cp govm /usr/local/bin/govm
if [ $? -eq 0 ]; then
    echo "Success : copying govm to /usr/local/bin"
else
    echo "Fail : unable to copy govm to /usr/local/bin"
    exit 0
fi
sudo cp src/amber /usr/local/bin/amber
if [ $? -eq 0 ]; then
    echo "Success : copying amber to /usr/local/bin"
else
    echo "Fail : unable to copy amber to /usr/local/bin"
    exit 0
fi
sudo cp -r src/_amber_modules /usr/local/bin/_amber_modules
if [ $? -eq 0 ]; then
    echo "Success : copying _amber_modules to /usr/local/bin"
else
    echo "Fail : unable to copy _amber_modules to /usr/local/bin"
    exit 0
fi
rm govm
if [ $? -eq 0 ]; then
    echo "Success : cleaning directory"
else
    echo "Fail : unable to remove govm bin file from this directory"
    echo "This won't stop installation"
fi
echo "Installation completed"
echo ""
echo "Usage :"
echo "  amber [input] [executable name]"
echo "  govm [executable name]"
echo ""