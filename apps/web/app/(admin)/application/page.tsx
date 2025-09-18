import { DataTable } from "@/app/_components/Datatable";
import Section, { SectionContent } from "@/app/_components/Section";
import React from "react";
import { applicationsListColumns } from "./_datatable/columns";
import { applicationListColumnsData } from "@/lib/constant/dummy";

import Header from "../_section/Header";
import { Metadata } from "next";
import Link from "next/link";
import { Plus } from "lucide-react";
import { Button } from "@workspace/ui/components/button";

export const metadata: Metadata = {
  title: "Application | Job Trakcer",
  description: "Job Tracker Web Application",
};

type Props = {};

export default function page({}: Props) {
  return (
    <Section>
      <SectionContent>
        <Header />
        <div className="flex items-center justify-between">
          <h5 className="font-light text-md mt-4 mb-2">Application List</h5>
          <Link href={"/application/add"} className="" passHref>
            <Button
              className="text-indigo-400 hover:text-indigo-600"
              variant={"outline"}
              size={"icon"}
            >
              <Plus />
            </Button>
          </Link>
        </div>
        <DataTable
          columns={applicationsListColumns}
          data={applicationListColumnsData}
        />
      </SectionContent>
    </Section>
  );
}
