while true
do
  ps -ef|grep "ssh -N -g -L 33023:192.168.133.110:6443 183.207.130.12 -p22"|grep -v grep
  if [ $? -eq 1 ]; then
    nohup ssh -N -g -L 33023:192.168.133.110:6443 183.207.130.12 -p22 > /dev/null 2>&1 &
  fi
  ps -ef|grep "ssh -N -g -L 33028:127.0.0.1:80 192.168.13.12"|grep -v grep
  if [ $? -eq 1 ]; then
    nohup ssh -N -g -L 33028:127.0.0.1:80 192.168.13.12 > /dev/null 2>&1 &
  fi
  ps -ef|grep "ssh -N -g -L 33030:192.168.22.34:32290 183.207.130.12 -p22"|grep -v grep
  if [ $? -eq 1 ]; then
    nohup ssh -N -g -L 33030:192.168.22.34:32290 183.207.130.12 -p22 > /dev/null 2>&1 &
  fi

  #ps -ef|grep "ssh -N -g -L 33021:127.0.0.1:33021 183.207.130.12 -p22"|grep -v grep
  #if [ $? -eq 1 ]; then
  #  nohup ssh -N -g -L 33021:127.0.0.1:33021 183.207.130.12 -p22 > /dev/null 2>&1 &
  #fi

  ps -ef|grep "ssh -N -g -L 33021:10.253.199.82:33021 183.207.130.12 -p22"|grep -v grep
  if [ $? -eq 1 ]; then
    nohup ssh -N -g -L 33021:10.253.199.82:33021 183.207.130.12 -p22 > /dev/null 2>&1 &
  fi

  ps -ef|grep "ssh -N -g -L 30371:10.253.199.46:30371 183.207.130.12 -p22"|grep -v grep
  if [ $? -eq 1 ]; then
    nohup ssh -N -g -L 30371:10.253.199.46:30371 183.207.130.12 -p22 > /dev/null 2>&1 &
  fi

  ps -ef|grep "ssh -N -g -L 33024:127.0.0.1:33024 183.207.130.12 -p22"|grep -v grep
  if [ $? -eq 1 ]; then
    nohup ssh -N -g -L 33024:127.0.0.1:33024 183.207.130.12 -p22 > /dev/null 2>&1 &
  fi

  sleep 3
done
