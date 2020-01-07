//const serverAddress = "10.69.29.69";
const serverAddress = "localhost";
const wsPort = '8082';
const projectName = 'processEdit';

const WSCmd = {
    Connect: 0,
    wsCMD:1
};

var wsConn = new WebSocket('ws://' + serverAddress + ':' + wsPort + '/' + projectName);
wsConn.onopen = function () {
    console.log("websocket connect success");
    var connectedJson = {};
    connectedJson["cmdID"] = WSCmd.Connect;
    wsConn.send(JSON.stringify(connectedJson));
};
wsConn.onmessage = function (ev) {
    var json = eval("(" + ev.data + ")");
    var cmdID = json.cmdID;

    //实体态势
    if (cmdID === WSCmd.wsCMD) {
        //todo
    }
};