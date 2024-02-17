import { MessageInput } from "./components/MessageInput";
import { MessageList } from "./components/MessageList";

const App = () => {
  return (
    <div>
      <h1>Chat App</h1>
      <MessageInput />
      <MessageList />
    </div>
  );
};

export default App;
