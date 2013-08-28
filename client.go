package roxanne

import (
  "net"
  "fmt"
  "bufio"
  "strings"
)

type Response map[string]string

type Client struct {
  c net.Conn
  response Response
}

func (r *Client) Connect(HostPort string) (e error) {
  r.response = make(Response)
  r.c, e = net.Dial("tcp", HostPort)
  return e
}

func (r *Client) callAndResponse (command, key string) {
  fmt.Fprintf(r.c, "%s %s\n", command, key)
  scanner := bufio.NewScanner(r.c)
  for i := 0; i < 2; i++ {
    scanner.Scan()
    line := strings.SplitN(scanner.Text(), ": ", 2)
    r.response[line[0]] = line[1]
  }
  scanner.Scan()
  r.response["BODY"] = scanner.Text()
}

func (r *Client) Read(key string) (Response) {
  r.callAndResponse("read", key)
  return r.response
}
func (r *Client) Create(key string) (Response) {
  r.callAndResponse("create", key)
  return r.response
}

func (r *Client) Delete(key string) (Response) {
  r.callAndResponse("delete", key)
  return r.response
}

func (r *Client) Keys(key string) (Response) {
  r.callAndResponse("keys", key)
  return r.response
}

