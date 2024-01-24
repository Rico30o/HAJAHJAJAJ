package routes

import (
	"instapay/Trace_alert/routes"

	"github.com/gofiber/fiber/v2"
)

func SetupUserRoutes(app *fiber.App) {

	//--------------------------- Endpoint ------------------------------//

	//trace
	trace := app.Group("/financial-crime")
	tracenetwork := trace.Group("/networks")
	tracevisualisations := trace.Group("/visualisations")

	//alerts
	alerts := app.Group("/alerts")
	alertaccount := alerts.Group("/account")
	// alerttransaction := alert.Group("/transactions")
	// alertnetworks := alert.Group("/networks")

	//------------------------- Feedback --------------------------------//
	//feedback
	app.Post("/feedback", routes.Feedback)

	//--------------------------- Trace ---------------------------------//

	//trace
	tracenetwork.Post("/tracenetwork", routes.Tracenetwork)
	tracevisualisations.Post("/network", routes.NetworkAlertID)
	//--------------------------- Trace ---------------------------------//

	//alert
	alerts.Post("/accounts", routes.Alertsaccount)
	alerts.Post("/transactions", routes.Alerttransaction)
	alerts.Post("/networks", routes.Alertnetwork)
	//alertID
	alertaccount.Post("/accountalertid", routes.GetAccInfo)
	alerts.Post("/id", routes.TransactionIDmatches)

	//--------------------------- Janus --------------------------------//
	app.Post("/janusloan", routes.CreateResponse)

}
