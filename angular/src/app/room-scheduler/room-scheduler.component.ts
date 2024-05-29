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
import { S, cl, co, s } from '@fullcalendar/core/internal-common';
import { MatInputModule } from '@angular/material/input';
import { MatIconModule } from '@angular/material/icon';
import {provideNativeDateAdapter} from '@angular/material/core';
import { UserService } from '../API/user.service';
import { AuthService } from '../API/auth.service';
import { ItemService } from '../API/item.service';
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
  start: string;
  end: string;
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
  constructor(private changeDetector: ChangeDetectorRef, private router: Router, private dialog: MatDialog, private userService: UserService, private authService: AuthService, private itemService: ItemService) {
  }
  fakeData: Room[] = [];//for testing
  fakeEvents: Event[] = [];//for testing
  fakeTags: CodeValue[] = [];//for testing
  fakeUsers: User[] = [];//for testing
  fakeOrganizer: User = {id: 1, username: 'user1', email: '1@1.com', role: 'user'};//for testing
  fakeTagsTable: {[tag_id: number]: Tag}=[];
  TagData: CodeValue[] = [];//get from api
  UserData: User[] = [];//get from api
  EventData: Event[] = [];//get from api
  roomTable: {[room_name:string]: number} = {'room1': 10, 'room2': 20, 'room3': 30, 'room4': 40, 'room5': 50};//get from api
  availableTags:String[] = ['projector', 'food'];//need to get from api
  selectedTags: number[] = [];//get from RoomData
  capacity:number = 0;//search by capacity
  filteredRooms: Room[] = [];//search by tags and capacity and time
  CurrentUser: any;//get from auth service
  selectedRoom = 'room1';//initial room
  isLogin = true;//get from auth service
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
      start: '7:00:00',
      end: '23:59:59'
    },
    slotMaxTime: '23:59:59',
    slotMinTime: '7:00:00',
    firstDay: 7,
    select: this.handleDateSelect.bind(this),
    eventClick: this.handleEventClick.bind(this),
    eventsSet: this.handleEvents.bind(this),
    datesSet: this.handleViewChange.bind(this),
    eventDrop: function(info){
      const minTime = '07:00';
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
    // this.CurrentUser = this.userService.getUser();
    // if(!this.CurrentUser){
    //   this.router.navigate(['/login']);
    // }
    // this.itemService.getAllRooms().subscribe((response: RoomResponse) => {
    //   if(response.status === 'success'){
    //     this.RoomData = response.data.rooms;
    //   }else{
    //     console.log('get room data failed');
    //   }
    // });
    // this.itemService.getAllTags().subscribe((response: CodeTypesResponse) => {
    //   if (response.status === 'success') {
    //     const filter:CodeType|undefined = response.data.code_types.find((code_type: CodeType) => code_type.id === 0);
    //     if(filter != undefined)
    //       this.TagData = filter.code_values;
    //   }else{
    //     console.log('get tag data failed');
    //   }
    // });
    // this.itemService.getAllUsers().subscribe((response: UserResponse) => {
    //   if(response.status === 'success'){
    //     this.UserData = response.data.users;
    //   }else{
    //     console.log('get user data failed');
    //   }
    // });
    // this.itemService.getMeetingByRoomIdAndTime(this.RoomData[0].id, '0000-01-01', '3000-01-01').subscribe((response: EventResponse) => {
    //   if(response.status === 'success'){
    //     this.EventData = response.data.meetings;
    //   }else{
    //     console.log('get event data failed');
    //   }
    // });
    this.fakeData = [
      {id: 1, room_name: 'room1', type: 'meeting', rules: [0], capacity: 10},
      {id: 2, room_name: 'room2', type: 'meeting', rules: [0, 1], capacity: 20},
      {id: 3, room_name: 'room3', type: 'meeting', rules: [1], capacity: 30},
      {id: 4, room_name: 'room4', type: 'meeting', rules: [0], capacity: 40},
      {id: 5, room_name: 'room5', type: 'meeting', rules: [0, 1], capacity: 50},
    ];
    this.fakeEvents = [
      {id: '1', room_id: 1, title: 'meeting1', description: 'meeting1',organizer: 1, start: TODAY_STR + 'T09:00:00', end: TODAY_STR + 'T13:00:00', participants: [1, 2], status_type: 'approved'},
      {id: '2', room_id: 2, title: 'meeting2', description: 'meeting2',organizer: 2, start: TODAY_STR + 'T10:00:00', end: TODAY_STR + 'T14:00:00', participants: [2, 3], status_type: 'approved'},
      {id: '3', room_id: 3, title: 'meeting3', description: 'meeting3',organizer: 3, start: TODAY_STR + 'T11:00:00', end: TODAY_STR + 'T15:00:00', participants: [3, 4], status_type: 'approved'},
      {id: '4', room_id: 4, title: 'meeting4', description: 'meeting4',organizer: 4, start: TODAY_STR + 'T12:00:00', end: TODAY_STR + 'T16:00:00', participants: [5, 1], status_type: 'approved'},
      {id: '5', room_id: 5, title: 'meeting5', description: 'meeting5',organizer: 5, start: TODAY_STR + 'T13:00:00', end: TODAY_STR + 'T17:00:00', participants: [1, 3], status_type: 'approved'},
    ];
    this.fakeTags = [
      {code_type_id: 0, code_value: 'projector', code_value_desc: 'This room is equipped with a projector, which can be used for presentations or visual displays.', id: 0},
      {code_type_id: 0, code_value: 'food', code_value_desc: 'can eat in the room.', id: 1}
    ];
    this.fakeUsers = [
      {id: 1, username: 'user1', email: '1@1.com', role: 'user'},
      {id: 2, username: 'user2', email: '2@2.com', role: 'user'},
      {id: 3, username: 'user3', email: '3@3.com', role: 'user'},
      {id: 4, username: 'user4', email: '4@4.com', role: 'user'},
      {id: 5, username: 'user5', email: '5@5.com', role: 'user'},
    ];
    this.fakeOrganizer = {id: 1, username: 'user1', email: '1@1.com', role: 'user'};
    this.roomTable = {'room1': 1, 'room2': 2, 'room3': 3, 'room4': 4, 'room5': 5};
    //by function
    this.fakeTags.forEach(tag => {
      this.fakeTagsTable[tag.id] = {tag_name: tag.code_value, tag_desc: tag.code_value_desc};
    });
    this.filteredRooms = this.fakeData;
    this.calendarOptions.events = this.CALLEvents(1);
  }
  applyFilter() {
    //filter by tags, capacity and time is valid between startDateTime and endDateTime
    this.filteredRooms = this.fakeData.filter(room => {
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
      const events = this.fakeEvents.filter(event => event.room_id === this.roomTable[room.room_name]);
      const isAvailable = events.every(event => {
        const eventStart = new Date(event.start);
        const eventEnd = new Date(event.end);
        return startDateTime >= eventEnd || endDateTime <= eventStart;
      });

      return tagsMatch && capacityMatch && isAvailable;
    });

    //filter by time
    this.selectedRoom = this.filteredRooms[0].room_name;
    console.log(this.selectedRoom);
    this.handleRoomChange(this.selectedRoom);
  }
  //event filter by room_id
  CALLEvents(room_id: number){
    console.log(this.fakeEvents.filter((event) => event.room_id === room_id));
    return this.fakeEvents.filter((event) => event.room_id === room_id);
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
      data: {title: '', organizer: 1, description: '', startTime: selectInfo.startStr, endTime: selectInfo.endStr, participants: [1]},
    });

    dialogRef.afterClosed().subscribe(data => {
      if (data && data.title !== undefined && data.title !== '') {
        console.log(data);
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
        clickInfo.event.setExtendedProp('description', data.description);
        clickInfo.event.setExtendedProp('participants', data.participants);
        //add put request
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
        event.remove();
        //add delete request
      }
    });
  }
  handleEvents(events: EventApi[]) {
    this.currentEvents.set(events);
    this.changeDetector.detectChanges();
  }
  handleRoomChange(room: string) {
      this.selectedRoom = room;
      let eventsToDisplay: any[] = [];
      console.log("change" + room);
      eventsToDisplay = this.CALLEvents(this.roomTable[room]);
      this.calendarOptions = {
        ...this.calendarOptions,
        events: eventsToDisplay
      };
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
    this.filteredRooms = this.fakeData;
    this.selectedRoom = this.filteredRooms[0].room_name;
    this.handleRoomChange(this.selectedRoom);
  }
}


