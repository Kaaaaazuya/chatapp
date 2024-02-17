import { useMessageList } from "../hooks/useMessageList";

export const MessageList = () => {
  const messageList = useMessageList();

  return (
    <div>
      {messageList.map((message, idx) => (
        <div key={idx}>{message.content}</div>
      ))}
    </div>
  );
};
