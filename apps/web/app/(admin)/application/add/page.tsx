import Section, {
  SectionContent,
  SectionTitle,
} from "@/app/_components/Section";
import React from "react";

type Props = {};

export default function page({}: Props) {
  return (
    <Section>
      <SectionContent>
        <SectionTitle title="Add Job Application" />
      </SectionContent>
    </Section>
  );
}
