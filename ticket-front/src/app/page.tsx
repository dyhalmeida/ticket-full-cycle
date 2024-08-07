import { Title } from "@/components/Title";
import { EventCard } from "@/components/EventCard";
import { TEvent } from "@/models";

export async function getEvents(): Promise<TEvent[]> {
  const response = await fetch(`${process.env.GOLANG_API_URL}/events`, {
    headers: {
      "apikey": process.env.GOLANG_API_TOKEN as string
    },
    cache: "no-store",
    // next: {
    //   tags: ["events"],
    // }
  });

  return (await response.json()).events as TEvent[];
}

export default async function HomePage() {

  const events = await getEvents()

  return (
    <main className="mt-10 flex flex-col">
      <Title>Eventos disponíveis</Title>
      <div className="mt-8 sm:grid sm:grid-cols-auto-fit-cards flex flex-wrap justify-center gap-x-2 gap-y-4">
        {events.map((event) => (
          <EventCard key={event.id} event={event} />
        ))}
      </div>
    </main>
  );
}
