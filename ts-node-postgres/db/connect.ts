const { Pool } = require('pg')

const conDB = new Pool({
	user: process.env.PGUSER,
	host: process.env.PGHOST,
	database: process.env.PGDB,
	password: process.env.PGPASS,
	port: process.env.PGPORT
})

export { conDB };