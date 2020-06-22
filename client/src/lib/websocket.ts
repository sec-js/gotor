let c: WebSocket;
const initConn = (address: string): Promise<void> =>
    new Promise((accept, reject) => {
        c = new WebSocket(address);
        c.onopen = () => accept();
        c.onerror = error => reject(error)
    });


const getConn = (): WebSocket => c;

export {
    getConn,
    initConn
};
