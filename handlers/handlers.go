package handlers

import (
	"fmt"
	"gambit_backend/auth"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
)

func Handlers(path, method, body string, headers map[string]string, request events.APIGatewayV2HTTPRequest) (int, string) {
	fmt.Println("processing" + path + ">" + method)

	id := request.PathParameters["id"]
	idn, _ := strconv.Atoi(id)

	isOk, statusCode, user := ValideAuthentication(path, method, headers)
	if !isOk {
		return statusCode, user
	}

	switch path[0:4] {
	case "user":
		return ProcessUser(body, path, method, user, id, request)

	case "prod":
		return ProcessProduct(body, path, method, user, idn, request)

	case "stoc":
		return ProcessStock(body, path, method, user, idn, request)

	case "addr":
		return ProcessAddress(body, path, method, user, idn, request)
	case "cate":
		return ProcessCategory(body, path, method, user, idn, request)
	case "orde":
		return ProcessOrders(body, path, method, user, idn, request)

	}

	return 400, "Method Invalid"

}

func ValideAuthentication(path, method string, headers map[string]string) (bool, int, string) {
	if (path == "product" && method == "GET") ||
		(path == "category" && method == "GET") {
		return true, 200, "OK"
	}

	token := headers["authorization"]
	if len(token) == 0 {
		return false, 401, "Token required"
	}

	requestOk, msg, err := auth.ValidToken(token)

	if !requestOk {
		if err != nil {
			fmt.Println("Error valide token " + err.Error())
			return false, 401, err.Error()
		} else {
			fmt.Println("Error valid token " + msg)
		}
		return false, 401, msg
	}

	fmt.Println("Token OK")
	return true, 200, msg
}

func ProcessUser(body, path, method, user, id string, request events.APIGatewayV2HTTPRequest) (int, string) {
	return 400, "Invalid"
}

func ProcessProduct(body, path, method, productuser string, id int, request events.APIGatewayV2HTTPRequest) (int, string) {
	return 400, "Invalid"
}

func ProcessCategory(body, path, method, productuser string, id int, request events.APIGatewayV2HTTPRequest) (int, string) {
	return 400, "Invalid"
}

func ProcessStock(body, path, method, productuser string, id int, request events.APIGatewayV2HTTPRequest) (int, string) {
	return 400, "Invalid"
}

func ProcessAddress(body, path, method, productuser string, id int, request events.APIGatewayV2HTTPRequest) (int, string) {
	return 400, "Invalid"
}

func ProcessOrders(body, path, method, productuser string, id int, request events.APIGatewayV2HTTPRequest) (int, string) {
	return 400, "Invalid"
}
