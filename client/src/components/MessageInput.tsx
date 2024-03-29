import { useSendMessage } from "../hooks/useSendMessage";

export const MessageInput = () => {
  const { input, setInput, send } = useSendMessage();

  return (
    <div>
      <input
        onChange={(e) => setInput(e.target.value)}
        value={input}
        placeholder="new message"
      />
      <button onClick={send}>Send</button>
    </div>
  );
};
