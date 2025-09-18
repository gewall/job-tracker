
export type NavLinkProps = {
  variant?: "Link" | "Dropdown";
  list?: SubNavLinkProps[];
  icon: React.ReactNode;
  label: string;
  href: string;
  pathname: string;
};

export type SubNavLinkProps = {
  label: string;
  icon: React.ReactNode;
  href: string;
  pathname: string;
};
