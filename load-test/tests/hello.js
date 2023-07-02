import ws from 'k6/ws';

export let options = {
  scenarios: {
    helloWorld: {
      executor: 'constant-vus',
      vus: 40,
      duration: '60s',
      gracefulStop: '1s',
      tags: { test_type: 'helloWorld' },
      exec: 'helloWorld',
    },
  },
  discardResponseBodies: true,
  thresholds: {},
};

export function helloWorld() {
  const url = 'wss://api.zossi.xyz/api/world-journey/?id=b01848b1-1047-4831-889d-b2bcbaf4ed8b';

  ws.connect(url, {}, function (socket) {
    socket.on('open', function open() {
      let tick = 0;
      let direction = 0;
      socket.setInterval(function timeout() {
        if (tick % 2 === 0) {
            direction = parseInt(Math.random() * 4, 10) % 4;
        }
        socket.send(
          JSON.stringify({
            type: 'MOVE',
            direction,
          }),
        );
        tick += 1;
      }, 1000);
    });
  });
}
