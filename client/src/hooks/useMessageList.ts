import { useWebsocketAtom } from "../state/websocket";
import { useMessageListAtom } from "../state/messages";
import { Message } from "../models/message";

/**
 * 新規のメッセージを受け取るたびに messageListAtom に追加・更新する
 * @returns {messageList}
 */
export const useMessageList = (): Message[] => {
  const socket = useWebsocketAtom();
  const { messageList, updateMessageList } = useMessageListAtom();

  socket.onmessage = (msg) => {
    const content = JSON.parse(msg.data as string);
    const message: Message = { content: content };
    updateMessageList(message);
  };

  return messageList;
};
