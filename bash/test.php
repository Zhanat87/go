<?php
// docker inspect --format '{{.NetworkSettings.IPAddress}}' postgresql1
$conn_string = "host=172.18.0.5 port=5432 dbname=go_restful user=postgres password=postgres";
try {
    $dbconn = pg_connect($conn_string);
    echo 'success connected' . PHP_EOL;
    $result = pg_query($dbconn, "SELECT id, title FROM news_shard_1");
    if (!$result) {
        echo "An error occurred.\n";
        exit;
    }

    while ($row = pg_fetch_row($result)) {
        echo "id: $row[0], title: $row[1]\r\n";
    }
} catch (Exception $e) {
    var_dump($e->getMessage(), $e->getCode());
}

// docker inspect --format '{{.NetworkSettings.IPAddress}}' postgresql2
$conn_string = "host=172.18.0.6 port=5432 dbname=go_restful user=postgres password=postgres";
try {
    $dbconn = pg_connect($conn_string);
    echo 'success connected 2' . PHP_EOL;
    $result = pg_query($dbconn, "SELECT id, title FROM news_shard_2");
    if (!$result) {
        echo "An error occurred 2.\n";
        exit;
    }

    while ($row = pg_fetch_row($result)) {
        echo "id: $row[0], title: $row[1]\r\n";
    }
} catch (Exception $e) {
    var_dump($e->getMessage(), $e->getCode());
}