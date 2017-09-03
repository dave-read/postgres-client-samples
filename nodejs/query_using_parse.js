const pg = require('pg');
const parse = require('pg-connection-string').parse;

// uri from environment
// NOTE needing to add SSL for this to run
const uri = process.env.uri + "?ssl=true"
// parse to config
var config = parse(uri)

//  using separate properties
const cfgClient = new pg.Client(config);
cfgClient.connect(err => {
    if (err) throw err;
    else { queryDatabase(cfgClient); }
});

function queryDatabase(client) {

    const query = 'SELECT * FROM inventory;';

    console.log(`Running query to PostgreSQL server using config: ${JSON.stringify(config)}`);
    
    client.query(query)
        .then(res => {
            const rows = res.rows;

            rows.map(row => {
                console.log(`Read: ${JSON.stringify(row)}`);
            });
            process.exit();
        })
        .catch(err => {
            console.log(err);
        });
}
