import React from "react";
import ReactDOM from "react-dom/client";
import App from "./App";
import { createClient as createWSClient } from "graphql-ws";

const wsClient = createWSClient({
  url: "ws://localhost:5000/subscription",
});

import {
  Provider,
  createClient,
  defaultExchanges,
  subscriptionExchange,
} from "urql";
import "./index.css";

const client = createClient({
  url: "http://localhost:5000/query",
  exchanges: [
    ...defaultExchanges,
    subscriptionExchange({
      forwardSubscription: (operation) => ({
        subscribe: (sink) => ({
          unsubscribe: wsClient.subscribe(operation, sink),
        }),
      }),
    }),
  ],
});

ReactDOM.createRoot(document.getElementById("root") as HTMLElement).render(
  <React.StrictMode>
    <Provider value={client}>
      <App />
    </Provider>
  </React.StrictMode>
);
