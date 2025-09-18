import { NavLinkProps } from "@/app/_types/navbar.type";
import { Home, NotepadText } from "lucide-react";

export const navbar: NavLinkProps[] = [
  {
    href: "/",
    label: "Dashboard",
    pathname: "",
    icon: <Home className="h-5 w-5 font-light" />,
  },
  {
    href: "/application",
    label: "Application",
    pathname: "application",
    icon: <NotepadText className="h-5 w-5 font-light" />,
  },
];
