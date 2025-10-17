"use client";

import React, { useState } from "react";

import {
  Field,
  FieldError,
  FieldGroup,
  FieldLabel,
  FieldSet,
} from "@workspace/ui/components/field";
import { Input } from "@workspace/ui/components/input";
import {
  InputGroup,
  InputGroupAddon,
  InputGroupButton,
  InputGroupInput,
} from "@workspace/ui/components/input-group";
import { Signin, User } from "@/lib/types/auth,type";
import { Controller, useForm } from "react-hook-form";
import { EyeClosed, EyeIcon, LoaderIcon } from "lucide-react";
import { Button } from "@workspace/ui/components/button";
import { useRouter } from "next/navigation";
import { useAppStore } from "@/lib/store/appstore";

type Props = {};

export default function Form({}: Props) {
  const { setUser } = useAppStore((s) => s);
  const route = useRouter();
  const [loading, setLoading] = useState<boolean>(false);
  const [showPass, setShowPass] = useState<boolean>(false);
  const form = useForm<Signin>({
    defaultValues: {
      email: "",
      password: "",
    },
  });

  const onSubmit = async (e: Signin) => {
    setLoading(true);
    let user: User;
    const req = await fetch("http://localhost:8080/auth/signin", {
      method: "POST",
      body: JSON.stringify(e),
      credentials: "include",
      headers: { "Content-Type": "application/json" },
    });
    const res = await req.json();

    if (!req) {
      console.log("Error");
      setLoading(false);
      return;
    }

    if (!res?.status) {
      setLoading(false);
      return;
    }

    user = {
      access_token: res?.data?.accessToken,
      userId: undefined,
    };

    setUser(user);

    setLoading(false);
    route.push("/");
  };

  return (
    <FieldSet>
      <form
        id="form-signin"
        action=""
        className="space-y-8"
        onSubmit={form.handleSubmit(onSubmit)}
      >
        <FieldGroup>
          <Controller
            control={form.control}
            name="email"
            render={({ field, fieldState }) => (
              <Field data-invalid={fieldState.invalid}>
                <FieldLabel htmlFor="email">Email</FieldLabel>
                <Input
                  {...field}
                  id="email"
                  aria-invalid={fieldState.invalid}
                  type="text"
                  placeholder="abc...@xxx.com"
                />
                {/* <FieldDescription>
                          Choose a unique username for your account.
                        </FieldDescription> */}
                {fieldState.invalid && (
                  <FieldError errors={[fieldState.error]} />
                )}
              </Field>
            )}
          />
          <Controller
            control={form.control}
            name="password"
            render={({ field, fieldState }) => (
              <Field data-invalid={fieldState.invalid}>
                <FieldLabel htmlFor="password">Password</FieldLabel>
                <InputGroup>
                  <InputGroupInput
                    {...field}
                    id="password"
                    type={showPass ? "text" : "password"}
                    placeholder="********"
                  />
                  <InputGroupAddon align={"inline-end"}>
                    <InputGroupButton onClick={() => setShowPass((e) => !e)}>
                      {showPass ? <EyeIcon /> : <EyeClosed />}
                    </InputGroupButton>
                  </InputGroupAddon>
                </InputGroup>

                {/* <FieldDescription>
                          Choose a unique username for your account.
                        </FieldDescription> */}
                {fieldState.invalid && (
                  <FieldError errors={[fieldState.error]} />
                )}
              </Field>
            )}
          />
        </FieldGroup>
        <div className="flex justify-center items-center">
          <Button variant={"outline"} type="submit">
            Sign In
            {loading && <LoaderIcon />}
          </Button>
        </div>
      </form>
    </FieldSet>
  );
}
