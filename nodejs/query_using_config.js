const pg = require('pg');

const config = {
    host: process.env.hostname,
    // Do not hard code your username and password.
    // Consider using Node environment variables.
    user: process.env.username,     
    password: process.env.password,
    database: process.env.database,
    port: 5432,
    ssl: true
};

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
