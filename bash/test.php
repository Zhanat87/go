<?php
$conn_string = "host=172.17.0.2 port=5432 dbname=go_restful user=postgres password=postgres";
try {
    $dbconn = pg_connect($conn_string);
    echo 'success connected' . PHP_EOL;
} catch (Exception $e) {
    var_dump($e->getMessage(), $e->getCode());
}