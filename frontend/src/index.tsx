import React from 'react';
import ReactDOM from 'react-dom';
import * as Sentry from "@sentry/react";
import { Integrations } from "@sentry/tracing";
import ReactGA from 'react-ga';
import './index.css';
import App from './App';


if (!process.env.NODE_ENV || process.env.NODE_ENV === 'development') {
  // dev code
  ReactGA.initialize("UA-201437458-2")
} else {
  Sentry.init({
    dsn: "https://1e4cb505b7a2479898cb19d788774ab8@o37323.ingest.sentry.io/5851622",
    integrations: [new Integrations.BrowserTracing()],

    // Set tracesSampleRate to 1.0 to capture 100%
    // of transactions for performance monitoring.
    // We recommend adjusting this value in production
    tracesSampleRate: 0.2,
  });

  ReactGA.initialize("UA-201437458-1")
}

ReactGA.pageview(window.location.pathname + window.location.search);

ReactDOM.render(
  <React.StrictMode>
    <App />
  </React.StrictMode>,
  document.getElementById('root')
);
