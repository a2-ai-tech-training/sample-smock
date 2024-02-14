#!/bin/bash

echo This is stdout for $0
echo This is too!! Below this is stderr
echo text > text.txt
echo text >> text.txt
sleep 1
cp text.txt /text.txt
echo after the error
