let c;
const initConn = address =>
    new Promise((accept, reject) => {
        c = new WebSocket(address);
        c.onopen = () => accept();
        c.onerror = error => reject(error)
    });


const getConn = () => c;

export {
    getConn,
    initConn
};
