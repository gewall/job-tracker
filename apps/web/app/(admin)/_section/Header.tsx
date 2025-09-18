"use client";

import { SectionTitle } from "@/app/_components/Section";
import { useAppStore } from "@/lib/store/appstore";
import { ApplicationStatus } from "@/lib/types/applications.type";
import { Button } from "@workspace/ui/components/button";
import { useSearchParams } from "next/navigation";
import React, { useEffect } from "react";

type Props = {};

export default function Header({}: Props) {
  const {
    applicationStatusFilter: status,
    applicationStatusFilterChange: setStatus,
  } = useAppStore((s) => s);
  const param = useSearchParams();

  const appStatus = [
    "all",
    "applied",
    "interview",
    "offer",
    "rejected",
    "saved",
  ] as ApplicationStatus[] | "all"[];

  const statusParam = param.get("status") as ApplicationStatus;

  useEffect(() => {
    if (statusParam) setStatus(statusParam);
  }, [param]);

  return (
    <>
      <SectionTitle title="Applications" />
      <div className="flex gap-2 bg-neutral-100 rounded-md py-2 px-4 mt-2 w-fit flex-wrap">
        {appStatus?.map((_, i) => (
          <Button
            key={i}
            className="cursor-pointer"
            variant={_ === status ? "outline" : "ghost"}
            onClick={() => setStatus(_)}
          >
            {_.slice(0, 1).toUpperCase() + _.slice(1, _.length)}
          </Button>
        ))}
      </div>
    </>
  );
}
