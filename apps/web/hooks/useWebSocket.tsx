import { useAppStore } from "@/lib/store/appstore";
import { useEffect, useState } from "react";

type WebSocketType = {
  uri?: string;
  onOpen?: () => void;
  onError?: (e: Event) => void;
};

export default function useWebSocket({ uri, onOpen, onError }: WebSocketType) {
  const { notifications, addNotification } = useAppStore((s) => s);

  const [socket, setSocket] = useState<WebSocket | null>(null);

  const sendNotification = (msg: string) => {
    if (socket?.readyState === WebSocket.OPEN) {
      socket.send(msg);
    }
  };

  useEffect(() => {
    const ws = new WebSocket(uri ? uri : "ws://localhost:8080/socket/notif");
    ws.onopen = () => {
      if (onOpen) onOpen();
    };

    ws.onmessage = (e) => {
      addNotification(e.data);
    };

    ws.onerror = (err) => {
      if (onError) onError(err);
    };

    setSocket(ws);

    return () => {
      ws.close();
    };
  }, []);

  return {
    notifications,
    socket,
    sendNotification,
  };
}
