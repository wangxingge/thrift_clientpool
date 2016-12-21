// Autogenerated by Thrift Compiler (0.9.3)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package main

import (
	"flag"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"math"
	"net"
	"net/url"
	"os"
	"strconv"
	"strings"
	"github.com/wangxingge/thrift_clientpool/examples/bookservice"
)

func Usage() {
	fmt.Fprintln(os.Stderr, "Usage of ", os.Args[0], " [-h host:port] [-u url] [-f[ramed]] function [arg1 [arg2...]]:")
	flag.PrintDefaults()
	fmt.Fprintln(os.Stderr, "\nFunctions:")
	fmt.Fprintln(os.Stderr, "   GetUserBooks(string userId)")
	fmt.Fprintln(os.Stderr, "  User GetUserInfo(string userId)")
	fmt.Fprintln(os.Stderr, "   GetAllUserInfo()")
	fmt.Fprintln(os.Stderr, "  bool AddUser(User userInfo)")
	fmt.Fprintln(os.Stderr, "  bool RemoveUser(string userId)")
	fmt.Fprintln(os.Stderr, "  bool UpdateUserAvatar(string userId, string avatar)")
	fmt.Fprintln(os.Stderr)
	os.Exit(0)
}

func main() {
	flag.Usage = Usage
	var host string
	var port int
	var protocol string
	var urlString string
	var framed bool
	var useHttp bool
	var parsedUrl url.URL
	var trans thrift.TTransport
	_ = strconv.Atoi
	_ = math.Abs
	flag.Usage = Usage
	flag.StringVar(&host, "h", "localhost", "Specify host and port")
	flag.IntVar(&port, "p", 9090, "Specify port")
	flag.StringVar(&protocol, "P", "binary", "Specify the protocol (binary, compact, simplejson, json)")
	flag.StringVar(&urlString, "u", "", "Specify the url")
	flag.BoolVar(&framed, "framed", false, "Use framed transport")
	flag.BoolVar(&useHttp, "http", false, "Use http")
	flag.Parse()

	if len(urlString) > 0 {
		parsedUrl, err := url.Parse(urlString)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error parsing URL: ", err)
			flag.Usage()
		}
		host = parsedUrl.Host
		useHttp = len(parsedUrl.Scheme) <= 0 || parsedUrl.Scheme == "http"
	} else if useHttp {
		_, err := url.Parse(fmt.Sprint("http://", host, ":", port))
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error parsing URL: ", err)
			flag.Usage()
		}
	}

	cmd := flag.Arg(0)
	var err error
	if useHttp {
		trans, err = thrift.NewTHttpClient(parsedUrl.String())
	} else {
		portStr := fmt.Sprint(port)
		if strings.Contains(host, ":") {
			host, portStr, err = net.SplitHostPort(host)
			if err != nil {
				fmt.Fprintln(os.Stderr, "error with host:", err)
				os.Exit(1)
			}
		}
		trans, err = thrift.NewTSocket(net.JoinHostPort(host, portStr))
		if err != nil {
			fmt.Fprintln(os.Stderr, "error resolving address:", err)
			os.Exit(1)
		}
		if framed {
			trans = thrift.NewTFramedTransport(trans)
		}
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error creating transport", err)
		os.Exit(1)
	}
	defer trans.Close()
	var protocolFactory thrift.TProtocolFactory
	switch protocol {
	case "compact":
		protocolFactory = thrift.NewTCompactProtocolFactory()
		break
	case "simplejson":
		protocolFactory = thrift.NewTSimpleJSONProtocolFactory()
		break
	case "json":
		protocolFactory = thrift.NewTJSONProtocolFactory()
		break
	case "binary", "":
		protocolFactory = thrift.NewTBinaryProtocolFactoryDefault()
		break
	default:
		fmt.Fprintln(os.Stderr, "Invalid protocol specified: ", protocol)
		Usage()
		os.Exit(1)
	}
	client := bookservice.NewUserServiceClientFactory(trans, protocolFactory)
	if err := trans.Open(); err != nil {
		fmt.Fprintln(os.Stderr, "Error opening socket to ", host, ":", port, " ", err)
		os.Exit(1)
	}

	switch cmd {
	case "GetUserBooks":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "GetUserBooks requires 1 args")
			flag.Usage()
		}
		argvalue0 := flag.Arg(1)
		value0 := argvalue0
		fmt.Print(client.GetUserBooks(value0))
		fmt.Print("\n")
		break
	case "GetUserInfo":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "GetUserInfo requires 1 args")
			flag.Usage()
		}
		argvalue0 := flag.Arg(1)
		value0 := argvalue0
		fmt.Print(client.GetUserInfo(value0))
		fmt.Print("\n")
		break
	case "GetAllUserInfo":
		if flag.NArg()-1 != 0 {
			fmt.Fprintln(os.Stderr, "GetAllUserInfo requires 0 args")
			flag.Usage()
		}
		fmt.Print(client.GetAllUserInfo())
		fmt.Print("\n")
		break
	case "AddUser":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "AddUser requires 1 args")
			flag.Usage()
		}
		arg40 := flag.Arg(1)
		mbTrans41 := thrift.NewTMemoryBufferLen(len(arg40))
		defer mbTrans41.Close()
		_, err42 := mbTrans41.WriteString(arg40)
		if err42 != nil {
			Usage()
			return
		}
		factory43 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt44 := factory43.GetProtocol(mbTrans41)
		argvalue0 := bookservice.NewUser()
		err45 := argvalue0.Read(jsProt44)
		if err45 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.AddUser(value0))
		fmt.Print("\n")
		break
	case "RemoveUser":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "RemoveUser requires 1 args")
			flag.Usage()
		}
		argvalue0 := flag.Arg(1)
		value0 := argvalue0
		fmt.Print(client.RemoveUser(value0))
		fmt.Print("\n")
		break
	case "UpdateUserAvatar":
		if flag.NArg()-1 != 2 {
			fmt.Fprintln(os.Stderr, "UpdateUserAvatar requires 2 args")
			flag.Usage()
		}
		argvalue0 := flag.Arg(1)
		value0 := argvalue0
		argvalue1 := []byte(flag.Arg(2))
		value1 := argvalue1
		fmt.Print(client.UpdateUserAvatar(value0, value1))
		fmt.Print("\n")
		break
	case "":
		Usage()
		break
	default:
		fmt.Fprintln(os.Stderr, "Invalid function ", cmd)
	}
}
