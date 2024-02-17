import { useCallback, useState } from "react";
import { useWebsocketAtom } from "../state/websocket";
/**
 * チャットを送信するためのカスタムフック
 * @returns {input, setInput, send}
 */
export const useSendMessage = () => {
  const socket = useWebsocketAtom();
  const [input, setInput] = useState<string>("");

  const send = useCallback(() => {
    if (input.length === 0) return;
    socket.send(JSON.stringify(input));
    setInput("");
  }, [input]);

  return { input, setInput, send };
};
