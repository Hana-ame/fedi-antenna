# next.js

**always do not forrget learn next.js is for build fedi-antenna's frontend.**

**for the purpose of finishing it as soon as possible, just do the following pages**
- register
- loggin
- timeline
  - home timeline
  - local timeline
  - global timeline
- profile
  - with replies
  - without replies
- settings

located at fedi-antenna -b frontend/next-js/readme

[docs](https://nextjs.org/docs)

later
https://nextjs.org/learn/react-foundations

routing
https://nextjs.org/docs/app/building-your-application/routing/route-handlers
https://nextjs.org/docs/app/api-reference/file-conventions/route
https://nextjs.org/docs/pages/building-your-application/routing/api-routes

all follows https://nextjs.org/learn/dashboard-app/getting-started

## Getting Started

### Creating a new project

run
```sh
npx create-next-app@latest nextjs-dashboard --use-npm --example "https://github.com/vercel/next-learn/tree/main/dashboard/starter-example"
```

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

### Exploring the project

#### Folder structure

- `/app`: Contains all the routes, components, and logic for your application, this is where you'll be mostly working from.
  - `/app/lib`: Contains functions used in your application, such as reusable utility functions and data fetching functions.
   - `/app/ui`: Contains all the UI components for your application, such as cards, tables, and forms. To save time, we've pre-styled these components for you.
- `/public`: Contains all the static assets for your application, such as images.
- `/scripts`: Contains a seeding script that you'll use to populate your database in a later chapter.
- Config Files: You'll also notice config files such as `next.config.js` at the root of your application. Most of these files are created and pre-configured when you start a new project using create-next-app. You will not need to modify them in this course.

#### Placeholder data
in `/app/lib`

#### Running the development server
```sh
cd nextjs-dashboard
npm i
npm run dev
```

## CSS Styling

- How to add a global CSS file to your application.
- Two different ways of styling: Tailwind and CSS modules.
- How to conditionally add class names with the clsx utility package.

### Global styles
import css file.
```tsx
import '@/app/ui/global.css'; // add this line to import global CSS
```

it makes a trangle on it.
```tsx
<div
  className="h-0 w-0 border-b-[30px] border-l-[20px] border-r-[20px] border-b-black border-l-transparent border-r-transparent"
/>
```
this is about **[tailwind](https://tailwindcss.com/)**

### CSS Modules

```css
.shape {
  height: 0;
  width: 0;
  border-bottom: 30px solid black;
  border-left: 20px solid transparent;
  border-right: 20px solid transparent;
}
```
原理是四周的border在遇到其他border时会按对角线平分。
显示出来是三角形

this makes modules's visiable scope

### Using the clsx library to toggle class names

输入一个Object:
- value truly 时, key加入class
输入一个Array:
- 当中truly的变量会加入class
- 可以嵌套
```tsx
className={clsx(
  'inline-flex items-center rounded-full px-2 py-1 text-xs',
  {
    'bg-gray-100 text-gray-500': status === 'pending',
    'bg-green-500 text-white': status === 'paid',
  },
)}
```
### Other styling solutions
In addition to the approaches we've discussed, you can also style your Next.js application with:
- Sass which allows you to import .css and .scss files.
- CSS-in-JS libraries such as [styled-jsx](https://github.com/vercel/styled-jsx), [styled-components](https://github.com/vercel/next.js/tree/canary/examples/with-styled-components), and [emotion](https://github.com/vercel/next.js/tree/canary/examples/with-emotion).

## Optimizing Fonts and Images

- How to add custom fonts with next/font.
- How to add images with next/image.
- How fonts and images are optimized in Next.js.

### Why optimize fonts?

问题描述

后加载的内容会改变布局

### Adding a primary font

字体有优先级

_把body删掉了。会发生找不到root的error加载不出任何东西_

加入了
```css
.__className_725fdb {
    font-family: '__Inter_Fallback_725fdb', '__Inter_Fallback_Fallback_725fdb';
    font-style: normal;
}
```

### [Practice: Adding a secondary font](https://nextjs.org/learn/dashboard-app/optimizing-fonts-images#practice-adding-a-secondary-font)

```log
`next/font` error:
Preload is enabled but no subsets were specified for font `Lusitana`. Please specify subsets or disable preloading if your intended subset can't be preloaded.
Available subsets: `latin`

https://nextjs.org/docs/messages/google-fonts-missing-subsets
```

NextFont 类
```ts
import { Inter, Lusitana } from 'next/font/google';
 
export const inter = Inter({ subsets: ['latin'] });
export const lusitana = Lusitana({ weight: "400", subsets: ['latin'] });
```
这么得到

使用的时候在在class中加入`[fontInstance].className`

编译出来看不懂的。

[查询fonts](https://fonts.google.com/specimen/Lusitana?query=Lusitana+)

### Why optimize images?

问题描述：

- 各种设备屏幕不一样
- 指定不同的size
- 不准乱动
- lazy load

### The <Image> component

功能描述：

- 不会乱动
- autosize
- lazy load by default (as they enter the view)
- mordern formats

### Adding the desktop hero image

usage

```tsx
import Image from 'next/image';

<Image
  src="/hero-desktop.png"
  width={1000}
  height={760}
  className="hidden md:block"
  alt="Screenshots of the dashboard project showing desktop version"
/>
```

### Practice: Adding the mobile hero image
```tsx
<Image
  src="/hero-desktop.png"
  width={1000}
  height={760}
  className="block md:hidden"
  {/* className="hidden max-md:block"  */}
  {/* 有个白痴这么用。 */}
  alt="Screenshots of the dashboard project showing desktop version"
/>
```

hidden 和 block 分别是css的显示方式

_记得补基础_

### Recommended reading
There's a lot more to learn about these topics, including optimizing remote images and using local font files. If you'd like to dive deeper into fonts and images, see:

- [Image Optimization Docs](https://nextjs.org/docs/app/building-your-application/optimizing/images)
- [Font Optimization Docs](https://nextjs.org/docs/app/building-your-application/optimizing/fonts)
- [Improving Web Performance with Images (MDN)](https://developer.mozilla.org/en-US/docs/Learn/Performance/Multimedia)
- [Web Fonts (MDN)](https://developer.mozilla.org/en-US/docs/Learn/CSS/Styling_text/Web_fonts)

## Creating Layouts and Pages

- Create the dashboard routes using file-system routing.
- Understand the role of folders and files when creating new route segments.
- Create a nested layout that can be shared between multiple dashboard pages.
- Understand what colocation, partial rendering, and the root layout are.

能跟着路径走真的是

### Nested routing

You can create separate UIs for each route using `layout.tsx` and `page.tsx` files.

`page.tsx` is a special Next.js file that exports a React component, and it's required for the route to be accessible. In your application, you already have a page file: `/app/page.tsx` - this is the home page associated with the route `/`.

To create a nested route, you can nest folders inside each other and add `page.tsx` files inside them. For example:

![Diagram showing how adding a folder called dashboard creates a new route '/dashboard'](https://nextjs.org/_next/image?url=%2Flearn%2Fdark%2Fdashboard-route.png&w=3840&q=75&dpl=dpl_9VWEnshqwJ4TWYcNesdqAEpT4iWx)

`/app/dashboard/page.tsx` is associated with the `/dashboard` path. Let's create the page to see how it works!

**`page.tsx`会被导出为这个path的组件**

### Creating the dashboard page

By having a special name for page files, Next.js allows you to colocate UI components, test files, and other related code with your routes. Only the content inside the page file will be publicly accessible. For example, the /ui and /lib folders are colocated inside the /app folder along with your routes.

说你可以把东西丢一起。

### [Practice: Creating the dashboard pages](https://nextjs.org/learn/dashboard-app/creating-layouts-and-pages#practice-creating-the-dashboard-pages)

创建文件夹
创建`page.tsx`
```tsx
export default function Page() {
  return <></>;
}
```

### Creating the dashboard layout

```tsx
import SideNav from '@/app/ui/dashboard/sidenav';
 
export default function Layout({ children }: { children: React.ReactNode }) {
  return (
    <div className="flex h-screen flex-col md:flex-row md:overflow-hidden">
      <div className="w-full flex-none md:w-64">
        <SideNav />
      </div>
      <div className="flex-grow p-6 md:overflow-y-auto md:p-12">{children}</div>
    </div>
  );
}
```

children是这个和子文件夹的`page.tsx`导出的组件

![Folder structure with dashboard layout nesting the dashboard pages as children](https://nextjs.org/_next/image?url=%2Flearn%2Fdark%2Fshared-layout.png&w=3840&q=75&dpl=dpl_9VWEnshqwJ4TWYcNesdqAEpT4iWx)

One benefit of using layouts in Next.js is that on navigation, only the page components update while the layout won't re-render. This is called [partial rendering](https://nextjs.org/docs/app/building-your-application/routing/linking-and-navigating#3-partial-rendering):

记得 partial rendering 是大坑。

### Root layout

This is called a root layout and is required. Any UI you add to the root layout will be shared across all pages in your application. You can use the root layout to modify your <html> and <body> tags, and add metadata (you'll learn more about metadata in a later chapter).

```tsx
import '@/app/ui/global.css';
import { inter } from '@/app/ui/fonts';
 
export default function RootLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <html lang="en">
      <body className={`${inter.className} antialiased`}>{children}</body>
    </html>
  );
}
```

比vue挂载在`<div id="app"></div>`上面好吧。

## Navigating Between Pages

- How to use the next/link component.
- How to show an active link with the usePathname() hook.
- How navigation works in Next.js.

链接怎么做的

### The <Link> component

用 `<Link>` 替换 `<a>`

```tsx
import Link from 'next/link';

<Link
  key={link.name}
  href={link.href}
  className="flex h-[48px] grow items-center justify-center gap-2 rounded-md bg-gray-50 p-3 text-sm font-medium hover:bg-sky-100 hover:text-blue-600 md:flex-none md:justify-start md:p-2 md:px-3"
>
  <LinkIcon className="w-6" />
  <p className="hidden md:block">{link.name}</p>
</Link>
```

~~_懒加载呢_~~

Futhermore, in production, whenever <Link> components appear in the browser's viewport, Next.js automatically prefetches the code for the linked route in the background. By the time the user clicks the link, the code for the destination page will already be loaded in the background, and this is what makes the page transition near-instant!

看到Link时加载后面的东西。

Learn more about [how navigation works](https://nextjs.org/docs/app/building-your-application/routing/linking-and-navigating#how-routing-and-navigation-works).

### Pattern: Showing active links

Since `usePathname()` is a hook, you'll need to turn `nav-links.tsx` into a Client Component. Add React's `"use client"` directive to the top of the file, then import `usePathname()` from next/navigation:

https://nextjs.org/docs/app/api-reference/functions/use-pathname

_Client Component_
_hook_

## Setting Up Your Database

Before you can continue working on your dashboard, you'll need some data. In this chapter, you'll be setting up a PostgreSQL database using @vercel/postgres. If you're already familiar with PostgreSQL and would prefer to use your own provider, you can skip this chapter and set it up on your own. Otherwise, let's continue!

为啥使用postgres
最近流行这个么

- Push your project to GitHub.
- Set up a Vercel account and link your GitHub repo for instant previews and deployments.
- Create and link your project to a Postgres database.
- Seed the database with initial data.

跳过

```json
  "seed": "node -r dotenv/config ./scripts/seed.js"
```
```ts
const { db } = require('@vercel/postgres');
```

难绷

## Fetching Data

- Learn about some approaches to fetching data: APIs, ORMs, SQL, etc.
- How Server Components can help you access back-end resources more securely.
- What network waterfalls are.
- How to implement parallel data fetching using a JavaScript Pattern.

### Choosing how to fetch data

#### API layer

In Next.js, you can create API endpoints using [Route Handlers](https://nextjs.org/docs/app/building-your-application/routing/route-handlers).

是接收。

#### Database queries

When you're creating a full-stack application, you'll also need to write logic to interact with your database. For [relational databases](https://aws.amazon.com/relational-database/) like Postgres, you can do this with SQL, or an [ORM](https://vercel.com/docs/storage/vercel-postgres/using-an-orm#) like [Prisma](https://www.prisma.io/).

#### Using Server Components to fetch data

By default, Next.js applications use React Server Components. Fetching data with Server Components is a relatively new approach and there are a few benefits of using them:

- Server Components support promises, providing a simpler solution for asynchronous tasks like data fetching. You can use async/await syntax without reaching out for useEffect, useState or data fetching libraries.
- Server Components execute on the server, so you can keep expensive data fetches and logic on the server and only send the result to the client.
- As mentioned before, since Server Components execute on the server, you can query the database directly without an additional API layer.

**默认是server组件**
**这个是next.js的概念**

_怎么编译的_
_放在什么地方_

#### Using SQL

**extra**

```sh
npm install prisma @prisma/client
```

[prisma](https://vercel.com/docs/storage/vercel-postgres/using-an-orm#prisma)

see [prostgress.md](postgres.md)

#### Fetching data for the dashboard overview page

**stopping here**

应该用prisma的rawquery很容易替换sql
但还没试。

不是
哥
你怎么把组件里面注释掉了。

queryRaw都拉出了什么东西啊。

参考`app/lib/data.ts`
自己改了的部分

## Static and Dynamic Rendering

- What static rendering is and how it can improve your application's performance.
- What dynamic rendering is and when to use it.
- Different approaches to make your dashboard dynamic.
- Simulate a slow data fetch to see what happens.

### What is Static Rendering?

就是静态资源

### What is Dynamic Rendering?
With dynamic rendering, content is rendered on the server for each user at request time (when the user visits the page). There are a couple of benefits of dynamic rendering:

- Real-Time Data - Dynamic rendering allows your application to display real-time or frequently updated data. This is ideal for applications where data changes often.
- User-Specific Content - It's easier to serve personalized content, such as dashboards or user profiles, and update the data based on user interaction.
- Request Time Information - Dynamic rendering allows you to access information that can only be known at request time, such as cookies or the URL search parameters.

为啥cookie和param要在request的时候才知道。

### Making the dashboard dynamic

_`noCache()`是啥_

### Simulating a Slow Data Fetch

会卡在加载的时候

## Streaming

In the previous chapter, you made your dashboard page dynamic, however, we discussed how the slow data fetches can impact the performance of your application. Let's look at how you can improve the user experience when there are slow data requests.

草所以`noCache()`是不要编译进去的意思啊。
_所以server side和client side怎么分的_

- What streaming is and when you might use it.
- How to implement streaming with loading.tsx and Suspense.
- What loading skeletons are.
- What route groups are, and when you might use them.
- Where to place Suspense boundaries in your application.

### What is streaming?

Streaming is a data transfer technique that allows you to break down a route into smaller "chunks" and progressively stream them from the server to the client as they become ready.

![Diagram showing time with sequential data fetching and parallel data fetching](https://nextjs.org/_next/image?url=%2Flearn%2Fdark%2Fserver-rendering-with-streaming.png&w=3840&q=75&dpl=dpl_BbSpPdzv9Yrsi74LnqWRCSDNSUNs)

By streaming, you can prevent slow data requests from blocking your whole page. This allows the user to see and interact with parts of the page without waiting for all the data to load before any UI can be shown to the user.

![Diagram showing time with sequential data fetching and parallel data fetching](https://nextjs.org/_next/image?url=%2Flearn%2Fdark%2Fserver-rendering-with-streaming-chart.png&w=3840&q=75&dpl=dpl_BbSpPdzv9Yrsi74LnqWRCSDNSUNs)

Streaming works well with React's component model, as each component can be considered a chunk.

There are two ways you implement streaming in Next.js:

- At the page level, with the `loading.tsx` file.
- For specific components, with `<Suspense>`.
Let's see how this works.

### Streaming a whole page with `loading.tsx`

怎么这样的。

- 先返回`loading.tsx`
- 等待`page.tsx`加载完成
- 返回`page.tsx`

A few things are happening here:

loading.tsx is a special Next.js file built on top of Suspense, it allows you to create fallback UI to show as a replacement while page content loads.
- Since <Sidebar> is static, so it's shown immediately. The user can interact with <Sidebar> while the dynamic content is loading.
- The user doesn't have to wait for the page to finish loading before navigating away (this is called interruptable navigation).
- Congratulations! You've just implemented streaming. But we can do more to improve the user experience. Let's show a loading skeleton instead of the Loading… text.

#### Adding loading skeletons

改`loading.tsx`

挺好看的

#### Fixing the loading skeleton bug with route groups

有啥错

Right now, your loading skeleton will apply to the invoices and customers pages as well.

下面都能用的。
_这个规则对`page.tsx`和`layout.tsx`是怎么工作的来着。。。_

Since loading.tsx is a level higher than `/invoices/page.tsx` and `/customers/page.tsx` in the file system, it's also applied to those pages.

We can change this with [Route Groups](https://nextjs.org/docs/app/building-your-application/routing/route-groups). Create a new folder called `/(overview)` inside the dashboard folder. Then, move your `loading.tsx` and `page.tsx` files inside the folder:

**在`./(overview)`文件夹下的`loading.tsx`, `page.tsx`只对当前目录有效果(对子目录没效果)**

_修改了。next下面的page。ts_

Here, you're using a route group to ensure loading.tsx only applies to your dashboard overview page. However, you can also use route groups to separate your application into sections (e.g. (marketing) routes and (shop) routes) or by teams for larger applications.

_啥意思_

#### Streaming a component

绷。你管这叫streaming。
好吧你说了算

难绷
和自己悟出来的思路差不多

### Grouping components

为啥不好看
我觉得挺好看
：你觉得。。

### Deciding where to place your Suspense boundaries

爱咋用咋用

### Looking ahead

Streaming and Server Components give us new ways to handle data fetching and loading states, ultimately with the goal of improving the end user experience.

In the next chapter, you'll learn about Partial Prerendering, a new Next.js rendering model built with streaming in mind.
