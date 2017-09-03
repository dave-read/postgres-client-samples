const pg = require('pg');

// NOTE needing to add SSL for this to run
const uri = process.env.uri + "?ssl=true"

//  using uri directly
const cfgClient = new pg.Client(uri);
cfgClient.connect(err => {
    if (err) throw err;
    else { queryDatabase(cfgClient); }
});

function queryDatabase(client) {

    const query = 'SELECT * FROM inventory;';

    console.log(`Running query to PostgreSQL server using uri: ${uri}`);
    
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
