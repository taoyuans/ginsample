package httpreq

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"

	"ginsample/lib/factory"
)

type contentType struct {
	MIMEApplicationJSON                  string
	MIMEApplicationJSONCharsetUTF8       string
	MIMEApplicationJavaScript            string
	MIMEApplicationJavaScriptCharsetUTF8 string
	MIMEApplicationXML                   string
	MIMEApplicationXMLCharsetUTF8        string
	MIMETextXML                          string
	MIMETextXMLCharsetUTF8               string
	MIMEApplicationForm                  string
	MIMEApplicationFormUTF8              string
	MIMEApplicationProtobuf              string
	MIMEApplicationMsgpack               string
	MIMETextHTML                         string
	MIMETextHTMLCharsetUTF8              string
	MIMETextPlain                        string
	MIMETextPlainCharsetUTF8             string
	MIMEMultipartForm                    string
	MIMEOctetStream                      string
}

const (
	charsetUTF8 = "charset=UTF-8"
)

type Header struct {
	ContentType string
	Token       string
}

var ContentType = contentType{
	MIMEApplicationJSON:                  "application/json",
	MIMEApplicationJSONCharsetUTF8:       "application/json" + "; " + charsetUTF8,
	MIMEApplicationJavaScript:            "application/javascript",
	MIMEApplicationJavaScriptCharsetUTF8: "application/javascript" + "; " + charsetUTF8,
	MIMEApplicationXML:                   "application/xml",
	MIMEApplicationXMLCharsetUTF8:        "application/xml" + "; " + charsetUTF8,
	MIMETextXML:                          "text/xml",
	MIMETextXMLCharsetUTF8:               "text/xml" + "; " + charsetUTF8,
	MIMEApplicationForm:                  "application/x-www-form-urlencoded",
	MIMEApplicationFormUTF8:              "application/x-www-form-urlencoded" + "; " + charsetUTF8,
	MIMEApplicationProtobuf:              "application/protobuf",
	MIMEApplicationMsgpack:               "application/msgpack",
	MIMETextHTML:                         "text/html",
	MIMETextHTMLCharsetUTF8:              "text/html" + "; " + charsetUTF8,
	MIMETextPlain:                        "text/plain",
	MIMETextPlainCharsetUTF8:             "text/plain" + "; " + charsetUTF8,
	MIMEMultipartForm:                    "multipart/form-data",
	MIMEOctetStream:                      "application/octet-stream",
}

func Get(ctx context.Context, url string, header *Header, transport *http.Transport) ([]byte, error) {
	httpReq, err := http.NewRequest("GET", url, nil)
	if header != nil {
		httpReq.Header.Set("Content-Type", header.ContentType)
	} else {
		httpReq.Header.Set("Content-Type", ContentType.MIMEApplicationJSONCharsetUTF8)
	}

	if requestId := factory.RequestId(ctx); len(requestId) > 0 {
		httpReq.Header.Set("X-Request-Id", requestId)
	}
	if token := factory.Token(ctx); len(token) > 0 {
		httpReq.Header.Set("Authorization", "Bearer "+token)
	}

	var client *http.Client
	if transport != nil {
		client = &http.Client{Transport: transport}
	} else {
		client = &http.Client{}
	}
	resp, err := client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, err
	}

	respData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return respData, nil
}

func Post(ctx context.Context, url string, param []byte, header *Header, transport *http.Transport) ([]byte, error) {
	httpReq, err := http.NewRequest("POST", url, bytes.NewBuffer(param))

	if header != nil {
		httpReq.Header.Set("Content-Type", header.ContentType)
	} else {
		httpReq.Header.Set("Content-Type", ContentType.MIMEApplicationJSONCharsetUTF8)
	}

	if requestId := factory.RequestId(ctx); len(requestId) > 0 {
		httpReq.Header.Set("X-Request-Id", requestId)
	}
	if token := factory.Token(ctx); len(token) > 0 {
		httpReq.Header.Set("Authorization", "Bearer "+token)
	}

	var client *http.Client
	if transport != nil {
		client = &http.Client{Transport: transport}
	} else {
		client = &http.Client{}
	}
	resp, err := client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, err
	}
	respData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return respData, nil
}
