(function() {
    var conn = new WebSocket("ws://192.168.1.3:8080/ws");
    document.addEventListener("input", function(evt) {
        alert("starting")
        console.log(evt.target)
        s = evt.data;
        conn.send(s);
    });
})();
