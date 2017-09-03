import java.sql.*;
import java.util.Properties;
import java.net.URI;


public class ReadTable {
    public static void main(String[] args) throws Exception {

        // check that the driver is installed
        try {
            Class.forName("org.postgresql.Driver");
        } catch (ClassNotFoundException e) {
            throw new ClassNotFoundException("PostgreSQL JDBC driver NOT detected in library path.", e);
        }
        System.out.println("PostgreSQL JDBC driver detected in library path.");

        System.out.println("Query using properties");
        doQuery(
            getEnvVarOrExit("hostname"), 
            getEnvVarOrExit("database"), 
            getEnvVarOrExit("username"), 
            getEnvVarOrExit("password")
        );

        // get properties from URI
        String uriFromEnv = getEnvVarOrExit("uri");
        URI uri = null;
        try {
            uri = new URI(uriFromEnv);
        } catch (Exception e) {
            e.printStackTrace();
            System.out.println("Error parsing uri:"+uriFromEnv);
            System.exit(1);
        }
        String uriHost = uri.getHost();
        String uriDatabase = uri.getPath();
        uriDatabase = uriDatabase.startsWith("/") ? uriDatabase.substring(1) : uriDatabase;

        String userInfo = uri.getUserInfo();
        String[] userInfoParts = userInfo.split(":",2);
        String uriUser =userInfoParts[0];
        String uriPassword = userInfoParts[1];    
        System.out.println("Query using properties from uri:"+uri);
        doQuery(uriHost, uriDatabase, uriUser, uriPassword);
      
    }

    private static void doQuery(String host, String database, String user, String password) {
        Connection connection = null;
        try {         
            String url = String.format("jdbc:postgresql://%s/%s", host, database);
            System.out.println("jdbc url:"+url);
            Properties properties = new Properties();
            properties.setProperty("user", user);
            properties.setProperty("password", password);
            properties.setProperty("ssl", "true");   

            connection = DriverManager.getConnection(url, properties);
            if(connection != null) { 
                System.out.println("Successfully created connection to database.");
            
                Statement statement = connection.createStatement();
                ResultSet results = statement.executeQuery("SELECT * from inventory;");
                while (results.next()) {
                    String outputString = 
                        String.format("Data row = (%s, %s, %s)", results.getString(1),
                    results.getString(2), results.getString(3));
                    System.out.println(outputString);
                }
                System.out.println("");
            }
        }
        catch (SQLException e) {
            throw new RuntimeException("Encountered an error when executing given sql statement.", e);
        } 
        finally {
            if(connection != null) {
                try { connection.close();} catch(Exception e) {}
            }
        }
    }    
    private static String getEnvVarOrExit(String varName) {
        String val = System.getenv(varName);
        if(val ==null) {
            System.out.println("No value for environment variable:" + varName);
            System.exit(1);
        }
        return val;
    }


}