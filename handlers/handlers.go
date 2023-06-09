package handlers

import (
	"fmt"
	"gambit_backend/auth"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
)

// Handlers maneja las solicitudes HTTP entrantes
func Handlers(path, method, body string, headers map[string]string, request events.APIGatewayV2HTTPRequest) (int, string) {
	// Registra que se está procesando la solicitud
	fmt.Println("processing" + path + ">" + method)

	// Extrae el parámetro 'id' de la ruta de la solicitud
	id := request.PathParameters["id"]
	// Convierte 'id' a un número entero (ignora el error si no se puede convertir)
	idn, _ := strconv.Atoi(id)

	// Valida si la solicitud es auténtica
	isOk, statusCode, user := ValideAuthentication(path, method, headers)
	// Si la solicitud no es auténtica, devuelve el código de estado y el mensaje de error
	if !isOk {
		return statusCode, user
	}

	// Dependiendo del prefijo de la ruta, procesa la solicitud correspondiente
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

	default:
		// Si el prefijo de la ruta no corresponde a ninguno conocido, devuelve un error 400 "Método inválido"
		return 400, "Method Invalid"
	}
}

// ValideAuthentication verifica la autenticación de una solicitud
func ValideAuthentication(path, method string, headers map[string]string) (bool, int, string) {
	// Permite el acceso sin autenticación a los métodos GET para 'product' y 'category'
	if (path == "product" && method == "GET") ||
		(path == "category" && method == "GET") {
		return true, 200, "OK"
	}

	// Extrae el token de los encabezados de la solicitud
	token := headers["authorization"]
	// Si no se proporciona un token, devuelve un error 401 "Token requerido"
	if len(token) == 0 {
		return false, 401, "Token required"
	}

	// Valida el token
	requestOk, msg, err := auth.ValidToken(token)
	// Si la validación del token falla, imprime un mensaje de error y devuelve el error
	if !requestOk {
		if err != nil {
			fmt.Println("Error valide token " + err.Error())
			return false, 401, err.Error()
		} else {
			fmt.Println("Error valid token " + msg)
		}
		return false, 401, msg
	}

	// Si el token es válido, imprime "Token OK" y devuelve un 200 "OK"
	fmt.Println("Token OK")
	return true, 200, msg
}

// Las siguientes funciones procesan diferentes tipos de solicitudes
// Por ahora, simplemente devuelven un error 400 "Inválido"
// Deberían ser reemplazadas por funciones que realmente manejen las solicitudes correspondientes

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
