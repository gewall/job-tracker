"use client";

import React, { useEffect, useState } from "react";
import { NavLink } from "../_components/NavLink";
import { navbar } from "@/lib/constant/navbar";
import Section from "../_components/Section";
import { Button } from "@workspace/ui/components/button";
import { Bell, ChevronsLeft, ChevronsRight } from "lucide-react";
import { useAppStore } from "@/lib/store/appstore";
import clsx from "clsx";

type Props = {};

export default function Navbar({}: Props) {
  const { sidebarOpen, sidebarToggle, notifications } = useAppStore((s) => s);

  return (
    <>
      <nav
        className={clsx(
          "flex flex-col fixed top-0  h-screen  border-r bg-white z-50 w-64 transition-all  overflow-auto",
          {
            "left-0": sidebarOpen,
            "-left-64": !sidebarOpen,
          }
        )}
      >
        <div className=" flex-none px-4 py-12 w-full">
          <h2 className="text-center font-light text-2xl">Job Tracker</h2>
        </div>
        <div className=" flex-none px-8">
          <p className="text-sm font-extralight">Dashboard</p>
        </div>

        <div className="flex flex-1 overflow-y-auto flex-col mt-4 px-4 space-y-2 transition-all duration-1000">
          {navbar.map((_, i) => (
            <NavLink {..._} key={i} />
          ))}
        </div>
        <div className="flex flex-none justify-between p-8 items-center">
          <p className="text-sm font-extralight">
            &copy;gingrha {new Date().getFullYear()}
          </p>
          <div className="md:hidden block">
            <Button variant={"ghost"} size={"sm"} onClick={sidebarToggle}>
              <ChevronsLeft />
            </Button>
          </div>
        </div>
      </nav>
      <nav className="fixed w-full bg-white max-h-80 top-0 left-0 z-10">
        <Section classname=" py-4 flex items-center">
          <Button variant={"outline"} size={"sm"} onClick={sidebarToggle}>
            {sidebarOpen ? <ChevronsLeft /> : <ChevronsRight />}
          </Button>
          <div className="ml-auto">{notifications.length > 0 && <Bell />}</div>
        </Section>
      </nav>
    </>
  );
}
