import React from "react";
import ServerStatus from "./components/ServerStatus";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";

const queryClient = new QueryClient();

function App() {
  return (
    <QueryClientProvider client={queryClient}>
      <div className="p-5">
        <h1 className="text-xl font-bold">Load Balancer Dashboard</h1>
        <ServerStatus />
      </div>
    </QueryClientProvider>
  );
}

export default App;