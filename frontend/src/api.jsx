import { useQuery, useMutation, useQueryClient } from "@tanstack/react-query";

const API_URL = "http://localhost:8080";

export const fetchServers = async () => {
  const response = await fetch(`${API_URL}/servers`);
  return response.json();
};

export const toggleServer = async (url) => {
  await fetch(`${API_URL}/toggle-server`, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ url }),
  });
};

export const useServers = () => {
  return useQuery({ queryKey: ["servers"], queryFn: fetchServers, refetchInterval: 5000 });
};

export const useToggleServer = () => {
  const queryClient = useQueryClient();
  return useMutation({
    mutationFn: toggleServer,
    onSuccess: () => queryClient.invalidateQueries(["servers"]),
  });
};