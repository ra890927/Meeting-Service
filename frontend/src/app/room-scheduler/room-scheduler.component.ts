import { Component, signal, ChangeDetectorRef, OnInit,inject } from '@angular/core';
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
import { R, S, cl, co, ez, s } from '@fullcalendar/core/internal-common';
import { MatInputModule } from '@angular/material/input';
import { MatIconModule } from '@angular/material/icon';
import {provideNativeDateAdapter} from '@angular/material/core';
import { UserService } from '../API/user.service';
import { AuthService } from '../API/auth.service';
import { ItemService } from '../API/item.service';
import { interval, startWith, Subscription, switchMap } from 'rxjs';
import { F } from '@angular/cdk/keycodes';
import { FileService } from '../API/file.service';
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
  id: number;
  tag: string;
  description: string;
  codeTypeId: number;
}
interface fileUpload{
  file: File;
  file_name: string;
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
  //inject
  is = inject(ItemService);
  constructor(private changeDetector: ChangeDetectorRef, private router: Router, private dialog: MatDialog, private userService: UserService, private authService: AuthService, private itemService: ItemService, private fileService: FileService) {
  }
  isadmin = false;
  error = signal<string | null>(null);
  success = signal<string | null>(null);
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
    slotLabelFormat: {
      hour: '2-digit',
      minute: '2-digit',
      hour12: false
    },
    defaultAllDay: false,
    nowIndicator: true,
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
      end: '23:30:00'
    },
    slotMaxTime: '23:30:00',
    slotMinTime: '8:00:00',
    firstDay: 7,
    
    select: this.handleDateSelect.bind(this),
    eventClick: this.handleEventClick.bind(this),
    eventsSet: this.handleEvents.bind(this),
    datesSet: this.handleViewChange.bind(this),
    eventDrop: (info) =>{
      const minTime = '08:00';
      const maxTime = '23:30';
      console.log(info.event.id);
      if (info.event.start&&info.event.end) {
        const eventStartTime = info.event.start.toLocaleTimeString('en-US', { hour12: false, hour: '2-digit', minute: '2-digit' });
        const eventEndTime = info.event.end.toLocaleTimeString('en-US', { hour12: false, hour: '2-digit', minute: '2-digit' });
        if (eventEndTime < eventStartTime||eventStartTime < minTime || eventEndTime > maxTime || info.event.start < new Date() || info.event.extendedProps['organizer'] !== this.CurrentUser.id) {
          this.error.set('time is invalid or you are not the organizer.');
          setTimeout(() => {
            this.error.set(null);
          }, 1000); 
          this.success.set(null);
          info.revert();
        }else{
          this.itemService.putMeeting(info.event.id, info.event.extendedProps['description'], info.event.endStr, info.event.extendedProps['organizer'], info.event.extendedProps['participants'], this.selectedRoom.id, info.event.startStr, 'approved', info.event.title).subscribe((response: any) => {
            if(response.status === 'success'){
              info.event.setStart(info.event.startStr);
              info.event.setEnd(info.event.endStr);
              this.success.set('Event updated successfully.');
              setTimeout(() => {
                this.success.set(null);
              }, 1000);  
              this.error.set(null);
            }else{
              info.revert();
              this.error.set('Failed to update event.');
              setTimeout(() => {
                this.error.set(null);
              }, 1000); 
              this.success.set(null);
            }
          });
        }
      }else{
        this.error.set('time is invalid');
        setTimeout(() => {
          this.error.set(null);
        }, 1000); 
        this.success.set(null);
        info.revert();
      }
    },
    eventResize:(info) =>{
        const minTime = '08:00';
        const maxTime = '23:30';
        console.log(info.event.id);
        if (info.event.start&&info.event.end) {
          if(info.event.start.getDate() !== info.event.end.getDate()){
            this.error.set('Event must be within the same day');
            setTimeout(() => {
              this.error.set(null);
            }, 1000);
            this.success.set(null);
            info.revert();
            return;
          }
          const eventStartTime = info.event.start.toLocaleTimeString('en-US', { hour12: false, hour: '2-digit', minute: '2-digit' });
          const eventEndTime = info.event.end.toLocaleTimeString('en-US', { hour12: false, hour: '2-digit', minute: '2-digit' });
          if (eventEndTime < eventStartTime||eventStartTime < minTime || eventEndTime > maxTime || info.event.start < new Date() || info.event.extendedProps['organizer'] !== this.CurrentUser.id) {
            this.error.set('time is invalid or you are not the organizer.');
            setTimeout(() => {
              this.error.set(null);
            }, 1000); 
            info.revert();
          }else{
            console.log(info.event.id);
            this.itemService.putMeeting(info.event.id, info.event.extendedProps['description'], info.event.endStr, info.event.extendedProps['organizer'], info.event.extendedProps['participants'], this.selectedRoom.id, info.event.startStr, 'approved', info.event.title).subscribe((response: any) => {
              if(response.status === 'success'){
                info.event.setStart(info.event.startStr);
                info.event.setEnd(info.event.endStr);
                this.success.set('Event updated successfully.');
                setTimeout(() => {
                  this.success.set(null);
                }, 1000); 
                this.error.set(null);
              }else{
                this.error.set('Failed to update event.');
                setTimeout(() => {
                  this.error.set(null);
                }, 1000); 
                this.success.set(null);
                info.revert();
              }
            });
          }
        }else{
          this.error.set('time is invalid');
          setTimeout(() => {
            this.error.set(null);
          }, 1000); 
          this.success.set(null);
          info.revert();
        }
    },
    selectAllow: function (info) {
      return (info.start >= new Date());
    }
  };
  ngOnInit(){
    if(!this.userService.isLoggedIn()){
      this.CurrentUser = {id: 0, username: '', email: '', role: ''};
    }
    else{
      this.CurrentUser = this.userService.getUser();
    }
    // this.CurrentUser = this.userService.getUser();
    this.isLogin = this.userService.isLoggedIn();
    this.calendarOptions.editable = this.isLogin;
    this.calendarOptions.selectable = this.isLogin;
    this.itemService.getAllRooms().subscribe((response: RoomResponse) => {
      if(response.status === 'success'){
        this.RoomData = response.data.rooms;
        this.selectedRoom = this.RoomData[0];
        this.filteredRooms = this.RoomData;
        this.CALLEvents(this.selectedRoom.id);
      }else{
        console.log('get room data failed');
      }
    });
    this.itemService.getAllTags().subscribe((response: CodeValue[]) => {
      if (response) {
          this.TagData = response
          this.TagData.forEach(tag => {
            this.TagsTable[tag.id] = {tag_name: tag.tag, tag_desc: tag.description};
          });
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
    this.authService.whoami().subscribe((response: any) => {
      if(response.status === 'success'){
        this.isadmin = this.CurrentUser.role === 'admin';
      }else{
        console.log('get user data failed');
      }
    });
  }
  errorClick(){
    // 設置一個超時計時器
  setTimeout(() => {
    this.error.set(null);
  }, 1000); // 10000 毫秒 = 10 秒
    // this.error.set(null);
  }
  successClick(){
    this.success.set(null);
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
    this.handleRoomChange(this.selectedRoom);
  }
  //event filter by room_id
  CALLEvents(room_id: number){
    //set polling to get event data
    this.pollingSubscription = interval(5000).pipe(
      startWith(0),
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
      selectInfo.startStr = temp + 'T08:00:00';
      selectInfo.endStr = temp + 'T11:30:00';
      selectInfo.allDay = false;
    }
    if(selectInfo.start.getDate() !== selectInfo.end.getDate()){
      this.error.set('Event must be within the same day');
      setTimeout(() => {
        this.error.set(null);
      }, 1000); 
      this.success.set(null);
      calendarApi.unselect();
      return;
    }
    const dialogRef = this.dialog.open(PopUpFormComponent, {
      width: '50%',
      height: '50%',
      data: {title: '', organizer: this.CurrentUser.id, description: '', startTime: selectInfo.startStr, endTime: selectInfo.endStr, participants: [this.CurrentUser.id], files: []},
    });
    dialogRef.afterClosed().subscribe(data => {
      if (data && data.title !== undefined && data.title !== '') {
        console.log(data);
        this.itemService.postMeeting( data.description, data.endTime, data.organizer, data.participants, this.selectedRoom.id, data.startTime, 'approved', data.title).subscribe((response: any) => {
          const meeting_id = response.data.meeting.id;

          if(response.status === 'success'){
            if(data.files.length > 0){
              data.files.forEach((file: File) => {
                this.fileService.uploadFile(file, meeting_id).subscribe((response: any) => {
                  if(response.status === 'success'){
                    console.log('upload file success');
                  }else{
                    console.log('upload file failed');
                  }
                });
              });
            }
            this.fileService
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
            this.success.set('Event created successfully.');
            setTimeout(() => {
              this.success.set(null);
            }, 1000); 
            this.error.set(null);
          }else{
            this.error.set('Failed to create event.');
            setTimeout(() => {
              this.error.set(null);
            }, 1000); 
            this.success.set(null);
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
      data: {id: clickInfo.event.id, title: clickInfo.event.title,organizer: clickInfo.event.extendedProps['organizer'],description: clickInfo.event.extendedProps['description'], startTime: clickInfo.event.startStr, endTime: clickInfo.event.endStr, participants: clickInfo.event.extendedProps['participants']},
    });
    dialogRef.afterClosed().subscribe(data => {
      if (data && data.title !== undefined && data.title !== '') {
        clickInfo.event.setProp('title', data.title);
        this.itemService.putMeeting(clickInfo.event.id, data.description, data.endTime, data.organizer, data.participants, this.selectedRoom.id, data.startTime, 'approved', data.title).subscribe((response: any) => {
          if(response.status === 'success'){
            this.success.set('Event updated successfully.');
            setTimeout(() => {
              this.success.set(null);
            }, 1000); 
            this.error.set(null);
            clickInfo.event.setExtendedProp('description', data.description);
            clickInfo.event.setExtendedProp('participants', data.participants);
            clickInfo.event.setStart(data.startTime);
            clickInfo.event.setEnd(data.endTime);
          }else{
            this.error.set('Failed to update event.');
            setTimeout(() => {
              this.error.set(null);
            }, 1000); 
            this.success.set(null);
            console.log('put event failed');
          }
        });
      }
    });
  }
  deleteEvent(event: EventApi, $event: MouseEvent) {
    $event.stopPropagation(); 
    console.log(event.extendedProps['organizer'] + ' ' + this.CurrentUser.id + ' ' + this.isadmin);
    if(event.extendedProps['organizer'] !== this.CurrentUser.id && !this.isadmin){
      return;
    }
    const dialogRef = this.dialog.open(PopUpDeleteConfirmComponent, {
      width: '250px',
    });
    dialogRef.afterClosed().subscribe(data => {
      if (data) {
        this.itemService.deleteMeeting(event.id).subscribe((response: any) => {
          if (response.status === 'success') {
            console.log('delete event success');
            this.error.set(null);
            this.success.set('Event deleted successfully.');
            setTimeout(() => {
              this.success.set(null);
            }, 1000); 
            event.remove();
          } else {
            this.error.set('Failed to delete event.');
            setTimeout(() => {
              this.error.set(null);
            }, 1000); 
            this.success.set(null);
            console.log('delete event failed');
          }
        });
        }
    });
  }
  handleEvents(events: EventApi[]) {
    console.log(this.currentEvents());
    this.currentEvents.set(events);
    this.changeDetector.detectChanges();
  }
  handleRoomChange(room: Room) {
      if(this.pollingSubscription){
        this.pollingSubscription.unsubscribe();
      }
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


