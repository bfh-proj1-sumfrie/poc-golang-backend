# Setup

Import the Sample Sql file into you Maria DB.
If you want to use a bigger test set you can download another file
here: https://sample-videos.com/download-sample-sql.php

This POC uses the root user with an emtpy password, which
is the default after the installation.

run it from the root folder: `./poc-golang-backend`
This starts a server on localhost:800

Make the following query: `curl localhost:8000/query%2FSELECT%20*%20FROM%20user_details`

The url decode query looks like this: 
`localhost:8000/query/SELECT * FROM user_details`


