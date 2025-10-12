"use client";

import useWebSocket from "@/hooks/useWebSocket";
import { Button } from "@workspace/ui/components/button";
import React from "react";

type Props = {};

export default function Socket({}: Props) {
  const { notifications, sendNotification } = useWebSocket({
    onOpen: () => console.log("Socket Connected"),
    onError: (err) => {
      console.log(err);
      return;
    },
  });

  return (
    <div>
      <Button onClick={() => sendNotification("HALLOO")}>Add MSG</Button>
      <ul>
        {notifications.map((m, i) => (
          <li key={i}>{m}</li>
        ))}
      </ul>
    </div>
  );
}
