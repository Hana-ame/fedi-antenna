# install
follow
- https://reactrouter.com/en/main/start/tutorial#setup
- https://tailwindcss.com/docs/installation/using-postcss

```sh
npm create vite@latest name-of-your-project -- --template react
# follow prompts
cd name-of-your-project
npm install react-router-dom localforage match-sorter sort-by

# tailwind
npm install -D tailwindcss postcss autoprefixer
npx tailwindcss init

npm run dev
```

install in vscode 
https://stackoverflow.com/questions/70597844/unknown-at-rule-tailwind-css-in-reactjs  
https://marketplace.visualstudio.com/items?itemName=csstools.postcss  

# vite
https://vitejs.dev/config/server-options  

# router
https://reactrouter.com/en/main/start/tutorial#setup

## https://reactrouter.com/en/main/start/tutorial#adding-a-router 
in this step, it makes `/` can be accessed and return soemthing.  
an undefined rotuer returns error.

## https://reactrouter.com/en/main/start/tutorial#the-root-route
page added in router list, prop = element.

## https://reactrouter.com/en/main/start/tutorial#handling-not-found-errors
error page added in router list, prop = errorElement.

## https://reactrouter.com/en/main/start/tutorial#the-contact-route-ui
> However, it's not inside of our root layout 😠

~~this means that it was hooked directly on the app root not under the `<App />` component.~~  
ps: cat icon is unavaliable.

## https://reactrouter.com/en/main/start/tutorial#nested-routes
oops, it means that the side bar is not appeared.  

## https://reactrouter.com/en/main/start/tutorial#client-side-routing
use `<Link to={'/path/to/file'}></Link>` instad `<a href="/path/to/file"></a>`

## https://reactrouter.com/en/main/start/tutorial#loading-data
datas are stored at localstorage.

## https://reactrouter.com/en/main/start/tutorial#data-writes--html-forms
is that saying POST is not avaliable?

## https://reactrouter.com/en/main/start/tutorial#creating-contacts
how to refresh the side bar.

## https://reactrouter.com/en/main/start/tutorial#url-params-in-loaders
hook is located in `contact.jsx`, registed in `main.tsx`

