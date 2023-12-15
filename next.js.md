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
