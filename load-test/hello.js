import { sleep } from 'k6';
import http from 'k6/http';
import ws from 'k6/ws';

export let options = {
  scenarios: {
    helloWorld: {
      executor: 'constant-vus',
      vus: 15,
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
  // const url = 'wss://api.zossi.xyz/api/world-journey/?id=b01848b1-1047-4831-889d-b2bcbaf4ed8b';
  const url = 'ws://web:8080/api/world-journey/?id=dc4ef400-5f8f-45cb-b373-68ed789b5d46';

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
      }, 100);
    });
  });
}

// export function helloWorld() {
//   const url = 'https://api.zossi.xyz/api/items';
//   // const url = 'https://api.zossi.xyz/api/worlds?limit=20&offset=0';
//   // const url = 'http://web:8080/api/worlds?limit=20&offset=0';

//   const res = http.get(url);
//   console.log(res.status);
//   console.log(res.error);
//   sleep(1);
// }
