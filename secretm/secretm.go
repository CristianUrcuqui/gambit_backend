package secretm

import (
	"encoding/json"
	"fmt"
	"gambit_backend/awsgo"
	"gambit_backend/models"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
)

// Define la función GetSecret que toma un nombre de secreto como argumento
func GetSecret(nameSecret string) (models.SecretRDSJson, error) {
	var dataSecret models.SecretRDSJson        // Define la variable dataSecret de tipo SecretRDSJson
	fmt.Println(" > get Secret " + nameSecret) // Imprime el nombre del secreto que vamos a obtener

	svc := secretsmanager.NewFromConfig(awsgo.Cfg) // Crea un nuevo cliente de Secrets Manager utilizando la configuración de AWS
	// Intenta obtener el valor del secreto utilizando el nombre del secreto proporcionado
	keys, err := svc.GetSecretValue(awsgo.Ctx, &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(nameSecret),
	})
	// Si hay un error, lo imprime y lo devuelve junto con un SecretRDSJson vacío
	if err != nil {
		fmt.Println(err.Error())
		return dataSecret, err
	}

	// Deserializa el valor del secreto (cadena JSON) en un objeto SecretRDSJson
	json.Unmarshal([]byte(*keys.SecretString), &dataSecret)
	fmt.Println(">reading secret Ok" + nameSecret) // Imprime un mensaje que indica que la lectura del secreto fue exitosa

	return dataSecret, nil // Devuelve el objeto SecretRDSJson y nil para el error
}
