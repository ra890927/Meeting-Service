import { Component, signal, ChangeDetectorRef } from '@angular/core';
import { FullCalendarModule } from '@fullcalendar/angular';
import { CalendarOptions, DateSelectArg, EventClickArg, EventApi  } from '@fullcalendar/core';
import dayGridPlugin from '@fullcalendar/daygrid';
import { Router } from '@angular/router';
import timeGridPlugin from '@fullcalendar/timegrid';
import interactionPlugin from '@fullcalendar/interaction';


import { CommonModule } from '@angular/common';
// import listPlugin from '@fullcalendar/list';
import { INITIAL_EVENTS, createEventId } from './event-utils';


@Component({
  selector: 'app-room-scheduler',
  standalone: true,
  imports: [FullCalendarModule],
  templateUrl: './room-scheduler.component.html',
  styleUrl: './room-scheduler.component.css'
})
export class RoomSchedulerComponent {
  calendarVisible = signal(true);

  calendarOptions: CalendarOptions = {
    plugins: [
      dayGridPlugin,
      timeGridPlugin,
      interactionPlugin
    ],
    headerToolbar: {
      left: 'prev,next today',
      center: 'title',
      right: 'dayGridMonth,timeGridWeek,timeGridDay'
    },

    // initialView: 'dayGridMonth',
    // editable: true,
    // // weekends: false,
    // firstDay: 1,
    // events: [
    //   { title: 'Meeting', start: new Date() }
    // ]
    initialView: 'dayGridMonth',
    initialEvents: INITIAL_EVENTS, // alternatively, use the `events` setting to fetch from a feed
    weekends: true,
    editable: true,
    selectable: true,
    selectMirror: true,
    dayMaxEvents: true,
    firstDay: 1,
    select: this.handleDateSelect.bind(this),
    eventClick: this.handleEventClick.bind(this),
    eventsSet: this.handleEvents.bind(this)
  };

  currentEvents = signal<EventApi[]>([]);

  constructor(private changeDetector: ChangeDetectorRef, private router: Router) {
  }

  handleCalendarToggle() {
    this.calendarVisible.update((bool) => !bool);
  }

  // handleWeekendsToggle() {
  //   this.calendarOptions.update((options) => ({
  //     ...options,
  //     weekends: !options.weekends,
  //   }));
  // }

  handleDateSelect(selectInfo: DateSelectArg) {
    const title = prompt('Please enter a new title for your event');
    const calendarApi = selectInfo.view.calendar;

    calendarApi.unselect(); // clear date selection

    if (title) {
      calendarApi.addEvent({
        id: createEventId(),
        title,
        start: selectInfo.startStr,
        end: selectInfo.endStr,
        allDay: selectInfo.allDay
      });
    }
  }

  handleEventClick(clickInfo: EventClickArg) {
    if (confirm(`Are you sure you want to delete the event '${clickInfo.event.title}'`)) {
      clickInfo.event.remove();
    }
  }

  handleEvents(events: EventApi[]) {
    this.currentEvents.set(events);
    this.changeDetector.detectChanges(); // workaround for pressionChangedAfterItHasBeenCheckedError
  }
}


