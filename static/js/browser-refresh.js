let socket;
let retries = 0;
const maxRetries = 20;

function connect() {
    if (retries >= maxRetries) {
        console.log("Max retries reached, could not establish WebSocket connection");
        return;
    }
    socket = new WebSocket("ws://localhost:8080/ws");

    socket.onopen = function() {
        retries = 0;
    };

    socket.onmessage = function(event) {
        if (event.data === "refresh") {
            location.reload();
        }
    };

    socket.onerror = function(error) {
        console.error("WebSocket error:", error);
    };

    socket.onclose = function() {
        retries++;
        setTimeout(connect, 100);
    };
}

connect();
