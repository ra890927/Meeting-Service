import { EventInput } from '@fullcalendar/core';

let eventGuid = 0;
export function createEventId() {
  return String(eventGuid++);
}
