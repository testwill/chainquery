/*
 * Chain Query
 *
 * The LBRY blockchain is read into SQL where important structured information can be extracted through the Chain Query API.
 *
 * API version: 0.1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

import (
	"net/http"

	"github.com/lbryio/chainquery/config"
	"github.com/lbryio/chainquery/db"
	sw "github.com/lbryio/chainquery/swagger/apiserver/go"

	"github.com/sirupsen/logrus"
)

func InitApiServer(hostAndPort string) {
	logrus.Info("API Server started")
	//API Chainquery DB connection
	chainqueryInstance, err := db.InitAPIQuery(config.GetAPIMySQLDSN(), false)
	if err != nil {
		logrus.Error("unable to connect to chainquery database instance for API Server: ", err)
	}
	defer db.CloseDB(chainqueryInstance)
	router := sw.NewRouter()

	logrus.Fatal(http.ListenAndServe(hostAndPort, router))
}
