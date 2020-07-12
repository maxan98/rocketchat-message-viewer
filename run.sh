#!/bin/bash

ssh -vvv -N -L 27017:127.0.0.1:27017 root@10.228.18.132
./goreadmongo "$*"