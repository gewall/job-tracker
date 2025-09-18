"use client";

import { useAppStore } from "@/lib/store/appstore";
import clsx from "clsx";
import React from "react";

type Props = {
  children: React.ReactNode;
  classname?: string;
};

export default function Section({ children, classname }: Props) {
  const { sidebarOpen } = useAppStore((s) => s);
  return (
    <section
      className={clsx("transition-all md:px-12 px-4", classname, {
        "md:ml-64": sidebarOpen,
        "md:ml-0": !sidebarOpen,
      })}
    >
      {children}
    </section>
  );
}

export function SectionContent({ children, classname }: Props) {
  return (
    <div
      className={clsx(
        "px-8 py-4 bg-white rounded-md border border-neutral-200",
        classname
      )}
    >
      {children}
    </div>
  );
}

export function SectionTitle({ title }: { title: string }) {
  return <h4 className="font-light text-lg">{title}</h4>;
}
