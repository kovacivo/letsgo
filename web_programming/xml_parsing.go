package main

import(
"encoding/xml"
"fmt"
"io/ioutil"
"net/http"
"os"
)

type Sitemapindex struct {
  Locations []Location `xml:"sitemap"`
}

type Location struct {
  Loc string `xml:"loc"`
}

func (l Location) String() string {
  return fmt.Sprintf(l.Loc)
}

func main() {
	resp, _ := http.Get("https://www.oracle.com/sitemap.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)
	var s Sitemapindex
	xml.Unmarshal(bytes, &s)
	fmt.Println(s.Locations)
	fmt.Println("Status code was ",resp.Status) 
	fmt.Fprintf(os.Stdout,"Status code was %s and protocol used was %s\n",resp.Status, resp.Proto) 
	resp.Body.Close()
	//for _, Location := range s.Locations {
//		fmt.Printf("%s\n", Location)
//	}
}
