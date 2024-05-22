package getpeers

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"net/url"
)

const (
	ENODE_FMT         = "enode://%s@%s:%d"
	ADMIN_ADDPEER_FMT = `admin.addPeer("enode://%s@%s:%d")`
	NET_MAINNET       = 1
	NET_TESTNET       = 2
)

type NodeData struct {
	Id            string `json:"id"`
	Host          string `json:"host"`
	Port          int    `json:"port"`
	ClientId      string `json:"clientId"`
	Client        string `json:"client"`
	ClientVersion string `json:"clientVersion"`
	Os            string `json:"os"`
	Country       string `json:"country"`
}

type Result struct {
	Data []NodeData `json:"data"`
}

func GetEthnodes(OutputFmt string, net, start, length int) ([]string, error) {

	res := make([]string, 0)

	// url := "https://www.ethernodes.org/data?draw=1&columns%5B0%5D%5Bdata%5D=id&columns%5B0%5D%5Bname%5D=&columns%5B0%5D%5Bsearchable%5D=true&columns%5B0%5D%5Borderable%5D=true&columns%5B0%5D%5Bsearch%5D%5Bvalue%5D=&columns%5B0%5D%5Bsearch%5D%5Bregex%5D=false&columns%5B1%5D%5Bdata%5D=host&columns%5B1%5D%5Bname%5D=&columns%5B1%5D%5Bsearchable%5D=true&columns%5B1%5D%5Borderable%5D=true&columns%5B1%5D%5Bsearch%5D%5Bvalue%5D=&columns%5B1%5D%5Bsearch%5D%5Bregex%5D=false&columns%5B2%5D%5Bdata%5D=isp&columns%5B2%5D%5Bname%5D=&columns%5B2%5D%5Bsearchable%5D=true&columns%5B2%5D%5Borderable%5D=true&columns%5B2%5D%5Bsearch%5D%5Bvalue%5D=&columns%5B2%5D%5Bsearch%5D%5Bregex%5D=false&columns%5B3%5D%5Bdata%5D=country&columns%5B3%5D%5Bname%5D=&columns%5B3%5D%5Bsearchable%5D=true&columns%5B3%5D%5Borderable%5D=true&columns%5B3%5D%5Bsearch%5D%5Bvalue%5D=&columns%5B3%5D%5Bsearch%5D%5Bregex%5D=false&columns%5B4%5D%5Bdata%5D=client&columns%5B4%5D%5Bname%5D=&columns%5B4%5D%5Bsearchable%5D=true&columns%5B4%5D%5Borderable%5D=true&columns%5B4%5D%5Bsearch%5D%5Bvalue%5D=&columns%5B4%5D%5Bsearch%5D%5Bregex%5D=false&columns%5B5%5D%5Bdata%5D=clientVersion&columns%5B5%5D%5Bname%5D=&columns%5B5%5D%5Bsearchable%5D=true&columns%5B5%5D%5Borderable%5D=true&columns%5B5%5D%5Bsearch%5D%5Bvalue%5D=&columns%5B5%5D%5Bsearch%5D%5Bregex%5D=false&columns%5B6%5D%5Bdata%5D=os&columns%5B6%5D%5Bname%5D=&columns%5B6%5D%5Bsearchable%5D=true&columns%5B6%5D%5Borderable%5D=true&columns%5B6%5D%5Bsearch%5D%5Bvalue%5D=&columns%5B6%5D%5Bsearch%5D%5Bregex%5D=false&columns%5B7%5D%5Bdata%5D=lastUpdate&columns%5B7%5D%5Bname%5D=&columns%5B7%5D%5Bsearchable%5D=true&columns%5B7%5D%5Borderable%5D=true&columns%5B7%5D%5Bsearch%5D%5Bvalue%5D=&columns%5B7%5D%5Bsearch%5D%5Bregex%5D=false&columns%5B8%5D%5Bdata%5D=inSync&columns%5B8%5D%5Bname%5D=&columns%5B8%5D%5Bsearchable%5D=true&columns%5B8%5D%5Borderable%5D=true&columns%5B8%5D%5Bsearch%5D%5Bvalue%5D=&columns%5B8%5D%5Bsearch%5D%5Bregex%5D=false&order%5B0%5D%5Bcolumn%5D=0&order%5B0%5D%5Bdir%5D=asc&start=0&length=10&search%5Bvalue%5D=&search%5Bregex%5D=false&_=1716189654844"
	// resp, err := http.Get(url)
	// if err != nil {
	// 	return res, err
	// }

	// // Close the response once we return from the function.
	// defer resp.Body.Close()

	baseURL := "https://www.ethernodes.org/data"

	queryParams := url.Values{}

	queryParams.Add("draw", "7")
	queryParams.Add("columns[0][data]", "id")
	queryParams.Add("columns[0][name]", "")
	queryParams.Add("columns[0][searchable]", "true")
	queryParams.Add("columns[0][orderable]", "true")
	queryParams.Add("columns[0][search][value]", "")
	queryParams.Add("columns[0][search][regex]", "false")
	queryParams.Add("columns[1][data]", "host")
	queryParams.Add("columns[1][name]", "")
	queryParams.Add("columns[1][searchable]", "true")
	queryParams.Add("columns[1][orderable]", "true")
	queryParams.Add("columns[1][search][value]", "")
	queryParams.Add("columns[1][search][regex]", "false")
	queryParams.Add("columns[2][data]", "isp")
	queryParams.Add("columns[2][name]", "")
	queryParams.Add("columns[2][searchable]", "true")
	queryParams.Add("columns[2][orderable]", "true")
	queryParams.Add("columns[2][search][value]", "")
	queryParams.Add("columns[2][search][regex]", "false")
	queryParams.Add("columns[3][data]", "country")
	queryParams.Add("columns[3][name]", "")
	queryParams.Add("columns[3][searchable]", "true")
	queryParams.Add("columns[3][orderable]", "true")
	queryParams.Add("columns[3][search][value]", "")
	queryParams.Add("columns[3][search][regex]", "false")
	queryParams.Add("columns[4][data]", "client")
	queryParams.Add("columns[4][name]", "")
	queryParams.Add("columns[4][searchable]", "true")
	queryParams.Add("columns[4][orderable]", "true")
	queryParams.Add("columns[4][search][value]", "")
	queryParams.Add("columns[4][search][regex]", "false")
	queryParams.Add("columns[5][data]", "clientVersion")
	queryParams.Add("columns[5][name]", "")
	queryParams.Add("columns[5][searchable]", "true")
	queryParams.Add("columns[5][orderable]", "true")
	queryParams.Add("columns[5][search][value]", "")
	queryParams.Add("columns[5][search][regex]", "false")
	queryParams.Add("columns[6][data]", "os")
	queryParams.Add("columns[6][name]", "")
	queryParams.Add("columns[6][searchable]", "true")
	queryParams.Add("columns[6][orderable]", "true")
	queryParams.Add("columns[6][search][value]", "")
	queryParams.Add("columns[6][search][regex]", "false")
	queryParams.Add("columns[7][data]", "lastUpdate")
	queryParams.Add("columns[7][name]", "")
	queryParams.Add("columns[7][searchable]", "true")
	queryParams.Add("columns[7][orderable]", "true")
	queryParams.Add("columns[7][search][value]", "")
	queryParams.Add("columns[7][search][regex]", "false")
	queryParams.Add("columns[8][data]", "inSync")
	queryParams.Add("columns[8][name]", "")
	queryParams.Add("columns[8][searchable]", "true")
	queryParams.Add("columns[8][orderable]", "true")
	queryParams.Add("columns[8][search][value]", "")
	queryParams.Add("columns[8][search][regex]", "false")
	queryParams.Add("order[0][column]", "3")
	queryParams.Add("order[0][dir]", "asc")
	queryParams.Add("start", "0")
	queryParams.Add("length", "100")
	queryParams.Add("search[value]", "chi")
	queryParams.Add("search[regex]", "false")
	queryParams.Add("_", "1716190551952")

	fullURL := baseURL + "?" + queryParams.Encode()

	method := "GET"
	client := &http.Client{}
	req, err := http.NewRequest(method, fullURL, nil)

	if err != nil {
		fmt.Println(err)
		return res, err
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return res, err
	}
	defer resp.Body.Close()

	// Check the status code for a 200 so we know we have received a
	// proper response.
	if resp.StatusCode != 200 {
		return res, fmt.Errorf("HTTP Response Error %d\n", resp.StatusCode)
	}

	// Decode the rss feed document into our struct type.
	// We don't need to check for errors, the caller can do this.
	var result Result
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return res, err
	}
	if len(result.Data) == 0 {
		return res, fmt.Errorf("Error data format of result")
	}

	for _, x := range result.Data {
		res = append(res, fmt.Sprintf(OutputFmt, x.Id, x.Host, x.Port))
	}

	return res, nil
}

var flagStart int
var flagLength int
var flagFmt string
var flagNet string

func init() {
	flag.StringVar(&flagNet, "net", "mainnet", "mainnet or testnet")
	flag.IntVar(&flagStart, "start", 0, "start index")
	flag.IntVar(&flagLength, "length", 10, "length ")
	flag.StringVar(&flagFmt, "fmt", "addpeer", "enode or addpeer")
}

func GetPeers() {
	flag.Parse()

	fmtString := ADMIN_ADDPEER_FMT
	if flagFmt == "enode" {
		fmtString = ENODE_FMT
	}

	net := NET_MAINNET
	if flagNet == "testnet" {
		net = NET_TESTNET
	}

	l, err := GetEthnodes(fmtString, net, flagStart, flagLength)
	if err != nil {
		fmt.Printf("%v", err)
	}

	for _, e := range l {
		fmt.Println(e)
	}

}
