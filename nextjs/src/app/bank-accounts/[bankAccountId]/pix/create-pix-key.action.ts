"use server";

import { revalidateTag } from "next/cache";

export async function createPixKeyAction(
  bankAccountId: string,
  formData: FormData
) {
  const pixKeyKind = formData.get("kind");
  const pixKeyKey = formData.get("key");
  const response = await fetch(
    `http://host.docker.internal:3000/bank-accounts/${bankAccountId}/pix-keys`,
    {
      method: "POST",
      body: JSON.stringify({
        kind: pixKeyKind,
        key: pixKeyKey,
      }),
      headers: {
        "Content-Type": "application/json",
      },
    }
  );

  if (!response.ok) {
    throw new Error(JSON.stringify(await response.json()));
  }

  const result = await response.json();
  revalidateTag(`pix-keys-${bankAccountId}`);
  return result;
}