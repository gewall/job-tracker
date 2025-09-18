import React from "react";
import { SectionContent } from "./Section";
import { Button } from "@workspace/ui/components/button";
import { ExternalLink } from "lucide-react";
import { ApplicationStatus } from "@/lib/types/applications.type";
import Link from "next/link";

export type CardProps = {
  label: ApplicationStatus;
  count: number;
};

export default function Card({ label, count }: CardProps) {
  return (
    <div className="flex-1 relative">
      <SectionContent classname="">
        <div className="flex">
          <div className="">
            <h5 className="font-light uppercase text-sm text-neutral-400">
              {label}
            </h5>
            <h5 className="font-semibold text-3xl">{count}</h5>
          </div>
        </div>
        <div className="ml-auto absolute top-2 right-2">
          <Link href={"/application?status=" + label} passHref>
            <Button
              variant={"outline"}
              size={"icon"}
              className="rounded-full text-indigo-400 hover:text-indigo-600"
            >
              <ExternalLink />
            </Button>
          </Link>
        </div>
      </SectionContent>
    </div>
  );
}
