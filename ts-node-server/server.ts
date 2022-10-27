import express = require("express")
const app: express.Application = express();
require("dotenv").config();

app.get("/", (request: express.Request, response: express.Response) => {
	console.log("\"/\" url accessed");
	response.send({ "msg": "Hello world" })
})

let port: number | string = process.env.PORT || 3000;
let listener: any = app.listen(port, function () {
	console.log('App is listening on port ' + listener.address().port);
});
