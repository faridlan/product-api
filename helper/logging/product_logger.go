package logging

import (
	"fmt"
	"net"
	"net/http"

	"github.com/faridlan/product-api/model/web"
	"github.com/sirupsen/logrus"
)

func ProductLogger(webRespone web.WebResponse, writer http.ResponseWriter, request *http.Request) {

	logger := logrus.New()
	logger.SetFormatter(&logrus.TextFormatter{})
	ip, _, err := net.SplitHostPort(request.RemoteAddr)
	if err != nil {
		fmt.Fprintf(writer, "userip: %q is not IP:port", request.RemoteAddr)
	}

	userIP := net.ParseIP(ip)
	if userIP == nil {
		fmt.Fprintf(writer, "userip: %q is not IP:port", request.RemoteAddr)
		return
	}

	logger.WithFields(logrus.Fields{
		"user_ip": ip,
		"code":    webRespone.Code,
		"status":  webRespone.Status,
		// "data":    data,
		"method": request.Method,
	}).Infof("Product")

}

func ProductLoggerError(webRespone web.WebResponse, writer http.ResponseWriter, request *http.Request, errMessage any) {

	logger := logrus.New()
	logger.SetFormatter(&logrus.TextFormatter{})
	ip, _, err := net.SplitHostPort(request.RemoteAddr)
	if err != nil {
		fmt.Fprintf(writer, "userip: %q is not IP:port", request.RemoteAddr)
	}

	userIP := net.ParseIP(ip)
	if userIP == nil {
		fmt.Fprintf(writer, "userip: %q is not IP:port", request.RemoteAddr)
		return
	}

	logger.WithFields(logrus.Fields{
		"user_ip": ip,
		"code":    webRespone.Code,
		"status":  webRespone.Status,
		// "data":    data,
		"method":        request.Method,
		"error_message": errMessage,
	}).Error("Product")

}
