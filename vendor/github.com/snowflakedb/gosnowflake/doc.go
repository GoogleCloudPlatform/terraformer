/*
Package gosnowflake is a pure Go Snowflake driver for the database/sql package.

Clients can use the database/sql package directly. For example:

	import (
		"database/sql"

		_ "github.com/snowflakedb/gosnowflake"
	)

	func main() {
		db, err := sql.Open("snowflake", "user:password@myaccount/mydb")
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()
		...
	}

Connection String

Use Open to create a database handle with connection parameters:

	db, err := sql.Open("snowflake", "<connection string>")

The Go Snowflake Driver supports the following connection syntaxes (or data source name formats):

	* username[:password]@accountname/dbname/schemaname[?param1=value&...&paramN=valueN
	* username[:password]@accountname/dbname[?param1=value&...&paramN=valueN
	* username[:password]@hostname:port/dbname/schemaname?account=<your_account>[&param1=value&...&paramN=valueN]

where all parameters must be escaped or use `Config` and `DSN` to construct a DSN string.

The following example opens a database handle with the Snowflake account
myaccount where the username is jsmith, password is mypassword, database is
mydb, schema is testschema, and warehouse is mywh:

	db, err := sql.Open("snowflake", "jsmith:mypassword@myaccount/mydb/testschema?warehouse=mywh")

Connection Parameters

The following connection parameters are supported:

	* account <string>: Specifies the name of your Snowflake account, where string is the name
		assigned to your account by Snowflake. In the URL you received from
		Snowflake, your account name is the first segment in the domain (e.g.
		abc123 in https://abc123.snowflakecomputing.com). This parameter is
		optional if your account is specified after the @ character. If you
		are not on us-west-2 region or AWS deployment, then append the region
		after the account name, e.g. “<account>.<region>”. If you are not on
		AWS deployment, then append not only the region, but also the platform,
		e.g., “<account>.<region>.<platform>”. Account, region, and platform
		should be separated by a period (“.”), as shown above.

	* region <string>: DEPRECATED. You may specify a region, such as
		“eu-central-1”, with this parameter. However, since this parameter
		is deprecated, it is best to specify the region as part of the
		account parameter. For details, see the description of the account
		parameter.

	* database: Specifies the database to use by default in the client session
		(can be changed after login).

	* schema: Specifies the database schema to use by default in the client
		session (can be changed after login).

	* warehouse: Specifies the virtual warehouse to use by default for queries,
		loading, etc. in the client session (can be changed after login).

	* role: Specifies the role to use by default for accessing Snowflake
		objects in the client session (can be changed after login).

	* passcode: Specifies the passcode provided by Duo when using MFA for login.

	* passcodeInPassword: false by default. Set to true if the MFA passcode is
		embedded in the login password. Appends the MFA passcode to the end of the
		password.

	* loginTimeout: Specifies the timeout, in seconds, for login. The default
		is 60 seconds. The login request gives up after the timeout length if the
		HTTP response is success.

	* authenticator: Specifies the authenticator to use for authenticating user credentials:
		- To use the internal Snowflake authenticator, specify snowflake (Default).
		- To authenticate through Okta, specify https://<okta_account_name>.okta.com (URL prefix for Okta).
		- To authenticate using your IDP via a browser, specify externalbrowser.
		- To authenticate via OAuth, specify oauth and provide an OAuth Access Token (see the token parameter below).

	* application: Identifies your application to Snowflake Support.

	* insecureMode: false by default. Set to true to bypass the Online
		Certificate Status Protocol (OCSP) certificate revocation check.
		IMPORTANT: Change the default value for testing or emergency situations only.

	* token: a token that can be used to authenticate. Should be used in conjunction with the "oauth" authenticator.

	* client_session_keep_alive: Set to true have a heartbeat in the background every hour to keep the connection alive
		such that the connection session will never expire. Care should be taken in using this option as it opens up
		the access forever as long as the process is alive.

	* ocspFailOpen: true by default. Set to false to make OCSP check fail closed mode.

	* validateDefaultParameters: true by default. Set to false to disable checks on existence and privileges check for
								 Database, Schema, Warehouse and Role when setting up the connection

All other parameters are taken as session parameters. For example, TIMESTAMP_OUTPUT_FORMAT session parameter can be
set by adding:

	...&TIMESTAMP_OUTPUT_FORMAT=MM-DD-YYYY...

Proxy

The Go Snowflake Driver honors the environment variables HTTP_PROXY, HTTPS_PROXY and NO_PROXY for the forward proxy setting.

Logging

By default, the driver's builtin logger is NOP; no output is generated. This is
intentional for those applications that use the same set of logger parameters
not to conflict with glog, which is incorporated in the driver logging
framework.

In order to enable debug logging for the driver, add a build tag sfdebug to the
go tool command lines, for example:

	go build -tags=sfdebug

For tests, run the test command with the tag along with glog parameters. For
example, the following command will generate all acitivty logs in the standard
error.

	go test -tags=sfdebug -v . -vmodule=*=2 -stderrthreshold=INFO

Likewise, if you build your application with the tag, you may specify the same
set of glog parameters.

	your_go_program -vmodule=*=2 -stderrthreshold=INFO

To get the logs for a specific module, use the -vmodule option. For example, to
retrieve the driver.go and connection.go module logs:

	your_go_program -vmodule=driver=2,connection=2 -stderrthreshold=INFO

Note: If your request retrieves no logs, call db.Close() or glog.flush() to flush the glog buffer.

Note: The logger may be changed in the future for better logging. Currently if
the applications use the same parameters as glog, you cannot collect both
application and driver logs at the same time.

Canceling Query by CtrlC

From 0.5.0, a signal handling responsibility has moved to the applications. If you want to cancel a
query/command by Ctrl+C, add a os.Interrupt trap in context to execute methods that can take the context parameter,
e.g., QueryContext, ExecContext.

	// handle interrupt signal
	ctx, cancel := context.WithCancel(context.Background())
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	defer func() {
		signal.Stop(c)
	}()
	go func() {
		select {
		case <-c:
			cancel()
		case <-ctx.Done():
		}
	}()
	... (connection)
	// execute a query
	rows, err := db.QueryContext(ctx, query)
	... (Ctrl+C to cancel the query)

See cmd/selectmany.go for the full example.

Supported Data Types

Queries return SQL column type information in the ColumnType type. The
DatabaseTypeName method returns the following strings representing Snowflake
data types:

	String Representation	Snowflake Data Type
	FIXED	                NUMBER/INT
	REAL	                REAL
	TEXT	                VARCHAR/STRING
	DATE	                DATE
	TIME	                TIME
	TIMESTAMP_LTZ	        TIMESTAMP_LTZ
	TIMESTAMP_NTZ	        TIMESTAMP_NTZ
	TIMESTAMP_TZ	        TIMESTAMP_TZ
	VARIANT	                VARIANT
	OBJECT	                OBJECT
	ARRAY	                ARRAY
	BINARY	                BINARY
	BOOLEAN	                BOOLEAN

Binding Time Type

Go's database/sql package limits Go's data types to the following for binding and fetching:

	int64
	float64
	bool
	[]byte
	string
	time.Time

Fetching data isn't an issue since the database data type is provided along
with the data so the Go Snowflake Driver can translate Snowflake data types to
Go native data types.

When the client binds data to send to the server, however, the driver cannot
determine the date/timestamp data types to associate with binding parameters.
For example:

	dbt.mustExec("CREATE OR REPLACE TABLE tztest (id int, ntz, timestamp_ntz, ltz timestamp_ltz)")
	// ...
	stmt, err :=dbt.db.Prepare("INSERT INTO tztest(id,ntz,ltz) VALUES(1, ?, ?)")
	// ...
	tmValue time.Now()
	// ... Is tmValue a TIMESTAMP_NTZ or TIMESTAMP_LTZ?
	_, err = stmt.Exec(tmValue, tmValue)

To resolve this issue, a binding parameter flag is introduced that associates
any subsequent time.Time type to the DATE, TIME, TIMESTAMP_LTZ, TIMESTAMP_NTZ
or BINARY data type. The above example could be rewritten as follows:

	import (
		sf "github.com/snowflakedb/gosnowflake"
	)
	dbt.mustExec("CREATE OR REPLACE TABLE tztest (id int, ntz, timestamp_ntz, ltz timestamp_ltz)")
	// ...
	stmt, err :=dbt.db.Prepare("INSERT INTO tztest(id,ntz,ltz) VALUES(1, ?, ?)")
	// ...
	tmValue time.Now()
	// ...
	_, err = stmt.Exec(sf.DataTypeTimestampNtz, tmValue, sf.DataTypeTimestampLtz, tmValue)

Timestamps with Time Zones

The driver fetches TIMESTAMP_TZ (timestamp with time zone) data using the
offset-based Location types, which represent a collection of time offsets in
use in a geographical area, such as CET (Central European Time) or UTC
(Coordinated Universal Time). The offset-based Location data is generated and
cached when a Go Snowflake Driver application starts, and if the given offset
is not in the cache, it is generated dynamically.

Currently, Snowflake doesn't support the name-based Location types, e.g.,
America/Los_Angeles.

For more information about Location types, see the Go documentation for https://golang.org/pkg/time/#Location.

Binary Data

Internally, this feature leverages the []byte data type. As a result, BINARY
data cannot be bound without the binding parameter flag. In the following
example, sf is an alias for the gosnowflake package:

	var b = []byte{0x01, 0x02, 0x03}
	_, err = stmt.Exec(sf.DataTypeBinary, b)

Maximum number of Result Set Chunk Downloader

The driver directly downloads a result set from the cloud storage if the size is large. It is
required to shift workloads from the Snowflake database to the clients for scale. The download takes place by goroutine
named "Chunk Downloader" asynchronously so that the driver can fetch the next result set while the application can
consume the current result set.

The application may change the number of result set chunk downloader if required. Note this doesn't help reduce
memory footprint by itself. Consider Custom JSON Decoder.

	import (
		sf "github.com/snowflakedb/gosnowflake"
	)
	sf.MaxChunkDownloadWorkers = 2


Experimental: Custom JSON Decoder for parsing Result Set

The application may have the driver use a custom JSON decoder that incrementally parses the result set as follows.

	import (
		sf "github.com/snowflakedb/gosnowflake"
	)
	sf.CustomJSONDecoderEnabled = true
	...

This option will reduce the memory footprint to half or even quarter, but it can significantly degrade the
performance depending on the environment. The test cases running on Travis Ubuntu box show five times less memory
footprint while four times slower. Be cautious when using the option.

(Private Preview) JWT authentication

** Not recommended for production use until GA

Now JWT token is supported when compiling with a golang version of 1.10 or higher. Binary compiled with lower version
of golang would return an error at runtime when users try to use JWT authentication feature.

To enable this feature, one can construct DSN with fields "authenticator=SNOWFLAKE_JWT&privateKey=<your_private_key>",
or using Config structure specifying:

	config := &Config{
		...
		Authenticator: "SNOWFLAKE_JWT"
		PrivateKey:   "<your_private_key_struct in *rsa.PrivateKey type>"
	}

The <your_private_key> should be a base64 URL encoded PKCS8 rsa private key string. One way to encode a byte slice to URL
base 64 URL format is through base64.URLEncoding.EncodeToString() function.

On the server side, one can alter the public key with the SQL command:

	ALTER USER <your_user_name> SET RSA_PUBLIC_KEY='<your_public_key>';

The <your_public_key> should be a base64 Standard encoded PKI public key string. One way to encode a byte slice to base
64 Standard format is through base64.StdEncoding.EncodeToString() function.

To generate the valid key pair, one can do the following command on the shell script:

	# generate 2048-bit pkcs8 encoded RSA private key
	openssl genpkey -algorithm RSA \
    	-pkeyopt rsa_keygen_bits:2048 \
    	-pkeyopt rsa_keygen_pubexp:65537 | \
  		openssl pkcs8 -topk8 -nocrypt -outform der > rsa-2048-private-key.p8

	# extravt 2048-bit PKI encoded RSA public key from the private key
	openssl pkey -pubout -inform der -outform der \
    	-in rsa-2048-private-key.p8 \
    	-out rsa-2048-public-key.spki

Limitations

GET and PUT operations are unsupported.
*/
package gosnowflake
