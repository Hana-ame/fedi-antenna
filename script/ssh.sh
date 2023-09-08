#!/bin/bash

while ((1 < 2))
do
  ssh -R 3000:localhost:3000 root@[2a07:d884::130d] "bash script/kill_process.sh 3000";
  ssh -R 3000:localhost:3000 root@[2a07:d884::130d] -o ServerAliveInterval=1 
  # 断开了会自动冲脸么不知大。
done
