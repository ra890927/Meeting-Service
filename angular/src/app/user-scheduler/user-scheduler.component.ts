import { Component, signal, ChangeDetectorRef, OnInit } from '@angular/core';
import { FullCalendarModule } from '@fullcalendar/angular';
import { CalendarOptions, DateSelectArg, EventClickArg, EventApi, DatesSetArg} from '@fullcalendar/core';
import {MatFormFieldModule, MatHint} from '@angular/material/form-field';
import dayGridPlugin from '@fullcalendar/daygrid';
import { Router } from '@angular/router';
import timeGridPlugin from '@fullcalendar/timegrid';
import interactionPlugin from '@fullcalendar/interaction';
import { MatButtonModule } from '@angular/material/button';
import {MatCardModule} from '@angular/material/card';
import { FormsModule } from '@angular/forms';
import {MatTooltipModule} from '@angular/material/tooltip';
import {MatSelectModule} from '@angular/material/select';
import { MatDialog } from '@angular/material/dialog';
import { CommonModule } from '@angular/common';
import {MatDatepickerModule} from '@angular/material/datepicker';
// import listPlugin from '@fullcalendar/list';
import {createEventId} from './event-utils';//test use
import { S, cl, co, s } from '@fullcalendar/core/internal-common';
import { MatInputModule } from '@angular/material/input';
import { MatIconModule } from '@angular/material/icon';
import {provideNativeDateAdapter} from '@angular/material/core';
import { MeetingDetailComponent } from './meeting-detail/meeting-detail.component';
import { UserService } from '../API/user.service';
//need to save in a interface file
interface Room {
  id: number;
  room_name: string;
  type: string;
  rules: number[];
  capacity: number;
}
interface RoomInfo{
  room_name: string;
  rules: Tag[];
} 
interface RoomEvent {
  id: string;
  room_id: number;
  title: string;
  description: string;
  participants: number[];
  start: string;
  end: string;
  OrganizerID: number;
}

