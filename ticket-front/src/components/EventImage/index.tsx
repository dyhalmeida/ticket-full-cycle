import Image from "next/image";

export type TEventImageProps = {
  src: string;
  alt: string;
};

export function EventImage(props: TEventImageProps) {
  return (
    <Image
      src={props.src}
      alt={props.alt}
      width={277}
      height={277}
      priority
      className="rounded-2xl"
    />
  );
}
