import { toast } from "@zerodevx/svelte-toast";

export const failure = (m) =>
  toast.push(m, {
    theme: {
      "--toastBackground": "rgb(239, 68, 68)", // bg-red-500
      "--toastProgressBackground": "rgb(220, 38, 38)", // bg-red-600
    },
  });
