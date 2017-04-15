<?php
/*
POSTGRESQL_PORT_5432_TCP_ADDR=172.17.0.2
POSTGRESQL_ENV_POSTGRES_PASSWORD=postgres
POSTGRESQL_PORT_5432_TCP_PORT=5432
POSTGRESQL_PORT_5432_TCP_PROTO=tcp
POSTGRESQL_ENV_POSTGRES_USER=postgres
POSTGRESQL_PORT=tcp://172.17.0.2:5432
POSTGRESQL_PORT_5432_TCP=tcp://172.17.0.2:5432
POSTGRESQL_NAME=/restful/postgresql
POSTGRESQL_ENV_POSTGRES_DB=go_restful
 */
$conn_string = "host=" . $_SERVER["POSTGRESQL_PORT_5432_TCP_ADDR"] .
    " port=" . $_SERVER["POSTGRESQL_PORT_5432_TCP_PORT"] .
    " dbname=" . $_SERVER["POSTGRESQL_ENV_POSTGRES_DB"] .
    " user=" . $_SERVER["POSTGRESQL_ENV_POSTGRES_USER"] .
    " password=" . $_SERVER["POSTGRESQL_ENV_POSTGRES_DB"];
try {
    $dbconn = pg_connect($conn_string);
    echo 'success connected to: ' . $conn_string . PHP_EOL;
} catch (Exception $e) {
    var_dump($e->getMessage(), $e->getCode());
}