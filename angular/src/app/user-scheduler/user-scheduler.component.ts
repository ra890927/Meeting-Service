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
import { C, S, cl, co, s } from '@fullcalendar/core/internal-common';
import { MatInputModule } from '@angular/material/input';
import { MatIconModule } from '@angular/material/icon';
import {provideNativeDateAdapter} from '@angular/material/core';
import { MeetingDetailComponent } from './meeting-detail/meeting-detail.component';
import { UserService } from '../API/user.service';
import { ItemService } from '../API/item.service';
import { forkJoin, map } from 'rxjs';
//need to save in a interface file
interface Room {
  id: number;
  room_name: string;
  type: string;
  rules: number[];
  capacity: number;
}
interface EventResponse {
  data: {
    meetings: Event[];
  };
  message: string;
  status: string;
}
interface RoomResponse {
  data: {
    rooms: Room[];
  };
  message: string;
  status: string;
}
interface RoomInfo{
  room_name: string;
  rules: CodeValue[];
} 
interface Event{
  id: string;
  title: string;
  description: string;
  start_time: string;
  end_time: string;
  organizer: number;
  participants: number[];
  room_id: number;
  status_type: string;
}
interface UserResponse {
  data: {
    users: User[];
  };
  message: string;
  status: string;
}
interface CodeValue {
  id: number;
  tag: string;
  description: string;
  codeTypeId: number;
}
interface User{
  id: number;
  username: string;
  email: string;
  role?: string
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
  constructor(private changeDetector: ChangeDetectorRef, private router: Router, private dialog: MatDialog, private userService: UserService, private itemService: ItemService) {
  }
  RoomData: Room[] = [];
  UserData: User[] = [];
  EventData: Event[] = [];
  TagData: CodeValue[] = [];
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
    slotLabelFormat: {
      hour: '2-digit',
      minute: '2-digit',
      hour12: false
    },
    nowIndicator: true,
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
    if(!this.userService.isLoggedIn()){
      this.router.navigate(['/login']);
    }
    else{
      this.User = this.userService.getUser();
    }
    forkJoin({
      rooms: this.itemService.getAllRooms().pipe(
        map((response: RoomResponse) => {
          if (response.status === 'success') {
            return response.data.rooms;
          } else {
            console.log('get room data failed');
            return [] as Room[];
          }
        })
      ),
      users: this.itemService.getAllUsers().pipe(
        map((response: UserResponse) => {
          if (response.status === 'success') {
            return response.data.users;
          } else {
            console.log('get user data failed');
            return [] as User[];
          }
        })
      ),
      meetings: this.itemService.getMeetingByUserId(this.User.id).pipe(
        map((response: EventResponse) => {
          if (response.status === 'success') {
            console.log("response.data.meetings", response.data.meetings);
            return response.data.meetings;
          } else {
            console.log('get meeting data failed');
            return [] as Event[];
          }
        })
      )
    }).subscribe(({ rooms, users, meetings }) => {
      this.RoomData = rooms;
      this.UserData = users;
      this.EventData = meetings;
      this.itemService.getAllTags().subscribe((response: CodeValue[]) => {
        if (response) {
            this.TagData = response;
            this.generateDetailsInfo();
            console.log(this.detailsInfo);
            this.calendarOptions.events = this.detailsInfo;
        }else{
          console.log('get tag data failed');
        }
      });

    });
  }
  calendarVisible = signal(true);

  currentEvents = signal<EventApi[]>([]);
  generateDetailsInfo() {
    this.detailsInfo = this.EventData.map((event) => {
      return {
        id: event.id,
        title: event.title,
        description: event.description,
        participants: event.participants.map((id) => {
          const user = this.UserData.find((user) => user.id === id);
          if (!user) {
            throw new Error(`User with id ${id} not found`);
          }
          return user;
        }),
        start: event.start_time,
        end: event.end_time,
        Organizer:(() => {
          const organizer = this.UserData.find((user) => user.id === event.organizer);
          if (!organizer) {
            throw new Error(`Organizer with ID ${event.organizer} not found`);
          }
          return organizer;
        })(),

        RoomDetail: (() => {
          const room = this.RoomData.find((room) => room.id === event.room_id);
          console.log("room data", this.RoomData);
          console.log("event", event);
          if (!room) {
            return {
              room_name: 'Room not found',
              rules: [], 

              // room.rules.map((id) => {
              //   const tag = this.TagData.find((tag) => tag.id === id);
              //   console.log(this.TagData);
              //   if (!tag) {
              //     throw new Error(`Tag with ID ${id} not found`);
              //   }
              //   return tag;
              // }),
            };
            // throw new Error(`Room with ID ${event.room_id} not found`);
          }
          return {
            room_name: room.room_name,
            rules: room.rules.map((id) => {
              const tag = this.TagData.find((tag) => tag.id === id);
              console.log(this.TagData);
              if (!tag) {
                return {
                  id: 0,
                  tag: 'Tag not found',
                  description: 'Tag not found',
                  codeTypeId: 0
                };
                // throw new Error(`Tag with ID ${id} not found`);
              }
              return tag;
            }),
          };
        })(),
      };
    });
    console.log("detailsInfo", this.detailsInfo);
  }
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