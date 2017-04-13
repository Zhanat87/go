export class CommonHelperObject {
    static array_pluck(list, key, value) {
        let result = {};

        for (let row of list) {
            if (row[key] && row[value]) {
                result[row[key]] = row[value];
            }
        }

        return result;
    }

    static prepareObjectForSelect(rows, value, label) {
        let result = [];

        for (let row of rows) {
            result.push({
                'value': row[value].toString(),
                'label': row[label].toString()
            });
        }

        return result;
    }

    static propertyToKey(rows, property) {
        let result = {};

        for (let row of rows) {
            result[row[property]] = row;
        }

        return result;
    }
}
