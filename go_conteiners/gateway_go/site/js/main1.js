
function lol() {
    let settings = {
        "url": "http://localhost:8080/api/group/",
        "method": "POST",
        "timeout": 0,
        "headers": {
            "Content-Type": "application/json"
        },
        "data": JSON.stringify({
            "name": "lol23445"
        }),
    };

    $.ajax(settings).done(function (response) {
        console.log(response);
    });
}
function lol2() {
    let settings = {
        "url": "http://localhost:8080/api/group/1",
        "method": "GET",
        "timeout": 0,
    };

    $.ajax(settings).done(function (response) {
        console.log(response);
    });
}