"use client";

import clsx from "clsx";
import { ChevronDown, ChevronRight, Home } from "lucide-react";
import Link from "next/link";
import { usePathname } from "next/navigation";
import { useState } from "react";
import { NavLinkProps } from "../_types/navbar.type";

export function NavLink({
  variant,
  list,
  icon,
  href,
  label,
  pathname,
}: NavLinkProps) {
  const baseStyle =
    "px-8 py-2 text-md font-light w-full flex gap-2 items-center rounded-lg  cursor-pointer transition-all ease-in-ou duration-300";

  const [listActive, setListActive] = useState<boolean>(false);
  const path = usePathname();

  const checkPathname = (e: string) => {
    const res =
      path
        .split("/")
        .filter((_) => _ !== "")
        .find((_) => _ === e) === pathname || path === pathname + "/";
    if (res && variant !== undefined) setListActive(true);

    return res;
  };

  if (variant === "Dropdown")
    return (
      <div className="transition-all ease-in-out">
        <div
          className={clsx(baseStyle, {
            "bg-indigo-500 text-neutral-50 hover:bg-indigo-600":
              checkPathname(pathname),
          })}
          onClick={() => setListActive(!listActive)}
        >
          {icon}
          <p>Asd</p>
          <div className="ml-auto">
            {!listActive ? <ChevronRight /> : <ChevronDown />}
          </div>
        </div>
        {list && (
          <div
            className={clsx(
              "overflow-hidden flex flex-col space-y-2 transition-all  duration-500",
              {
                "max-h-60 mt-2": listActive,
                "max-h-0 mt-0": !listActive,
              }
            )}
          >
            {list.map((_, i) => (
              <Link
                key={i}
                href={_.href}
                className={clsx("px-12 flex py-2 gap-2 text-sm items-center ", {
                  "text-indigo-400 hover:text-indigo-600": checkPathname(
                    _.pathname
                  ),
                  "hover:text-indigo-400 text-primary": !checkPathname(
                    _.pathname
                  ),
                })}
              >
                {_.icon}
                {_.label}
              </Link>
            ))}
          </div>
        )}
      </div>
    );

  return (
    <Link
      className={clsx(baseStyle, {
        "bg-indigo-500 text-neutral-50 hover:bg-indigo-600":
          checkPathname(pathname),
        "hover:text-indigo-400": !checkPathname(pathname),
      })}
      href={href}
    >
      {icon}
      {label}
    </Link>
  );
}
