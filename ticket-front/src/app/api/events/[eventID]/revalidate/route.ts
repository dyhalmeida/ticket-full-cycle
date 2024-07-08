import { revalidateTag } from "next/cache";
import { NextRequest } from "next/server";

export function POST(
  request: NextRequest,
  { params }: { params: { eventID: string } }
) {
  revalidateTag("events");
  revalidateTag(`events/${params.eventID}`);

  return new Response(null, { status: 204 });
}
