import React from "react";
import { useServers, useToggleServer } from "../api";

const ServerStatus = () => {
  const { data: servers = [], isLoading } = useServers();
  const toggleServer = useToggleServer();

  if (isLoading) return <p>Loading servers...</p>;

  return (
    <div className="p-4">
      <h2 className="text-lg font-bold">Server Status</h2>
      <ul>
        {servers.map((server, index) => (
          <li key={index} className="p-2 flex justify-between items-center">
            <span className={`${server.active ? "text-green-600" : "text-red-600"}`}>
              {server.url} - {server.active ? "Active" : "Down"}
            </span>
            <button
              onClick={() => toggleServer.mutate(server.url)}
              className="px-3 py-1 bg-blue-500 text-white rounded hover:bg-blue-600"
            >
              {server.active ? "Disable" : "Enable"}
            </button>
          </li>
        ))}
      </ul>
    </div>
  );
};

export default ServerStatus;
