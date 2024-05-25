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
import { PopUpFormComponent } from './pop-up-form/pop-up-form.component';
import { PopUpDetailsComponent } from './pop-up-details/pop-up-details.component';
import { PopUpDeleteConfirmComponent } from './pop-up-delete-confirm/pop-up-delete-confirm.component';
import { MatDialog } from '@angular/material/dialog';
import { CommonModule } from '@angular/common';
import {MatDatepickerModule} from '@angular/material/datepicker';
// import listPlugin from '@fullcalendar/list';
import { INITIAL_EVENTS, createEventId, SECOND_EVENTS} from './event-utils';//test use
import { cl, co, s } from '@fullcalendar/core/internal-common';
import { RoomAPIService } from './room-api.service';
import { MatInputModule } from '@angular/material/input';
import { MatIconModule } from '@angular/material/icon';
import {provideNativeDateAdapter} from '@angular/material/core';
//need to save in a interface file
interface Room {
  room_name: string;
  type: string;
  rules: string[];
  capacity: number;
}

interface RoomEvent {
  id: string;
  room_id: number;
  title: string;
  description: string;
  start: string;
  end: string;
}

interface Tag{
  type_name: string;
  type_desc: string;
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
     PopUpDetailsComponent,
     MatTooltipModule,
     MatSelectModule, 
     MatButtonModule,
     MatIconModule,
     PopUpDeleteConfirmComponent,
     MatDatepickerModule],
  templateUrl: './room-scheduler.component.html',
  styleUrl: './room-scheduler.component.css',
  providers: [provideNativeDateAdapter()]
})
export class RoomSchedulerComponent implements OnInit{
  constructor(private changeDetector: ChangeDetectorRef, private router: Router, private dialog: MatDialog, private roomAPIService: RoomAPIService) {
  }
  fakeData: Room[] = [];
  fakeEvents: RoomEvent[] = [];
  roomTable: {[room_name:string]: number} = {'room1': 10, 'room2': 20, 'room3': 30, 'room4': 40, 'room5': 50};//get from api
  availableTags:String[] = ['projector', 'food'];//need to get from api
  selectedTags: string[] = [];//get from RoomData
  capacity:number = 0;//search by capacity
  filteredRooms: Room[] = [];//search by tags and capacity and time
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
    events: [],
    weekends: true,
    editable: this.isLogin,
    selectable: this.isLogin,
    selectMirror: true,
    dayMaxEvents: 2,
    allDaySlot: false,
    selectOverlap: false,
    slotMaxTime: '24:00:00',
    slotMinTime: '7:00:00',
    firstDay: 7,
    select: this.handleDateSelect.bind(this),
    eventClick: this.handleEventClick.bind(this),
    eventsSet: this.handleEvents.bind(this),
    datesSet: this.handleViewChange.bind(this),
    selectAllow: function (info) {
      return (info.start >= new Date());
    }
  };
  ngOnInit(){
    this.roomAPIService.getRooms().subscribe({
      next: (data: any) => {
        console.log(data);
      },
      error: (error: any) => {
        console.log(error);
      }
    });
    this.fakeData = [
      {room_name: 'room1', type: 'meeting', rules: ['projector'], capacity: 10},
      {room_name: 'room2', type: 'meeting', rules: ['projector', 'food'], capacity: 20},
      {room_name: 'room3', type: 'meeting', rules: ['food'], capacity: 30},
      {room_name: 'room4', type: 'meeting', rules: ['projector'], capacity: 40},
      {room_name: 'room5', type: 'meeting', rules: ['projector', 'food'], capacity: 50},
    ];
    this.fakeEvents = [
      {id: '1', room_id: 1, title: 'meeting1', description: 'meeting1', start: TODAY_STR + 'T09:00:00', end: TODAY_STR + 'T13:00:00'},
      {id: '2', room_id: 2, title: 'meeting2', description: 'meeting2', start: TODAY_STR + 'T12:00:00', end: TODAY_STR + 'T15:00:00'},
      {id: '3', room_id: 3, title: 'meeting3', description: 'meeting3', start: TODAY_STR + 'T14:00:00', end: TODAY_STR + 'T17:00:00'},
      {id: '4', room_id: 4, title: 'meeting4', description: 'meeting4', start: TODAY_STR + 'T16:00:00', end: TODAY_STR + 'T19:00:00'},
      {id: '5', room_id: 5, title: 'meeting5', description: 'meeting5', start: TODAY_STR + 'T18:00:00', end: TODAY_STR + 'T21:00:00'},
    ];
    this.roomTable = {'room1': 1, 'room2': 2, 'room3': 3, 'room4': 4, 'room5': 5};
    this.filteredRooms = this.fakeData;
    this.calendarOptions.events = this.filterEvents(this.roomTable[this.selectedRoom]);
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
  filterEvents(room_id: number){
    console.log(this.fakeEvents.filter((event) => event.room_id === room_id));
    return this.fakeEvents.filter((event) => event.room_id === room_id);
  }
  calendarVisible = signal(true);
  title:String|undefined;
  description:String|undefined;

  currentEvents = signal<EventApi[]>([]);

  handleCalendarToggle() {
    this.calendarVisible.update((bool) => !bool);
  }
  handleDateSelect(selectInfo: DateSelectArg) {
    const calendarApi = selectInfo.view.calendar;

    const dialogRef = this.dialog.open(PopUpFormComponent, {
      width: '50%',
      height: '50%',
      data: {title: this.title, description: this.description, startTime: selectInfo.startStr, endTime: selectInfo.endStr},
    });

    dialogRef.afterClosed().subscribe(data => {
      if (data && data.title !== undefined && data.title !== '') {
        calendarApi.addEvent({
          id: createEventId(),
          title: data.title,
          description: data.description,
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
    const dialogRef = this.dialog.open(PopUpDetailsComponent, {
      width: '50%',
      height: '50%',
      data: {title: clickInfo.event.title, description: clickInfo.event.extendedProps['description'], startTime: clickInfo.event.startStr, endTime: clickInfo.event.endStr},
    });
    dialogRef.afterClosed().subscribe(data => {
      if (data && data.title !== undefined && data.title !== '') {
        clickInfo.event.setProp('title', data.title);
        clickInfo.event.setExtendedProp('description', data.description);
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
      eventsToDisplay = this.filterEvents(this.roomTable[room]);
      this.calendarOptions = {
        ...this.calendarOptions,
        events: eventsToDisplay
      };
  }
  toggleSearchContainer() {
    this.isSearchContainerOpen = !this.isSearchContainerOpen;
  }
}


