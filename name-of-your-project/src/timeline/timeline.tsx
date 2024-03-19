import render from "../html/html";
import { Statuses } from "./statuses";

export function Timeline() {
  const child = render('<p>Take Ownership of Your Future Self<br />----<br />- 18 minutes ago | 14 points | 0 comments<br />- URL: <a href="https://hbr.org/2020/08/take-ownership-of-your-future-self" target="_blank" rel="nofollow noopener noreferrer" translate="no"><span class="invisible">https://</span><span class="ellipsis">hbr.org/2020/08/take-ownership</span><span class="invisible">-of-your-future-self</span></a><br />- Discussions: <a href="https://news.ycombinator.com/item?id=39742349" target="_blank" rel="nofollow noopener noreferrer" translate="no"><span class="invisible">https://</span><span class="ellipsis">news.ycombinator.com/item?id=3</span><span class="invisible">9742349</span></a><br />- Summary: This article discusses the concept of the &quot;end of history illusion,&quot; where people overestimate the stability of their current selves and underestimate the potential for personal change. The article suggests three strategies to take ownership of your future self: recognizing the differences between your former, current, and future selves; imagining and planning for your desired future self; and changing your identity narrative to align with your future self. The article emphasizes the importance of being intentional about personal growth and development.</p>')
  return (
    <Statuses></Statuses>
  )
}