// El paquete bd se encarga de establecer la conexión con la base de datos.
package bd

import (
	"database/sql" // Para operar con SQL
	"fmt"
	"gambit_backend/models"
	"gambit_backend/secretm"
	"os"

	_ "github.com/go-sql-driver/mysql" // Driver de MySQL para el paquete database/sql
)

// SecretModel es una variable global que guarda los detalles del secreto recuperado de AWS Secrets Manager.
// Es de tipo SecretRDSJson definido en el paquete models.
var SecretModel models.SecretRDSJson

// err es una variable global que guarda cualquier error que pueda surgir en las funciones del paquete.
var err error

// Db es un puntero a una instancia de sql.DB, que representa una agrupación de conexiones de base de datos.
var Db *sql.DB

// ReadSecret es una función que recupera un secreto de AWS Secrets Manager.
// Obtiene el nombre del secreto de la variable de entorno "SecretName" y utiliza la función GetSecret del paquete secretm para obtenerlo.
// El resultado se almacena en la variable global SecretModel.
// En caso de error, se asigna a la variable global err y se retorna.
// Si no hay errores, retorna nil.
func ReadSecret() error {
	SecretModel, err = secretm.GetSecret(os.Getenv("SecretName"))
	return err
}

// DbConnect establece una conexión a la base de datos.
// Genera una cadena de conexión utilizando los detalles de SecretModel y la utiliza para abrir una nueva conexión.
// Luego intenta hacer ping a la base de datos para verificar que la conexión está activa.
// Si hay un error, lo imprime y lo retorna.
// Si no hay errores, imprime un mensaje de éxito y retorna nil.
func DbConnect() error {
	Db, err = sql.Open("mysql", ConnStr(SecretModel))
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	err = Db.Ping()
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("Successful connection DB")
	return nil
}

// ConnStr genera una cadena de conexión a la base de datos a partir de los detalles almacenados en SecretModel.
// Se deben tener en cuenta las mejores prácticas de seguridad al imprimir la cadena de conexión,
// especialmente en un entorno de producción.
func ConnStr(keys models.SecretRDSJson) string {
	var dbUser, authToken, dbEndPoint, dbName string
	dbUser = keys.Username
	authToken = keys.Password
	dbEndPoint = keys.Host
	dbName = "gambitsas"
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?allowCleartextPasswords=true", dbUser, authToken, dbEndPoint, dbName)
	return dsn
}
