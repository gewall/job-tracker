"use client";

import { useAppStore } from "@/lib/store/appstore";
import { useRouter } from "next/navigation";
import { useJwt } from "react-jwt";
import React, { useEffect } from "react";

type Props = {
  children: React.ReactNode;
};

export default function PrivateRoute({ children }: Props) {
  const router = useRouter();
  const { user, setUser } = useAppStore((s) => s);
  const { decodedToken, isExpired } = useJwt(user.access_token as string);

  const decoded = decodedToken as {
    userId: string;
  };

  useEffect(() => {
    if (!user.access_token || isExpired) {
      router.push("/signin");
      return;
    }
    setUser({
      ...user,
      userId: decoded?.userId,
    });
  }, [decodedToken]);
  console.log(user);

  return <div>{children}</div>;
}
