package main

import(  "fmt"
         "net/http"
         "log"
         "os/exec"
         "strings"
          "os" )

func Hello() {
    fmt.Println("testing Go");
}

func runcommand(command string) int { 
     splitted := strings.Split(command, " ");
     fmt.Println(splitted);
     out,error := exec.Command(splitted[0],splitted[1]).CombinedOutput();
     if error == nil{
        version := strings.Split(string(out)," ")[2];
        fmt.Println(version)
        if version == "go1.13" {
                   fmt.Println("got it..."); 
                   return 0   }
    }else{
        log.Fatal(error);

    }
      return 1
}

func main(){
     Hello();
     _tasks_ := 0
     fmt.Println("sending request to google......");
     resp,error := http.Get("http://www.google.com");
     if error == nil{
      status := int(resp.StatusCode);
      if  status == 200{
           _tasks_ = _tasks_ + 1 
      log.Println(resp.StatusCode);
      }
     _go_v := int(runcommand("go version")); 
     if _go_v == 0{
          _tasks_ += 1 
    }
    if _tasks_ == 2{
       os.Exit(0)
    }else{
      os.Exit(1)
}
}
}
