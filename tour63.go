package main

import (
    "io"
    "os"
    "strings"
)

type rot13Reader struct {
    r io.Reader
}

func (reader rot13Reader) Read(p []byte) (n int, e error){
    
    count, err:= reader.r.Read(p)
    for index,b := range p{
        ascii := uint8(b)
        if ascii>=109{
            p[index]=ascii-13
        } else if ascii>=96{
        	p[index]=ascii+13
        } else if ascii>=78{
        	p[index]=ascii-13
        } else{
            p[index]=ascii+13
        }
    }
    
    return count, err 
}

func main() {
    s := strings.NewReader(
        "Lbh penpxrq gur pbqr!")
    r := rot13Reader{s}
    io.Copy(os.Stdout, &r)
}
