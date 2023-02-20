package logging

import (
	"fmt"
	"net"
	"net/http"
	"os"

	"github.com/faridlan/product-api/helper"
	"github.com/faridlan/product-api/model/web"
	"github.com/sirupsen/logrus"
)

func ProductLogger(webRespone web.WebResponse, writer http.ResponseWriter, request *http.Request) {

	logger := logrus.New()
	logger.SetFormatter(&logrus.TextFormatter{})

	logger.SetOutput(os.Stdout)

	file, err := os.OpenFile("helper/logging/logger.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	helper.PanicIfErr(err)

	defer file.Close()

	logger.SetOutput(file)

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
		"data":    webRespone.Data,
		"method":  request.Method,
	}).Infof("Product")

}

func ProductLoggerError(webRespone web.WebResponse, writer http.ResponseWriter, request *http.Request, errMessage any) {

	logger := logrus.New()
	logger.SetFormatter(&logrus.TextFormatter{})

	logger.SetOutput(os.Stdout)

	file, err := os.OpenFile("helper/logging/logger.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	helper.PanicIfErr(err)

	defer file.Close()

	logger.SetOutput(file)

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
		"user_ip":       ip,
		"code":          webRespone.Code,
		"status":        webRespone.Status,
		"method":        request.Method,
		"error_message": errMessage,
	}).Error("Product")

}
