import { Message } from "../models/message";
import { atom, useRecoilCallback, useRecoilValue } from "recoil";

export const messageListAtom = atom<Message[]>({
  key: "messageList",
  default: [],
});

/**
 * messageListAtom を外部で利用するための hook
 * @returns { messageList, updateMessageList}
 */
export const useMessageListAtom = () => {
  const messageList = useRecoilValue(messageListAtom);
  const updateMessageList = useRecoilCallback(
    ({ set }) =>
      (message: Message) => {
        set(messageListAtom, [...messageList, message]);
      }
  );

  return { messageList, updateMessageList };
};
