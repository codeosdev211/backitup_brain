package db

import (
    "database/sql"
)

func CallDatabase(isSelect bool, query *string) ([]map[string]interface{}, error) {
    /* connecting to database */
    db, err := sql.Open("mysql", "root:omkara@211@tcp(localhost:3306)/backitup")
    if err != nil {
        return nil, err
    }
    defer db.Close()

    /* getting rows from result */
    rows, err := db.Query(*query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    /* returning if insert, alter, update, drop, truncate */
    if !isSelect {
        return nil, nil
    }

    /* getting columns of rows */
    columns, err := rows.Columns()
    if err != nil {
        return nil, err
    }

    var colCount int = len(columns)
    /* creating jsonarray structure using nested map */
    response := make([]map[string]interface{}, 0)
    keys := make([]interface{}, colCount)
    ptrKeys := make([]interface{}, colCount)

    /* looping through each row */
    for rows.Next() {
        for cell := 0; cell < colCount; cell++ {
            ptrKeys[cell] = &keys[cell]
        }
        rows.Scan(ptrKeys...)

        /* creating jsonobject like structure */
        each := make(map[string]interface{})
        for col, key := range columns {
            var value interface{}
            bytes, ok := keys[col].([]byte)
            if ok {
                value = string(bytes)
            }else{
                value = keys[col]
            }
            each[key] = value
        }
        response = append(response, each)
    }
    return response, nil
}
