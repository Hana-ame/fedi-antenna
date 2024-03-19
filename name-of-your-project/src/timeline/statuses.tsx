import { useLoaderData } from "react-router-dom";
import { IStatus } from "../api/@types";
import Status from "../components/status";
import loadJSON from "../assets/loadjson";
import useSWR from "swr";

interface LoaderArgs {
  request: Request;
  params: any;
}

export async function loader({ request, params }: LoaderArgs) {
  const url = new URL(request.url);
  const page = url.searchParams.get("page") ?? "0";
  const statuses = loadJSON(page)
  return { statuses };
}

export function Statuses({ maxId, page }: { maxId?: number, page?: number }) {
  const { data, error, isLoading } = useSWR('statuses' + (page?.toString() ?? '0'), loadJSON)
  if (isLoading) {
    return <>loading</>
  }
  if (error) {
    console.log(error)
    return <>{JSON.stringify(error)}</>
  }
  const statuses = data as IStatus[]
  console.log(statuses)
  return (
    <div>
      {statuses.map(status => <Status status={status}></Status>)}
    </div>
  )
}