package auth

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type TokenJSON struct {
	Sub       string
	Event_ID  string
	Token_use string
	Scope     string
	Auth_time int
	Iss       string
	Exp       int
	Iat       int
	Client_id string
	Username  string
}

// ValidToken valida un token JWT
func ValidToken(token string) (bool, string, error) {
	// El token JWT se divide en tres partes: Encabezado, Carga útil y Firma
	parts := strings.Split(token, ".")

	// Si el token no tiene exactamente tres partes, no es un token JWT válido
	if len(parts) != 3 {
		fmt.Println("Invalid token")
		// Retorna falso, un mensaje de error y ningún error
		return false, "Invalid token", nil
	}

	// Decodifica la carga útil del token (que está en formato Base64)
	userInfo, err := base64.StdEncoding.DecodeString(parts[1])
	// Si ocurre un error al decodificar, retorna falso, el mensaje de error y el error
	if err != nil {
		fmt.Println("decode fail" + err.Error())
		return false, err.Error(), err
	}

	// Crea un objeto TokenJSON vacío para almacenar la carga útil decodificada
	var tkj TokenJSON
	// Deserializa la carga útil decodificada en el objeto TokenJSON
	err = json.Unmarshal(userInfo, &tkj)
	// Si ocurre un error al deserializar, retorna falso, el mensaje de error y el error
	if err != nil {
		fmt.Println("decode fail" + err.Error())
		return false, err.Error(), err
	}

	// Obtiene la hora actual
	here := time.Now()
	// Convierte el tiempo de expiración del token a un objeto Time
	tm := time.Unix(int64(tkj.Exp), 0)

	// Si el tiempo de expiración del token es anterior a la hora actual, el token ha expirado
	if tm.Before(here) {
		fmt.Println("Date expiration token = " + tm.String())
		fmt.Println("Token expiration !")
		// Retorna falso, un mensaje de error y ningún error
		return false, "Token expiration", err
	}

	// Si el token es válido y no ha expirado, retorna verdadero, el nombre de usuario y ningún error
	return true, string(tkj.Username), nil
}
