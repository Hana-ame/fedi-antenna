**(this log file is informal)**

# environment

## proxy

```sh
npm config set proxy http://127.0.0.1:10809
npm config set https-proxy http://127.0.0.1:10809
```

## npm and npx

for some reason i don't even know what is `npm` and what is `npx`

i don't know why 
```sh
npx create-next-app@latest nextjs-dashboard --use-npm --example "https://github.com/vercel/next-learn/tree/main/dashboard/starter-example"
```
and why 
```sh
npm create-next-app@latest nextjs-dashboard --example "https://github.com/vercel/next-learn/tree/main/dashboard/starter-example"
```
doesn't work

## npx

after install it gives such message. 
```sh
  npm run dev
    Starts the development server.

  npm run build
    Builds the app for production.

  npm start
    Runs the built app in production mode.

We suggest that you begin by typing:

  cd nextjs-dashboard
  npm run dev
```

## use Inter()
```log
Retrying 3/3...
AbortError: The user aborted a request.
    at abort (C:\workplace\fedi\next-js\learn-1\nextjs-dashboard\node_modules\next\dist\compiled\node-fetch\index.js:1:65594)
    at EventTarget.abortAndFinalize (C:\workplace\fedi\next-js\learn-1\nextjs-dashboard\node_modules\next\dist\compiled\node-fetch\index.js:1:65814)
    at [nodejs.internal.kHybridDispatch] (node:internal/event_target:735:20)
    at EventTarget.dispatchEvent (node:internal/event_target:677:26)
    at abortSignal (node:internal/abort_controller:308:10)
    at AbortController.abort (node:internal/abort_controller:338:5)
    at Timeout.<anonymous> (C:\workplace\fedi\next-js\learn-1\nextjs-dashboard\node_modules\next\dist\compiled\@next\font\dist\google\fetch-css-from-google-fonts.js:44:47)
    at listOnTimeout (node:internal/timers:569:17)
  type: 'aborted'
}
 ⨯ Failed to download `Inter` from Google Fonts. Using fallback font instead.

Failed to fetch `Inter` from Google Fonts.}

```