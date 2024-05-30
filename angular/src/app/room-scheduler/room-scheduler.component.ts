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
import { PopUpFormComponent} from './pop-up-form/pop-up-form.component';
import { PopUpDeleteConfirmComponent } from './pop-up-delete-confirm/pop-up-delete-confirm.component';
import { MatDialog } from '@angular/material/dialog';
import { CommonModule } from '@angular/common';
import {MatDatepickerModule} from '@angular/material/datepicker';
// import listPlugin from '@fullcalendar/list';
import {createEventId} from './event-utils';//test use
import { R, S, cl, co, s } from '@fullcalendar/core/internal-common';
import { MatInputModule } from '@angular/material/input';
import { MatIconModule } from '@angular/material/icon';
import {provideNativeDateAdapter} from '@angular/material/core';
import { UserService } from '../API/user.service';
import { AuthService } from '../API/auth.service';
import { ItemService } from '../API/item.service';
import { interval, Subscription, switchMap } from 'rxjs';
//need to save in a interface file
interface Room {
  id: number;
  room_name: string;
  type: string;
  rules: number[];
  capacity: number;
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
interface CodeValue {
  code_type_id: number;
  code_value: string;
  code_value_desc: string;
  id: number;
}

interface CodeType {
  code_values: CodeValue[];
  id: number;
  type_desc: string;
  type_name: string;
}

interface CodeTypesResponse {
  data: {
    code_types: CodeType[];
  };
  message: string;
  status: string;
}

interface UserResponse {
  data: {
    users: User[];
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
interface EventResponse {
  data: {
    meetings: Event[];
  };
  message: string;
  status: string;
}
interface User{
  id: number;
  username: string;
  email: string;
  role?: string
}
interface Tag{
  tag_name: string;
  tag_desc: string;
}
interface Userinfo{
  username: string;
  email: string;
}

const TODAY_STR = new Date().toISOString().replace(/T.*$/, ''); // YYYY-MM-DD of today
//=======================================================
@Component({
  selector: 'app-room-scheduler',
  standalone: true,
  imports: [FullCalendarModule,
     FormsModule, 
     PopUpFormComponent, 
     MatFormFieldModule,
     MatInputModule, 
     MatCardModule, 
     CommonModule,
     MatTooltipModule,
     MatSelectModule, 
     MatButtonModule,
     MatIconModule,
     PopUpDeleteConfirmComponent,
     MatDatepickerModule],
  providers: [provideNativeDateAdapter()],
  templateUrl: './room-scheduler.component.html',
  styleUrl: './room-scheduler.component.css',
})
export class RoomSchedulerComponent implements OnInit{
  private pollingSubscription: Subscription | null = null;
  constructor(private changeDetector: ChangeDetectorRef, private router: Router, private dialog: MatDialog, private userService: UserService, private authService: AuthService, private itemService: ItemService) {
  }
  TagsTable: {[tag_id: number]: Tag} = {};//get from api
  UserTable: {[user_id: number]: Userinfo} = {};//get from api
  TagData: CodeValue[] = [];//get from api
  UserData: User[] = [];//get from api
  RoomData: Room[] = [];//get from api
  EventData: Event[] = [];//get from api
  selectedTags: number[] = [];//get from RoomData
  capacity:number = 0;//search by capacity
  filteredRooms: Room[] = [];//search by tags and capacity and time
  CurrentUser: any;//get from auth service
  selectedRoom:Room = {id: 0, room_name: '', type: '', rules: [], capacity: 0};//get from api
  isLogin = false;//get from auth service
  isMonthView: boolean = false;//initial view is week
  startDate: Date= new Date();
  startTime: string =  '';
  endDate: Date= new Date();
  endTime: string = ''; 
  isSearchContainerOpen = false;
  roomDescriptions: { [code: string]: string } = {'projector':'This room is equipped with a projector, which can be used for presentations or visual displays.','food':'can eat in the room.'};//get from api
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
    editable: this.isLogin,
    selectable: this.isLogin,
    selectMirror: true,
    dayMaxEvents: 2,
    allDaySlot: false,
    selectOverlap: false,
    eventOverlap: false,
    validRange: {
      start: '8:00:00',
      end: '23:59:59'
    },
    slotMaxTime: '23:59:59',
    slotMinTime: '8:00:00',
    firstDay: 7,
    select: this.handleDateSelect.bind(this),
    eventClick: this.handleEventClick.bind(this),
    eventsSet: this.handleEvents.bind(this),
    datesSet: this.handleViewChange.bind(this),
    eventDrop: function(info){
      const minTime = '08:00';
      const maxTime = '24:00';

      if (info.event.start&&info.event.end) {
        const eventStartTime = info.event.start.toLocaleTimeString('en-US', { hour12: false, hour: '2-digit', minute: '2-digit' });
        const eventEndTime = info.event.end.toLocaleTimeString('en-US', { hour12: false, hour: '2-digit', minute: '2-digit' });
        console.log(eventStartTime);
        console.log(eventEndTime);
        if (eventEndTime < eventStartTime||eventStartTime < minTime || eventEndTime > maxTime) {
          info.revert();
        }
      }else{
        //put request
      }
    },
    eventResize: function(info){
        //put request and add modify event
    },
    selectAllow: function (info) {
      return (info.start >= new Date());
    }
  };
  ngOnInit(){
    this.CurrentUser = this.userService.getUser();
    console.log(this.CurrentUser);
    this.isLogin = this.userService.isLoggedIn();
    console.log(this.isLogin);
    this.calendarOptions.editable = this.isLogin;
    this.calendarOptions.selectable = this.isLogin;
    this.itemService.getAllRooms().subscribe((response: RoomResponse) => {
      if(response.status === 'success'){
        this.RoomData = response.data.rooms;
        this.selectedRoom = this.RoomData[0];
        this.filteredRooms = this.RoomData;
      }else{
        console.log('get room data failed');
      }
    });
    this.itemService.getAllTags().subscribe((response: CodeTypesResponse) => {
      if (response.status === 'success') {
        const filter:CodeType|undefined = response.data.code_types.find((code_type: CodeType) => code_type.id === 1);
        console.log(response);
        if(filter != undefined){
          console.log(filter);
          this.TagData = filter.code_values;
          this.TagData.forEach(tag => {
            this.TagsTable[tag.id] = {tag_name: tag.code_value, tag_desc: tag.code_value_desc};
          });
        }
      }else{
        console.log('get tag data failed');
      }
    });
    this.itemService.getAllUsers().subscribe((response: UserResponse) => {
      if(response.status === 'success'){
        this.UserData = response.data.users;
        this.UserData.forEach(user => {
          this.UserTable[user.id] = {username: user.username, email: user.email};
        });
      }else{
        console.log('get user data failed');
      }
    });
    this.CALLEvents(1);
  }
  applyFilter() {
    //filter by tags, capacity and time is valid between startDateTime and endDateTime
    this.filteredRooms = this.RoomData.filter(room => {
      // Check if the room's rules include all selected tags
      const tagsMatch = this.selectedTags.every(tag => room.rules.includes(tag));
      // Check if the room's capacity is sufficient
      const capacityMatch = room.capacity >= this.capacity;

      //check if the room is available between startDate: Date= new Date();startTime: string =  '';endDate: Date= new Date();endTime: string = '';
      //if time is not filled, return true 
      if(this.startDate === undefined || this.startTime === '' || this.endDate === undefined || this.endTime === ''){
        return tagsMatch && capacityMatch;
      }
      //if time is filled, check if the room is available
      const startDateTime = new Date(this.startDate.toDateString() + ' ' + this.startTime);
      const endDateTime = new Date(this.endDate.toDateString() + ' ' + this.endTime);
      // Check if the room is available between startDateTime and endDateTime 
      //by checking if the room has any events that overlap with the selected time
      const events = this.EventData.filter(event => event.room_id === room.id);
      const isAvailable = events.every(event => {
        const eventStart = new Date(event.start_time);
        const eventEnd = new Date(event.end_time);
        return startDateTime >= eventEnd || endDateTime <= eventStart;
      });

      return tagsMatch && capacityMatch && isAvailable;
    });

    //filter by time
    this.selectedRoom = this.filteredRooms[0];
    console.log(this.selectedRoom);
    this.handleRoomChange(this.selectedRoom);
  }
  //event filter by room_id
  CALLEvents(room_id: number){
    //set polling to get event data
    this.pollingSubscription = interval(1000).pipe(
      switchMap(() => this.itemService.getMeetingByRoomIdAndTime(room_id, '2000-01-01', '3000-01-01'))
    ).subscribe((response: EventResponse) => {
      if (response.status === 'success') {
        this.EventData = response.data.meetings;
        this.calendarOptions.events = this.EventData.map(event => {
          return {
            id: event.id,
            title: event.title,
            organizer: event.organizer,
            description: event.description,
            participants: event.participants,
            start: event.start_time,
            end: event.end_time,
            allDay: false
          };
        });
      } else {
        console.log('get event data failed');
      }
    });
  }
  ngOnDestroy() {
      if (this.pollingSubscription) {
        this.pollingSubscription.unsubscribe();
      }
  }
  calendarVisible = signal(true);

  currentEvents = signal<EventApi[]>([]);

  handleCalendarToggle() {
    this.calendarVisible.update((bool) => !bool);
  }
  handleDateSelect(selectInfo: DateSelectArg) {
    const calendarApi = selectInfo.view.calendar;
    if (selectInfo.allDay) {
      //set allday false and set start time to 7:00 and end time to 23:59
      let temp:String = selectInfo.startStr;
      selectInfo.startStr = temp + 'T07:00:00';
      selectInfo.endStr = temp + 'T24:00:00';
      selectInfo.allDay = false;
    }
    const dialogRef = this.dialog.open(PopUpFormComponent, {
      width: '50%',
      height: '50%',
      data: {title: '', organizer: this.CurrentUser.id, description: '', startTime: selectInfo.startStr, endTime: selectInfo.endStr, participants: [this.CurrentUser.id]},
    });
    dialogRef.afterClosed().subscribe(data => {
      if (data && data.title !== undefined && data.title !== '') {
        console.log(data);
        this.itemService.postMeeting( data.description, data.endTime, data.organizer, data.participants, this.selectedRoom.id, data.startTime, 'approved', data.title).subscribe((response: any) => {
          if(response.status === 'success'){
            calendarApi.addEvent({
              // id: createEventId(),
              title: data.title,
              organizer: data.organizer,
              description: data.description,
              participants: data.participants,
              start: selectInfo.startStr,
              end: selectInfo.endStr,
              allDay: selectInfo.allDay
            });
          }else{
            console.log('post event failed');
          }
        });
      }else{
        calendarApi.unselect();
      }
    });
  }
  handleViewChange(viewInfo: DatesSetArg) {
    this.isMonthView = viewInfo.view.type === 'dayGridMonth';
  }
  handleEventClick(clickInfo: EventClickArg) {
    const dialogRef = this.dialog.open(PopUpFormComponent, {
      width: '50%',
      height: '50%',
      data: {title: clickInfo.event.title,organizer: clickInfo.event.extendedProps['organizer'],description: clickInfo.event.extendedProps['description'], startTime: clickInfo.event.startStr, endTime: clickInfo.event.endStr, participants: clickInfo.event.extendedProps['participants']},
    });
    dialogRef.afterClosed().subscribe(data => {
      if (data && data.title !== undefined && data.title !== '') {
        clickInfo.event.setProp('title', data.title);
        this.itemService.putMeeting(clickInfo.event.id, data.description, data.endTime, data.organizer, data.participants, this.selectedRoom.id, data.startTime, 'approved', data.title).subscribe((response: any) => {
          if(response.status === 'success'){
            clickInfo.event.setExtendedProp('description', data.description);
            clickInfo.event.setExtendedProp('participants', data.participants);
            clickInfo.event.setStart(data.startTime);
            clickInfo.event.setEnd(data.endTime);
          }else{
            console.log('put event failed');
          }
        });
      }
    });
  }
  deleteEvent(event: EventApi, $event: MouseEvent) {
    $event.stopPropagation(); 
    const dialogRef = this.dialog.open(PopUpDeleteConfirmComponent, {
      width: '50%',
      height: '50%',
    });
    dialogRef.afterClosed().subscribe(data => {
      if (data) {
        this.itemService.deleteMeeting(event.id).subscribe((response: any) => {
          if (response.status === 'success') {
            event.remove();
          } else {
            console.log('delete event failed');
          }
        });
        }
    });
  }
  handleEvents(events: EventApi[]) {
    this.currentEvents.set(events);
    this.changeDetector.detectChanges();
  }
  handleRoomChange(room: Room) {
      this.selectedRoom = room;
      console.log("change" + room);
      this.CALLEvents(this.selectedRoom.id);
  }
  toggleSearchContainer() {
    this.isSearchContainerOpen = !this.isSearchContainerOpen;
    //reset search
    this.selectedTags = [];
    this.capacity = 0;
    this.startDate = new Date();
    this.startTime = '';
    this.endDate = new Date();
    this.endTime = '';
    this.filteredRooms = this.RoomData;
    this.selectedRoom = this.filteredRooms[0];
    this.handleRoomChange(this.selectedRoom);
  }
}


