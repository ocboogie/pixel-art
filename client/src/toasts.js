import { toast } from "@zerodevx/svelte-toast";

export const failure = (m) =>
  toast.push(m, {
    theme: {
      "--toastBackground": "rgb(239, 68, 68)", // bg-red-500
      "--toastProgressBackground": "rgb(220, 38, 38)", // bg-red-600
    },
  });

export const success = (m) =>
  toast.push(m, {
    theme: {
      "--toastBackground": "#10B981", // bg-green-500
      "--toastProgressBackground": "#059669", // bg-green-600
    },
  });
