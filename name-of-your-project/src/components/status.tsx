import { IStatus } from "../api/@types";
import P from "../html/p";

export default function Status({ status }: { status: IStatus } ) {
  return (
    <P>
      <div>
        {status.account.acct}  
      </div>
      <div>
        {status.spoiler_text}  
        {status.content}  
      </div>
      <div>
        <span>123</span>
      </div>
      <br />
    </P>
  )  
}