interface Tag{
  id: number;
  type_name: string;
  type_desc: string;
}
interface User{
  id: number;
  username: string;
  email: string;
}
interface needData{
  id: string;
  title: string;
  description: string;
  participants: User[];
  start: string;
  end: string;
  Organizer: User;
  RoomDetail: RoomInfo;
}
const TODAY_STR = new Date().toISOString().replace(/T.*$/, ''); // YYYY-MM-DD of today
@Component({
  selector: 'app-user-scheduler',
  standalone: true,
  imports: [FullCalendarModule,
     FormsModule, 
     MatFormFieldModule,
     MatInputModule, 
     MatCardModule, 
     CommonModule,
     MatTooltipModule,
     MatSelectModule, 
     MatButtonModule,
     MatIconModule,
     MatDatepickerModule],
  templateUrl: './user-scheduler.component.html',
  styleUrl: './user-scheduler.component.css'
})
export class UserSchedulerComponent {
  constructor(private changeDetector: ChangeDetectorRef, private router: Router, private dialog: MatDialog, private userService: UserService) {
  }
  fakeData: Room[] = [];
  fakeEvents: RoomEvent[] = [];
  fakeTags: Tag[] = [];
  fakeUsers: User[] = [];
  User: any;//get from session storage
  isMonthView: boolean = false;//initial view is week
  detailsInfo: needData[] = [];
  calendarOptions: CalendarOptions = {
    plugins: [
      dayGridPlugin,
      timeGridPlugin,
      interactionPlugin
    ],
    headerToolbar: {
      left: 'prev,next today',
      center: 'title',
      right: 'dayGridMonth,timeGridWeek'
    },
    initialView: 'timeGridWeek',
    eventTimeFormat: {
      hour: '2-digit',
      minute: '2-digit',
      hour12: false
    },
    defaultAllDay: false,
    events: [],
    weekends: true,
    editable: false,
    selectable: false,
    selectMirror: true,
    dayMaxEvents: 2,
    allDaySlot: false,
    selectOverlap: false,
    eventOverlap: false,
    validRange: {
      start: '7:00:00',
      end: '23:59:59'
    },
    slotMaxTime: '23:59:59',
    slotMinTime: '7:00:00',
    firstDay: 7,
    eventClick: this.handleEventClick.bind(this),
    eventsSet: this.handleEvents.bind(this),
    datesSet: this.handleViewChange.bind(this),
  };
  ngOnInit(){
    // if(!this.userService.isLoggedIn()){
    //   this.router.navigate(['/login']);
    // }
    this.fakeData = [//getAllRooms from backend
      {id: 1, room_name: 'room1', type: 'meeting', rules: [1], capacity: 10},
      {id: 2, room_name: 'room2', type: 'meeting', rules: [1, 2], capacity: 20},
      {id: 3, room_name: 'room3', type: 'meeting', rules: [2], capacity: 30},
      {id: 4, room_name: 'room4', type: 'meeting', rules: [1], capacity: 40},
      {id: 5, room_name: 'room5', type: 'meeting', rules: [1, 2], capacity: 50},
    ];
    this.fakeTags = [//getALLTags from backend
      {id: 1,type_name: 'projector', type_desc: 'can use projector'},
      {id: 2,type_name: 'food', type_desc: 'can eat food'},
    ];
    this.fakeEvents = [//getEvents from backend
      {id: '1', room_id: 1, title: 'meeting1', description: 'meeting1', start: TODAY_STR + 'T09:00:00', end: TODAY_STR + 'T13:00:00', participants: [1,2,3], OrganizerID: 1},
      {id: '2', room_id: 2, title: 'meeting2', description: 'meeting2', start: TODAY_STR + 'T12:00:00', end: TODAY_STR + 'T15:00:00', participants: [2,3], OrganizerID: 2},
      {id: '3', room_id: 3, title: 'meeting3', description: 'meeting3', start: TODAY_STR + 'T14:00:00', end: TODAY_STR + 'T17:00:00', participants: [1,2], OrganizerID: 3},
      {id: '4', room_id: 4, title: 'meeting4', description: 'meeting4', start: TODAY_STR + 'T16:00:00', end: TODAY_STR + 'T19:00:00', participants: [1,3,5], OrganizerID: 4},
      {id: '5', room_id: 5, title: 'meeting5', description: 'meeting5', start: TODAY_STR + 'T18:00:00', end: TODAY_STR + 'T21:00:00', participants: [2,4], OrganizerID: 5},
    ];
    this.fakeUsers = [//getAllUsers from backend
      {id: 1, username: 'user1', email: '1@1.com'},
      {id: 2, username: 'user2', email: '2@2.com'},
      {id: 3, username: 'user3', email: '3@3.com'},
      {id: 4, username: 'user4', email: '4@4.com'},
      {id: 5, username: 'user5', email: '5@5.com'},
    ];
    this.User = this.userService.getUser();
    if(!this.User){ 
      this.User = {id: 1, username: 'user1', email: 'admin@admin.com'};
    }
    this.detailsInfo = this.fakeEvents.map((event) => {
      return {
        id: event.id,
        title: event.title,
        description: event.description,
        participants: event.participants.map((id) => {
          const user = this.fakeUsers.find((user) => user.id === id);
          if (!user) {
            throw new Error(`User with id ${id} not found`);
          }
          return user;
        }),
        start: event.start,
        end: event.end,
        Organizer:(() => {
          const organizer = this.fakeUsers.find((user) => user.id === event.OrganizerID);
          if (!organizer) {
            throw new Error(`Organizer with ID ${event.OrganizerID} not found`);
          }
          return organizer;
        })(),

        RoomDetail: (() => {
          const room = this.fakeData.find((room) => room.id === event.room_id);
          if (!room) {
            throw new Error(`Room with ID ${event.room_id} not found`);
          }
          return {
            room_name: room.room_name,
            rules: room.rules.map((id) => {
              const tag = this.fakeTags.find((tag) => tag.id === id);
              if (!tag) {
                throw new Error(`Tag with ID ${id} not found`);
              }
              return tag;
            }),
          };
        })(),
      };
    });
    console.log(this.detailsInfo);
    this.calendarOptions.events = this.detailsInfo;
  }
  calendarVisible = signal(true);

  currentEvents = signal<EventApi[]>([]);

  handleCalendarToggle() {
    this.calendarVisible.update((bool) => !bool);
  }
  handleViewChange(viewInfo: DatesSetArg) {
    this.isMonthView = viewInfo.view.type === 'dayGridMonth';
  }
  handleEventClick(clickInfo: EventClickArg) {
    this.openDialog(clickInfo.event.id);
  }
  handleEvents(events: EventApi[]) {
    this.currentEvents.set(events);
    this.changeDetector.detectChanges();
  }
  openDialog(id: string) {
    const dialogRef = this.dialog.open(MeetingDetailComponent, {
      width: '50%',
      height: '50%',
      data: this.getEventById(id),
    });
  }
  //get event by id
  getEventById(id: string){
    return this.detailsInfo.find((event) => event.id === id);
  }
}