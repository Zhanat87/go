/**
 * @link https://www.typescriptlang.org/docs/handbook/modules.html
 */

/**
 * @param needle
 * @param haystack
 * @param argStrict
 * @returns {boolean}
 */
export function in_array(needle, haystack, argStrict): boolean {
    var key = ''
    var strict = !!argStrict

    if (strict) {
        for (key in haystack) {
            if (haystack[key] === needle) {
                return true
            }
        }
    } else {
        for (key in haystack) {
            if (haystack[key] == needle) {
                return true
            }
        }
    }

    return false
}

/**
 * @param str
 * @param charlist
 * @returns {string}
 */
export function trim(str: string, charlist: any): string {
    //  discuss at: http://locutus.io/php/trim/
    // original by: Kevin van Zonneveld (http://kvz.io)
    // improved by: mdsjack (http://www.mdsjack.bo.it)
    // improved by: Alexander Ermolaev (http://snippets.dzone.com/user/AlexanderErmolaev)
    // improved by: Kevin van Zonneveld (http://kvz.io)
    // improved by: Steven Levithan (http://blog.stevenlevithan.com)
    // improved by: Jack
    //    input by: Erkekjetter
    //    input by: DxGx
    // bugfixed by: Onno Marsman (https://twitter.com/onnomarsman)
    //   example 1: trim('    Kevin van Zonneveld    ')
    //   returns 1: 'Kevin van Zonneveld'
    //   example 2: trim('Hello World', 'Hdle')
    //   returns 2: 'o Wor'
    //   example 3: trim(16, 1)
    //   returns 3: '6'
    var whitespace = [
        ' ',
        '\n',
        '\r',
        '\t',
        '\f',
        '\x0b',
        '\xa0',
        '\u2000',
        '\u2001',
        '\u2002',
        '\u2003',
        '\u2004',
        '\u2005',
        '\u2006',
        '\u2007',
        '\u2008',
        '\u2009',
        '\u200a',
        '\u200b',
        '\u2028',
        '\u2029',
        '\u3000'
    ].join('')
    var l = 0
    var i = 0
    str += ''
    if (charlist) {
        whitespace = (charlist + '').replace(/([[\]().?/*{}+$^:])/g, '$1')
    }
    l = str.length
    for (i = 0; i < l; i++) {
        if (whitespace.indexOf(str.charAt(i)) === -1) {
            str = str.substring(i)
            break
        }
    }
    l = str.length
    for (i = l - 1; i >= 0; i--) {
        if (whitespace.indexOf(str.charAt(i)) === -1) {
            str = str.substring(0, i + 1)
            break
        }
    }
    return whitespace.indexOf(str.charAt(0)) === -1 ? str : ''
}

export function var_dump(variable: any): void {
    // "undefined", "object", "boolean", "number", "string", "symbol", Implementation-dependent, "function", "object"
    var typeOf = typeof variable;
    console.log('var_dump, typeOf is ', typeOf);
    switch (typeOf) {
        case 'object':
            console.log(Object.getOwnPropertyNames(variable).sort());
            Object.getOwnPropertyNames(variable).forEach(function (val, idx, array) {
                console.log(val + ' -> ' + variable[val]);
            });
            break;
        default:
            console.log(variable);
            break;
    }
}

// Return an array of the selected option values
// select is an HTML select element
export function getSelectedValues(select): any {
    var result = [];
    var options = select && select.options;
    var opt;

    for (var i = 0, iLen = options.length; i < iLen; i++) {
        opt = options[i];

        if (opt.selected) {
            result.push(opt.value || opt.text);
        }
    }
    return result;
}

/**
 * @link http://stackoverflow.com/questions/3710204/how-to-check-if-a-string-is-a-valid-json-string-in-javascript-without-using-try
 *
 * @param jsonString
 * @returns {any}
 */
export function tryParseJSON(jsonString: string, defaultValue?: any): any {
    try {
        var o = JSON.parse(jsonString);

        // Handle non-exception-throwing cases:
        // Neither JSON.parse(false) or JSON.parse(1234) throw errors, hence the type-checking,
        // but... JSON.parse(null) returns null, and typeof null === "object",
        // so we must check for that, too. Thankfully, null is falsey, so this suffices:
        if (o && typeof o === "object") {
            return o;
        }
    } catch (e) {
    }
    return defaultValue !== 'undefined' ? defaultValue : false;
}

/**
 * @link http://stackoverflow.com/questions/4994201/is-object-empty
 * @param obj
 * @returns {boolean}
 */
export function isEmpty(obj: any): boolean {

    // null and undefined are "empty"
    if (obj == null) return true;

    // Assume if it has a length property with a non-zero value
    // that that property is correct.
    if (obj.length > 0)    return false;
    if (obj.length === 0)  return true;

    // If it isn't an object at this point
    // it is empty, but it can't be anything *but* empty
    // Is it empty?  Depends on your application.
    if (typeof obj !== "object") return true;

    // Otherwise, does it have any properties of its own?
    // Note that this doesn't handle
    // toString and valueOf enumeration bugs in IE < 9
    // Speed up calls to hasOwnProperty
    let hasOwnProperty = Object.prototype.hasOwnProperty;
    for (var key in obj) {
        if (hasOwnProperty.call(obj, key)) return false;
    }

    return true;
}

/*
 sleep time expects milliseconds
 @link http://stackoverflow.com/questions/951021/what-is-the-javascript-version-of-sleep
 */
export function sleep(time) {
    return new Promise((resolve) => setTimeout(resolve, time));
}

// note: may be move to separate file with dates functions
export function today(): string {
    var today = new Date();
    return formatDate(today);
}

export function addDays(date, days): string {
    try {
        date = createDate(date);
        date.setDate(date.getDate() + days);
        return formatDate(date);
    } catch (err) {
        console.info('addDays', date, days, err);
    }
}

export function subDays(date, days): string {
    try {
        date = createDate(date);
        date.setDate(date.getDate() - days);
        return formatDate(date);
    } catch (err) {
        console.info('subDays', date, days, err);
    }
}

export function createDate(date): Date {
    return new Date(date.substr(3, 2) + '/' + date.substr(0, 2) + '/' + date.substr(6, 4));
}

export function formatDate(date): string {
    var dd = date.getDate();
    var mm = date.getMonth() + 1;

    var yyyy = date.getFullYear();

    return (dd < 10 ? '0' + dd : dd) + '/' + (mm < 10 ? '0' + mm : mm) + '/' + yyyy;
}

export function compareDates(date, date2): number { // spaceship operator
    var res;
    date = createDate(date).getTime();
    date2 = createDate(date2).getTime();
    if (date > date2) {
        res = 1;
    } else if (date == date2) {
        res = 0;
    } else if (date < date2) {
        res = -1;
    }
    return res;
}

/**
 * @link http://stackoverflow.com/questions/105034/create-guid-uuid-in-javascript
 *
 * @returns {string}
 */
export function guid() {
    function s4() {
        return Math.floor((1 + Math.random()) * 0x10000)
            .toString(16)
            .substring(1);
    }

    return s4() + s4() + '-' + s4() + '-' + s4() + '-' + s4() + '-' + s4() + s4() + s4();
}