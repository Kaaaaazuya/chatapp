import { atom, selector, useRecoilValue } from "recoil";
import * as WebSocket from "websocket";

const connect = (): Promise<WebSocket.w3cwebsocket> => {
  return new Promise((resolve, reject) => {
    // TODO: url の管理を修正する
    const socket = new WebSocket.w3cwebsocket("ws://localhost:80/ws");
    socket.onopen = () => {
      console.log("connected");
      resolve(socket);
    };
    socket.onclose = () => {
      console.log("reconnecting...");
      connect();
    };
    socket.onerror = (err) => {
      console.log("connection error:", err);
      reject(err);
    };
  });
};

const connectWebsocketSelector = selector({
  key: "connectWebsocket",
  get: async (): Promise<WebSocket.w3cwebsocket> => {
    return await connect();
  },
});

const websocketAtom = atom<WebSocket.w3cwebsocket>({
  key: "websocket",
  default: connectWebsocketSelector,
});

/**
 * websocketAtom を外部で利用するための hook
 * @returns { websocketAtom }
 */
export const useWebsocketAtom = () => {
  return useRecoilValue(websocketAtom);
};
