import express = require("express")
const router: express.Router = express.Router();
import { conDB } from "../db/connect"

router.route("/get/:table").get((req: express.Request, res: express.Response) => {

	conDB.query(`SELECT * FROM ${req.params.table}`,
		(err: any, data: any) => {
			console.log(`\"/api/get/${req.params.table}\" accessed`)
			res.send(data.rows);
		})
});

module.exports = router;