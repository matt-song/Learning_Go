package main

import (
    "net/http"
    //"github.com/davecgh/go-spew/spew"
    "fmt"
    "io/ioutil"
)

func main() {

    url := "https://sgp-ping.vultr.com/vultr.com.100MB.bin"

    /* ---------- Genertate a http request ---------- */
    fmt.Println("Generating the http request")
    request, err := http.NewRequest("GET", url, nil)
    if err != nil {
        fmt.Println("Failed to download from url ["+ url + "]")
    }
    //spew.Dump(request)

    /* ---------- send the request to http server ---------- */
    fmt.Println("Send the http request to server")
    response, err := http.DefaultClient.Do(request)
    //spew.Dump(response)

    /* ---------- write the file to disk ---------- */
    fmt.Println("Write the file to /tmp/test.file")
    content, _ := ioutil.ReadAll(response.Body)   // return content []byte, err error
    err = ioutil.WriteFile("/tmp/test.file", content, 0666)
    if err != nil {
        fmt.Println("Failed to write file to [/tmp/test.file]")
    }
}

/*

(*http.Request)(0xc000122000)({
 Method: (string) (len=3) "GET",
 URL: (*url.URL)(0xc000120000)(https://sgp-ping.vultr.com/vultr.com.100MB.bin),
 Proto: (string) (len=8) "HTTP/1.1",
 ProtoMajor: (int) 1,
 ProtoMinor: (int) 1,
 Header: (http.Header) {
 },
 Body: (io.ReadCloser) <nil>,
 GetBody: (func() (io.ReadCloser, error)) <nil>,
 ContentLength: (int64) 0,
 TransferEncoding: ([]string) <nil>,
 Close: (bool) false,
 Host: (string) (len=18) "sgp-ping.vultr.com",
 Form: (url.Values) <nil>,
 PostForm: (url.Values) <nil>,
 MultipartForm: (*multipart.Form)(<nil>),
 Trailer: (http.Header) <nil>,
 RemoteAddr: (string) "",
 RequestURI: (string) "",
 TLS: (*tls.ConnectionState)(<nil>),
 Cancel: (<-chan struct {}) <nil>,
 Response: (*http.Response)(<nil>),
 ctx: (context.Context) <nil>
})


*/